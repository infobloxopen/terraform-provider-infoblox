package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// ZoneforwardforwardingserversForwardToModel is the Terraform model for ZoneforwardforwardingserversForwardTo
type ZoneforwardforwardingserversForwardToModel struct {
	Address                      iptypes.IPAddress `tfsdk:"address"`
	Name                         types.String      `tfsdk:"name"`
	SharedWithMsParentDelegation types.Bool        `tfsdk:"shared_with_ms_parent_delegation"`
	Stealth                      types.Bool        `tfsdk:"stealth"`
}

// ZoneforwardforwardingserversForwardToAttrTypes contains the attribute types for ZoneforwardforwardingserversForwardToModel
var ZoneforwardforwardingserversForwardToAttrTypes = map[string]attr.Type{
	"address":                          iptypes.IPAddressType{},
	"name":                             types.StringType,
	"shared_with_ms_parent_delegation": types.BoolType,
	"stealth":                          types.BoolType,
}

// ZoneforwardforwardingserversForwardToResourceSchemaAttributes contains the schema attributes for ZoneforwardforwardingserversForwardToModel
var ZoneforwardforwardingserversForwardToResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPAddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address or IPv6 Address of the server.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "A resolvable domain name for the external DNS server.",
	},
	"shared_with_ms_parent_delegation": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "This flag represents whether the name server is shared with the parent Microsoft primary zone's delegation server.",
	},
	"stealth": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Set this flag to hide the NS record for the primary name server from DNS queries.",
	},
}

// ExpandZoneforwardforwardingserversForwardTo converts a Terraform Object to SDK type
func ExpandZoneforwardforwardingserversForwardTo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.ZoneforwardforwardingserversForwardTo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ZoneforwardforwardingserversForwardToModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ZoneforwardforwardingserversForwardToModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.ZoneforwardforwardingserversForwardTo {
	if m == nil {
		return nil
	}
	to := &niosdns.ZoneforwardforwardingserversForwardTo{
		Address:                      flex.ExpandIPAddress(m.Address),
		Name:                         flex.ExpandStringPointerNullAsEmpty(m.Name),
		SharedWithMsParentDelegation: flex.ExpandBoolPointer(m.SharedWithMsParentDelegation),
		Stealth:                      flex.ExpandBoolPointer(m.Stealth),
	}
	return to
}

// FlattenZoneforwardforwardingserversForwardTo converts an SDK type to Terraform Object
func FlattenZoneforwardforwardingserversForwardTo(ctx context.Context, from *niosdns.ZoneforwardforwardingserversForwardTo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ZoneforwardforwardingserversForwardToAttrTypes)
	}
	m := &ZoneforwardforwardingserversForwardToModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ZoneforwardforwardingserversForwardToAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ZoneforwardforwardingserversForwardToModel) Flatten(ctx context.Context, from *niosdns.ZoneforwardforwardingserversForwardTo, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenIPAddress(from.Address)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.SharedWithMsParentDelegation = flex.FlattenBoolPointer(from.SharedWithMsParentDelegation)
	m.Stealth = flex.FlattenBoolPointer(from.Stealth)
}
