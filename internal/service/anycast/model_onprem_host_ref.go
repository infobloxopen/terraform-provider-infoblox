package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	int64planmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddianycast "github.com/infobloxopen/universal-ddi-go-client/anycast"
)

// OnpremHostRefModel is the Terraform model for OnpremHostRef
type OnpremHostRefModel struct {
	Id            types.Int64         `tfsdk:"id"`
	IpAddress     iptypes.IPv4Address `tfsdk:"ip_address"`
	Ipv6Address   iptypes.IPv6Address `tfsdk:"ipv6_address"`
	Name          types.String        `tfsdk:"name"`
	Ophid         types.String        `tfsdk:"ophid"`
	RuntimeStatus types.String        `tfsdk:"runtime_status"`
}

// OnpremHostRefAttrTypes contains the attribute types for OnpremHostRefModel
var OnpremHostRefAttrTypes = map[string]attr.Type{
	"id":             types.Int64Type,
	"ip_address":     iptypes.IPv4AddressType{},
	"ipv6_address":   iptypes.IPv6AddressType{},
	"name":           types.StringType,
	"ophid":          types.StringType,
	"runtime_status": types.StringType,
}

// OnpremHostRefResourceSchemaAttributes contains the schema attributes for OnpremHostRefModel
var OnpremHostRefResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.Int64Attribute{
		Required: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"ip_address": schema.StringAttribute{
		Optional:            true,
		CustomType:          iptypes.IPv4AddressType{},
		MarkdownDescription: "IPv4 address of the host in string format",
	},
	"ipv6_address": schema.StringAttribute{
		Optional:            true,
		CustomType:          iptypes.IPv6AddressType{},
		MarkdownDescription: "IPv6 address of the host in string format",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the anycast.",
	},
	"ophid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Unique 32-character string identifier assigned to the host",
	},
	"runtime_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The runtime status of the host",
	},
}

// ExpandOnpremHostRef converts a Terraform Object to SDK type
func ExpandOnpremHostRef(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddianycast.OnpremHostRef {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m OnpremHostRefModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *OnpremHostRefModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddianycast.OnpremHostRef {
	if m == nil {
		return nil
	}
	to := &uddianycast.OnpremHostRef{
		Id:            flex.ExpandInt64Pointer(m.Id),
		IpAddress:     flex.ExpandIPv4Address(m.IpAddress),
		Ipv6Address:   flex.ExpandIPv6Address(m.Ipv6Address),
		Name:          flex.ExpandStringPointer(m.Name),
		Ophid:         flex.ExpandStringPointer(m.Ophid),
		RuntimeStatus: flex.ExpandStringPointer(m.RuntimeStatus),
	}
	return to
}

// FlattenOnpremHostRef converts an SDK type to Terraform Object
func FlattenOnpremHostRef(ctx context.Context, from *uddianycast.OnpremHostRef, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(OnpremHostRefAttrTypes)
	}
	m := &OnpremHostRefModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, OnpremHostRefAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *OnpremHostRefModel) Flatten(ctx context.Context, from *uddianycast.OnpremHostRef, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.IpAddress = flex.FlattenIPv4Address(from.IpAddress)
	m.Ipv6Address = flex.FlattenIPv6Address(from.Ipv6Address)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.RuntimeStatus = flex.FlattenStringPointer(from.RuntimeStatus)
}
