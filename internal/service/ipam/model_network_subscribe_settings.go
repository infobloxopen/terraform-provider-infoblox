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

// NetworkSubscribeSettingsModel is the Terraform model for NetworkSubscribeSettings
type NetworkSubscribeSettingsModel struct {
	EnabledAttributes  types.List `tfsdk:"enabled_attributes"`
	MappedEaAttributes types.List `tfsdk:"mapped_ea_attributes"`
}

// NetworkSubscribeSettingsAttrTypes contains the attribute types for NetworkSubscribeSettingsModel
var NetworkSubscribeSettingsAttrTypes = map[string]attr.Type{
	"enabled_attributes":   types.ListType{ElemType: types.StringType},
	"mapped_ea_attributes": types.ListType{ElemType: types.ObjectType{AttrTypes: NetworksubscribesettingsMappedEaAttributesAttrTypes}},
}

// NetworkSubscribeSettingsResourceSchemaAttributes contains the schema attributes for NetworkSubscribeSettingsModel
var NetworkSubscribeSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled_attributes": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			customvalidator.StringsInSlice([]string{"DOMAINNAME", "ENDPOINT_PROFILE", "SECURITY_GROUP", "SESSION_STATE", "SSID", "USERNAME", "VLAN"}),
		},
		MarkdownDescription: "The list of Cisco ISE attributes allowed for subscription.",
	},
	"mapped_ea_attributes": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworksubscribesettingsMappedEaAttributesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of NIOS extensible attributes to Cisco ISE attributes mappings.",
	},
}

// ExpandNetworkSubscribeSettings converts a Terraform Object to SDK type
func ExpandNetworkSubscribeSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkSubscribeSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkSubscribeSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkSubscribeSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkSubscribeSettings {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkSubscribeSettings{
		EnabledAttributes:  flex.ExpandFrameworkListString(ctx, m.EnabledAttributes, diags),
		MappedEaAttributes: flex.ExpandFrameworkListNestedBlock(ctx, m.MappedEaAttributes, diags, ExpandNetworksubscribesettingsMappedEaAttributes),
	}
	return to
}

// FlattenNetworkSubscribeSettings converts an SDK type to Terraform Object
func FlattenNetworkSubscribeSettings(ctx context.Context, from *niosipam.NetworkSubscribeSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkSubscribeSettingsAttrTypes)
	}
	m := &NetworkSubscribeSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkSubscribeSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkSubscribeSettingsModel) Flatten(ctx context.Context, from *niosipam.NetworkSubscribeSettings, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnabledAttributes = flex.FlattenFrameworkListString(ctx, from.EnabledAttributes, diags)
	m.MappedEaAttributes = flex.FlattenFrameworkListNestedBlock(ctx, from.MappedEaAttributes, NetworksubscribesettingsMappedEaAttributesAttrTypes, diags, FlattenNetworksubscribesettingsMappedEaAttributes)
}
