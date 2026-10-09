// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iptypes

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
	_ basetypes.StringValuable                   = (*IPv6Address)(nil)
	_ basetypes.StringValuableWithSemanticEquals = (*IPv6Address)(nil)
	_ xattr.ValidateableAttribute                = (*IPv6Address)(nil)
	_ function.ValidateableParameter             = (*IPv6Address)(nil)
)

// IPv6Address represents a valid IPv6 address string (RFC 4291). Semantic equality logic is defined for IPv6Address
// such that an address string with the zero bits `compressed` will be considered equivalent to the `non-compressed` string.
//
// Examples:
//   - `0:0:0:0:0:0:0:0` is semantically equal to `::`
//   - `2001:DB8:0:0:8:800:200C:417A` is semantically equal to `2001:DB8::8:800:200C:417A`
//   - `FF01:0:0:0:0:0:0:101` is semantically equal to `FF01::101`
//
// IPv6Address also supports IPv6 address strings with embedded IPv4 addresses, see RFC 4291 for more details: https://www.rfc-editor.org/rfc/rfc4291.html#section-2.5.5
type IPv6Address struct {
	basetypes.StringValue
}

// Type returns an IPv6AddressType.
func (v IPv6Address) Type(_ context.Context) attr.Type {
	return IPv6AddressType{}
}

// Equal returns true if the given value is equivalent.
func (v IPv6Address) Equal(o attr.Value) bool {
	other, ok := o.(IPv6Address)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

// StringSemanticEquals returns true if the given IPv6 address string value is semantically equal to the current IPv6 address string value.
// This comparison utilizes netip.ParseAddr and then compares the resulting netip.Addr representations. This means `compressed` IPv6 address values
// are considered semantically equal to `non-compressed` IPv6 address values.
//
// Examples:
//   - `0:0:0:0:0:0:0:0` is semantically equal to `::`
//   - `2001:DB8:0:0:8:800:200C:417A` is semantically equal to `2001:DB8::8:800:200C:417A`
//   - `FF01:0:0:0:0:0:0:101` is semantically equal to `FF01::101`
//
// See RFC 4291 for more details on IPv6 string format: https://www.rfc-editor.org/rfc/rfc4291.html#section-2.2
func (v IPv6Address) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(IPv6Address)
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

	// IPv6 addresses are already validated at this point, ignoring errors
	newIpAddr, _ := netip.ParseAddr(newValue.ValueString())
	currentIpAddr, _ := netip.ParseAddr(v.ValueString())

	return currentIpAddr == newIpAddr, diags
}

// ValidateAttribute implements attribute value validation. This type requires the value provided to be a String
// value that is a valid IPv6 address.
func (v IPv6Address) ValidateAttribute(ctx context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	if v.IsUnknown() || v.IsNull() {
		return
	}

	ipAddr, err := netip.ParseAddr(v.ValueString())
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv6 Address String Value",
			"A string value was provided that is not valid IPv6 string format (RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n"+
				"Error: "+err.Error(),
		)

		return
	}

	if ipAddr.Is4() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv6 Address String Value",
			"An IPv4 string format was provided, string value must be IPv6 string format or IPv4-Mapped IPv6 string format (RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}

	if !ipAddr.IsValid() || !ipAddr.Is6() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv6 Address String Value",
			"A string value was provided that is not valid IPv6 string format (RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}
}

// ValidateParameter implements provider-defined function parameter value validation. This type requires the value provided
// to be a String value that is a valid IPv6 address.
func (v IPv6Address) ValidateParameter(ctx context.Context, req function.ValidateParameterRequest, resp *function.ValidateParameterResponse) {
	if v.IsUnknown() || v.IsNull() {
		return
	}

	ipAddr, err := netip.ParseAddr(v.ValueString())
	if err != nil {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IPv6 Address String Value: "+
				"A string value was provided that is not valid IPv6 string format (RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n"+
				"Error: "+err.Error(),
		)

		return
	}

	if ipAddr.Is4() {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IPv6 Address String Value: "+
				"An IPv4 string format was provided, string value must be IPv6 string format or IPv4-Mapped IPv6 string format (RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}

	if !ipAddr.IsValid() || !ipAddr.Is6() {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IPv6 Address String Value: "+
				"A string value was provided that is not valid IPv6 string format (RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}
}

// ValueIPv6Address calls netip.ParseAddr with the IPv6Address StringValue. A null or unknown value will produce an error diagnostic.
func (v IPv6Address) ValueIPv6Address() (netip.Addr, diag.Diagnostics) {
	var diags diag.Diagnostics

	if v.IsNull() {
		diags.Append(diag.NewErrorDiagnostic("IPv6Address ValueIPv6Address Error", "IPv6 address string value is null"))
		return netip.Addr{}, diags
	}

	if v.IsUnknown() {
		diags.Append(diag.NewErrorDiagnostic("IPv6Address ValueIPv6Address Error", "IPv6 address string value is unknown"))
		return netip.Addr{}, diags
	}

	ipv6Addr, err := netip.ParseAddr(v.ValueString())
	if err != nil {
		diags.Append(diag.NewErrorDiagnostic("IPv6Address ValueIPv6Address Error", err.Error()))
		return netip.Addr{}, diags
	}

	return ipv6Addr, nil
}

// NewIPv6AddressNull creates an IPv6Address with a null value. Determine whether the value is null via IsNull method.
func NewIPv6AddressNull() IPv6Address {
	return IPv6Address{
		StringValue: basetypes.NewStringNull(),
	}
}

// NewIPv6AddressUnknown creates an IPv6Address with an unknown value. Determine whether the value is unknown via IsUnknown method.
func NewIPv6AddressUnknown() IPv6Address {
	return IPv6Address{
		StringValue: basetypes.NewStringUnknown(),
	}
}

// NewIPv6AddressValue creates an IPv6Address with a known value. Access the value via ValueString method.
func NewIPv6AddressValue(value string) IPv6Address {
	return IPv6Address{
		StringValue: basetypes.NewStringValue(value),
	}
}

// NewIPv6AddressPointerValue creates an IPv6Address with a null value if nil or a known value. Access the value via ValueStringPointer method.
func NewIPv6AddressPointerValue(value *string) IPv6Address {
	return IPv6Address{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}
