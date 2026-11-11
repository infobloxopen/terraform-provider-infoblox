package dhcp

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
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	coresvc "github.com/infobloxopen/terraform-provider-infoblox/internal/core/service/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/retry"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

var (
	_ resource.Resource                   = &FixedaddresstemplateResource{}
	_ resource.ResourceWithValidateConfig = &FixedaddresstemplateResource{}
	_ resource.ResourceWithConfigure      = &FixedaddresstemplateResource{}
	_ resource.ResourceWithImportState    = &FixedaddresstemplateResource{}
	_ resource.ResourceWithIdentity       = &FixedaddresstemplateResource{}
)

func NewFixedaddresstemplateResource() resource.Resource {
	return &FixedaddresstemplateResource{}
}

type FixedaddresstemplateResource struct {
	backend core.BackendType
	service coresvc.FixedaddresstemplateService
}

func (r *FixedaddresstemplateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fixed_address_template"
	resp.ResourceBehavior = resource.ResourceBehavior{
		MutableIdentity: true,
	}
}

func (r *FixedaddresstemplateResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				RequiredForImport: true,
			},
		},
	}
}

func (r *FixedaddresstemplateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Infoblox Fixedaddresstemplate in the NIOS backend.",
		Attributes:          FixedaddresstemplateResourceSchemaAttributes,
	}
}

func (r *FixedaddresstemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.service = coresvc.NewFixedaddresstemplateService(r.backend, client.NIOS, client.UDDI)
}

func (r *FixedaddresstemplateResource) retryPolicy(op retry.Operation) retry.Policy {
	return retry.For[coremodel.Fixedaddresstemplate](r.backend, op)
}

func (r *FixedaddresstemplateResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data FixedaddresstemplateModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Common backend block validations
	validator.ValidateBackendBlocks(r.backend, data.NIOS, types.ObjectNull(map[string]attr.Type{}), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	ValidateFixedaddresstemplate(ctx, data, resp)
}

func (r *FixedaddresstemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FixedaddresstemplateModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	obj := data.Expand(ctx, &resp.Diagnostics, true)
	if resp.Diagnostics.HasError() {
		return
	}

	if r.backend == core.BackendNIOS {
		ApplyFixedaddresstemplateNIOSUseFlags(ctx, req.Config, obj, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	var (
		apiResp  *coremodel.Fixedaddresstemplate
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpCreate), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Create(ctx, obj, &core.Options{
			ReturnFields: FixedaddresstemplateReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Fixedaddresstemplate: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *FixedaddresstemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FixedaddresstemplateModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var (
		apiResp  *coremodel.Fixedaddresstemplate
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpRead), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Read(ctx, data.Id.ValueString(), &core.Options{
			ReturnFields: FixedaddresstemplateReturnFields,
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Fixedaddresstemplate: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *FixedaddresstemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data FixedaddresstemplateModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags := req.State.GetAttribute(ctx, path.Root("id"), &data.Id)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	var planExtAttrs types.Map

	// Merge ext_attrs with state ext_attrs_all (inherited)
	if r.backend == core.BackendNIOS {
		var stateNIOSObj types.Object
		diags = req.State.GetAttribute(ctx, path.Root("nios"), &stateNIOSObj)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}

		planNIOS := flex.ExpandNestedObject[NIOSFixedaddresstemplateModel](ctx, data.NIOS, &resp.Diagnostics)
		stateNIOS := flex.ExpandNestedObject[NIOSFixedaddresstemplateModel](ctx, stateNIOSObj, &resp.Diagnostics)
		if planNIOS == nil {
			planNIOS = &NIOSFixedaddresstemplateModel{}
		}

		// Preserve the plan ext_attrs (without inherited EAs) for restore after Update
		planExtAttrs = planNIOS.ExtAttrs

		// Merge with state ext_attrs_all (inherited EAs)
		if stateNIOS != nil {
			planNIOS.ExtAttrs = flex.MergeEAs(planNIOS.ExtAttrs, stateNIOS.ExtAttrsAll)
		}
		data.NIOS = flex.FlattenNestedObject(ctx, planNIOS, NIOSFixedaddresstemplateAttrTypes, &resp.Diagnostics)
	}

	obj := data.Expand(ctx, &resp.Diagnostics, false)
	if resp.Diagnostics.HasError() {
		return
	}

	if r.backend == core.BackendNIOS {
		ApplyFixedaddresstemplateNIOSUseFlags(ctx, req.Config, obj, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	var (
		apiResp  *coremodel.Fixedaddresstemplate
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpUpdate), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Update(ctx, data.Id.ValueString(), obj, &core.Options{
			ReturnFields: FixedaddresstemplateReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Fixedaddresstemplate: %s", err))
		return
	}

	// Restore the plan ext_attrs (without inherited EAs) so Flatten preserves the user's input
	if r.backend == core.BackendNIOS {
		niosObj := flex.ExpandNestedObject[NIOSFixedaddresstemplateModel](ctx, data.NIOS, &resp.Diagnostics)
		if niosObj != nil {
			niosObj.ExtAttrs = planExtAttrs
			data.NIOS = flex.FlattenNestedObject(ctx, niosObj, NIOSFixedaddresstemplateAttrTypes, &resp.Diagnostics)
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
}

func (r *FixedaddresstemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FixedaddresstemplateModel

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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Fixedaddresstemplate: %s", err))
	}
}

func (r *FixedaddresstemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
