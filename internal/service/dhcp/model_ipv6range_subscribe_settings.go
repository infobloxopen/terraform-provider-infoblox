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

// Ipv6rangeSubscribeSettingsModel is the Terraform model for Ipv6rangeSubscribeSettings
type Ipv6rangeSubscribeSettingsModel struct {
	EnabledAttributes  types.List `tfsdk:"enabled_attributes"`
	MappedEaAttributes types.List `tfsdk:"mapped_ea_attributes"`
}

// Ipv6rangeSubscribeSettingsAttrTypes contains the attribute types for Ipv6rangeSubscribeSettingsModel
var Ipv6rangeSubscribeSettingsAttrTypes = map[string]attr.Type{
	"enabled_attributes":   types.ListType{ElemType: types.StringType},
	"mapped_ea_attributes": types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6rangesubscribesettingsMappedEaAttributesAttrTypes}},
}

// Ipv6rangeSubscribeSettingsResourceSchemaAttributes contains the schema attributes for Ipv6rangeSubscribeSettingsModel
var Ipv6rangeSubscribeSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled_attributes": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.ValueStringsAre(stringvalidator.OneOf("DOMAINNAME", "ENDPOINT_PROFILE", "SECURITY_GROUP", "SESSION_STATE", "SSID", "USERNAME", "VLAN")),
		},
		MarkdownDescription: "The list of Cisco ISE attributes allowed for subscription.",
	},
	"mapped_ea_attributes": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6rangesubscribesettingsMappedEaAttributesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of NIOS extensible attributes to Cisco ISE attributes mappings.",
	},
}

// ExpandIpv6rangeSubscribeSettings converts a Terraform Object to SDK type
func ExpandIpv6rangeSubscribeSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.Ipv6rangeSubscribeSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6rangeSubscribeSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6rangeSubscribeSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.Ipv6rangeSubscribeSettings {
	if m == nil {
		return nil
	}
	to := &niosdhcp.Ipv6rangeSubscribeSettings{
		EnabledAttributes:  flex.ExpandFrameworkListString(ctx, m.EnabledAttributes, diags),
		MappedEaAttributes: flex.ExpandFrameworkListNestedBlock(ctx, m.MappedEaAttributes, diags, ExpandIpv6rangesubscribesettingsMappedEaAttributes),
	}
	return to
}

// FlattenIpv6rangeSubscribeSettings converts an SDK type to Terraform Object
func FlattenIpv6rangeSubscribeSettings(ctx context.Context, from *niosdhcp.Ipv6rangeSubscribeSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangeSubscribeSettingsAttrTypes)
	}
	m := &Ipv6rangeSubscribeSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangeSubscribeSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6rangeSubscribeSettingsModel) Flatten(ctx context.Context, from *niosdhcp.Ipv6rangeSubscribeSettings, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnabledAttributes = flex.FlattenFrameworkListString(ctx, from.EnabledAttributes, diags)
	m.MappedEaAttributes = flex.FlattenFrameworkListNestedBlock(ctx, from.MappedEaAttributes, Ipv6rangesubscribesettingsMappedEaAttributesAttrTypes, diags, FlattenIpv6rangesubscribesettingsMappedEaAttributes)
}
