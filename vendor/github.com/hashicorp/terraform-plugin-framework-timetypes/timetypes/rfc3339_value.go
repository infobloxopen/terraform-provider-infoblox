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
	_ basetypes.StringValuable       = (*RFC3339)(nil)
	_ xattr.ValidateableAttribute    = (*RFC3339)(nil)
	_ function.ValidateableParameter = (*RFC3339)(nil)
)

// RFC3339 represents a valid RFC3339-formatted string. Semantic equality logic is defined for RFC3339
// such that inconsequential differences between the `Z` suffix and a `00:00` UTC offset are ignored.
type RFC3339 struct {
	basetypes.StringValue
}

// Type returns an RFC3339Type.
func (v RFC3339) Type(_ context.Context) attr.Type {
	return RFC3339Type{}
}

// Equal returns true if the given value is equivalent.
func (v RFC3339) Equal(o attr.Value) bool {
	other, ok := o.(RFC3339)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

// StringSemanticEquals returns true if the given RFC3339 string value is semantically equal to the current RFC3339 string value.
// This comparison utilizes time.Parse to create time.Time instances and then compares the resulting RFC 3339-formatted string representations.
// This means RFC3339 values utilizing the `Z` Zulu suffix as an offset are considered semantically equal to RFC3339 that define a `00:00` UTC offset.
//
// Examples:
//   - `2023-07-25T20:43:16+00:00` is semantically equal to `2023-07-25T20:43:16Z`
//   - `2023-07-25T20:43:16-00:00` is semantically equal to `2023-07-25T20:43:16Z` - while RFC 3339 defines an unknown local
//     offset (`-00:00`) to be different from an offset of `Z`, time.Parse converts `-00:00` to `+00:00` during parsing.
//
// Counterexamples:
//   - `2023-07-25T23:43:16+00:00` expresses the same time as `2023-07-25T20:43:16-03:00` but is NOT considered
//     to be semantically equal.
//
// See RFC 3339 for more details on the string format: https://www.rfc-editor.org/rfc/rfc3339.html.
func (v RFC3339) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(RFC3339)
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

	// RFC3339 strings are already validated at this point, ignoring errors
	newRFC3339time, _ := time.Parse(time.RFC3339, newValue.ValueString())
	currentRFC3339time, _ := time.Parse(time.RFC3339, v.ValueString())

	return currentRFC3339time.Format(time.RFC3339) == newRFC3339time.Format(time.RFC3339), diags
}

// ValidateAttribute implements attribute value validation. This type requires the value to be a String value that
// is valid RFC 3339 format. This utilizes the Go `time` library which does not strictly adhere to the RFC 3339
// standard and may allow strings that are not valid RFC 3339 strings
//
// See https://github.com/golang/go/issues/54580 for more info on the Go `time` library's RFC 3339 parsing differences.
func (v RFC3339) ValidateAttribute(ctx context.Context, req xattr.ValidateAttributeRequest, resp *xattr.ValidateAttributeResponse) {
	if v.IsUnknown() || v.IsNull() {
		return
	}

	if _, err := time.Parse(time.RFC3339, v.ValueString()); err != nil {
		resp.Diagnostics.Append(diag.WithPath(req.Path, rfc3339InvalidStringDiagnostic(v.ValueString(), err)))

		return
	}
}

// ValidateParameter implements provider-defined function parameter value validation. This type requires the value to
// be a String value that is valid RFC 3339 format. This utilizes the Go `time` library which does not strictly
// adhere to the RFC 3339 standard and may allow strings that are not valid RFC 3339 strings
//
// See https://github.com/golang/go/issues/54580 for more info on the Go `time` library's RFC 3339 parsing differences.
func (v RFC3339) ValidateParameter(ctx context.Context, req function.ValidateParameterRequest, resp *function.ValidateParameterResponse) {
	if v.IsUnknown() || v.IsNull() {
		return
	}

	if _, err := time.Parse(time.RFC3339, v.ValueString()); err != nil {
		resp.Error = function.NewArgumentFuncError(
			req.Position,
			"Invalid RFC3339 String Value: "+
				"A string value was provided that is not valid RFC3339 string format.\n\n"+
				"Given Value: "+v.ValueString()+"\n"+
				"Error: "+err.Error(),
		)

		return
	}
}

// ValueRFC3339Time creates a new time.Time instance with the RFC3339 StringValue. A null or unknown value will produce an error diagnostic.
func (v RFC3339) ValueRFC3339Time() (time.Time, diag.Diagnostics) {
	var diags diag.Diagnostics

	if v.IsNull() {
		diags.Append(diag.NewErrorDiagnostic("RFC3339 ValueRFC3339Time Error", "RFC3339 string value is null"))
		return time.Time{}, diags
	}

	if v.IsUnknown() {
		diags.Append(diag.NewErrorDiagnostic("RFC3339 ValueRFC3339Time Error", "RFC3339 string value is unknown"))
		return time.Time{}, diags
	}

	rfc3339Time, err := time.Parse(time.RFC3339, v.ValueString())
	if err != nil {
		diags.Append(diag.NewErrorDiagnostic("RFC3339 ValueRFC3339Time Error", err.Error()))
		return time.Time{}, diags
	}

	return rfc3339Time, nil
}

// NewRFC3339Null creates an RFC3339 with a null value. Determine whether the value is null via IsNull method.
func NewRFC3339Null() RFC3339 {
	return RFC3339{
		StringValue: basetypes.NewStringNull(),
	}
}

// NewRFC3339Unknown creates an RFC3339 with an unknown value. Determine whether the value is unknown via IsUnknown method.
func NewRFC3339Unknown() RFC3339 {
	return RFC3339{
		StringValue: basetypes.NewStringUnknown(),
	}
}

// NewRFC3339TimeValue creates an RFC3339 with a known value.
func NewRFC3339TimeValue(value time.Time) RFC3339 {
	return RFC3339{
		StringValue: basetypes.NewStringValue(value.Format(time.RFC3339)),
	}
}

// NewRFC3339TimePointerValue creates an RFC3339 with a null value if nil or
// a known value.
func NewRFC3339TimePointerValue(value *time.Time) RFC3339 {
	if value == nil {
		return NewRFC3339Null()
	}

	return RFC3339{
		StringValue: basetypes.NewStringValue(value.Format(time.RFC3339)),
	}
}

// NewRFC3339Value creates an RFC3339 with a known value or raises an error
// diagnostic if the string is not RFC3339 format.
func NewRFC3339Value(value string) (RFC3339, diag.Diagnostics) {
	_, err := time.Parse(time.RFC3339, value)

	if err != nil {
		// Returning an unknown value will guarantee that, as a last resort,
		// Terraform will return an error if attempting to store into state.
		return NewRFC3339Unknown(), diag.Diagnostics{rfc3339InvalidStringDiagnostic(value, err)}
	}

	return RFC3339{
		StringValue: basetypes.NewStringValue(value),
	}, nil
}

// NewRFC3339ValueMust creates an RFC3339 with a known value or raises a panic
// if the string is not RFC3339 format.
//
// This creation function is only recommended to create RFC3339 values which
// either will not potentially affect practitioners, such as testing, or within
// exhaustively tested provider logic.
func NewRFC3339ValueMust(value string) RFC3339 {
	_, err := time.Parse(time.RFC3339, value)

	if err != nil {
		panic(fmt.Sprintf("Invalid RFC3339 String Value (%s): %s", value, err))
	}

	return RFC3339{
		StringValue: basetypes.NewStringValue(value),
	}
}

// NewRFC3339PointerValue creates an RFC3339 with a null value if nil, a known
// value, or raises an error diagnostic if the string is not RFC3339 format.
func NewRFC3339PointerValue(value *string) (RFC3339, diag.Diagnostics) {
	if value == nil {
		return NewRFC3339Null(), nil
	}

	return NewRFC3339Value(*value)
}

// NewRFC3339PointerValueMust creates an RFC3339 with a null value if nil, a
// known value, or raises a panic if the string is not RFC3339 format.
//
// This creation function is only recommended to create RFC3339 values which
// either will not potentially affect practitioners, such as testing, or within
// exhaustively tested provider logic.
func NewRFC3339PointerValueMust(value *string) RFC3339 {
	if value == nil {
		return NewRFC3339Null()
	}

	return NewRFC3339ValueMust(*value)
}
