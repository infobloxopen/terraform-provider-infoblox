package types

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.StringTypable = (*IPNameType)(nil)
)

type IPNameType struct {
	basetypes.StringType
}

// String returns a human readable string of the type name.
func (t IPNameType) String() string {
	return "iptypes.IPAddressNameType"
}

// ValueType returns the Value type.
func (t IPNameType) ValueType(ctx context.Context) attr.Value {
	return IPName{}
}

// Equal returns true if the given type is equivalent.
func (t IPNameType) Equal(o attr.Type) bool {
	other, ok := o.(IPNameType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

// ValueFromString returns a StringValuable type given a StringValue.
func (t IPNameType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return IPName{
		StringValue: in,
	}, nil
}

// ValueFromTerraform returns a Value given a tftypes.Value. This is meant to convert the tftypes.Value into a more convenient Go type
// for the provider to consume the data with.
func (t IPNameType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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
