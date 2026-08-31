package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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

// NetworkviewDdnsZonePrimariesModel is the Terraform model for NetworkviewDdnsZonePrimaries
type NetworkviewDdnsZonePrimariesModel struct {
	ZoneMatch      types.String `tfsdk:"zone_match"`
	DnsGridZone    types.Object `tfsdk:"dns_grid_zone"`
	DnsGridPrimary types.String `tfsdk:"dns_grid_primary"`
	DnsExtZone     types.String `tfsdk:"dns_ext_zone"`
	DnsExtPrimary  types.String `tfsdk:"dns_ext_primary"`
}

// NetworkviewDdnsZonePrimariesAttrTypes contains the attribute types for NetworkviewDdnsZonePrimariesModel
var NetworkviewDdnsZonePrimariesAttrTypes = map[string]attr.Type{
	"zone_match":       types.StringType,
	"dns_grid_zone":    types.ObjectType{AttrTypes: NetworkviewDdnsZonePrimariesDnsGridZoneAttrTypes},
	"dns_grid_primary": types.StringType,
	"dns_ext_zone":     types.StringType,
	"dns_ext_primary":  types.StringType,
}

// NetworkviewDdnsZonePrimariesResourceSchemaAttributes contains the schema attributes for NetworkviewDdnsZonePrimariesModel
var NetworkviewDdnsZonePrimariesResourceSchemaAttributes = map[string]schema.Attribute{
	"zone_match": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("GRID", "EXTERNAL", "ANY_EXTERNAL", "ANY_GRID"),
		},
		Required:            true,
		MarkdownDescription: "Indicate matching type.",
	},
	"dns_grid_zone": schema.SingleNestedAttribute{
		Attributes:          NetworkviewDdnsZonePrimariesDnsGridZoneResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "",
	},
	"dns_grid_primary": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of a Grid member.",
	},
	"dns_ext_zone": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "The name of external zone in FQDN format.",
	},
	"dns_ext_primary": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IP address of the External server. Valid when zone_match is \"EXTERNAL\" or \"ANY_EXTERNAL\".",
	},
}

// ExpandNetworkviewDdnsZonePrimaries converts a Terraform Object to SDK type
func ExpandNetworkviewDdnsZonePrimaries(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkviewDdnsZonePrimaries {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkviewDdnsZonePrimariesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkviewDdnsZonePrimariesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkviewDdnsZonePrimaries {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkviewDdnsZonePrimaries{
		ZoneMatch:      flex.ExpandStringPointerNullAsEmpty(m.ZoneMatch),
		DnsGridZone:    ExpandNetworkviewDdnsZonePrimariesDnsGridZone(ctx, m.DnsGridZone, diags),
		DnsGridPrimary: flex.ExpandStringPointerNullAsEmpty(m.DnsGridPrimary),
		DnsExtZone:     flex.ExpandStringPointerNullAsEmpty(m.DnsExtZone),
		DnsExtPrimary:  flex.ExpandStringPointer(m.DnsExtPrimary),
	}
	return to
}

// FlattenNetworkviewDdnsZonePrimaries converts an SDK type to Terraform Object
func FlattenNetworkviewDdnsZonePrimaries(ctx context.Context, from *niosipam.NetworkviewDdnsZonePrimaries, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkviewDdnsZonePrimariesAttrTypes)
	}
	m := &NetworkviewDdnsZonePrimariesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkviewDdnsZonePrimariesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkviewDdnsZonePrimariesModel) Flatten(ctx context.Context, from *niosipam.NetworkviewDdnsZonePrimaries, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.ZoneMatch = flex.FlattenStringPointerEmptyAsNull(from.ZoneMatch)
	m.DnsGridZone = FlattenNetworkviewDdnsZonePrimariesDnsGridZone(ctx, from.DnsGridZone, diags)
	m.DnsGridPrimary = flex.FlattenStringPointerEmptyAsNull(from.DnsGridPrimary)
	m.DnsExtZone = flex.FlattenStringPointerEmptyAsNull(from.DnsExtZone)
	m.DnsExtPrimary = flex.FlattenStringPointerEmptyAsNull(from.DnsExtPrimary)
}
