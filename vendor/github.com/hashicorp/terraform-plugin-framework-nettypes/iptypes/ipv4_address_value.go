// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iptypes

import (
	"context"
	"net/netip"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ basetypes.StringValuable       = (*IPv4Address)(nil)
	_ xattr.ValidateableAttribute    = (*IPv4Address)(nil)
	_ function.ValidateableParameter = (*IPv4Address)(nil)
)

// IPv4Address represents a valid IPv4 address string (dotted decimal, no leading zeroes). No semantic equality
// logic is defined for IPv4Address, so it will follow Terraform's data-consistency rules for strings, which must match byte-for-byte.
type IPv4Address struct {
	basetypes.StringValue
}

// Type returns an IPv4AddressType.
func (v IPv4Address) Type(_ context.Context) attr.Type {
	return IPv4AddressType{}
}

// Equal returns true if the given value is equivalent.
func (v IPv4Address) Equal(o attr.Value) bool {
	other, ok := o.(IPv4Address)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

// ValidateAttribute implements attribute value validation. This type requires the value provided to be a String
// value that is a valid IPv4 address. This utilizes the Go `net/netip` library for parsing so leading zeroes
// will be rejected as invalid.
func (v IPv4Address) ValidateAttribute(ctx context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	if v.IsUnknown() || v.IsNull() {
		return
	}

	ipAddr, err := netip.ParseAddr(v.ValueString())
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv4 Address String Value",
			"A string value was provided that is not valid IPv4 string format.\n\n"+
				"Given Value: "+v.ValueString()+"\n"+
				"Error: "+err.Error(),
		)

		return
	}

	if ipAddr.Is6() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv4 Address String Value",
			"An IPv6 string format was provided, string value must be IPv4 format.\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}

	if !ipAddr.IsValid() || !ipAddr.Is4() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IPv4 Address String Value",
			"A string value was provided that is not valid IPv4 string format.\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}
}

// ValidateParameter implements provider-defined function parameter value validation. This type requires the value
// provided to be a String value that is a valid IPv4 address. This utilizes the Go `net/netip` library for
// parsing so leading zeroes will be rejected as invalid.
func (v IPv4Address) ValidateParameter(ctx context.Context, req function.ValidateParameterRequest, resp *function.ValidateParameterResponse) {
	if v.IsUnknown() || v.IsNull() {
		return
	}

	ipAddr, err := netip.ParseAddr(v.ValueString())
	if err != nil {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IPv4 Address String Value: "+
				"A string value was provided that is not valid IPv4 string format.\n\n"+
				"Given Value: "+v.ValueString()+"\n"+
				"Error: "+err.Error(),
		)

		return
	}

	if ipAddr.Is6() {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IPv4 Address String Value: "+
				"An IPv6 string format was provided, string value must be IPv4 format.\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}

	if !ipAddr.IsValid() || !ipAddr.Is4() {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid IPv4 Address String Value: "+
				"A string value was provided that is not valid IPv4 string format.\n\n"+
				"Given Value: "+v.ValueString()+"\n",
		)

		return
	}
}

// ValueIPv4Address calls netip.ParseAddr with the IPv4Address StringValue. A null or unknown value will produce an error diagnostic.
func (v IPv4Address) ValueIPv4Address() (netip.Addr, diag.Diagnostics) {
	var diags diag.Diagnostics

	if v.IsNull() {
		diags.Append(diag.NewErrorDiagnostic("IPv4Address ValueIPv4Address Error", "IPv4 address string value is null"))
		return netip.Addr{}, diags
	}

	if v.IsUnknown() {
		diags.Append(diag.NewErrorDiagnostic("IPv4Address ValueIPv4Address Error", "IPv4 address string value is unknown"))
		return netip.Addr{}, diags
	}

	ipv4Addr, err := netip.ParseAddr(v.ValueString())
	if err != nil {
		diags.Append(diag.NewErrorDiagnostic("IPv4Address ValueIPv4Address Error", err.Error()))
		return netip.Addr{}, diags
	}

	return ipv4Addr, nil
}

// NewIPv4AddressNull creates an IPv4Address with a null value. Determine whether the value is null via IsNull method.
func NewIPv4AddressNull() IPv4Address {
	return IPv4Address{
		StringValue: basetypes.NewStringNull(),
	}
}

// NewIPv4AddressUnknown creates an IPv4Address with an unknown value. Determine whether the value is unknown via IsUnknown method.
func NewIPv4AddressUnknown() IPv4Address {
	return IPv4Address{
		StringValue: basetypes.NewStringUnknown(),
	}
}

// NewIPv4AddressValue creates an IPv4Address with a known value. Access the value via ValueString method.
func NewIPv4AddressValue(value string) IPv4Address {
	return IPv4Address{
		StringValue: basetypes.NewStringValue(value),
	}
}

// NewIPv4AddressPointerValue creates an IPv4Address with a null value if nil or a known value. Access the value via ValueStringPointer method.
func NewIPv4AddressPointerValue(value *string) IPv4Address {
	return IPv4Address{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}
