package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// Ipv6networkSubscribeSettingsModel is the Terraform model for Ipv6networkSubscribeSettings
type Ipv6networkSubscribeSettingsModel struct {
	EnabledAttributes  types.List `tfsdk:"enabled_attributes"`
	MappedEaAttributes types.List `tfsdk:"mapped_ea_attributes"`
}

// Ipv6networkSubscribeSettingsAttrTypes contains the attribute types for Ipv6networkSubscribeSettingsModel
var Ipv6networkSubscribeSettingsAttrTypes = map[string]attr.Type{
	"enabled_attributes":   types.ListType{ElemType: types.StringType},
	"mapped_ea_attributes": types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networksubscribesettingsMappedEaAttributesAttrTypes}},
}

// Ipv6networkSubscribeSettingsResourceSchemaAttributes contains the schema attributes for Ipv6networkSubscribeSettingsModel
var Ipv6networkSubscribeSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled_attributes": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			customvalidator.StringsInSlice([]string{"DOMAINNAME", "ENDPOINT_PROFILE", "SECURITY_GROUP", "SESSION_STATE", "SSID", "USERNAME", "VLAN"}),
		},
		MarkdownDescription: "The list of Cisco ISE attributes allowed for subscription.",
	},
	"mapped_ea_attributes": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networksubscribesettingsMappedEaAttributesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of NIOS extensible attributes to Cisco ISE attributes mappings.",
	},
}

// ExpandIpv6networkSubscribeSettings converts a Terraform Object to SDK type
func ExpandIpv6networkSubscribeSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.Ipv6networkSubscribeSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkSubscribeSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6networkSubscribeSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.Ipv6networkSubscribeSettings {
	if m == nil {
		return nil
	}
	to := &niosipam.Ipv6networkSubscribeSettings{
		EnabledAttributes:  flex.ExpandFrameworkListString(ctx, m.EnabledAttributes, diags),
		MappedEaAttributes: flex.ExpandFrameworkListNestedBlock(ctx, m.MappedEaAttributes, diags, ExpandIpv6networksubscribesettingsMappedEaAttributes),
	}
	return to
}

// FlattenIpv6networkSubscribeSettings converts an SDK type to Terraform Object
func FlattenIpv6networkSubscribeSettings(ctx context.Context, from *niosipam.Ipv6networkSubscribeSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkSubscribeSettingsAttrTypes)
	}
	m := &Ipv6networkSubscribeSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkSubscribeSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6networkSubscribeSettingsModel) Flatten(ctx context.Context, from *niosipam.Ipv6networkSubscribeSettings, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnabledAttributes = flex.FlattenFrameworkListString(ctx, from.EnabledAttributes, diags)
	m.MappedEaAttributes = flex.FlattenFrameworkListNestedBlock(ctx, from.MappedEaAttributes, Ipv6networksubscribesettingsMappedEaAttributesAttrTypes, diags, FlattenIpv6networksubscribesettingsMappedEaAttributes)
}
