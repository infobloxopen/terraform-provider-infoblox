// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iptypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ basetypes.StringTypable = (*IPAddressType)(nil)

// IPAddressType is an attribute type that represents a valid IPv4 or IPv6 address string (RFC 791, RFC 4291). Semantic equality logic is defined for IPAddressType
// such that an address string with the zero bits `compressed` will be considered equivalent to the `non-compressed` string.
//
// Examples:
//   - `0:0:0:0:0:0:0:0` is semantically equal to `::`
//   - `2001:DB8:0:0:8:800:200C:417A` is semantically equal to `2001:DB8::8:800:200C:417A`
//   - `FF01:0:0:0:0:0:0:101` is semantically equal to `FF01::101`
//
// Note that IPv4 addresses generally do not have a compressed form, so the semantic equality logic is primarily for IPv6 addresses.
//
// IPAddressType also supports IP address strings with embedded IPv4 addresses, see RFC 4291 for more details: https://www.rfc-editor.org/rfc/rfc4291.html#section-2.5.5
// Also see RFC 791 for more details on IP address string format: https://www.rfc-editor.org/rfc/rfc791.html#section-3.2
type IPAddressType struct {
	basetypes.StringType
}

// String returns a human readable string of the type name.
func (t IPAddressType) String() string {
	return "iptypes.IPAddressType"
}

// ValueType returns the Value type.
func (t IPAddressType) ValueType(ctx context.Context) attr.Value {
	return IPAddress{}
}

// Equal returns true if the given type is equivalent.
func (t IPAddressType) Equal(o attr.Type) bool {
	other, ok := o.(IPAddressType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

// ValueFromString returns a StringValuable type given a StringValue.
func (t IPAddressType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return IPAddress{
		StringValue: in,
	}, nil
}

// ValueFromTerraform returns a Value given a tftypes.Value.  This is meant to convert the tftypes.Value into a more convenient Go type
// for the provider to consume the data with.
func (t IPAddressType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}
