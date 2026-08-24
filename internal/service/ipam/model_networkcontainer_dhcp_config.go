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

// NetworkcontainerDHCPConfigModel is the Terraform model for DHCPConfig
type NetworkcontainerDHCPConfigModel struct {
	AllowUnknown          types.Bool                       `tfsdk:"allow_unknown"`
	AuthoritativeDhcp     types.Bool                       `tfsdk:"authoritative_dhcp"`
	Filters               types.List                       `tfsdk:"filters"`
	FiltersLargeSelection types.List                       `tfsdk:"filters_large_selection"`
	IgnoreClientUid       types.Bool                       `tfsdk:"ignore_client_uid"`
	IgnoreList            internaltypes.UnorderedListValue `tfsdk:"ignore_list"`
	LeaseTime             types.Int64                      `tfsdk:"lease_time"`
}

// NetworkcontainerDHCPConfigAttrTypes contains the attribute types for NetworkcontainerDHCPConfigModel
var NetworkcontainerDHCPConfigAttrTypes = map[string]attr.Type{
	"allow_unknown":           types.BoolType,
	"authoritative_dhcp":      types.BoolType,
	"filters":                 types.ListType{ElemType: types.StringType},
	"filters_large_selection": types.ListType{ElemType: types.StringType},
	"ignore_client_uid":       types.BoolType,
	"ignore_list":             internaltypes.UnorderedList{ListType: types.ListType{ElemType: types.ObjectType{AttrTypes: IgnoreItemAttrTypes}}},
	"lease_time":              types.Int64Type,
}

// NetworkcontainerDHCPConfigResourceSchemaAttributes contains the schema attributes for NetworkcontainerDHCPConfigModel
var NetworkcontainerDHCPConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"allow_unknown": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Disable to allow leases only for known IPv4 clients, those for which a fixed address is configured.",
	},
	"authoritative_dhcp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Set DHCP server as authoritative.",
	},
	"filters": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The resource identifier.",
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
	"lease_time": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(3600),
		MarkdownDescription: "The lease duration in seconds.",
	},
}

// ExpandNetworkcontainerDHCPConfig converts a Terraform Object to SDK type
func ExpandNetworkcontainerDHCPConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.DHCPConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerDHCPConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkcontainerDHCPConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.DHCPConfig {
	if m == nil {
		return nil
	}
	to := &uddiipam.DHCPConfig{
		AllowUnknown:          flex.ExpandBoolPointer(m.AllowUnknown),
		AuthoritativeDhcp:     flex.ExpandBoolPointer(m.AuthoritativeDhcp),
		Filters:               flex.ExpandFrameworkListString(ctx, m.Filters, diags),
		FiltersLargeSelection: flex.ExpandFrameworkListString(ctx, m.FiltersLargeSelection, diags),
		IgnoreClientUid:       flex.ExpandBoolPointer(m.IgnoreClientUid),
		IgnoreList:            flex.ExpandFrameworkListNestedBlock(ctx, m.IgnoreList, diags, ExpandIgnoreItem),
		LeaseTime:             flex.ExpandInt64Pointer(m.LeaseTime),
	}
	return to
}

// FlattenNetworkcontainerDHCPConfig converts an SDK type to Terraform Object
func FlattenNetworkcontainerDHCPConfig(ctx context.Context, from *uddiipam.DHCPConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerDHCPConfigAttrTypes)
	}
	m := &NetworkcontainerDHCPConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerDHCPConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkcontainerDHCPConfigModel) Flatten(ctx context.Context, from *uddiipam.DHCPConfig, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AllowUnknown = flex.FlattenBoolPointer(from.AllowUnknown)
	m.AuthoritativeDhcp = flex.FlattenBoolPointer(from.AuthoritativeDhcp)
	m.Filters = flex.FlattenFrameworkListString(ctx, from.Filters, diags)
	m.FiltersLargeSelection = flex.FlattenFrameworkListString(ctx, from.FiltersLargeSelection, diags)
	m.IgnoreClientUid = flex.FlattenBoolPointer(from.IgnoreClientUid)
	m.IgnoreList = flex.FlattenFrameworkUnorderedListNestedBlock(ctx, from.IgnoreList, IgnoreItemAttrTypes, diags, FlattenIgnoreItem)
	m.LeaseTime = flex.FlattenInt64Pointer(from.LeaseTime)
}
