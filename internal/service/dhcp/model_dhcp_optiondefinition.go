package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type DhcpOptiondefinitionModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var DhcpOptiondefinitionAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSDhcpOptiondefinitionAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIDhcpOptiondefinitionAttrTypes},
}

type NIOSDhcpOptiondefinitionModel struct {
	Code  types.Int64  `tfsdk:"code"`
	Name  types.String `tfsdk:"name"`
	Space types.String `tfsdk:"space"`
	Type  types.String `tfsdk:"type"`
}

var NIOSDhcpOptiondefinitionAttrTypes = map[string]attr.Type{
	"code":  types.Int64Type,
	"name":  types.StringType,
	"space": types.StringType,
	"type":  types.StringType,
}

type UDDIDhcpOptiondefinitionModel struct {
	Array       types.Bool   `tfsdk:"array"`
	Code        types.Int64  `tfsdk:"code"`
	Comment     types.String `tfsdk:"comment"`
	Name        types.String `tfsdk:"name"`
	OptionSpace types.String `tfsdk:"option_space"`
	Type        types.String `tfsdk:"type"`
}

var UDDIDhcpOptiondefinitionAttrTypes = map[string]attr.Type{
	"array":        types.BoolType,
	"code":         types.Int64Type,
	"comment":      types.StringType,
	"name":         types.StringType,
	"option_space": types.StringType,
	"type":         types.StringType,
}

const (
	DhcpOptiondefinitionReturnFields = "code,name,space,type"
)

var DhcpOptiondefinitionResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          DhcpOptiondefinitionResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          DhcpOptiondefinitionResourceUddiSchemaAttributes,
	},
}

var DhcpOptiondefinitionResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"code": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 254),
		},
		MarkdownDescription: "The code of a DHCP option definition object. An option code number is used to identify the DHCP option.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of a DHCP option definition object.",
	},
	"space": schema.StringAttribute{
		Default:  stringdefault.StaticString("DHCP"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The space of a DHCP option definition object.",
	},
	"type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("16-bit signed integer", "16-bit unsigned integer", "32-bit signed integer", "32-bit unsigned integer", "64-bit unsigned integer", "8-bit signed integer", "8-bit unsigned integer", "8-bit unsigned integer (1,2,4,8)", "array of 16-bit integer", "array of 16-bit unsigned integer", "array of 32-bit integer", "array of 32-bit unsigned integer", "array of 64-bit unsigned integer", "array of 8-bit integer", "array of 8-bit unsigned integer", "array of ip-address", "array of ip-address pair", "array of string", "binary", "boolean", "boolean array of ip-address", "boolean-text", "domain-list", "domain-name", "encapsulated", "ip-address", "string", "text"),
		},
		Required:            true,
		MarkdownDescription: "The data type of the Grid DHCP option.",
	},
}

var DhcpOptiondefinitionResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"array": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates whether the option value is an array of the type or not.",
	},
	"code": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The option code.",
	},
	"comment": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the option code. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 256),
		},
		MarkdownDescription: "The name of the option code. Must contain 1 to 256 characters. Can include UTF-8.",
	},
	"option_space": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The option type for the option code.  Valid values are: * _address4_ * _address6_ * _boolean_ * _empty_ * _fqdn_ * _int8_ * _int16_ * _int32_ * _text_ * _uint8_ * _uint16_ * _uint32_",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *DhcpOptiondefinitionModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.DhcpOptiondefinition {
	if m == nil {
		return nil
	}

	obj := &coremodel.DhcpOptiondefinition{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSDhcpOptiondefinitionModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIDhcpOptiondefinitionModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSDhcpOptiondefinitionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSDhcpOptiondefinitionExt {
	return &coremodel.NIOSDhcpOptiondefinitionExt{
		Code:  flex.ExpandInt64Pointer(m.Code),
		Name:  flex.ExpandStringPointerNullAsEmpty(m.Name),
		Space: flex.ExpandStringPointerNullAsEmpty(m.Space),
		Type:  flex.ExpandStringPointerNullAsEmpty(m.Type),
	}
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIDhcpOptiondefinitionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIDhcpOptiondefinitionExt {
	return &coremodel.UDDIDhcpOptiondefinitionExt{
		Array:       flex.ExpandBoolPointer(m.Array),
		Code:        flex.ExpandInt64(m.Code),
		Comment:     flex.ExpandStringPointer(m.Comment),
		Name:        flex.ExpandString(m.Name),
		OptionSpace: flex.ExpandString(m.OptionSpace),
		Type:        flex.ExpandString(m.Type),
	}
}

// Flatten populates the TF model from a core response.
func (m *DhcpOptiondefinitionModel) Flatten(ctx context.Context, resp *coremodel.DhcpOptiondefinition, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSDhcpOptiondefinitionModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSDhcpOptiondefinitionModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSDhcpOptiondefinitionAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSDhcpOptiondefinitionAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIDhcpOptiondefinitionModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIDhcpOptiondefinitionModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIDhcpOptiondefinitionAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIDhcpOptiondefinitionAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSDhcpOptiondefinitionModel) Flatten(ctx context.Context, from *coremodel.NIOSDhcpOptiondefinitionExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Code = flex.FlattenInt64Pointer(from.Code)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Space = flex.FlattenStringPointerEmptyAsNull(from.Space)
	m.Type = flex.FlattenStringPointerEmptyAsNull(from.Type)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIDhcpOptiondefinitionModel) Flatten(ctx context.Context, from *coremodel.UDDIDhcpOptiondefinitionExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Array = flex.FlattenBoolPointer(from.Array)
	m.Code = flex.FlattenInt64(from.Code)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Name = flex.FlattenString(from.Name)
	m.OptionSpace = flex.FlattenString(from.OptionSpace)
	m.Type = flex.FlattenString(from.Type)
}
