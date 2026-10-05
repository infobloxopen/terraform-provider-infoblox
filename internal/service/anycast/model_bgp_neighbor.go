package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddianycast "github.com/infobloxopen/universal-ddi-go-client/anycast"
)

// BgpNeighborModel is the Terraform model for BgpNeighbor
type BgpNeighborModel struct {
	Asn         types.Int64  `tfsdk:"asn"`
	AsnText     types.String `tfsdk:"asn_text"`
	IpAddress   types.String `tfsdk:"ip_address"`
	MaxHopCount types.Int64  `tfsdk:"max_hop_count"`
	Multihop    types.Bool   `tfsdk:"multihop"`
	Password    types.String `tfsdk:"password"`
}

// BgpNeighborAttrTypes contains the attribute types for BgpNeighborModel
var BgpNeighborAttrTypes = map[string]attr.Type{
	"asn":           types.Int64Type,
	"asn_text":      types.StringType,
	"ip_address":    types.StringType,
	"max_hop_count": types.Int64Type,
	"multihop":      types.BoolType,
	"password":      types.StringType,
}

// BgpNeighborResourceSchemaAttributes contains the schema attributes for BgpNeighborModel
var BgpNeighborResourceSchemaAttributes = map[string]schema.Attribute{
	"asn": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "Autonomous system number of this BGP/anycast enabled on-prem host.",
	},
	"asn_text": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Autonomous system as text (supported in ASDOT or ASPLAIN format) Optional, requires the asn field to be set to the equivalent integer value of the ASDOT/ASPLAIN string contained in this field or be unset/zero.",
	},
	"ip_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IPv4 address of the BGP neighbor",
	},
	"max_hop_count": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Max hop count, if BGP multihop is enabled.",
	},
	"multihop": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "BGP multihop enabled or not.",
	},
	"password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "BGP protocol access password for this BGP neighbor, max 25 characters long.",
	},
}

// ExpandBgpNeighbor converts a Terraform Object to SDK type
func ExpandBgpNeighbor(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddianycast.BgpNeighbor {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m BgpNeighborModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *BgpNeighborModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddianycast.BgpNeighbor {
	if m == nil {
		return nil
	}
	to := &uddianycast.BgpNeighbor{
		Asn:         flex.ExpandInt64Pointer(m.Asn),
		AsnText:     flex.ExpandStringPointer(m.AsnText),
		IpAddress:   flex.ExpandStringPointer(m.IpAddress),
		MaxHopCount: flex.ExpandInt64Pointer(m.MaxHopCount),
		Multihop:    flex.ExpandBoolPointer(m.Multihop),
		Password:    flex.ExpandStringPointer(m.Password),
	}
	return to
}

// FlattenBgpNeighbor converts an SDK type to Terraform Object
func FlattenBgpNeighbor(ctx context.Context, from *uddianycast.BgpNeighbor, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(BgpNeighborAttrTypes)
	}
	m := &BgpNeighborModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, BgpNeighborAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *BgpNeighborModel) Flatten(ctx context.Context, from *uddianycast.BgpNeighbor, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Asn = flex.FlattenInt64Pointer(from.Asn)
	m.AsnText = flex.FlattenStringPointer(from.AsnText)
	m.IpAddress = flex.FlattenStringPointer(from.IpAddress)
	m.MaxHopCount = flex.FlattenInt64Pointer(from.MaxHopCount)
	m.Multihop = flex.FlattenBoolPointer(from.Multihop)
	m.Password = flex.FlattenStringPointer(from.Password)
}
