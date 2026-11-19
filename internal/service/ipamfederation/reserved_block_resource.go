package ipamfederation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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

var (
	_ resource.Resource                   = &ReservedBlockResource{}
	_ resource.ResourceWithValidateConfig = &ReservedBlockResource{}
	_ resource.ResourceWithConfigure      = &ReservedBlockResource{}
	_ resource.ResourceWithImportState    = &ReservedBlockResource{}
	_ resource.ResourceWithIdentity       = &ReservedBlockResource{}
)

func NewReservedBlockResource() resource.Resource {
	return &ReservedBlockResource{}
}

type ReservedBlockResource struct {
	backend core.BackendType
	service coresvc.ReservedBlockService
	client  *uddiclient.APIClient
}

func (r *ReservedBlockResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_reserved_block"
	resp.ResourceBehavior = resource.ResourceBehavior{
		MutableIdentity: true,
	}
}

func (r *ReservedBlockResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				RequiredForImport: true,
			},
		},
	}
}

func (r *ReservedBlockResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Infoblox ReservedBlock in the UDDI backend.",
		Attributes:          ReservedBlockResourceSchemaAttributes,
	}
}

func (r *ReservedBlockResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.service = coresvc.NewReservedBlockService(r.backend, client.NIOS, client.UDDI)
	r.client = client.UDDI
}

func (r *ReservedBlockResource) retryPolicy(op retry.Operation) retry.Policy {
	return retry.For[coremodel.ReservedBlock](r.backend, op)
}

func (r *ReservedBlockResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data ReservedBlockModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Common backend block validations
	validator.ValidateBackendBlocks(r.backend, types.ObjectNull(map[string]attr.Type{}), data.UDDI, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	ValidateReservedBlock(ctx, data, resp)
}

func (r *ReservedBlockResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ReservedBlockModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	unlock := LockReservedBlockAllocation(ctx, data.UDDI, &resp.Diagnostics)
	defer unlock()
	if resp.Diagnostics.HasError() {
		return
	}

	if uddiModel := flex.ExpandNestedObject[UDDIReservedBlockModel](ctx, data.UDDI, &resp.Diagnostics); uddiModel != nil &&
		!uddiModel.DynamicAllocation.IsNull() && !uddiModel.DynamicAllocation.IsUnknown() {
		r.createNextAvailable(ctx, &data, uddiModel, resp)
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}

	obj := data.Expand(ctx, &resp.Diagnostics, true)
	if resp.Diagnostics.HasError() {
		return
	}

	var (
		apiResp  *coremodel.ReservedBlock
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpCreate), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Create(ctx, obj, &core.Options{
			ReturnFields: ReservedBlockReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ReservedBlock: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

// createNextAvailable allocates the block under a parent Federated Block instead of at a
// static address, porting the legacy bloxone provider. The create-next-available endpoint
// does not persist federated_pool_id and rejects tags, so those are patched in with a
// follow-up Update; if that fails the freshly allocated block is deleted again so nothing
// is left dangling (federated_pool_id is immutable, so a later apply could never patch it in).
func (r *ReservedBlockResource) createNextAvailable(ctx context.Context, data *ReservedBlockModel, uddiModel *UDDIReservedBlockModel, resp *resource.CreateResponse) {
	var alloc dynamicallocation.NextAvailableReservedBlockModel
	resp.Diagnostics.Append(uddiModel.DynamicAllocation.As(ctx, &alloc, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	one := int64(1)
	body := uddiipamfederation.NextAvailableBlockRequest{
		Cidr:  flex.ExpandInt64Pointer(uddiModel.Cidr),
		Count: &one,
	}
	// tags are deliberately not sent here: the create-next-available endpoint rejects any
	// request carrying a tags field (even an empty map) with a misleading 409 "Not sufficient
	// Address space available". They are patched in with the follow-up Update below instead.
	tags := flex.ExpandMapStringAny(ctx, uddiModel.Tags, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !uddiModel.Name.IsNull() && !uddiModel.Name.IsUnknown() {
		body.Name = uddiModel.Name.ValueStringPointer()
	}
	if !uddiModel.Comment.IsNull() && !uddiModel.Comment.IsUnknown() {
		body.Comment = uddiModel.Comment.ValueStringPointer()
	}

	apiRes, _, err := r.client.IPAMFederationAPI.
		NextAvailableReservedBlockAPI.
		CreateNextAvailableReservedBlocks(ctx, alloc.NextAvailableId.ValueString()).
		Body(body).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create next available ReservedBlock: %s", err))
		return
	}
	results := apiRes.GetResults()
	if len(results) == 0 {
		resp.Diagnostics.AddError("Client Error", "Unable to create next available ReservedBlock, no results were returned")
		return
	}
	newId := results[0].GetId()

	// The create-next-available response shape differs from the canonical Read response,
	// so the new block is read back through the normal Read path.
	created, _, err := r.service.Read(ctx, newId, &core.Options{ReturnFields: ReservedBlockReturnFields})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read newly created ReservedBlock: %s", err))
		return
	}

	final := created
	poolConfigured := !uddiModel.FederatedPoolId.IsNull() && !uddiModel.FederatedPoolId.IsUnknown()
	if poolConfigured || len(tags) > 0 {
		// The update is a full PUT, so carry over the mutable fields as the API returned
		// them (address is immutable and rejected on update) and change only
		// federated_pool_id and tags.
		patch := &coremodel.ReservedBlock{UDDI: &coremodel.UDDIReservedBlockExt{}}
		if created.UDDI != nil {
			patch.UDDI.Cidr = created.UDDI.Cidr
			patch.UDDI.Comment = created.UDDI.Comment
			patch.UDDI.Name = created.UDDI.Name
			patch.UDDI.FederatedRealm = created.UDDI.FederatedRealm
			patch.UDDI.Tags = created.UDDI.Tags
		}
		if poolConfigured {
			patch.UDDI.FederatedPoolId = uddiModel.FederatedPoolId.ValueStringPointer()
		}
		if len(tags) > 0 {
			patch.UDDI.Tags = tags
		}
		updated, _, updErr := r.service.Update(ctx, newId, patch, &core.Options{ReturnFields: ReservedBlockReturnFields})
		if updErr != nil {
			if _, delErr := r.service.Delete(ctx, newId); delErr != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf(
					"Unable to set federated_pool_id/tags on next available ReservedBlock (got error: %s), and failed to roll back the created block (got error: %s); it must be deleted manually: %s",
					updErr, delErr, newId))
				return
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to set federated_pool_id/tags on next available ReservedBlock, got error: %s", updErr))
			return
		}
		final = updated
	}

	data.Flatten(ctx, final, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *ReservedBlockResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ReservedBlockModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var (
		apiResp  *coremodel.ReservedBlock
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpRead), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Read(ctx, data.Id.ValueString(), &core.Options{
			ReturnFields: ReservedBlockReturnFields,
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ReservedBlock: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *ReservedBlockResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ReservedBlockModel

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
		apiResp  *coremodel.ReservedBlock
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpUpdate), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Update(ctx, data.Id.ValueString(), obj, &core.Options{
			ReturnFields: ReservedBlockReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ReservedBlock: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *ReservedBlockResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ReservedBlockModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var httpResp *http.Response

	err := retry.Do(ctx, r.retryPolicy(retry.OpDelete), func(ctx context.Context) (int, error) {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ReservedBlock: %s", err))
	}
}

func (r *ReservedBlockResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
