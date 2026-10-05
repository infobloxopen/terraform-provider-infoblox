package fw

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/fw"
	coresvc "github.com/infobloxopen/terraform-provider-infoblox/internal/core/service/fw"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/retry"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

var (
	_ resource.Resource                   = &ApplicationFilterResource{}
	_ resource.ResourceWithValidateConfig = &ApplicationFilterResource{}
	_ resource.ResourceWithConfigure      = &ApplicationFilterResource{}
	_ resource.ResourceWithImportState    = &ApplicationFilterResource{}
	_ resource.ResourceWithIdentity       = &ApplicationFilterResource{}
)

func NewApplicationFilterResource() resource.Resource {
	return &ApplicationFilterResource{}
}

type ApplicationFilterResource struct {
	backend core.BackendType
	service coresvc.ApplicationFilterService
}

func (r *ApplicationFilterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_filter"
	resp.ResourceBehavior = resource.ResourceBehavior{
		MutableIdentity: true,
	}
}

func (r *ApplicationFilterResource) IdentitySchema(_ context.Context, _ resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"id": identityschema.Int32Attribute{
				RequiredForImport: true,
			},
		},
	}
}

func (r *ApplicationFilterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Infoblox ApplicationFilter in the UDDI backend.",
		Attributes:          ApplicationFilterResourceSchemaAttributes,
	}
}

func (r *ApplicationFilterResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.service = coresvc.NewApplicationFilterService(r.backend, client.NIOS, client.UDDI)
}

func (r *ApplicationFilterResource) retryPolicy(op retry.Operation) retry.Policy {
	return retry.For[coremodel.ApplicationFilter](r.backend, op)
}

func (r *ApplicationFilterResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data ApplicationFilterModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Common backend block validations
	validator.ValidateBackendBlocks(r.backend, types.ObjectNull(map[string]attr.Type{}), data.UDDI, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	ValidateApplicationFilter(ctx, data, resp)
}

func (r *ApplicationFilterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApplicationFilterModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	obj := data.Expand(ctx, &resp.Diagnostics, true)
	if resp.Diagnostics.HasError() {
		return
	}

	var (
		apiResp  *coremodel.ApplicationFilter
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpCreate), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Create(ctx, obj, &core.Options{
			ReturnFields: ApplicationFilterReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ApplicationFilter: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *ApplicationFilterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApplicationFilterModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var (
		apiResp  *coremodel.ApplicationFilter
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpRead), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Read(ctx, data.Id.ValueInt32(), &core.Options{
			ReturnFields: ApplicationFilterReturnFields,
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ApplicationFilter: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *ApplicationFilterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ApplicationFilterModel

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
		apiResp  *coremodel.ApplicationFilter
		httpResp *http.Response
	)

	err := retry.Do(ctx, r.retryPolicy(retry.OpUpdate), func(ctx context.Context) (int, error) {
		var apiErr error
		apiResp, httpResp, apiErr = r.service.Update(ctx, data.Id.ValueInt32(), obj, &core.Options{
			ReturnFields: ApplicationFilterReturnFields,
		})
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ApplicationFilter: %s", err))
		return
	}

	data.Flatten(ctx, apiResp, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), &data.Id)...)
}

func (r *ApplicationFilterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ApplicationFilterModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var httpResp *http.Response

	err := retry.Do(ctx, r.retryPolicy(retry.OpDelete), func(ctx context.Context) (int, error) {
		var apiErr error
		httpResp, apiErr = r.service.Delete(ctx, data.Id.ValueInt32())
		if httpResp != nil {
			return httpResp.StatusCode, apiErr
		}
		return 0, apiErr
	})
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ApplicationFilter: %s", err))
	}
}

func (r *ApplicationFilterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.Identity != nil && req.Identity.Raw.IsKnown() && !req.Identity.Raw.IsNull() {
		var identityID types.Int32
		diags := req.Identity.GetAttribute(ctx, path.Root("id"), &identityID)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
		req.ID = strconv.FormatInt(int64(identityID.ValueInt32()), 10)
	}

	parsedID, err := strconv.ParseInt(req.ID, 10, 32)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("Expected a numeric ApplicationFilter ID, got %q: %s", req.ID, err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), int32(parsedID))...)
	resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), int32(parsedID))...)
}
