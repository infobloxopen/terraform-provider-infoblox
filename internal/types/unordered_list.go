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

var UnorderedListOfStringType = UnorderedList{basetypes.ListType{ElemType: basetypes.StringType{}}}

var _ basetypes.ListValuableWithSemanticEquals = UnorderedListValue{}
var _ basetypes.ListTypable = UnorderedList{}

// UnorderedList is a type representing a list of values where the order of the elements is not significant.
type UnorderedList struct {
	basetypes.ListType
}

func (t UnorderedList) Equal(o attr.Type) bool {
	other, ok := o.(UnorderedList)

	if !ok {
		return false
	}

	return t.ListType.Equal(other.ListType)
}

func (UnorderedList) String() string {
	return "UnorderedList"
}

func (t UnorderedList) ValueFromList(ctx context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NewUnorderedListValueNull(t.ElemType), diags
	}

	if in.IsUnknown() {
		return NewUnorderedListValueUnknown(t.ElemType), diags
	}

	v, d := basetypes.NewListValue(t.ElemType, in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return NewUnorderedListValueUnknown(t.ElemType), diags
	}

	return UnorderedListValue{ListValue: v}, diags
}

func (t UnorderedList) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ListType.ValueFromTerraform(ctx, in)

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

func (t UnorderedList) ValueType(ctx context.Context) attr.Value {
	return UnorderedListValue{
		ListValue: t.ListType.ValueType(ctx).(basetypes.ListValue),
	}
}

func NewUnorderedList(elemType attr.Type) UnorderedList {
	return UnorderedList{basetypes.ListType{ElemType: elemType}}
}

type UnorderedListValue struct {
	basetypes.ListValue
}

func (v UnorderedListValue) ListSemanticEquals(ctx context.Context, newValuable basetypes.ListValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(UnorderedListValue)
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

func (v UnorderedListValue) Equal(o attr.Value) bool {
	other, ok := o.(UnorderedListValue)

	if !ok {
		return false
	}

	return v.ListValue.Equal(other.ListValue)
}

func (v UnorderedListValue) Type(ctx context.Context) attr.Type {
	return NewUnorderedList(v.ElementType(ctx))
}

func NewUnorderedListValueNull(elemType attr.Type) UnorderedListValue {
	return UnorderedListValue{ListValue: basetypes.NewListNull(elemType)}
}

func NewUnorderedListValueUnknown(elemType attr.Type) UnorderedListValue {
	return UnorderedListValue{ListValue: basetypes.NewListUnknown(elemType)}
}
func NewUnorderedListValue(elemType attr.Type, elements []attr.Value) (UnorderedListValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	v, d := basetypes.NewListValue(elemType, elements)
	diags.Append(d...)
	if diags.HasError() {
		return NewUnorderedListValueUnknown(elemType), diags
	}

	return UnorderedListValue{ListValue: v}, diags
}

func NewUnorderedListValueFrom(ctx context.Context, elemType attr.Type, elements any) (UnorderedListValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	v, d := basetypes.NewListValueFrom(ctx, elemType, elements)
	diags.Append(d...)
	if diags.HasError() {
		return NewUnorderedListValueUnknown(elemType), diags
	}

	return UnorderedListValue{ListValue: v}, diags
}
