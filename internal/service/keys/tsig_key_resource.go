package keys

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
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/keys"
	coresvc "github.com/infobloxopen/terraform-provider-infoblox/internal/core/service/keys"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/retry"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

var (
	_ resource.Resource                   = &TsigKeyResource{}
	_ resource.ResourceWithValidateConfig = &TsigKeyResource{}
	_ resource.ResourceWithConfigure      = &TsigKeyResource{}
	_ resource.ResourceWithImportState    = &TsigKeyResource{}
	_ resource.ResourceWithIdentity       = &TsigKeyResource{}
)

func NewTsigKeyResource() resource.Resource {
	return &TsigKeyResource{}
}

type TsigKeyResource struct {
	backend core.BackendType
	service coresvc.TsigKeyService
}

func (r *TsigKeyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tsig_key"
	resp.ResourceBehavior = resource.ResourceBehavior{
		MutableIdentity: true,
	}
}

func (r *TsigKeyResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.StringAttribute{
				RequiredForImport: true,
			},
		},
	}
}

func (r *TsigKeyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Infoblox TsigKey in the UDDI backend.",
		Attributes:          TsigKeyResourceSchemaAttributes,
	}
}

func (r *TsigKeyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.service = coresvc.NewTsigKeyService(r.backend, client.NIOS, client.UDDI)
}

func (r *TsigKeyResource) retryPolicy(op retry.Operation) retry.Policy {
	return retry.For[coremodel.TsigKey](r.backend, op)
}

func (r *TsigKeyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data TsigKeyModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Common backend block validations
	validator.ValidateBackendBlocks(r.backend, types.ObjectNull(map[string]attr.Type{}), data.UDDI, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	ValidateTsigKey(ctx, data, resp)
}

func (r *TsigKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TsigKeyModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	obj := data.Expand(ctx, &resp.Diagnostics, true)
	if resp.Diagnostics.HasError() {
		return
	}

	var (
		apiResp  *coremodel.TsigKey
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpCreate), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Create(ctx, obj, &core.Options{
			ReturnFields: TsigKeyReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create TsigKey: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *TsigKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TsigKeyModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var (
		apiResp  *coremodel.TsigKey
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpRead), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Read(ctx, data.Id.ValueString(), &core.Options{
			ReturnFields: TsigKeyReturnFields,
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read TsigKey: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *TsigKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TsigKeyModel

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
		apiResp  *coremodel.TsigKey
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpUpdate), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Update(ctx, data.Id.ValueString(), obj, &core.Options{
			ReturnFields: TsigKeyReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update TsigKey: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *TsigKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TsigKeyModel

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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete TsigKey: %s", err))
	}
}

func (r *TsigKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
