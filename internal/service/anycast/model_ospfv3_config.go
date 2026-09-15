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

// Ospfv3ConfigModel is the Terraform model for Ospfv3Config
type Ospfv3ConfigModel struct {
	Area               types.String `tfsdk:"area"`
	Cost               types.Int64  `tfsdk:"cost"`
	DeadInterval       types.Int64  `tfsdk:"dead_interval"`
	HelloInterval      types.Int64  `tfsdk:"hello_interval"`
	Interface          types.String `tfsdk:"interface"`
	RetransmitInterval types.Int64  `tfsdk:"retransmit_interval"`
	TransmitDelay      types.Int64  `tfsdk:"transmit_delay"`
}

// Ospfv3ConfigAttrTypes contains the attribute types for Ospfv3ConfigModel
var Ospfv3ConfigAttrTypes = map[string]attr.Type{
	"area":                types.StringType,
	"cost":                types.Int64Type,
	"dead_interval":       types.Int64Type,
	"hello_interval":      types.Int64Type,
	"interface":           types.StringType,
	"retransmit_interval": types.Int64Type,
	"transmit_delay":      types.Int64Type,
}

// Ospfv3ConfigResourceSchemaAttributes contains the schema attributes for Ospfv3ConfigModel
var Ospfv3ConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"area": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "OSPF area identifier; usually in the format of an IPv4 address (although not an address itself)",
	},
	"cost": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"dead_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"hello_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"interface": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Name of the interface that is configured with external IP address of the host",
	},
	"retransmit_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"transmit_delay": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "",
	},
}

// ExpandOspfv3Config converts a Terraform Object to SDK type
func ExpandOspfv3Config(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddianycast.Ospfv3Config {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ospfv3ConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ospfv3ConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddianycast.Ospfv3Config {
	if m == nil {
		return nil
	}
	to := &uddianycast.Ospfv3Config{
		Area:               flex.ExpandStringPointer(m.Area),
		Cost:               flex.ExpandInt64Pointer(m.Cost),
		DeadInterval:       flex.ExpandInt64Pointer(m.DeadInterval),
		HelloInterval:      flex.ExpandInt64Pointer(m.HelloInterval),
		Interface:          flex.ExpandStringPointer(m.Interface),
		RetransmitInterval: flex.ExpandInt64Pointer(m.RetransmitInterval),
		TransmitDelay:      flex.ExpandInt64Pointer(m.TransmitDelay),
	}
	return to
}

// FlattenOspfv3Config converts an SDK type to Terraform Object
func FlattenOspfv3Config(ctx context.Context, from *uddianycast.Ospfv3Config, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ospfv3ConfigAttrTypes)
	}
	m := &Ospfv3ConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ospfv3ConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ospfv3ConfigModel) Flatten(ctx context.Context, from *uddianycast.Ospfv3Config, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Area = flex.FlattenStringPointer(from.Area)
	m.Cost = flex.FlattenInt64Pointer(from.Cost)
	m.DeadInterval = flex.FlattenInt64Pointer(from.DeadInterval)
	m.HelloInterval = flex.FlattenInt64Pointer(from.HelloInterval)
	m.Interface = flex.FlattenStringPointer(from.Interface)
	m.RetransmitInterval = flex.FlattenInt64Pointer(from.RetransmitInterval)
	m.TransmitDelay = flex.FlattenInt64Pointer(from.TransmitDelay)
}
