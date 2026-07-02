// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timetypes

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ basetypes.StringValuableWithSemanticEquals = (*GoDuration)(nil)
	_ xattr.ValidateableAttribute                = (*GoDuration)(nil)
	_ function.ValidateableParameter             = (*GoDuration)(nil)
)

// GoDuration represents a valid Go time duration string.
// See https://pkg.go.dev/time#ParseDuration for more details
type GoDuration struct {
	basetypes.StringValue
}

// Type returns an RFC3339Type.
func (d GoDuration) Type(_ context.Context) attr.Type {
	return GoDurationType{}
}

// Equal returns true if the given value is equivalent.
func (d GoDuration) Equal(o attr.Value) bool {
	other, ok := o.(GoDuration)

	if !ok {
		return false
	}

	return d.StringValue.Equal(other.StringValue)
}

// ValidateAttribute implements attribute value validation. This type requires the value to be a String value that
// is valid Go time duration and utilizes the Go `time` library
func (d GoDuration) ValidateAttribute(ctx context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	if d.IsUnknown() || d.IsNull() {
		return
	}

	if _, err := time.ParseDuration(d.ValueString()); err != nil {
		resp.Diagnostics.Append(diag.WithPath(req.Path, goDurationInvalidStringDiagnostic(d.ValueString(), err)))

		return
	}
}

// ValidateParameter implements provider-defined function parameter value validation. This type requires the value to
// be a String value that is a valid time duration and utilizes the Go `time` library
func (d GoDuration) ValidateParameter(ctx context.Context, req function.ValidateParameterRequest, resp *function.ValidateParameterResponse) {
	if d.IsUnknown() || d.IsNull() {
		return
	}

	if _, err := time.ParseDuration(d.ValueString()); err != nil {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid Go Time Duration String Value: "+
				"A string value was provided that is not a valid Go Time Duration string format. "+
				`A duration string is a sequence of numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". `+
				`Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".\n\n`+
				"Given Value: "+d.ValueString()+"\n"+
				"Error: "+err.Error(),
		)

		return
	}
}

// ValueGoDuration creates a new time.Duration instance with the time duration StringValue. A null or unknown value will produce an error diagnostic.
func (d GoDuration) ValueGoDuration() (time.Duration, diag.Diagnostics) {
	var diags diag.Diagnostics

	if d.IsNull() {
		diags.Append(diag.NewErrorDiagnostic("Go Duration ValueDuration Error", "Duration string value is null"))
		return time.Duration(0), diags
	}

	if d.IsUnknown() {
		diags.Append(diag.NewErrorDiagnostic("Go Duration ValueDuration Error", "Duration string value is unknown"))
		return time.Duration(0), diags
	}

	duration, err := time.ParseDuration(d.ValueString())
	if err != nil {
		diags.Append(diag.NewErrorDiagnostic("Go Duration ValueDuration Error", err.Error()))
		return time.Duration(0), diags
	}

	return duration, nil
}

// StringSemanticEquals returns true if the given GoDuration string value is semantically equal to the current GoDuration string value.
// It ensures that two duration values are semantically equal even if their string representations are different.
func (d GoDuration) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(GoDuration)
	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", d)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)

		return false, diags
	}

	priorDuration, _ := d.ValueGoDuration()
	newDuration, _ := newValue.ValueGoDuration()

	return priorDuration == newDuration, diags
}

// NewGoDurationNull creates an Duration with a null value. Determine whether the value is null via IsNull method.
func NewGoDurationNull() GoDuration {
	return GoDuration{
		StringValue: basetypes.NewStringNull(),
	}
}

// NewGoDurationUnknown creates an Duration with an unknown value. Determine whether the value is unknown via IsUnknown method.
func NewGoDurationUnknown() GoDuration {
	return GoDuration{
		StringValue: basetypes.NewStringUnknown(),
	}
}

// NewGoDurationValue creates an Duration with a known value.
func NewGoDurationValue(value time.Duration) GoDuration {
	return GoDuration{
		StringValue: basetypes.NewStringValue(value.String()),
	}
}

// NewGoDurationPointerValue creates an Duration with a null value if nil or
// a known value.
func NewGoDurationPointerValue(value *time.Duration) GoDuration {
	if value == nil {
		return NewGoDurationNull()
	}

	return GoDuration{
		StringValue: basetypes.NewStringValue(value.String()),
	}
}

// NewGoDurationValueFromString creates an Duration with a known value or raises an error
// diagnostic if the string is not Duration format.
func NewGoDurationValueFromString(value string) (GoDuration, diag.Diagnostics) {
	_, err := time.ParseDuration(value)

	if err != nil {
		// Returning an unknown value will guarantee that, as a last resort,
		// Terraform will return an error if attempting to store into state.
		return NewGoDurationUnknown(), diag.Diagnostics{goDurationInvalidStringDiagnostic(value, err)}
	}

	return GoDuration{
		StringValue: basetypes.NewStringValue(value),
	}, nil
}

// NewGoDurationValueFromStringMust creates an Duration with a known value or raises a panic
// if the string is not Duration format.
//
// This creation function is only recommended to create Duration values which
// either will not potentially affect practitioners, such as testing, or within
// exhaustively tested provider logic.
func NewGoDurationValueFromStringMust(value string) GoDuration {
	_, err := time.ParseDuration(value)

	if err != nil {
		panic(fmt.Sprintf("Invalid Go Duration String Value (%s): %s", value, err))
	}

	return GoDuration{
		StringValue: basetypes.NewStringValue(value),
	}
}

// NewGoDurationValueFromPointerString creates an Duration with a null value if nil, a known
// value, or raises an error diagnostic if the string is not Duration format.
func NewGoDurationValueFromPointerString(value *string) (GoDuration, diag.Diagnostics) {
	if value == nil {
		return NewGoDurationNull(), nil
	}

	return NewGoDurationValueFromString(*value)
}

// NewGoDurationValueFromPointerStringMust creates an Duration with a null value if nil, a
// known value, or raises a panic if the string is not Duration format.
//
// This creation function is only recommended to create Duration values which
// either will not potentially affect practitioners, such as testing, or within
// exhaustively tested provider logic.
func NewGoDurationValueFromPointerStringMust(value *string) GoDuration {
	if value == nil {
		return NewGoDurationNull()
	}

	return NewGoDurationValueFromStringMust(*value)
}
