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

var _ basetypes.StringTypable = (*CaseInsensitiveStringType)(nil)
var _ basetypes.StringValuableWithSemanticEquals = (*CaseInsensitiveString)(nil)

type CaseInsensitiveStringType struct {
	basetypes.StringType
}

func (t CaseInsensitiveStringType) Equal(o attr.Type) bool {
	other, ok := o.(CaseInsensitiveStringType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t CaseInsensitiveStringType) String() string {
	return "types.CaseInsensitiveStringType"
}

func (t CaseInsensitiveStringType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	// CaseInsensitiveString defined in the value type section
	value := CaseInsensitiveString{
		StringValue: in,
	}

	return value, nil
}

func (t CaseInsensitiveStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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

func (t CaseInsensitiveStringType) Validate(ctx context.Context, in tftypes.Value, path path.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	if in.Type() == nil {
		return diags
	}

	return diags
}

func (t CaseInsensitiveStringType) ValueType(ctx context.Context) attr.Value {
	return CaseInsensitiveString{}
}

type CaseInsensitiveString struct {
	basetypes.StringValue
}

// StringSemanticEquals implements the custom semantic equality hook for string-like
// custom types. The framework will call oldVal.StringSemanticEquals(ctx, newVal)
func (old CaseInsensitiveString) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
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

func (v CaseInsensitiveString) Type(ctx context.Context) attr.Type {
	return CaseInsensitiveStringType{}
}

func (v CaseInsensitiveString) ValueString() string {
	return v.StringValue.ValueString()
}

func (v CaseInsensitiveString) String() string {
	return v.StringValue.ValueString()
}

func (v CaseInsensitiveString) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	if in.IsNull() {
		return CaseInsensitiveString{basetypes.NewStringNull()}, diags
	}

	var value string
	err := in.As(&value)
	if err != nil {
		diags.AddError("Error converting value", err.Error())
		return nil, diags
	}

	return CaseInsensitiveString{basetypes.NewStringValue(value)}, diags
}

// ValueType returns the value type of the CaseInsensitiveString.
func (v CaseInsensitiveString) ValueType(_ context.Context) attr.Value {
	return CaseInsensitiveString{}
}

// NewCaseInsensitiveStringValueNull creates an CaseInsensitiveString with a null value. Determine whether the value is null via IsNull method.
func NewCaseInsensitiveStringValueNull() CaseInsensitiveString {
	return CaseInsensitiveString{
		StringValue: basetypes.NewStringNull(),
	}
}

// NewCaseInsensitiveStringValueUnknown creates an CaseInsensitiveString with an unknown value. Determine whether the value is unknown via IsUnknown method.
func NewCaseInsensitiveStringValueUnknown() CaseInsensitiveString {
	return CaseInsensitiveString{
		StringValue: basetypes.NewStringUnknown(),
	}
}

// NewCaseInsensitiveStringValue creates an CaseInsensitiveString with a known value. Access the value via ValueString method.
func NewCaseInsensitiveStringValue(value string) CaseInsensitiveString {
	return CaseInsensitiveString{
		StringValue: basetypes.NewStringValue(value),
	}
}

func NewCaseInsensitiveStringPointerValue(value *string) CaseInsensitiveString {
	return CaseInsensitiveString{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}
