package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// RangeSubscribeSettingsModel is the Terraform model for RangeSubscribeSettings
type RangeSubscribeSettingsModel struct {
	EnabledAttributes  types.List `tfsdk:"enabled_attributes"`
	MappedEaAttributes types.List `tfsdk:"mapped_ea_attributes"`
}

// RangeSubscribeSettingsAttrTypes contains the attribute types for RangeSubscribeSettingsModel
var RangeSubscribeSettingsAttrTypes = map[string]attr.Type{
	"enabled_attributes":   types.ListType{ElemType: types.StringType},
	"mapped_ea_attributes": types.ListType{ElemType: types.ObjectType{AttrTypes: RangesubscribesettingsMappedEaAttributesAttrTypes}},
}

// RangeSubscribeSettingsResourceSchemaAttributes contains the schema attributes for RangeSubscribeSettingsModel
var RangeSubscribeSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled_attributes": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.ValueStringsAre(stringvalidator.OneOf("DOMAINNAME", "ENDPOINT_PROFILE", "SECURITY_GROUP", "SESSION_STATE", "SSID", "USERNAME", "VLAN")),
		},
		MarkdownDescription: "The list of Cisco ISE attributes allowed for subscription.",
	},
	"mapped_ea_attributes": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangesubscribesettingsMappedEaAttributesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of NIOS extensible attributes to Cisco ISE attributes mappings.",
	},
}

// ExpandRangeSubscribeSettings converts a Terraform Object to SDK type
func ExpandRangeSubscribeSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.RangeSubscribeSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangeSubscribeSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RangeSubscribeSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.RangeSubscribeSettings {
	if m == nil {
		return nil
	}
	to := &niosdhcp.RangeSubscribeSettings{
		EnabledAttributes:  flex.ExpandFrameworkListString(ctx, m.EnabledAttributes, diags),
		MappedEaAttributes: flex.ExpandFrameworkListNestedBlock(ctx, m.MappedEaAttributes, diags, ExpandRangesubscribesettingsMappedEaAttributes),
	}
	return to
}

// FlattenRangeSubscribeSettings converts an SDK type to Terraform Object
func FlattenRangeSubscribeSettings(ctx context.Context, from *niosdhcp.RangeSubscribeSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangeSubscribeSettingsAttrTypes)
	}
	m := &RangeSubscribeSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangeSubscribeSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RangeSubscribeSettingsModel) Flatten(ctx context.Context, from *niosdhcp.RangeSubscribeSettings, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnabledAttributes = flex.FlattenFrameworkListString(ctx, from.EnabledAttributes, diags)
	m.MappedEaAttributes = flex.FlattenFrameworkListNestedBlock(ctx, from.MappedEaAttributes, RangesubscribesettingsMappedEaAttributesAttrTypes, diags, FlattenRangesubscribesettingsMappedEaAttributes)
}
