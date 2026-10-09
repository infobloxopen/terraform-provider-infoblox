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
	_ basetypes.StringValuable                   = (*IPv6Prefix)(nil)
	_ basetypes.StringValuableWithSemanticEquals = (*IPv6Prefix)(nil)
	_ xattr.ValidateableAttribute                = (*IPv6Prefix)(nil)
	_ function.ValidateableParameter             = (*IPv6Prefix)(nil)
)

// IPv6Prefix represents a valid IPv6 CIDR string (RFC 4291). Semantic equality logic is defined for IPv6Prefix
// such that a CIDR string with the address zero bits `compressed` will be considered equivalent to the `non-compressed` string.
//
// Examples:
//   - `0:0:0:0:0:0:0:0/128` is semantically equal to `::/128`
//   - `2001:0DB8:0:0:0:0:0:0CD3/60` is semantically equal to `2001:0DB8::CD30/60`
//   - `FF00:0:0:0:0:0:0:0/8` is semantically equal to `FF00::/8`
type IPv6Prefix struct {
	basetypes.StringValue
}

// Type returns an IPv6PrefixType.
func (v IPv6Prefix) Type(_ context.Context) attr.Type {
	return IPv6PrefixType{}
}

// Equal returns true if the given value is equivalent.
func (v IPv6Prefix) Equal(o attr.Value) bool {
	other, ok := o.(IPv6Prefix)

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
// See RFC 4291 for more details on IPv6 CIDR string format: https://www.rfc-editor.org/rfc/rfc4291.html#section-2.3
func (v IPv6Prefix) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(IPv6Prefix)
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

	// IPv6 CIDRs are already validated at this point, ignoring errors
	newIpPrefix, _ := netip.ParsePrefix(newValue.ValueString())
	currentIpPrefix, _ := netip.ParsePrefix(v.ValueString())

	cidrMatch := currentIpPrefix.Addr() == newIpPrefix.Addr() && currentIpPrefix.Bits() == newIpPrefix.Bits()

	return cidrMatch, diags
}

// ValidateAttribute implements attribute value validation. This type requires the value provided to be a String
// value that is a valid IPv6 CIDR.
func (v IPv6Prefix) ValidateAttribute(ctx context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	if v.IsUnknown() || v.IsNull() {
		return
	}

	ipPrefix, err := netip.ParsePrefix(v.ValueString())
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv6 CIDR String Value",
			"A string value was provided that is not valid IPv6 CIDR string format (RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n"+
				"Error: "+err.Error(),
		)

		return
	}

	if ipPrefix.Addr().Is4() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv6 CIDR String Value",
			"An IPv4 CIDR string format was provided, string value must be IPv6 CIDR string format (RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}

	if !ipPrefix.IsValid() || !ipPrefix.Addr().Is6() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv6 CIDR String Value",
			"A string value was provided that is not valid IPv6 CIDR string format (RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}
}

// ValidateParameter implements provider-defined function parameter value validation. This type requires the value
// provided to be a String value that is a valid IPv6 CIDR.
func (v IPv6Prefix) ValidateParameter(ctx context.Context, req function.ValidateParameterRequest, resp *function.ValidateParameterResponse) {
	if v.IsUnknown() || v.IsNull() {
		return
	}

	ipPrefix, err := netip.ParsePrefix(v.ValueString())
	if err != nil {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IPv6 CIDR String Value: "+
				"A string value was provided that is not valid IPv6 CIDR string format (RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n"+
				"Error: "+err.Error(),
		)

		return
	}

	if ipPrefix.Addr().Is4() {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IPv6 CIDR String Value: "+
				"An IPv4 CIDR string format was provided, string value must be IPv6 CIDR string format (RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}

	if !ipPrefix.IsValid() || !ipPrefix.Addr().Is6() {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IPv6 CIDR String Value: "+
				"A string value was provided that is not valid IPv6 CIDR string format (RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}
}

// ValueIPv6Prefix calls netip.ParsePrefix with the IPv6Prefix StringValue. A null or unknown value will produce an error diagnostic.
func (v IPv6Prefix) ValueIPv6Prefix() (netip.Prefix, diag.Diagnostics) {
	var diags diag.Diagnostics

	if v.IsNull() {
		diags.Append(diag.NewErrorDiagnostic("IPv6Prefix ValueIPv6Prefix Error", "IPv6 CIDR string value is null"))
		return netip.Prefix{}, diags
	}

	if v.IsUnknown() {
		diags.Append(diag.NewErrorDiagnostic("IPv6Prefix ValueIPv6Prefix Error", "IPv6 CIDR string value is unknown"))
		return netip.Prefix{}, diags
	}

	ipv6Prefix, err := netip.ParsePrefix(v.ValueString())
	if err != nil {
		diags.Append(diag.NewErrorDiagnostic("IPv6Prefix ValueIPv6Prefix Error", err.Error()))
		return netip.Prefix{}, diags
	}

	return ipv6Prefix, nil
}

// NewIPv6PrefixNull creates an IPv6Prefix with a null value. Determine whether the value is null via IsNull method.
func NewIPv6PrefixNull() IPv6Prefix {
	return IPv6Prefix{
		StringValue: basetypes.NewStringNull(),
	}
}

// NewIPv6PrefixUnknown creates an IPv6Prefix with an unknown value. Determine whether the value is unknown via IsUnknown method.
func NewIPv6PrefixUnknown() IPv6Prefix {
	return IPv6Prefix{
		StringValue: basetypes.NewStringUnknown(),
	}
}

// NewIPv6PrefixValue creates an IPv6Prefix with a known value. Access the value via ValueString method.
func NewIPv6PrefixValue(value string) IPv6Prefix {
	return IPv6Prefix{
		StringValue: basetypes.NewStringValue(value),
	}
}

// NewIPv6PrefixPointerValue creates an IPv6Prefix with a null value if nil or a known value. Access the value via ValueStringPointer method.
func NewIPv6PrefixPointerValue(value *string) IPv6Prefix {
	return IPv6Prefix{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}
