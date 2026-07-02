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
	_ basetypes.StringValuable                   = (*IPAddress)(nil)
	_ basetypes.StringValuableWithSemanticEquals = (*IPAddress)(nil)
	_ xattr.ValidateableAttribute                = (*IPAddress)(nil)
	_ function.ValidateableParameter             = (*IPAddress)(nil)
)

// IPAddress represents a valid IPv4 or IPv6 address string (RFC 791, RFC 4291). Semantic equality logic is defined for IPAddress
// such that an IPv6 address string with the zero bits `compressed` will be considered equivalent to the `non-compressed` string.
//
// Examples:
//   - `0:0:0:0:0:0:0:0` is semantically equal to `::`
//   - `2001:DB8:0:0:8:800:200C:417A` is semantically equal to `2001:DB8::8:800:200C:417A`
//   - `FF01:0:0:0:0:0:0:101` is semantically equal to `FF01::101`
//
// Note that IPv4 addresses generally do not have a compressed form, so the semantic equality logic is primarily for IPv6 addresses.
//
// IPAddress also supports IPv6 address strings with embedded IPv4 addresses, see RFC 4291 for more details: https://www.rfc-editor.org/rfc/rfc4291.html#section-2.5.5
// Also see RFC 791 for more details on IPv4 string format: https://www.rfc-editor.org/rfc/rfc791.html#section-3.2
type IPAddress struct {
	basetypes.StringValue
}

// Type returns an IPAddressType.
func (v IPAddress) Type(_ context.Context) attr.Type {
	return IPAddressType{}
}

// Equal returns true if the given value is equivalent.
func (v IPAddress) Equal(o attr.Value) bool {
	other, ok := o.(IPAddress)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

// StringSemanticEquals returns true if the given IPv6 address string value is semantically equal to the current IPv6 address string value.
// This comparison utilizes netip.ParseAddr and then compares the resulting netip.Addr representations. This means `compressed` IPv6 address values
// are considered semantically equal to `non-compressed` IPv6 address values.
//
// Note that IPv4 addresses generally do not have a compressed form, so the semantic equality logic is primarily for IPv6 addresses.
//
// Examples:
//   - `0:0:0:0:0:0:0:0` is semantically equal to `::`
//   - `2001:DB8:0:0:8:800:200C:417A` is semantically equal to `2001:DB8::8:800:200C:417A`
//   - `FF01:0:0:0:0:0:0:101` is semantically equal to `FF01::101`
//
// See RFC 4291 for more details on IPv6 string format: https://www.rfc-editor.org/rfc/rfc4291.html#section-2.2
// See RFC 791 for more details on IPv4 string format: https://www.rfc-editor.org/rfc/rfc791.html#section-3.2
func (v IPAddress) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(IPAddress)
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

	// IP addresses are already validated at this point, ignoring errors
	newIpAddr, _ := netip.ParseAddr(newValue.ValueString())
	currentIpAddr, _ := netip.ParseAddr(v.ValueString())

	return currentIpAddr == newIpAddr, diags
}

// ValidateAttribute implements attribute value validation. This type requires the value provided to be a String
// value that is a valid IP address.
func (v IPAddress) ValidateAttribute(ctx context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	if v.IsUnknown() || v.IsNull() {
		return
	}

	ipAddr, err := netip.ParseAddr(v.ValueString())
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IP Address String Value",
			"A string value was provided that is not valid IPv4 or IPv6 string format (RFC 791, RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n"+
				"Error: "+err.Error(),
		)

		return
	}

	if !ipAddr.IsValid() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IP Address String Value",
			"A string value was provided that is not valid IPv4 or IPv6 string format (RFC 791, RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}
}

// ValidateParameter implements provider-defined function parameter value validation. This type requires the value provided
// to be a String value that is a valid IP address.
func (v IPAddress) ValidateParameter(ctx context.Context, req function.ValidateParameterRequest, resp *function.ValidateParameterResponse) {
	if v.IsUnknown() || v.IsNull() {
		return
	}

	ipAddr, err := netip.ParseAddr(v.ValueString())
	if err != nil {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IP Address String Value: "+
				"A string value was provided that is not valid IPv4 or IPv6 string format (RFC 791, RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n"+
				"Error: "+err.Error(),
		)

		return
	}
	if !ipAddr.IsValid() {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IP Address String Value: "+
				"A string value was provided that is not valid IPv4 or IPv6 string format (RFC 791, RFC 4291).\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}
}

// ValueIPAddress calls netip.ParseAddr with the IPAddress StringValue. A null or unknown value will produce an error diagnostic.
func (v IPAddress) ValueIPAddress() (netip.Addr, diag.Diagnostics) {
	var diags diag.Diagnostics

	if v.IsNull() {
		diags.Append(diag.NewErrorDiagnostic("IPAddress ValueIPAddress Error", "IP address string value is null"))
		return netip.Addr{}, diags
	}

	if v.IsUnknown() {
		diags.Append(diag.NewErrorDiagnostic("IPAddress ValueIPAddress Error", "IP address string value is unknown"))
		return netip.Addr{}, diags
	}

	ipAddr, err := netip.ParseAddr(v.ValueString())
	if err != nil {
		diags.Append(diag.NewErrorDiagnostic("IPAddress ValueIPAddress Error", err.Error()))
		return netip.Addr{}, diags
	}

	return ipAddr, nil
}

// NewIPAddressNull creates an IPAddress with a null value. Determine whether the value is null via IsNull method.
func NewIPAddressNull() IPAddress {
	return IPAddress{
		StringValue: basetypes.NewStringNull(),
	}
}

// NewIPAddressUnknown creates an IPAddress with an unknown value. Determine whether the value is unknown via IsUnknown method.
func NewIPAddressUnknown() IPAddress {
	return IPAddress{
		StringValue: basetypes.NewStringUnknown(),
	}
}

// NewIPAddressValue creates an IPAddress with a known value. Access the value via ValueString method.
func NewIPAddressValue(value string) IPAddress {
	return IPAddress{
		StringValue: basetypes.NewStringValue(value),
	}
}

// NewIPAddressPointerValue creates an IPAddress with a null value if nil or a known value. Access the value via ValueStringPointer method.
func NewIPAddressPointerValue(value *string) IPAddress {
	return IPAddress{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}
