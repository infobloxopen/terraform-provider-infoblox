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
	_ resource.Resource                   = &RecordRpzCnameIpaddressResource{}
	_ resource.ResourceWithValidateConfig = &RecordRpzCnameIpaddressResource{}
	_ resource.ResourceWithConfigure      = &RecordRpzCnameIpaddressResource{}
	_ resource.ResourceWithImportState    = &RecordRpzCnameIpaddressResource{}
	_ resource.ResourceWithIdentity       = &RecordRpzCnameIpaddressResource{}
)

func NewRecordRpzCnameIpaddressResource() resource.Resource {
	return &RecordRpzCnameIpaddressResource{}
}

type RecordRpzCnameIpaddressResource struct {
	backend core.BackendType
	service coresvc.RecordRpzCnameIpaddressService
}

func (r *RecordRpzCnameIpaddressResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_record_rpz_cname_ipaddress"
	resp.ResourceBehavior = resource.ResourceBehavior{
		MutableIdentity: true,
	}
}

func (r *RecordRpzCnameIpaddressResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				RequiredForImport: true,
			},
		},
	}
}

func (r *RecordRpzCnameIpaddressResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Infoblox RecordRpzCnameIpaddress in the NIOS backend.",
		Attributes:          RecordRpzCnameIpaddressResourceSchemaAttributes,
	}
}

func (r *RecordRpzCnameIpaddressResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.service = coresvc.NewRecordRpzCnameIpaddressService(r.backend, client.NIOS, client.UDDI)
}

func (r *RecordRpzCnameIpaddressResource) retryPolicy(op retry.Operation) retry.Policy {
	return retry.For[coremodel.RecordRpzCnameIpaddress](r.backend, op)
}

func (r *RecordRpzCnameIpaddressResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data RecordRpzCnameIpaddressModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Common backend block validations
	validator.ValidateBackendBlocks(r.backend, data.NIOS, types.ObjectNull(map[string]attr.Type{}), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	ValidateRecordRpzCnameIpaddress(ctx, data, resp)
}

func (r *RecordRpzCnameIpaddressResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RecordRpzCnameIpaddressModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Add Terraform Internal ID to ext_attrs
	if r.backend == core.BackendNIOS {
		nios := flex.ExpandNestedObject[NIOSRecordRpzCnameIpaddressModel](ctx, data.NIOS, &resp.Diagnostics)
		if nios == nil {
			nios = &NIOSRecordRpzCnameIpaddressModel{}
		}
		nios.ExtAttrs = flex.SetInternalID(ctx, nios.ExtAttrs, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
		data.NIOS = flex.FlattenNestedObject(ctx, nios, NIOSRecordRpzCnameIpaddressAttrTypes, &resp.Diagnostics)
	}

	obj := data.Expand(ctx, &resp.Diagnostics, true)
	if resp.Diagnostics.HasError() {
		return
	}

	if r.backend == core.BackendNIOS {
		ApplyRecordRpzCnameIpaddressNIOSUseFlags(ctx, req.Config, obj, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	var (
		apiResp  *coremodel.RecordRpzCnameIpaddress
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpCreate), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Create(ctx, obj, &core.Options{
			ReturnFields: RecordRpzCnameIpaddressReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create RecordRpzCnameIpaddress: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *RecordRpzCnameIpaddressResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RecordRpzCnameIpaddressModel

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
		apiResp  *coremodel.RecordRpzCnameIpaddress
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpRead), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Read(ctx, data.Id.ValueString(), &core.Options{
			ReturnFields: RecordRpzCnameIpaddressReturnFields,
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read RecordRpzCnameIpaddress: %s", err))
		return
	}

	// For NIOS verify internal ID matches (handles case where ref changes but resource still exists)
	if r.backend == core.BackendNIOS && associateInternalId == nil && apiResp.NIOS != nil {
		// Get state internal ID
		stateNIOS := flex.ExpandNestedObject[NIOSRecordRpzCnameIpaddressModel](ctx, data.NIOS, &resp.Diagnostics)
		if stateNIOS == nil || stateNIOS.ExtAttrsAll.IsNull() || stateNIOS.ExtAttrsAll.IsUnknown() {
			resp.Diagnostics.AddError(
				"Missing Internal ID",
				"Unable to read RecordRpzCnameIpaddress because the internal ID (from ext_attrs_all) is missing or invalid.",
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

func (r *RecordRpzCnameIpaddressResource) ReadByExtAttrs(ctx context.Context, data *RecordRpzCnameIpaddressModel, resp *resource.ReadResponse) bool {
	// Only applicable for NIOS backend
	if r.backend != core.BackendNIOS {
		return false
	}

	nios := flex.ExpandNestedObject[NIOSRecordRpzCnameIpaddressModel](ctx, data.NIOS, &resp.Diagnostics)
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
		records  []*coremodel.RecordRpzCnameIpaddress
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpRead), func(ctx context.Context) (int, error) {
		var apiErr error
		records, httpResp, _, apiErr = r.service.List(ctx, &core.ListOptions{
			ReturnFields: RecordRpzCnameIpaddressReturnFields,
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to search RecordRpzCnameIpaddress by extattrs: %s", err))
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

func (r *RecordRpzCnameIpaddressResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data RecordRpzCnameIpaddressModel

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

		planNIOS := flex.ExpandNestedObject[NIOSRecordRpzCnameIpaddressModel](ctx, data.NIOS, &resp.Diagnostics)
		stateNIOS := flex.ExpandNestedObject[NIOSRecordRpzCnameIpaddressModel](ctx, stateNIOSObj, &resp.Diagnostics)
		if planNIOS == nil {
			planNIOS = &NIOSRecordRpzCnameIpaddressModel{}
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
		data.NIOS = flex.FlattenNestedObject(ctx, planNIOS, NIOSRecordRpzCnameIpaddressAttrTypes, &resp.Diagnostics)
	}

	obj := data.Expand(ctx, &resp.Diagnostics, false)
	if resp.Diagnostics.HasError() {
		return
	}

	if r.backend == core.BackendNIOS {
		ApplyRecordRpzCnameIpaddressNIOSUseFlags(ctx, req.Config, obj, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	var (
		apiResp  *coremodel.RecordRpzCnameIpaddress
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpUpdate), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Update(ctx, data.Id.ValueString(), obj, &core.Options{
			ReturnFields: RecordRpzCnameIpaddressReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update RecordRpzCnameIpaddress: %s", err))
		return
	}

	// Restore the plan ext_attrs (without inherited EAs) so Flatten preserves the user's input
	if r.backend == core.BackendNIOS {
		niosObj := flex.ExpandNestedObject[NIOSRecordRpzCnameIpaddressModel](ctx, data.NIOS, &resp.Diagnostics)
		if niosObj != nil {
			niosObj.ExtAttrs = planExtAttrs
			data.NIOS = flex.FlattenNestedObject(ctx, niosObj, NIOSRecordRpzCnameIpaddressAttrTypes, &resp.Diagnostics)
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

func (r *RecordRpzCnameIpaddressResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RecordRpzCnameIpaddressModel

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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete RecordRpzCnameIpaddress: %s", err))
	}
}

func (r *RecordRpzCnameIpaddressResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
