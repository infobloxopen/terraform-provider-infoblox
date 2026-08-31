package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// Ipv6NetworkcontainerDHCPConfigModel is the Terraform model for DHCPConfig
type Ipv6NetworkcontainerDHCPConfigModel struct {
	AllowUnknownV6        types.Bool                       `tfsdk:"allow_unknown_v6"`
	AuthoritativeDhcp     types.Bool                       `tfsdk:"authoritative_dhcp"`
	FiltersLargeSelection types.List                       `tfsdk:"filters_large_selection"`
	FiltersV6             types.List                       `tfsdk:"filters_v6"`
	IgnoreClientUid       types.Bool                       `tfsdk:"ignore_client_uid"`
	IgnoreList            internaltypes.UnorderedListValue `tfsdk:"ignore_list"`
	LeaseTimeV6           types.Int64                      `tfsdk:"lease_time_v6"`
}

// Ipv6NetworkcontainerDHCPConfigAttrTypes contains the attribute types for Ipv6NetworkcontainerDHCPConfigModel
var Ipv6NetworkcontainerDHCPConfigAttrTypes = map[string]attr.Type{
	"allow_unknown_v6":        types.BoolType,
	"authoritative_dhcp":      types.BoolType,
	"filters_large_selection": types.ListType{ElemType: types.StringType},
	"filters_v6":              types.ListType{ElemType: types.StringType},
	"ignore_client_uid":       types.BoolType,
	"ignore_list":             internaltypes.UnorderedList{ListType: types.ListType{ElemType: types.ObjectType{AttrTypes: IgnoreItemAttrTypes}}},
	"lease_time_v6":           types.Int64Type,
}

// Ipv6NetworkcontainerDHCPConfigResourceSchemaAttributes contains the schema attributes for Ipv6NetworkcontainerDHCPConfigModel
var Ipv6NetworkcontainerDHCPConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"allow_unknown_v6": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Disable to allow leases only for known IPV6 clients, those for which a fixed address is configured.",
	},
	"authoritative_dhcp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Set DHCP server as authoritative.",
	},
	"filters_large_selection": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"filters_v6": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"ignore_client_uid": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Enable to ignore the client UID when issuing a DHCP lease. Use this option to prevent assigning two IP addresses for a client which does not have a UID during one phase of PXE boot but acquires one for the other phase.",
	},
	"ignore_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IgnoreItemResourceSchemaAttributes,
		},
		Optional:   true,
		CustomType: internaltypes.UnorderedList{ListType: types.ListType{ElemType: types.ObjectType{AttrTypes: IgnoreItemAttrTypes}}},
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of clients to ignore requests from.",
	},
	"lease_time_v6": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(3600),
		MarkdownDescription: "The lease duration in seconds for IPV6 clients.",
	},
}

// ExpandIpv6NetworkcontainerDHCPConfig converts a Terraform Object to SDK type
func ExpandIpv6NetworkcontainerDHCPConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.DHCPConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6NetworkcontainerDHCPConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6NetworkcontainerDHCPConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.DHCPConfig {
	if m == nil {
		return nil
	}
	to := &uddiipam.DHCPConfig{
		AllowUnknownV6:        flex.ExpandBoolPointer(m.AllowUnknownV6),
		AuthoritativeDhcp:     flex.ExpandBoolPointer(m.AuthoritativeDhcp),
		FiltersLargeSelection: flex.ExpandFrameworkListString(ctx, m.FiltersLargeSelection, diags),
		FiltersV6:             flex.ExpandFrameworkListString(ctx, m.FiltersV6, diags),
		IgnoreClientUid:       flex.ExpandBoolPointer(m.IgnoreClientUid),
		IgnoreList:            flex.ExpandFrameworkListNestedBlock(ctx, m.IgnoreList, diags, ExpandIgnoreItem),
		LeaseTimeV6:           flex.ExpandInt64Pointer(m.LeaseTimeV6),
	}
	return to
}

// FlattenIpv6NetworkcontainerDHCPConfig converts an SDK type to Terraform Object
func FlattenIpv6NetworkcontainerDHCPConfig(ctx context.Context, from *uddiipam.DHCPConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6NetworkcontainerDHCPConfigAttrTypes)
	}
	m := &Ipv6NetworkcontainerDHCPConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6NetworkcontainerDHCPConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6NetworkcontainerDHCPConfigModel) Flatten(ctx context.Context, from *uddiipam.DHCPConfig, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AllowUnknownV6 = flex.FlattenBoolPointer(from.AllowUnknownV6)
	m.AuthoritativeDhcp = flex.FlattenBoolPointer(from.AuthoritativeDhcp)
	m.FiltersLargeSelection = flex.FlattenFrameworkListString(ctx, from.FiltersLargeSelection, diags)
	m.FiltersV6 = flex.FlattenFrameworkListString(ctx, from.FiltersV6, diags)
	m.IgnoreClientUid = flex.FlattenBoolPointer(from.IgnoreClientUid)
	m.IgnoreList = flex.FlattenFrameworkUnorderedListNestedBlock(ctx, from.IgnoreList, IgnoreItemAttrTypes, diags, FlattenIgnoreItem)
	m.LeaseTimeV6 = flex.FlattenInt64Pointer(from.LeaseTimeV6)
}
