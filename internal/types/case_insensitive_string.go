package types

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ basetypes.StringTypable = (*CaseInsensitiveString)(nil)
var _ basetypes.StringValuableWithSemanticEquals = (*CaseInsensitiveStringValue)(nil)

type CaseInsensitiveString struct {
	basetypes.StringType
}

func (t CaseInsensitiveString) Equal(o attr.Type) bool {
	other, ok := o.(CaseInsensitiveString)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t CaseInsensitiveString) String() string {
	return "customtypes.CaseInsensitiveStringType"
}

func (t CaseInsensitiveString) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	// CaseInsensitiveStringValue defined in the value type section
	value := CaseInsensitiveStringValue{
		StringValue: in,
	}

	return value, nil
}

func (t CaseInsensitiveString) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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

func (t CaseInsensitiveString) Validate(ctx context.Context, in tftypes.Value, path path.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	if in.Type() == nil {
		return diags
	}

	return diags
}

func (t CaseInsensitiveString) ValueType(ctx context.Context) attr.Value {
	return CaseInsensitiveStringValue{}
}

type CaseInsensitiveStringValue struct {
	basetypes.StringValue
}

// StringSemanticEquals implements the custom semantic equality hook for string-like
// custom types. The framework will call oldVal.StringSemanticEquals(ctx, newVal)
func (old CaseInsensitiveStringValue) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	if old.IsUnknown() || newValuable.IsUnknown() {
		return false, diags
	}

	if old.IsNull() && newValuable.IsNull() {
		return true, diags
	}

	if old.IsNull() || newValuable.IsNull() {
		return false, diags
	}

	oldStr := old.ValueString()
	newStr := newValuable.String()

	// Normalize using simple lower-case. Use more advanced normalization if needed.
	return strings.EqualFold(oldStr, newStr), diags
}

func (v CaseInsensitiveStringValue) Type(ctx context.Context) attr.Type {
	return CaseInsensitiveString{}
}

func (v CaseInsensitiveStringValue) ValueString() string {
	return v.StringValue.ValueString()
}

func (v CaseInsensitiveStringValue) String() string {
	return v.StringValue.ValueString()
}

func (v CaseInsensitiveStringValue) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	if in.IsNull() {
		return CaseInsensitiveStringValue{basetypes.NewStringNull()}, diags
	}

	var value string
	err := in.As(&value)
	if err != nil {
		diags.AddError("Error converting value", err.Error())
		return nil, diags
	}

	return CaseInsensitiveStringValue{basetypes.NewStringValue(value)}, diags
}

// ValueType returns the value type of the CaseInsensitiveStringValue.
func (v CaseInsensitiveStringValue) ValueType(_ context.Context) attr.Value {
	return CaseInsensitiveStringValue{}
}

// NewCaseInsensitiveStringNull creates an CaseInsensitiveStringValue with a null value. Determine whether the value is null via IsNull method.
func NewCaseInsensitiveStringNull() CaseInsensitiveStringValue {
	return CaseInsensitiveStringValue{
		StringValue: basetypes.NewStringNull(),
	}
}

// NewCaseInsensitiveStringUnknown creates an CaseInsensitiveStringValue with an unknown value. Determine whether the value is unknown via IsUnknown method.
func NewCaseInsensitiveStringUnknown() CaseInsensitiveStringValue {
	return CaseInsensitiveStringValue{
		StringValue: basetypes.NewStringUnknown(),
	}
}

// NewCaseInsensitiveStringValue creates an CaseInsensitiveStringValue with a known value. Access the value via ValueString method.
func NewCaseInsensitiveStringValue(value string) CaseInsensitiveStringValue {
	return CaseInsensitiveStringValue{
		StringValue: basetypes.NewStringValue(value),
	}
}

func NewCaseInsensitiveStringPointerValue(value *string) CaseInsensitiveStringValue {
	return CaseInsensitiveStringValue{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}
