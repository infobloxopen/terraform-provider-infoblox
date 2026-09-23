package redirect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/redirect"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

type CustomRedirectModel struct {
	Id            types.Int32  `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var CustomRedirectAttrTypes = map[string]attr.Type{
	"id":             types.Int32Type,
	"update_trigger": types.StringType,
	"uddi":           types.ObjectType{AttrTypes: UDDICustomRedirectAttrTypes},
}

type UDDICustomRedirectModel struct {
	Data types.String `tfsdk:"data"`
	Name types.String `tfsdk:"name"`
}

var UDDICustomRedirectAttrTypes = map[string]attr.Type{
	"data": types.StringType,
	"name": types.StringType,
}

const (
	CustomRedirectReturnFields = ""
)

var CustomRedirectResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.Int32Attribute{
		Computed:            true,
		MarkdownDescription: "The Custom Redirect object identifier.",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          CustomRedirectResourceUddiSchemaAttributes,
	},
}

var CustomRedirectResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"data": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The list of csv custom IPv4/IPv6 or a single domain redirect address.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the custom redirect.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *CustomRedirectModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.CustomRedirect {
	if m == nil {
		return nil
	}

	obj := &coremodel.CustomRedirect{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDICustomRedirectModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDICustomRedirectModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDICustomRedirectExt {
	return &coremodel.UDDICustomRedirectExt{
		Data: flex.ExpandStringPointer(m.Data),
		Name: flex.ExpandStringPointer(m.Name),
	}
}

// Flatten populates the TF model from a core response.
func (m *CustomRedirectModel) Flatten(ctx context.Context, resp *coremodel.CustomRedirect, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenInt32Pointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDICustomRedirectModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDICustomRedirectModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDICustomRedirectAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDICustomRedirectAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDICustomRedirectModel) Flatten(ctx context.Context, from *coremodel.UDDICustomRedirectExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Data = flex.FlattenStringPointer(from.Data)
	m.Name = flex.FlattenStringPointer(from.Name)
}
