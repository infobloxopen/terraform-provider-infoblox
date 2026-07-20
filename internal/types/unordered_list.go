package types

import (
	"context"
	"fmt"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.ListTypable = (*UnorderedListStringType)(nil)
)

// UnorderedListStringType is a custom list type whose elements are strings and
// whose element order is not significant. The zero value (UnorderedListStringType{})
// is fully usable because the element type is fixed to basetypes.StringType.
type UnorderedListStringType struct {
	basetypes.ListType
}

// String returns a human-friendly name for the type.
func (t UnorderedListStringType) String() string {
	return "types.UnorderedListStringType"
}

// ElementType returns the fixed element type (string).
func (t UnorderedListStringType) ElementType() attr.Type {
	return basetypes.StringType{}
}

// WithElementType returns the type unchanged; the element type is fixed to string.
func (t UnorderedListStringType) WithElementType(attr.Type) attr.TypeWithElementType {
	return t
}

// TerraformType returns the tftypes.Type for this type.
func (t UnorderedListStringType) TerraformType(ctx context.Context) tftypes.Type {
	return basetypes.ListType{ElemType: basetypes.StringType{}}.TerraformType(ctx)
}

// Equal returns true if the given type is also an UnorderedListStringType.
func (t UnorderedListStringType) Equal(o attr.Type) bool {
	_, ok := o.(UnorderedListStringType)
	return ok
}

// ApplyTerraform5AttributePathStep applies the given path step.
func (t UnorderedListStringType) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	return basetypes.ListType{ElemType: basetypes.StringType{}}.ApplyTerraform5AttributePathStep(step)
}

// ValueType returns the Value type.
func (t UnorderedListStringType) ValueType(_ context.Context) attr.Value {
	return UnorderedListString{
		ListValue: basetypes.NewListNull(basetypes.StringType{}),
	}
}

// ValueFromList converts a ListValue to an UnorderedListString.
func (t UnorderedListStringType) ValueFromList(_ context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NewUnorderedListStringValueNull(), diags
	}

	if in.IsUnknown() {
		return NewUnorderedListStringValueUnknown(), diags
	}

	v, d := basetypes.NewListValue(basetypes.StringType{}, in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return NewUnorderedListStringValueUnknown(), diags
	}

	return UnorderedListString{ListValue: v}, diags
}

// ValueFromTerraform converts a tftypes.Value to an attr.Value.
func (t UnorderedListStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := basetypes.ListType{ElemType: basetypes.StringType{}}.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	listValue, ok := attrValue.(basetypes.ListValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	listValuable, diags := t.ValueFromList(ctx, listValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
	}

	return listValuable, nil
}

var (
	_ basetypes.ListValuable                   = (*UnorderedListString)(nil)
	_ basetypes.ListValuableWithSemanticEquals = (*UnorderedListString)(nil)
)

// UnorderedListString is the value type for UnorderedListStringType.
type UnorderedListString struct {
	basetypes.ListValue
}

// Type returns an UnorderedListStringType.
func (v UnorderedListString) Type(_ context.Context) attr.Type {
	return UnorderedListStringType{}
}

// Equal returns true if the given value is an equivalent UnorderedListString.
func (v UnorderedListString) Equal(o attr.Value) bool {
	other, ok := o.(UnorderedListString)
	if !ok {
		return false
	}

	return v.ListValue.Equal(other.ListValue)
}

// ListSemanticEquals returns true if both lists contain the same elements,
// regardless of their order.
func (v UnorderedListString) ListSemanticEquals(ctx context.Context, newValuable basetypes.ListValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(UnorderedListString)
	if !ok {
		return false, diags
	}

	o, d := v.ToListValue(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return false, diags
	}

	n, d := newValue.ToListValue(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return false, diags
	}

	oldElems, newElems := o.Elements(), n.Elements()

	if len(oldElems) != len(newElems) {
		return false, diags
	}

	for _, newElem := range newElems {
		found := false
		for i, oldElem := range oldElems {
			if oldElem.Equal(newElem) {
				oldElems = slices.Delete(oldElems, i, i+1)
				found = true
				break
			}
		}
		if !found {
			return false, diags
		}
	}

	return len(oldElems) == 0, diags
}

// NewUnorderedListStringValueNull creates a null UnorderedListString.
func NewUnorderedListStringValueNull() UnorderedListString {
	return UnorderedListString{ListValue: basetypes.NewListNull(basetypes.StringType{})}
}

// NewUnorderedListStringValueUnknown creates an unknown UnorderedListString.
func NewUnorderedListStringValueUnknown() UnorderedListString {
	return UnorderedListString{ListValue: basetypes.NewListUnknown(basetypes.StringType{})}
}

// NewUnorderedListStringValue creates an UnorderedListString from string-typed elements.
func NewUnorderedListStringValue(elements []attr.Value) (UnorderedListString, diag.Diagnostics) {
	var diags diag.Diagnostics

	v, d := basetypes.NewListValue(basetypes.StringType{}, elements)
	diags.Append(d...)
	if diags.HasError() {
		return NewUnorderedListStringValueUnknown(), diags
	}

	return UnorderedListString{ListValue: v}, diags
}

// NewUnorderedListStringValueFrom creates an UnorderedListString from a Go value (e.g. []string).
func NewUnorderedListStringValueFrom(ctx context.Context, elements any) (UnorderedListString, diag.Diagnostics) {
	var diags diag.Diagnostics

	v, d := basetypes.NewListValueFrom(ctx, basetypes.StringType{}, elements)
	diags.Append(d...)
	if diags.HasError() {
		return NewUnorderedListStringValueUnknown(), diags
	}

	return UnorderedListString{ListValue: v}, diags
}
