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

// NsgroupforwardingmemberforwardingserversForwardToModel is the Terraform model for NsgroupforwardingmemberforwardingserversForwardTo
type NsgroupforwardingmemberforwardingserversForwardToModel struct {
	Address iptypes.IPv4Address `tfsdk:"address"`
	Name    types.String        `tfsdk:"name"`
}

// NsgroupforwardingmemberforwardingserversForwardToAttrTypes contains the attribute types for NsgroupforwardingmemberforwardingserversForwardToModel
var NsgroupforwardingmemberforwardingserversForwardToAttrTypes = map[string]attr.Type{
	"address": iptypes.IPv4AddressType{},
	"name":    types.StringType,
}

// NsgroupforwardingmemberforwardingserversForwardToResourceSchemaAttributes contains the schema attributes for NsgroupforwardingmemberforwardingserversForwardToModel
var NsgroupforwardingmemberforwardingserversForwardToResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPv4AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
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
}

// ExpandNsgroupforwardingmemberforwardingserversForwardTo converts a Terraform Object to SDK type
func ExpandNsgroupforwardingmemberforwardingserversForwardTo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.NsgroupforwardingmemberforwardingserversForwardTo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NsgroupforwardingmemberforwardingserversForwardToModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NsgroupforwardingmemberforwardingserversForwardToModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.NsgroupforwardingmemberforwardingserversForwardTo {
	if m == nil {
		return nil
	}
	to := &niosdns.NsgroupforwardingmemberforwardingserversForwardTo{
		Address: flex.ExpandIPv4Address(m.Address),
		Name:    flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
	return to
}

// FlattenNsgroupforwardingmemberforwardingserversForwardTo converts an SDK type to Terraform Object
func FlattenNsgroupforwardingmemberforwardingserversForwardTo(ctx context.Context, from *niosdns.NsgroupforwardingmemberforwardingserversForwardTo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NsgroupforwardingmemberforwardingserversForwardToAttrTypes)
	}
	m := &NsgroupforwardingmemberforwardingserversForwardToModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NsgroupforwardingmemberforwardingserversForwardToAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NsgroupforwardingmemberforwardingserversForwardToModel) Flatten(ctx context.Context, from *niosdns.NsgroupforwardingmemberforwardingserversForwardTo, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenIPv4Address(from.Address)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
