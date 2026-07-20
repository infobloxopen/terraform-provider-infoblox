// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cidrtypes

import (
	"context"
	"fmt"
	"net/netip"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ basetypes.StringValuable                   = (*IPPrefix)(nil)
	_ basetypes.StringValuableWithSemanticEquals = (*IPPrefix)(nil)
	_ xattr.ValidateableAttribute                = (*IPPrefix)(nil)
	_ function.ValidateableParameter             = (*IPPrefix)(nil)
)

// IPPrefix represents a valid IPv4 or IPv6 CIDR string (RFC 4632, RFC 4291). Semantic equality logic is defined for IPPrefix
// such that a CIDR string with the address zero bits `compressed` will be considered equivalent to the `non-compressed` string.
//
// Examples:
//   - `0:0:0:0:0:0:0:0/128` is semantically equal to `::/128`
//   - `2001:0DB8:0:0:0:0:0:0CD3/60` is semantically equal to `2001:0DB8::CD30/60`
//   - `FF00:0:0:0:0:0:0:0/8` is semantically equal to `FF00::/8`
//
// Note that IPv4 CIDRs generally do not have a compressed form, so the semantic equality logic is primarily for IPv6 CIDR.
type IPPrefix struct {
	basetypes.StringValue
}

// Type returns an IPPrefixType.
func (v IPPrefix) Type(_ context.Context) attr.Type {
	return IPPrefixType{}
}

// Equal returns true if the given value is equivalent.
func (v IPPrefix) Equal(o attr.Value) bool {
	other, ok := o.(IPPrefix)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

// StringSemanticEquals returns true if the given IPv6 CIDR string value is semantically equal to the current IPv6 CIDR string value.
// This comparison utilizes netip.ParsePrefix and then compares the resulting netip.Prefix representations (comparing (Prefix).Addr() and
// (Prefix).Bits() respectively). This means `compressed` IPv6 CIDR values are considered semantically equal to `non-compressed` IPv6 CIDR values.
//
// Examples:
//   - `0:0:0:0:0:0:0:0/128` is semantically equal to `::/128`
//   - `2001:0DB8:0:0:0:0:0:0CD3/60` is semantically equal to `2001:0DB8::CD30/60`
//   - `FF00:0:0:0:0:0:0:0/8` is semantically equal to `FF00::/8`
//
// Note that IPv4 CIDRs generally do not have a compressed form, so the semantic equality logic is primarily for IPv6 CIDRs.
//
// See RFC 4291 for more details on IPv6 CIDR string format: https://www.rfc-editor.org/rfc/rfc4291.html#section-2.3
func (v IPPrefix) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(IPPrefix)
	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", v)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)

		return false, diags
	}

	// IP CIDRs are already validated at this point, ignoring errors
	newIpPrefix, _ := netip.ParsePrefix(newValue.ValueString())
	currentIpPrefix, _ := netip.ParsePrefix(v.ValueString())

	cidrMatch := currentIpPrefix.Addr() == newIpPrefix.Addr() && currentIpPrefix.Bits() == newIpPrefix.Bits()

	return cidrMatch, diags
}

// ValidateAttribute implements attribute value validation. This type requires the value provided to be a String
// value that is a valid IP CIDR.
func (v IPPrefix) ValidateAttribute(ctx context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	if v.IsUnknown() || v.IsNull() {
		return
	}

	ipPrefix, err := netip.ParsePrefix(v.ValueString())
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IP CIDR String Value",
			"A string value was provided that is not valid IPv4 or IPv6 CIDR string format (RFC 4632, RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n"+
				"Error: "+err.Error(),
		)

		return
	}

	if !ipPrefix.IsValid() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IP CIDR String Value",
			"A string value was provided that is not valid IPv4 or IPv6 CIDR string format (RFC 4632, RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}
}

// ValidateParameter implements provider-defined function parameter value validation. This type requires the value
// provided to be a String value that is a valid IP CIDR.
func (v IPPrefix) ValidateParameter(ctx context.Context, req function.ValidateParameterRequest, resp *function.ValidateParameterResponse) {
	if v.IsUnknown() || v.IsNull() {
		return
	}

	ipPrefix, err := netip.ParsePrefix(v.ValueString())
	if err != nil {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IP CIDR String Value: "+
				"A string value was provided that is not valid IPv4 or IPv6 CIDR string format (RFC 4632, RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n"+
				"Error: "+err.Error(),
		)

		return
	}

	if !ipPrefix.IsValid() {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IP CIDR String Value: "+
				"A string value was provided that is not valid IPv4 or IPv6 CIDR string format (RFC 4632, RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}
}

// ValueIPPrefix calls netip.ParsePrefix with the IPPrefix StringValue. A null or unknown value will produce an error diagnostic.
func (v IPPrefix) ValueIPPrefix() (netip.Prefix, diag.Diagnostics) {
	var diags diag.Diagnostics

	if v.IsNull() {
		diags.Append(diag.NewErrorDiagnostic("IPPrefix ValueIPPrefix Error", "IP CIDR string value is null"))
		return netip.Prefix{}, diags
	}

	if v.IsUnknown() {
		diags.Append(diag.NewErrorDiagnostic("IPPrefix ValueIPPrefix Error", "IP CIDR string value is unknown"))
		return netip.Prefix{}, diags
	}

	ipv6Prefix, err := netip.ParsePrefix(v.ValueString())
	if err != nil {
		diags.Append(diag.NewErrorDiagnostic("IPPrefix ValueIPPrefix Error", err.Error()))
		return netip.Prefix{}, diags
	}

	return ipv6Prefix, nil
}

// NewIPPrefixNull creates an IPPrefix with a null value. Determine whether the value is null via IsNull method.
func NewIPPrefixNull() IPPrefix {
	return IPPrefix{
		StringValue: basetypes.NewStringNull(),
	}
}

// NewIPPrefixUnknown creates an IPPrefix with an unknown value. Determine whether the value is unknown via IsUnknown method.
func NewIPPrefixUnknown() IPPrefix {
	return IPPrefix{
		StringValue: basetypes.NewStringUnknown(),
	}
}

// NewIPPrefixValue creates an IPPrefix with a known value. Access the value via ValueString method.
func NewIPPrefixValue(value string) IPPrefix {
	return IPPrefix{
		StringValue: basetypes.NewStringValue(value),
	}
}

// NewIPPrefixPointerValue creates an IPPrefix with a null value if nil or a known value. Access the value via ValueStringPointer method.
func NewIPPrefixPointerValue(value *string) IPPrefix {
	return IPPrefix{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}
