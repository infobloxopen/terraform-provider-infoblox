package rpz

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

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/rpz"
	coresvc "github.com/infobloxopen/terraform-provider-infoblox/internal/core/service/rpz"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/retry"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

var (
	_ resource.Resource                   = &RecordRpzAaaaIpaddressResource{}
	_ resource.ResourceWithValidateConfig = &RecordRpzAaaaIpaddressResource{}
	_ resource.ResourceWithConfigure      = &RecordRpzAaaaIpaddressResource{}
	_ resource.ResourceWithImportState    = &RecordRpzAaaaIpaddressResource{}
	_ resource.ResourceWithIdentity       = &RecordRpzAaaaIpaddressResource{}
)

func NewRecordRpzAaaaIpaddressResource() resource.Resource {
	return &RecordRpzAaaaIpaddressResource{}
}

type RecordRpzAaaaIpaddressResource struct {
	backend core.BackendType
	service coresvc.RecordRpzAaaaIpaddressService
}

func (r *RecordRpzAaaaIpaddressResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_record_rpz_aaaa_ipaddress"
	resp.ResourceBehavior = resource.ResourceBehavior{
		MutableIdentity: true,
	}
}

func (r *RecordRpzAaaaIpaddressResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				RequiredForImport: true,
			},
		},
	}
}

func (r *RecordRpzAaaaIpaddressResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Infoblox RecordRpzAaaaIpaddress in the NIOS backend.",
		Attributes:          RecordRpzAaaaIpaddressResourceSchemaAttributes,
	}
}

func (r *RecordRpzAaaaIpaddressResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.service = coresvc.NewRecordRpzAaaaIpaddressService(r.backend, client.NIOS, client.UDDI)
}

func (r *RecordRpzAaaaIpaddressResource) retryPolicy(op retry.Operation) retry.Policy {
	return retry.For[coremodel.RecordRpzAaaaIpaddress](r.backend, op)
}

func (r *RecordRpzAaaaIpaddressResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data RecordRpzAaaaIpaddressModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Common backend block validations
	validator.ValidateBackendBlocks(r.backend, data.NIOS, types.ObjectNull(map[string]attr.Type{}), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	ValidateRecordRpzAaaaIpaddress(ctx, data, resp)
}

func (r *RecordRpzAaaaIpaddressResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RecordRpzAaaaIpaddressModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Add Terraform Internal ID to ext_attrs
	if r.backend == core.BackendNIOS {
		nios := flex.ExpandNestedObject[NIOSRecordRpzAaaaIpaddressModel](ctx, data.NIOS, &resp.Diagnostics)
		if nios == nil {
			nios = &NIOSRecordRpzAaaaIpaddressModel{}
		}
		nios.ExtAttrs = flex.SetInternalID(ctx, nios.ExtAttrs, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
		data.NIOS = flex.FlattenNestedObject(ctx, nios, NIOSRecordRpzAaaaIpaddressAttrTypes, &resp.Diagnostics)
	}

	obj := data.Expand(ctx, &resp.Diagnostics, true)
	if resp.Diagnostics.HasError() {
		return
	}

	if r.backend == core.BackendNIOS {
		ApplyRecordRpzAaaaIpaddressNIOSUseFlags(ctx, req.Config, obj, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	var (
		apiResp  *coremodel.RecordRpzAaaaIpaddress
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpCreate), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Create(ctx, obj, &core.Options{
			ReturnFields: RecordRpzAaaaIpaddressReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create RecordRpzAaaaIpaddress: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *RecordRpzAaaaIpaddressResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RecordRpzAaaaIpaddressModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if we need to associate internal ID (import flow)
	associateInternalId, diags := req.Private.GetKey(ctx, flex.AssociateInternalIDKey)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var (
		apiResp  *coremodel.RecordRpzAaaaIpaddress
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpRead), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Read(ctx, data.Id.ValueString(), &core.Options{
			ReturnFields: RecordRpzAaaaIpaddressReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			if r.backend == core.BackendNIOS {
				if r.ReadByExtAttrs(ctx, &data, resp) {
					return
				}
			} else {
				resp.State.RemoveResource(ctx)
				return
			}
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read RecordRpzAaaaIpaddress: %s", err))
		return
	}

	// For NIOS verify internal ID matches (handles case where ref changes but resource still exists)
	if r.backend == core.BackendNIOS && associateInternalId == nil && apiResp.NIOS != nil {
		// Get state internal ID
		stateNIOS := flex.ExpandNestedObject[NIOSRecordRpzAaaaIpaddressModel](ctx, data.NIOS, &resp.Diagnostics)
		if stateNIOS == nil || stateNIOS.ExtAttrsAll.IsNull() || stateNIOS.ExtAttrsAll.IsUnknown() {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read RecordRpzAaaaIpaddress because the internal ID (from ext_attrs_all) is missing or invalid.",
			)
			return
		}

		stateExtAttrsAll := stateNIOS.ExtAttrsAll.Elements()
		stateTFID := ""
		if stateTFIDVal, ok := stateExtAttrsAll[flex.TerraformInternalID]; ok {
			if strVal, ok := stateTFIDVal.(types.String); ok {
				stateTFID = strVal.ValueString()
			}
		}

		// Get API internal ID
		apiTFID := ""
		if apiResp.NIOS.ExtAttrs != nil {
			if apiTFIDVal, ok := apiResp.NIOS.ExtAttrs[flex.TerraformInternalID]; ok {
				apiTFID, _ = apiTFIDVal.(string)
			}
		}

		if apiTFID != stateTFID {
			// Mismatch in internal ID, try to find the record using ExtAttrs
			if r.ReadByExtAttrs(ctx, &data, resp) {
				return
			}
		}
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *RecordRpzAaaaIpaddressResource) ReadByExtAttrs(ctx context.Context, data *RecordRpzAaaaIpaddressModel, resp *resource.ReadResponse) bool {
	// Only applicable for NIOS backend
	if r.backend != core.BackendNIOS {
		return false
	}

	nios := flex.ExpandNestedObject[NIOSRecordRpzAaaaIpaddressModel](ctx, data.NIOS, &resp.Diagnostics)
	if nios == nil || nios.ExtAttrsAll.IsNull() || nios.ExtAttrsAll.IsUnknown() {
		return false
	}

	extAttrsAll := nios.ExtAttrsAll.Elements()
	tfInternalIDVal, ok := extAttrsAll[flex.TerraformInternalID]
	if !ok {
		return false
	}
	tfInternalID := tfInternalIDVal.(types.String).ValueString()
	if tfInternalID == "" {
		return false
	}

	// Search for the record using the Terraform Internal ID
	var (
		records  []*coremodel.RecordRpzAaaaIpaddress
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpRead), func(ctx context.Context) (int, error) {
		var apiErr error
		records, httpResp, _, apiErr = r.service.List(ctx, &core.ListOptions{
			ReturnFields: RecordRpzAaaaIpaddressReturnFields,
			ExtAttrFilter: map[string]string{
				flex.TerraformInternalID: tfInternalID,
			},
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to search RecordRpzAaaaIpaddress by extattrs: %s", err))
		return true
	}

	// If not found, remove from state
	if len(records) == 0 {
		resp.State.RemoveResource(ctx)
		return true
	}

	data.Flatten(ctx, records[0], &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return true
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
	return true
}

func (r *RecordRpzAaaaIpaddressResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data RecordRpzAaaaIpaddressModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags := req.State.GetAttribute(ctx, path.Root("id"), &data.Id)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// Check if we need to associate internal ID (import flow)
	associateInternalId, diags := req.Private.GetKey(ctx, flex.AssociateInternalIDKey)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var planExtAttrs types.Map

	// Merge ext_attrs with state ext_attrs_all (inherited + TF ID)
	if r.backend == core.BackendNIOS {
		var stateNIOSObj types.Object
		diags = req.State.GetAttribute(ctx, path.Root("nios"), &stateNIOSObj)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}

		planNIOS := flex.ExpandNestedObject[NIOSRecordRpzAaaaIpaddressModel](ctx, data.NIOS, &resp.Diagnostics)
		stateNIOS := flex.ExpandNestedObject[NIOSRecordRpzAaaaIpaddressModel](ctx, stateNIOSObj, &resp.Diagnostics)
		if planNIOS == nil {
			planNIOS = &NIOSRecordRpzAaaaIpaddressModel{}
		}

		// Preserve the plan ext_attrs (without inherited EAs) for restore after Update
		planExtAttrs = planNIOS.ExtAttrs

		// If this is post-import, add the Terraform Internal ID
		if associateInternalId != nil {
			planNIOS.ExtAttrs = flex.SetInternalID(ctx, planNIOS.ExtAttrs, &resp.Diagnostics)
			if resp.Diagnostics.HasError() {
				return
			}
		}

		// Merge with state ext_attrs_all (inherited EAs)
		if stateNIOS != nil {
			planNIOS.ExtAttrs = flex.MergeEAs(planNIOS.ExtAttrs, stateNIOS.ExtAttrsAll)
		}
		data.NIOS = flex.FlattenNestedObject(ctx, planNIOS, NIOSRecordRpzAaaaIpaddressAttrTypes, &resp.Diagnostics)
	}

	obj := data.Expand(ctx, &resp.Diagnostics, false)
	if resp.Diagnostics.HasError() {
		return
	}

	if r.backend == core.BackendNIOS {
		ApplyRecordRpzAaaaIpaddressNIOSUseFlags(ctx, req.Config, obj, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	var (
		apiResp  *coremodel.RecordRpzAaaaIpaddress
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpUpdate), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Update(ctx, data.Id.ValueString(), obj, &core.Options{
			ReturnFields: RecordRpzAaaaIpaddressReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update RecordRpzAaaaIpaddress: %s", err))
		return
	}

	// Restore the plan ext_attrs (without inherited EAs) so Flatten preserves the user's input
	if r.backend == core.BackendNIOS {
		niosObj := flex.ExpandNestedObject[NIOSRecordRpzAaaaIpaddressModel](ctx, data.NIOS, &resp.Diagnostics)
		if niosObj != nil {
			niosObj.ExtAttrs = planExtAttrs
			data.NIOS = flex.FlattenNestedObject(ctx, niosObj, NIOSRecordRpzAaaaIpaddressAttrTypes, &resp.Diagnostics)
			if resp.Diagnostics.HasError() {
				return
			}
		}
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)

	if associateInternalId != nil {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, flex.AssociateInternalIDKey, nil)...)
	}
}

func (r *RecordRpzAaaaIpaddressResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RecordRpzAaaaIpaddressModel

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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete RecordRpzAaaaIpaddress: %s", err))
	}
}

func (r *RecordRpzAaaaIpaddressResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.Identity != nil && req.Identity.Raw.IsKnown() && !req.Identity.Raw.IsNull() {
		diags := req.Identity.GetAttribute(ctx, path.Root("id"), &req.ID)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), req.ID)...)

	if r.backend == core.BackendNIOS {
		// For NIOS backend, set the associate_internal_id private key
		// This triggers the plan modifier to mark ext_attrs_all as unknown,
		// and the Update method will add the Terraform Internal ID to ext_attrs
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, flex.AssociateInternalIDKey, []byte("true"))...)
	}
}
