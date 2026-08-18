package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidhcp "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// HAGroupHeartbeatsModel is the Terraform model for HAGroupHeartbeats
type HAGroupHeartbeatsModel struct {
	Peer                  types.String `tfsdk:"peer"`
	SuccessfulHeartbeat   types.String `tfsdk:"successful_heartbeat"`
	SuccessfulHeartbeatV6 types.String `tfsdk:"successful_heartbeat_v6"`
}

// HAGroupHeartbeatsAttrTypes contains the attribute types for HAGroupHeartbeatsModel
var HAGroupHeartbeatsAttrTypes = map[string]attr.Type{
	"peer":                    types.StringType,
	"successful_heartbeat":    types.StringType,
	"successful_heartbeat_v6": types.StringType,
}

// HAGroupHeartbeatsResourceSchemaAttributes contains the schema attributes for HAGroupHeartbeatsModel
var HAGroupHeartbeatsResourceSchemaAttributes = map[string]schema.Attribute{
	"peer": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the peer.",
	},
	"successful_heartbeat": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The timestamp as a string of the last successful heartbeat received from the peer above.",
	},
	"successful_heartbeat_v6": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The timestamp as a string of the last successful DHCPv6 heartbeat received from the peer above.",
	},
}

// ExpandHAGroupHeartbeats converts a Terraform Object to SDK type
func ExpandHAGroupHeartbeats(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidhcp.HAGroupHeartbeats {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m HAGroupHeartbeatsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *HAGroupHeartbeatsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidhcp.HAGroupHeartbeats {
	if m == nil {
		return nil
	}
	to := &uddidhcp.HAGroupHeartbeats{
		Peer:                  flex.ExpandStringPointer(m.Peer),
		SuccessfulHeartbeat:   flex.ExpandStringPointer(m.SuccessfulHeartbeat),
		SuccessfulHeartbeatV6: flex.ExpandStringPointer(m.SuccessfulHeartbeatV6),
	}
	return to
}

// FlattenHAGroupHeartbeats converts an SDK type to Terraform Object
func FlattenHAGroupHeartbeats(ctx context.Context, from *uddidhcp.HAGroupHeartbeats, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(HAGroupHeartbeatsAttrTypes)
	}
	m := &HAGroupHeartbeatsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, HAGroupHeartbeatsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *HAGroupHeartbeatsModel) Flatten(ctx context.Context, from *uddidhcp.HAGroupHeartbeats, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Peer = flex.FlattenStringPointer(from.Peer)
	m.SuccessfulHeartbeat = flex.FlattenStringPointer(from.SuccessfulHeartbeat)
	m.SuccessfulHeartbeatV6 = flex.FlattenStringPointer(from.SuccessfulHeartbeatV6)
}
