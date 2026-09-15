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
		MarkdownDescription: "",
	},
	"asn_text": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Examples:     ASDOT        ASPLAIN     INTEGER     VALID/INVALID     0.1          1           1           Valid     1            1           1           Valid     65535        65535       65535       Valid     0.65535      65535       65535       Valid     1.0          65536       65536       Valid     1.1          65537       65537       Valid     1.65535      131071      131071      Valid     65535.0      4294901760  4294901760  Valid     65535.1      4294901761  4294901761  Valid     65535.65535  4294967295  4294967295  Valid      0.65536                              Invalid     65535.655536                         Invalid     65536.0                              Invalid     65536.65535                          Invalid                  4294967296              Invalid",
	},
	"ip_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IPv4 address of the BGP neighbor",
	},
	"max_hop_count": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"multihop": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"password": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "",
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
