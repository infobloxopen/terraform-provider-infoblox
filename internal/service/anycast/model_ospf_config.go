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

// OspfConfigModel is the Terraform model for OspfConfig
type OspfConfigModel struct {
	Area                types.String `tfsdk:"area"`
	AreaType            types.String `tfsdk:"area_type"`
	AuthenticationKey   types.String `tfsdk:"authentication_key"`
	AuthenticationKeyId types.Int64  `tfsdk:"authentication_key_id"`
	AuthenticationType  types.String `tfsdk:"authentication_type"`
	Cost                types.Int64  `tfsdk:"cost"`
	DeadInterval        types.Int64  `tfsdk:"dead_interval"`
	HelloInterval       types.Int64  `tfsdk:"hello_interval"`
	Interface           types.String `tfsdk:"interface"`
	Preamble            types.String `tfsdk:"preamble"`
	RetransmitInterval  types.Int64  `tfsdk:"retransmit_interval"`
	TransmitDelay       types.Int64  `tfsdk:"transmit_delay"`
}

// OspfConfigAttrTypes contains the attribute types for OspfConfigModel
var OspfConfigAttrTypes = map[string]attr.Type{
	"area":                  types.StringType,
	"area_type":             types.StringType,
	"authentication_key":    types.StringType,
	"authentication_key_id": types.Int64Type,
	"authentication_type":   types.StringType,
	"cost":                  types.Int64Type,
	"dead_interval":         types.Int64Type,
	"hello_interval":        types.Int64Type,
	"interface":             types.StringType,
	"preamble":              types.StringType,
	"retransmit_interval":   types.Int64Type,
	"transmit_delay":        types.Int64Type,
}

// OspfConfigResourceSchemaAttributes contains the schema attributes for OspfConfigModel
var OspfConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"area": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "OSPF area identifier; usually in the format of an IPv4 address (although not an address itself)",
	},
	"area_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"authentication_key": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"authentication_key_id": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "",
	},
	"authentication_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "",
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
	"preamble": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Any predefined OSPF configuration, with embedded new lines; the preamble will be prepended to the generated BGP configuration.",
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

// ExpandOspfConfig converts a Terraform Object to SDK type
func ExpandOspfConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddianycast.OspfConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m OspfConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *OspfConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddianycast.OspfConfig {
	if m == nil {
		return nil
	}
	to := &uddianycast.OspfConfig{
		Area:                flex.ExpandStringPointer(m.Area),
		AreaType:            flex.ExpandStringPointer(m.AreaType),
		AuthenticationKey:   flex.ExpandStringPointer(m.AuthenticationKey),
		AuthenticationKeyId: flex.ExpandInt64Pointer(m.AuthenticationKeyId),
		AuthenticationType:  flex.ExpandStringPointer(m.AuthenticationType),
		Cost:                flex.ExpandInt64Pointer(m.Cost),
		DeadInterval:        flex.ExpandInt64Pointer(m.DeadInterval),
		HelloInterval:       flex.ExpandInt64Pointer(m.HelloInterval),
		Interface:           flex.ExpandStringPointer(m.Interface),
		Preamble:            flex.ExpandStringPointer(m.Preamble),
		RetransmitInterval:  flex.ExpandInt64Pointer(m.RetransmitInterval),
		TransmitDelay:       flex.ExpandInt64Pointer(m.TransmitDelay),
	}
	return to
}

// FlattenOspfConfig converts an SDK type to Terraform Object
func FlattenOspfConfig(ctx context.Context, from *uddianycast.OspfConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(OspfConfigAttrTypes)
	}
	m := &OspfConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, OspfConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *OspfConfigModel) Flatten(ctx context.Context, from *uddianycast.OspfConfig, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Area = flex.FlattenStringPointer(from.Area)
	m.AreaType = flex.FlattenStringPointer(from.AreaType)
	m.AuthenticationKey = flex.FlattenStringPointer(from.AuthenticationKey)
	m.AuthenticationKeyId = flex.FlattenInt64Pointer(from.AuthenticationKeyId)
	m.AuthenticationType = flex.FlattenStringPointer(from.AuthenticationType)
	m.Cost = flex.FlattenInt64Pointer(from.Cost)
	m.DeadInterval = flex.FlattenInt64Pointer(from.DeadInterval)
	m.HelloInterval = flex.FlattenInt64Pointer(from.HelloInterval)
	m.Interface = flex.FlattenStringPointer(from.Interface)
	m.Preamble = flex.FlattenStringPointer(from.Preamble)
	m.RetransmitInterval = flex.FlattenInt64Pointer(from.RetransmitInterval)
	m.TransmitDelay = flex.FlattenInt64Pointer(from.TransmitDelay)
}
