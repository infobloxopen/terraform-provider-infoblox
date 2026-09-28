package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddianycast "github.com/infobloxopen/universal-ddi-go-client/anycast"
)

// BgpConfigModel is the Terraform model for BgpConfig
type BgpConfigModel struct {
	Asn           types.Int64  `tfsdk:"asn"`
	AsnText       types.String `tfsdk:"asn_text"`
	Fields        types.Object `tfsdk:"fields"`
	HolddownSecs  types.Int64  `tfsdk:"holddown_secs"`
	KeepAliveSecs types.Int64  `tfsdk:"keep_alive_secs"`
	LinkDetect    types.Bool   `tfsdk:"link_detect"`
	Neighbors     types.List   `tfsdk:"neighbors"`
	Preamble      types.String `tfsdk:"preamble"`
}

// BgpConfigAttrTypes contains the attribute types for BgpConfigModel
var BgpConfigAttrTypes = map[string]attr.Type{
	"asn":             types.Int64Type,
	"asn_text":        types.StringType,
	"fields":          types.ObjectType{AttrTypes: ProtobufFieldMaskAttrTypes},
	"holddown_secs":   types.Int64Type,
	"keep_alive_secs": types.Int64Type,
	"link_detect":     types.BoolType,
	"neighbors":       types.ListType{ElemType: types.ObjectType{AttrTypes: BgpNeighborAttrTypes}},
	"preamble":        types.StringType,
}

// BgpConfigResourceSchemaAttributes contains the schema attributes for BgpConfigModel
var BgpConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"asn": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "Autonomous system number of this BGP/anycast enabled on-prem host.",
	},
	"asn_text": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Autonomous system as text (supported in ASDOT or ASPLAIN format) Optional, requires the asn field to be set to the equivalent integer value of the ASDOT/ASPLAIN string contained in this field or be unset/zero.",
	},
	"fields": schema.SingleNestedAttribute{
		Attributes:          ProtobufFieldMaskResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Represents a set of symbolic field paths.",
	},
	"holddown_secs": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "BGP route hold-down timer.",
	},
	"keep_alive_secs": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "BGP keep-alive timer.",
	},
	"link_detect": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Enable/disable link detection.",
	},
	"neighbors": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: BgpNeighborResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "List of BgpNeighbor structs.",
	},
	"preamble": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Any predefined BGP configuration, with embedded new lines; the preamble will be prepended to the generated BGP configuration.",
	},
}

// ExpandBgpConfig converts a Terraform Object to SDK type
func ExpandBgpConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddianycast.BgpConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m BgpConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *BgpConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddianycast.BgpConfig {
	if m == nil {
		return nil
	}
	to := &uddianycast.BgpConfig{
		Asn:           flex.ExpandInt64Pointer(m.Asn),
		AsnText:       flex.ExpandStringPointer(m.AsnText),
		Fields:        ExpandProtobufFieldMask(ctx, m.Fields, diags),
		HolddownSecs:  flex.ExpandInt64Pointer(m.HolddownSecs),
		KeepAliveSecs: flex.ExpandInt64Pointer(m.KeepAliveSecs),
		LinkDetect:    flex.ExpandBoolPointer(m.LinkDetect),
		Neighbors:     flex.ExpandFrameworkListNestedBlock(ctx, m.Neighbors, diags, ExpandBgpNeighbor),
		Preamble:      flex.ExpandStringPointer(m.Preamble),
	}
	return to
}

// FlattenBgpConfig converts an SDK type to Terraform Object
func FlattenBgpConfig(ctx context.Context, from *uddianycast.BgpConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(BgpConfigAttrTypes)
	}
	m := &BgpConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, BgpConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *BgpConfigModel) Flatten(ctx context.Context, from *uddianycast.BgpConfig, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Asn = flex.FlattenInt64Pointer(from.Asn)
	m.AsnText = flex.FlattenStringPointer(from.AsnText)
	m.Fields = FlattenProtobufFieldMask(ctx, from.Fields, diags)
	m.HolddownSecs = flex.FlattenInt64Pointer(from.HolddownSecs)
	m.KeepAliveSecs = flex.FlattenInt64Pointer(from.KeepAliveSecs)
	m.LinkDetect = flex.FlattenBoolPointer(from.LinkDetect)
	m.Neighbors = flex.FlattenFrameworkListNestedBlock(ctx, from.Neighbors, BgpNeighborAttrTypes, diags, FlattenBgpNeighbor)
	m.Preamble = flex.FlattenStringPointer(from.Preamble)
}
