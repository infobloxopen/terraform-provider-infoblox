package ipamfederation

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipamfederation"
	coresvc "github.com/infobloxopen/terraform-provider-infoblox/internal/core/service/ipamfederation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/retry"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	uddiipamfederation "github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

// federatedPoolIdPrefix identifies a next_available_id that names a Federated Pool, as opposed to
// a Federated Block. Next-available allocation from a Federated Pool has no preview (GET) support
// and is routed through a different API service than allocation from a Federated Block.
const federatedPoolIdPrefix = "federation/federated_pool/"

var (
	_ resource.Resource                   = &FederatedBlockResource{}
	_ resource.ResourceWithValidateConfig = &FederatedBlockResource{}
	_ resource.ResourceWithConfigure      = &FederatedBlockResource{}
	_ resource.ResourceWithImportState    = &FederatedBlockResource{}
	_ resource.ResourceWithIdentity       = &FederatedBlockResource{}
	_ resource.ResourceWithModifyPlan     = &FederatedBlockResource{}
)

// resolvePoolId returns the Federated Pool id that governs this block, either from
// federated_pool_id directly, or from next_available_id when it names a Federated Pool
// rather than a Federated Block.
func resolvePoolId(ctx context.Context, m *UDDIFederatedBlockModel, diags *diag.Diagnostics) (string, bool) {
	if !m.FederatedPoolId.IsNull() && !m.FederatedPoolId.IsUnknown() {
		return m.FederatedPoolId.ValueString(), true
	}
	if m.DynamicAllocation.IsNull() || m.DynamicAllocation.IsUnknown() {
		return "", false
	}
	var alloc dynamicallocation.NextAvailableFederatedBlockModel
	diags.Append(m.DynamicAllocation.As(ctx, &alloc, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return "", false
	}
	id := alloc.NextAvailableId.ValueString()
	if strings.HasPrefix(id, federatedPoolIdPrefix) {
		return id, true
	}
	return "", false
}

func NewFederatedBlockResource() resource.Resource {
	return &FederatedBlockResource{}
}

type FederatedBlockResource struct {
	backend core.BackendType
	service coresvc.FederatedBlockService
	client  *uddiclient.APIClient
}

func (r *FederatedBlockResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_federated_block"
	resp.ResourceBehavior = resource.ResourceBehavior{
		MutableIdentity: true,
	}
}

func (r *FederatedBlockResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				RequiredForImport: true,
			},
		},
	}
}

func (r *FederatedBlockResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Infoblox FederatedBlock in the UDDI backend.",
		Attributes:          FederatedBlockResourceSchemaAttributes,
	}
}

func (r *FederatedBlockResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*core.InfobloxClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *core.InfobloxClient, got: %T.", req.ProviderData),
		)
		return
	}

	if client.NIOS != nil {
		r.backend = core.BackendNIOS
	} else {
		r.backend = core.BackendUDDI
	}

	r.service = coresvc.NewFederatedBlockService(r.backend, client.NIOS, client.UDDI)
	r.client = client.UDDI
}

func (r *FederatedBlockResource) retryPolicy(op retry.Operation) retry.Policy {
	return retry.For[coremodel.FederatedBlock](r.backend, op)
}

func (r *FederatedBlockResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data FederatedBlockModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Common backend block validations
	validator.ValidateBackendBlocks(r.backend, types.ObjectNull(map[string]attr.Type{}), data.UDDI, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	ValidateFederatedBlock(ctx, data, resp)
}

// ModifyPlan rejects network_compliance at plan time for a new block in an AWS pool. A
// block takes its provider type from its pool, so the pool is read. If the pool is not
// known yet (it is created in the same run), the block is checked during apply instead, in
// Create. Existing blocks are not checked here, so that a destroy is never blocked by a
// pending change in the configuration; the API rejects any update of an AWS block at apply.
func (r *FederatedBlockResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() || !req.State.Raw.IsNull() || r.client == nil {
		return
	}

	var config FederatedBlockModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	uddiModel := flex.ExpandNestedObject[UDDIFederatedBlockModel](ctx, config.UDDI, &resp.Diagnostics)
	if resp.Diagnostics.HasError() || uddiModel == nil ||
		uddiModel.NetworkCompliance.IsNull() || uddiModel.NetworkCompliance.IsUnknown() {
		return
	}

	poolId, ok := resolvePoolId(ctx, uddiModel, &resp.Diagnostics)
	if resp.Diagnostics.HasError() || !ok {
		return
	}

	poolRes, _, err := r.client.IPAMFederationAPI.
		FederatedPoolAPI.
		Read(ctx, poolId).
		Execute()
	if err != nil {
		// If the pool cannot be read, the create request reports the problem, so it is
		// not checked here.
		return
	}
	pool := poolRes.GetResult()

	if pool.GetProvider() == uddiipamfederation.PROVIDERTYPE_AWS {
		resp.Diagnostics.AddAttributeError(
			path.Root("uddi").AtName("network_compliance"),
			"Unsupported attribute for provider type",
			"network_compliance is not supported for blocks in a pool with provider_type \"AWS\".",
		)
	}
}

func (r *FederatedBlockResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FederatedBlockModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	unlock := LockFederatedBlockAllocation(ctx, data.UDDI, &resp.Diagnostics)
	defer unlock()
	if resp.Diagnostics.HasError() {
		return
	}

	uddiModel := flex.ExpandNestedObject[UDDIFederatedBlockModel](ctx, data.UDDI, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// network_compliance is not supported for blocks in an AWS pool. A block takes its
	// provider from its pool, so look the pool up and reject the attribute before the API
	// does. ModifyPlan already catches this when the pool is known at plan time; this
	// catches the case where the pool is created in the same apply.
	if uddiModel != nil && !uddiModel.NetworkCompliance.IsNull() && !uddiModel.NetworkCompliance.IsUnknown() {
		if poolId, ok := resolvePoolId(ctx, uddiModel, &resp.Diagnostics); ok && !resp.Diagnostics.HasError() {
			poolRes, _, err := r.client.IPAMFederationAPI.FederatedPoolAPI.Read(ctx, poolId).Execute()
			if err == nil {
				pool := poolRes.GetResult()
				if pool.GetProvider() == uddiipamfederation.PROVIDERTYPE_AWS {
					resp.Diagnostics.AddAttributeError(path.Root("uddi").AtName("network_compliance"), "Unsupported attribute for provider type",
						"network_compliance is not supported for blocks in a pool with provider_type \"AWS\".")
					return
				}
			}
		}
	}
	if resp.Diagnostics.HasError() {
		return
	}

	if uddiModel != nil && !uddiModel.DynamicAllocation.IsNull() && !uddiModel.DynamicAllocation.IsUnknown() {
		r.createNextAvailable(ctx, &data, uddiModel, resp)
		return
	}

	obj := data.Expand(ctx, &resp.Diagnostics, true)
	if resp.Diagnostics.HasError() {
		return
	}

	var (
		apiResp  *coremodel.FederatedBlock
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpCreate), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Create(ctx, obj, &core.Options{
			ReturnFields: FederatedBlockReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create FederatedBlock: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

// createNextAvailable allocates the block from next_available_id instead of a static
// address, porting the legacy bloxone provider's routing: allocation targeting a
// Federated Pool id goes through a different API service than allocation under a parent
// Federated Block, and network_compliance does not persist on the create-next-available
// endpoint for the block-routed case, so it is patched in with a follow-up Update.
func (r *FederatedBlockResource) createNextAvailable(ctx context.Context, data *FederatedBlockModel, uddiModel *UDDIFederatedBlockModel, resp *resource.CreateResponse) {
	var alloc dynamicallocation.NextAvailableFederatedBlockModel
	resp.Diagnostics.Append(uddiModel.DynamicAllocation.As(ctx, &alloc, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}
	nextAvailableId := alloc.NextAvailableId.ValueString()

	body := uddiipamfederation.NextAvailableBlockRequest{
		Cidr: flex.ExpandInt64Pointer(uddiModel.Cidr),
	}
	if !uddiModel.Name.IsNull() {
		body.Name = uddiModel.Name.ValueStringPointer()
	}
	if !uddiModel.Comment.IsNull() {
		body.Comment = uddiModel.Comment.ValueStringPointer()
	}
	if !uddiModel.FederatedPoolId.IsNull() && !uddiModel.FederatedPoolId.IsUnknown() {
		body.FederatedPoolId = uddiModel.FederatedPoolId.ValueStringPointer()
	}

	if strings.HasPrefix(nextAvailableId, federatedPoolIdPrefix) {
		// Allocating directly from a Federated Pool: the pool is already the parent, so
		// there is no separate federated_pool_id to reconcile afterwards.
		apiRes, _, err := r.client.IPAMFederationAPI.
			NextAvailablePoolBlockAPI.
			CreateNextAvailableBlocks(ctx, nextAvailableId).
			Body(body).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create next available FederatedBlock: %s", err))
			return
		}
		results := apiRes.GetResults()
		if len(results) == 0 {
			resp.Diagnostics.AddError("Client Error", "Unable to create next available FederatedBlock, no results were returned")
			return
		}
		r.readAndSetState(ctx, data, results[0].GetId(), resp)
		return
	}

	// Allocating under a parent Federated Block. network_compliance does not persist on
	// the create-next-available endpoint, so it is patched in with a follow-up Update
	// call if set.
	configuredNetworkCompliance := uddiModel.NetworkCompliance

	apiRes, _, err := r.client.IPAMFederationAPI.
		NextAvailableFederatedBlockAPI.
		CreateNextAvailableFederatedBlocks(ctx, nextAvailableId).
		Body(body).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create next available FederatedBlock: %s", err))
		return
	}
	results := apiRes.GetResults()
	if len(results) == 0 {
		resp.Diagnostics.AddError("Client Error", "Unable to create next available FederatedBlock, no results were returned")
		return
	}
	newId := results[0].GetId()

	if configuredNetworkCompliance.IsNull() || configuredNetworkCompliance.IsUnknown() {
		r.readAndSetState(ctx, data, newId, resp)
		return
	}

	patch := &coremodel.FederatedBlock{
		UDDI: &coremodel.UDDIFederatedBlockExt{
			NetworkCompliance: ExpandNetworkCompliance(ctx, configuredNetworkCompliance, &resp.Diagnostics),
		},
	}
	if resp.Diagnostics.HasError() {
		return
	}

	updateResp, _, updErr := r.service.Update(ctx, newId, patch, &core.Options{ReturnFields: FederatedBlockReturnFields})
	if updErr != nil {
		// The FederatedBlock was already created successfully above. Save it to state as
		// created (without network_compliance), so it isn't orphaned outside of
		// Terraform; a subsequent apply will retry setting network_compliance via Update.
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to set network_compliance on next available FederatedBlock: %s", updErr))
		r.readAndSetState(ctx, data, newId, resp)
		return
	}

	data.Flatten(ctx, updateResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

// readAndSetState re-fetches the created object by id and writes it into state. The
// create-next-available response shape differs from the canonical Read response, so the
// freshly allocated block is read back through the normal Read path instead.
func (r *FederatedBlockResource) readAndSetState(ctx context.Context, data *FederatedBlockModel, id string, resp *resource.CreateResponse) {
	apiResp, _, err := r.service.Read(ctx, id, &core.Options{ReturnFields: FederatedBlockReturnFields})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read newly created FederatedBlock: %s", err))
		return
	}
	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *FederatedBlockResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FederatedBlockModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var (
		apiResp  *coremodel.FederatedBlock
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpRead), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Read(ctx, data.Id.ValueString(), &core.Options{
			ReturnFields: FederatedBlockReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read FederatedBlock: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *FederatedBlockResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data FederatedBlockModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags := req.State.GetAttribute(ctx, path.Root("id"), &data.Id)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	obj := data.Expand(ctx, &resp.Diagnostics, false)
	if resp.Diagnostics.HasError() {
		return
	}

	var (
		apiResp  *coremodel.FederatedBlock
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpUpdate), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Update(ctx, data.Id.ValueString(), obj, &core.Options{
			ReturnFields: FederatedBlockReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update FederatedBlock: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *FederatedBlockResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FederatedBlockModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var httpResp *http.Response

	// Blocks in cloud provider backed pools (e.g. AWS) can fail to deprovision while the
	// provider is still working on the CIDR, so retry on that specific error for up to 5
	// minutes, on top of the usual transient-error retries.
	deleteRetryPolicy := r.retryPolicy(retry.OpDelete)
	baseRetryable := deleteRetryPolicy.Retryable
	deleteRetryPolicy.Timeout = 5 * time.Minute
	deleteRetryPolicy.Retryable = func(err error) bool {
		if strings.Contains(err.Error(), "Failed to deprovision") {
			return true
		}
		return baseRetryable != nil && baseRetryable(err)
	}

	err := retry.Do(ctx, deleteRetryPolicy, func(ctx context.Context) (int, error) {
		var apiErr error
		httpResp, apiErr = r.service.Delete(ctx, data.Id.ValueString())
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete FederatedBlock: %s", err))
	}
}

func (r *FederatedBlockResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.Identity != nil && req.Identity.Raw.IsKnown() && !req.Identity.Raw.IsNull() {
		diags := req.Identity.GetAttribute(ctx, path.Root("id"), &req.ID)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
