// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cidrtypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ basetypes.StringTypable = (*IPPrefixType)(nil)

// IPPrefixType is an attribute type that represents a valid IPv4 or IPv6 CIDR string (RFC 4632, RFC 4291). Semantic equality logic is defined for IPPrefixType
// such that a CIDR string with the address zero bits `compressed` will be considered equivalent to the `non-compressed` string.
//
// Examples:
//   - `0:0:0:0:0:0:0:0/128` is semantically equal to `::/128`
//   - `2001:0DB8:0:0:0:0:0:0CD3/60` is semantically equal to `2001:0DB8::CD30/60`
//   - `FF00:0:0:0:0:0:0:0/8` is semantically equal to `FF00::/8`
//
// Note that IPv4 CIDR generally do not have a compressed form, so the semantic equality logic is primarily for IPv6 CIDR.
type IPPrefixType struct {
	basetypes.StringType
}

// String returns a human readable string of the type name.
func (t IPPrefixType) String() string {
	return "cidrtypes.IPPrefixType"
}

// ValueType returns the Value type.
func (t IPPrefixType) ValueType(ctx context.Context) attr.Value {
	return IPPrefix{}
}

// Equal returns true if the given type is equivalent.
func (t IPPrefixType) Equal(o attr.Type) bool {
	other, ok := o.(IPPrefixType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

// ValueFromString returns a StringValuable type given a StringValue.
func (t IPPrefixType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return IPPrefix{
		StringValue: in,
	}, nil
}

// ValueFromTerraform returns a Value given a tftypes.Value.  This is meant to convert the tftypes.Value into a more convenient Go type
// for the provider to consume the data with.
func (t IPPrefixType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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
