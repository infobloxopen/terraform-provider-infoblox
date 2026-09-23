package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	int32planmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/int32planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/fw"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type NetworkListModel struct {
	Id            types.Int32  `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var NetworkListAttrTypes = map[string]attr.Type{
	"id":             types.Int32Type,
	"update_trigger": types.StringType,
	"uddi":           types.ObjectType{AttrTypes: UDDINetworkListAttrTypes},
}

type UDDINetworkListModel struct {
	AddrBlock   types.List   `tfsdk:"addr_block"`
	Description types.String `tfsdk:"description"`
	Name        types.String `tfsdk:"name"`
}

var UDDINetworkListAttrTypes = map[string]attr.Type{
	"addr_block":  types.ListType{ElemType: types.ObjectType{AttrTypes: AddrBlockAttrTypes}},
	"description": types.StringType,
	"name":        types.StringType,
}

const (
	NetworkListReturnFields = ""
)

var NetworkListResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.Int32Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int32{
			int32planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The Network List object identifier.",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          NetworkListResourceUddiSchemaAttributes,
	},
}

var NetworkListResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"addr_block": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AddrBlockResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of address blocks (CIDRs) in the network list, each with an optional end-user description. The plain `items` field is deprecated in favor of this field, since it allows adding a description to each address block; this provider does not expose `items` for that reason.",
	},
	"description": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The brief description for the network list.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the network list.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *NetworkListModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NetworkList {
	if m == nil {
		return nil
	}

	obj := &coremodel.NetworkList{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDINetworkListModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDINetworkListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDINetworkListExt {
	return &coremodel.UDDINetworkListExt{
		AddrBlock:   flex.ExpandFrameworkListNestedBlock(ctx, m.AddrBlock, diags, ExpandAddrBlock),
		Description: flex.ExpandStringPointer(m.Description),
		Name:        flex.ExpandStringPointer(m.Name),
	}
}

// Flatten populates the TF model from a core response.
func (m *NetworkListModel) Flatten(ctx context.Context, resp *coremodel.NetworkList, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenInt32Pointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDINetworkListModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDINetworkListModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDINetworkListAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDINetworkListAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDINetworkListModel) Flatten(ctx context.Context, from *coremodel.UDDINetworkListExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AddrBlock = flex.FlattenFrameworkListNestedBlock(ctx, from.AddrBlock, AddrBlockAttrTypes, diags, FlattenAddrBlock)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Name = flex.FlattenStringPointer(from.Name)
}
