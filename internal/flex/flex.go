package flex

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/cidrtypes"
	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

type FrameworkElementFlExFunc[T any, U any] func(context.Context, T, *diag.Diagnostics) U

// ApplyToAll returns a new slice containing the results of applying the function `f` to each element of the original slice `s`.
func ApplyToAll[T, U any](s []T, f func(T) U) []U {
	v := make([]U, len(s))

	for i, e := range s {
		v[i] = f(e)
	}

	return v
}

// Expand Helpers (TF -> API)

func ExpandString(s types.String) string {
	if s.IsNull() || s.IsUnknown() {
		return ""
	}
	return s.ValueString()
}

func ExpandStringPointer(s types.String) *string {
	if s.IsNull() || s.IsUnknown() {
		return nil
	}
	return s.ValueStringPointer()
}

func ExpandStringPointerNullAsEmpty(s types.String) *string {
	if s.IsNull() || s.IsUnknown() {
		v := ""
		return &v
	}
	return s.ValueStringPointer()
}

func ExpandBool(b types.Bool) bool {
	if b.IsNull() || b.IsUnknown() {
		return false
	}
	return b.ValueBool()
}

func ExpandBoolPointer(b types.Bool) *bool {
	if b.IsNull() || b.IsUnknown() {
		return nil
	}
	return b.ValueBoolPointer()
}

func ExpandInt64(i types.Int64) int64 {
	if i.IsNull() || i.IsUnknown() {
		return 0
	}
	return i.ValueInt64()
}

func ExpandInt64Pointer(i types.Int64) *int64 {
	if i.IsNull() || i.IsUnknown() {
		return nil
	}
	return i.ValueInt64Pointer()
}

func ExpandInt32(i types.Int32) int32 {
	if i.IsNull() || i.IsUnknown() {
		return 0
	}
	return i.ValueInt32()
}

func ExpandInt32Pointer(i types.Int32) *int32 {
	if i.IsNull() || i.IsUnknown() {
		return nil
	}
	return i.ValueInt32Pointer()
}

// ExpandMapStringAny expands types.Map to map[string]any
// Returns nil for null, unknown, or empty maps.
func ExpandMapStringAny(ctx context.Context, m types.Map, diags *diag.Diagnostics) map[string]any {
	if m.IsNull() || m.IsUnknown() || len(m.Elements()) == 0 {
		return map[string]any{}
	}
	strMap := make(map[string]string)
	d := m.ElementsAs(ctx, &strMap, false)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}
	result := make(map[string]any, len(strMap))
	for k, v := range strMap {
		result[k] = v
	}
	return result
}

// ExpandObjectWithFn expands a types.Object (SingleNestedAttribute) into the target API type D.
// If obj is null/unknown, returns nil. Otherwise parses to TF model T, then calls expandFn to transform to API type D.
func ExpandObjectWithFn[T any, D any](ctx context.Context, obj types.Object, diags *diag.Diagnostics, expandFn func(context.Context, *T, *diag.Diagnostics) *D) *D {
	if obj.IsNull() || obj.IsUnknown() {
		return nil
	}
	var model T
	d := obj.As(ctx, &model, basetypes.ObjectAsOptions{})
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}
	return expandFn(ctx, &model, diags)
}

// ExpandNestedObject expands a types.Object (SingleNestedAttribute) directly into the target model T.
// Use this when no transformation is needed (TF model == target type).
// For transformation between different types, use ExpandObjectWithFn.
func ExpandNestedObject[T any](ctx context.Context, obj types.Object, diags *diag.Diagnostics) *T {
	if obj.IsNull() || obj.IsUnknown() {
		return nil
	}
	var model T
	d := obj.As(ctx, &model, basetypes.ObjectAsOptions{})
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}
	return &model
}

func ExpandTimePointer(_ context.Context, dt timetypes.RFC3339, diags *diag.Diagnostics) *time.Time {
	if dt.IsNull() || dt.IsUnknown() {
		return nil
	}
	t, d := dt.ValueRFC3339Time()
	diags.Append(d...)
	return &t
}

// Note: Terraform Framework uses types.Float64 for both float32 and float64 SDK types.
// The following helpers convert appropriately.

func ExpandFloat32(v types.Float64) float32 {
	if v.IsNull() || v.IsUnknown() {
		return 0
	}
	return float32(v.ValueFloat64())
}

func ExpandFloat32Pointer(v types.Float64) *float32 {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	f := float32(v.ValueFloat64())
	return &f
}

func ExpandFloat64(v types.Float64) float64 {
	return v.ValueFloat64()
}

func ExpandFloat64Pointer(v types.Float64) *float64 {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	return v.ValueFloat64Pointer()
}

// ExpandFrameworkListString expands types.List to []string
// Returns nil for null or unknown lists.
func ExpandFrameworkListString(ctx context.Context, tfList interface {
	basetypes.ListValuable
	ElementsAs(ctx context.Context, target any, allowUnhandled bool) diag.Diagnostics
}, diags *diag.Diagnostics) []string {
	if tfList.IsNull() || tfList.IsUnknown() {
		return nil
	}
	var data []string
	diags.Append(tfList.ElementsAs(ctx, &data, false)...)
	return data
}

// ExpandFrameworkListInt32 expands types.List to []int32
// Returns nil for null or unknown lists.
func ExpandFrameworkListInt32(ctx context.Context, tfList interface {
	basetypes.ListValuable
	ElementsAs(ctx context.Context, target any, allowUnhandled bool) diag.Diagnostics
}, diags *diag.Diagnostics) []int32 {
	if tfList.IsNull() || tfList.IsUnknown() {
		return nil
	}
	var data []int32
	diags.Append(tfList.ElementsAs(ctx, &data, false)...)
	return data
}

// ExpandFrameworkListInt64 expands types.List to []int64
// Returns nil for null or unknown lists.
func ExpandFrameworkListInt64(ctx context.Context, tfList interface {
	basetypes.ListValuable
	ElementsAs(ctx context.Context, target any, allowUnhandled bool) diag.Diagnostics
}, diags *diag.Diagnostics) []int64 {
	if tfList.IsNull() || tfList.IsUnknown() {
		return nil
	}
	var data []int64
	diags.Append(tfList.ElementsAs(ctx, &data, false)...)
	return data
}

func ExpandIPv4Address(ipv4addr iptypes.IPv4Address) *string {
	if ipv4addr.IsNull() || ipv4addr.IsUnknown() {
		return nil
	}
	v := ipv4addr.ValueString()
	return &v
}

func ExpandIPv6Address(ipv6addr iptypes.IPv6Address) *string {
	if ipv6addr.IsNull() || ipv6addr.IsUnknown() {
		return nil
	}
	v := ipv6addr.ValueString()
	return &v
}

func ExpandIPv4AddressValue(ipv4addr iptypes.IPv4Address) string {
	if ipv4addr.IsNull() || ipv4addr.IsUnknown() {
		return ""
	}
	return ipv4addr.ValueString()
}

func ExpandIPv6AddressValue(ipv6addr iptypes.IPv6Address) string {
	if ipv6addr.IsNull() || ipv6addr.IsUnknown() {
		return ""
	}
	return ipv6addr.ValueString()
}

func ExpandIPAddress(ipaddr iptypes.IPAddress) *string {
	if ipaddr.IsNull() || ipaddr.IsUnknown() {
		return nil
	}
	return ExpandStringPointer(ipaddr.StringValue)
}

func ExpandIPv4Prefix(ipv4addr cidrtypes.IPv4Prefix) *string {
	if ipv4addr.IsNull() || ipv4addr.IsUnknown() {
		return nil
	}
	return ExpandStringPointer(ipv4addr.StringValue)
}

func ExpandIPv6Prefix(ipv6addr cidrtypes.IPv6Prefix) *string {
	if ipv6addr.IsNull() || ipv6addr.IsUnknown() {
		return nil
	}
	return ExpandStringPointer(ipv6addr.StringValue)
}

func ExpandMACAddress(mac internaltypes.MACAddress) *string {
	if mac.IsNull() || mac.IsUnknown() {
		return nil
	}
	return ExpandStringPointer(mac.StringValue)
}

// ExpandTimeToUnix converts a naive datetime string (utils.NaiveDatetimeLayout,
// interpreted as UTC) into a Unix epoch seconds pointer for SDKs that only
// accept the wire value as an integer timestamp.
func ExpandTimeToUnix(v types.String, diags *diag.Diagnostics) *int64 {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	t, err := time.Parse(utils.NaiveDatetimeLayout, v.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid time value",
			fmt.Sprintf("Expected format: %s, got: %s (%s)", utils.NaiveDatetimeLayout, v.ValueString(), err),
		)
		return nil
	}
	unix := t.UTC().Unix()
	return &unix
}

func ExpandRFC3339(dt timetypes.RFC3339, diags *diag.Diagnostics) *time.Time {
	if dt.IsNull() || dt.IsUnknown() {
		return nil
	}
	t, d := dt.ValueRFC3339Time()
	diags.Append(d...)
	return &t
}

func ExpandFrameworkListNestedBlock[T any, U any](ctx context.Context, tfList interface {
	basetypes.ListValuable
	ElementsAs(ctx context.Context, target any, allowUnhandled bool) diag.Diagnostics
}, diags *diag.Diagnostics, f FrameworkElementFlExFunc[T, *U]) []U {
	if tfList.IsNull() || tfList.IsUnknown() {
		return make([]U, 0)
	}

	var data []T
	diags.Append(tfList.ElementsAs(ctx, &data, false)...)

	expanded := make([]U, 0, len(data))
	for _, t := range data {
		v := f(ctx, t, diags)
		if v == nil {
			// Skip unknown/null nested objects safely.
			continue
		}
		expanded = append(expanded, *v)
	}

	return expanded
}

// Flatten Helpers (API -> TF)

func FlattenString(s string) types.String {
	return types.StringValue(s)
}

func FlattenStringPointer(s *string) types.String {
	if s == nil {
		return types.StringNull()
	}
	return types.StringValue(*s)
}

func FlattenStringPointerEmptyAsNull(s *string) types.String {
	if s == nil || *s == "" {
		return types.StringNull()
	}
	return types.StringValue(*s)
}

func FlattenBool(b bool) types.Bool {
	return types.BoolValue(b)
}

func FlattenBoolPointer(b *bool) types.Bool {
	if b == nil {
		return types.BoolNull()
	}
	return types.BoolValue(*b)
}

func FlattenInt64(i int64) types.Int64 {
	return types.Int64Value(i)
}

func FlattenInt64Pointer(i *int64) types.Int64 {
	if i == nil {
		return types.Int64Null()
	}
	return types.Int64Value(*i)
}

func FlattenInt64PointerZeroAsNull(i *int64) types.Int64 {
	if i == nil || *i == 0 {
		return types.Int64Null()
	}
	return types.Int64Value(*i)
}

func FlattenInt32(i int32) types.Int32 {
	return types.Int32Value(i)
}

func FlattenInt32Pointer(i *int32) types.Int32 {
	if i == nil {
		return types.Int32Null()
	}
	return types.Int32Value(*i)
}

func FlattenFloat32(f float32) types.Float64 {
	return types.Float64Value(float64(f))
}

func FlattenFloat32Pointer(f *float32) types.Float64 {
	if f == nil {
		return types.Float64Null()
	}
	return types.Float64Value(float64(*f))
}

func FlattenFloat32PointerZeroAsNull(f *float32) types.Float64 {
	if f == nil || *f == 0 {
		return types.Float64Null()
	}
	return types.Float64Value(float64(*f))
}

func FlattenFloat64(f float64) types.Float64 {
	if f == 0 {
		return types.Float64Null()
	}
	return types.Float64Value(f)
}

func FlattenFloat64Pointer(f *float64) types.Float64 {
	if f == nil {
		return types.Float64Null()
	}
	return FlattenFloat64(*f)
}

// FlattenMapStringAny flattens map[string]any to types.Map with string values
// TODO: consider supporting other types in the future if needed apart from map[string]string
func FlattenMapStringAny(ctx context.Context, m map[string]any, diags *diag.Diagnostics) types.Map {
	if len(m) == 0 {
		return types.MapNull(types.StringType)
	}
	strMap := make(map[string]string, len(m))
	for k, v := range m {
		strMap[k] = fmt.Sprintf("%v", v)
	}
	mapVal, d := types.MapValueFrom(ctx, types.StringType, strMap)
	diags.Append(d...)
	return mapVal
}

// FlattenObjectWithFn flattens an API type to types.Object (SingleNestedAttribute) using a transform function.
// If src is nil, returns ObjectNull. Otherwise calls flattenFn to build the TF model.
func FlattenObjectWithFn[S any](ctx context.Context, src *S, attrTypes map[string]attr.Type, diags *diag.Diagnostics, flattenFn func(context.Context, *S, *diag.Diagnostics) any) types.Object {
	if src == nil {
		return types.ObjectNull(attrTypes)
	}
	model := flattenFn(ctx, src, diags)
	obj, d := types.ObjectValueFrom(ctx, attrTypes, model)
	diags.Append(d...)
	return obj
}

// FlattenNestedObject flattens a model struct directly to types.Object (SingleNestedAttribute).
// Use this when no transformation is needed (source type == TF model).
// For transformation between different types, use FlattenObjectWithFn.
func FlattenNestedObject[T any](ctx context.Context, model *T, attrTypes map[string]attr.Type, diags *diag.Diagnostics) types.Object {
	if model == nil {
		return types.ObjectNull(attrTypes)
	}
	obj, d := types.ObjectValueFrom(ctx, attrTypes, model)
	diags.Append(d...)
	return obj
}

// ExpandMapString expands types.Map to map[string]string
// Returns empty map for null or unknown maps.
func ExpandMapString(ctx context.Context, m types.Map, diags *diag.Diagnostics) map[string]string {
	if m.IsNull() || m.IsUnknown() {
		return map[string]string{}
	}
	strMap := make(map[string]string)
	d := m.ElementsAs(ctx, &strMap, false)
	diags.Append(d...)
	if diags.HasError() {
		return map[string]string{}
	}
	return strMap
}

// FlattenFrameworkListString flattens []string to types.List
// Returns null list if input is nil or empty.
func FlattenFrameworkListString(ctx context.Context, l []string, diags *diag.Diagnostics) types.List {
	if len(l) == 0 {
		return types.ListNull(types.StringType)
	}
	tfList, d := types.ListValueFrom(ctx, types.StringType, l)
	diags.Append(d...)
	return tfList
}

// FlattenFrameworkListInt32 flattens []int32 to types.List
// Returns null list if input is nil or empty.
func FlattenFrameworkListInt32(ctx context.Context, l []int32, diags *diag.Diagnostics) types.List {
	if len(l) == 0 {
		return types.ListNull(types.Int32Type)
	}
	tfList, d := types.ListValueFrom(ctx, types.Int32Type, l)
	diags.Append(d...)
	return tfList
}

// FlattenFrameworkListInt64 flattens []int64 to types.List
// Returns null list if input is nil or empty.
func FlattenFrameworkListInt64(ctx context.Context, l []int64, diags *diag.Diagnostics) types.List {
	if len(l) == 0 {
		return types.ListNull(types.Int64Type)
	}
	tfList, d := types.ListValueFrom(ctx, types.Int64Type, l)
	diags.Append(d...)
	return tfList
}

func FlattenIPv4Address(ipv4addr *string) iptypes.IPv4Address {
	if ipv4addr == nil || *ipv4addr == "" {
		return iptypes.NewIPv4AddressNull()
	}
	return iptypes.NewIPv4AddressValue(*ipv4addr)
}

func FlattenIPv6Address(ipv6addr *string) iptypes.IPv6Address {
	if ipv6addr == nil || *ipv6addr == "" {
		return iptypes.NewIPv6AddressNull()
	}
	return iptypes.NewIPv6AddressValue(*ipv6addr)
}

func FlattenIPv4AddressValue(ipv4addr string) iptypes.IPv4Address {
	if ipv4addr == "" {
		return iptypes.NewIPv4AddressNull()
	}
	return iptypes.NewIPv4AddressValue(ipv4addr)
}

func FlattenIPv6AddressValue(ipv6addr string) iptypes.IPv6Address {
	if ipv6addr == "" {
		return iptypes.NewIPv6AddressNull()
	}
	return iptypes.NewIPv6AddressValue(ipv6addr)
}

func FlattenIPAddress(ipaddr *string) iptypes.IPAddress {
	if ipaddr == nil || *ipaddr == "" {
		return iptypes.NewIPAddressNull()
	}
	return iptypes.IPAddress{
		StringValue: FlattenStringPointer(ipaddr),
	}
}

func FlattenIPv4Prefix(ipv4addr *string) cidrtypes.IPv4Prefix {
	if ipv4addr == nil || *ipv4addr == "" {
		return cidrtypes.NewIPv4PrefixNull()
	}
	return cidrtypes.IPv4Prefix{
		StringValue: FlattenStringPointer(ipv4addr),
	}
}

func FlattenIPv6Prefix(ipv6addr *string) cidrtypes.IPv6Prefix {
	if ipv6addr == nil || *ipv6addr == "" {
		return cidrtypes.NewIPv6PrefixNull()
	}
	return cidrtypes.IPv6Prefix{
		StringValue: FlattenStringPointer(ipv6addr),
	}
}

func FlattenMACAddress(mac *string) internaltypes.MACAddress {
	if mac == nil {
		return internaltypes.NewMACAddressNull()
	}
	return internaltypes.MACAddress{
		StringValue: FlattenStringPointer(mac),
	}
}

// FlattenUnixTime converts a Unix epoch seconds pointer back into a naive
// datetime string (utils.NaiveDatetimeLayout, UTC), the inverse of ExpandTimeToUnix.
func FlattenUnixTime(v *int64, diags *diag.Diagnostics) types.String {
	if v == nil {
		return types.StringNull()
	}
	return types.StringValue(time.Unix(*v, 0).UTC().Format(utils.NaiveDatetimeLayout))
}

func FlattenRFC3339(t *time.Time) timetypes.RFC3339 {
	if t == nil || t.IsZero() {
		return timetypes.NewRFC3339Null()
	}
	return timetypes.NewRFC3339TimeValue(*t)
}

func FlattenFrameworkUnorderedListString(ctx context.Context, data []string, diags *diag.Diagnostics) internaltypes.UnorderedListValue {
	if len(data) == 0 {
		return internaltypes.NewUnorderedListValueNull(types.StringType)
	}
	tfList, d := internaltypes.NewUnorderedListValueFrom(ctx, types.StringType, data)
	diags.Append(d...)
	return tfList
}

func FlattenFrameworkUnorderedListNestedBlock[T any, U any](ctx context.Context, data []T, attrTypes map[string]attr.Type, diags *diag.Diagnostics, f FrameworkElementFlExFunc[*T, U]) internaltypes.UnorderedListValue {
	if len(data) == 0 {
		return internaltypes.NewUnorderedListValueNull(types.ObjectType{AttrTypes: attrTypes})
	}

	tfData := ApplyToAll(data, func(t T) U {
		return f(ctx, &t, diags)
	})

	tfList, d := internaltypes.NewUnorderedListValueFrom(ctx, types.ObjectType{AttrTypes: attrTypes}, tfData)

	diags.Append(d...)
	return tfList
}

func FlattenFrameworkListNestedBlock[T any, U any](ctx context.Context, data []T, attrTypes map[string]attr.Type, diags *diag.Diagnostics, f FrameworkElementFlExFunc[*T, U]) types.List {
	if len(data) == 0 {
		return types.ListNull(types.ObjectType{AttrTypes: attrTypes})
	}

	tfData := ApplyToAll(data, func(t T) U {
		return f(ctx, &t, diags)
	})

	tfList, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: attrTypes}, tfData)

	diags.Append(d...)
	return tfList
}

// RDataStringPtr coerces an untyped map value to *string
func RDataStringPtr(v any) *string {
	if s, ok := v.(string); ok && s != "" {
		return &s
	}
	return nil
}

// RDataBoolPtr coerces an untyped map value to *bool.
func RDataBoolPtr(v any) *bool {
	switch t := v.(type) {
	case bool:
		return &t
	case string:
		if b, err := strconv.ParseBool(t); err == nil {
			return &b
		}
	}
	return nil
}

// RDataInt64Ptr coerces an untyped map value to *int64. rdata is a decoded JSON
// object, so a numeric subfield is a float64 rather than an integer type.
func RDataInt64Ptr(v any) *int64 {
	switch t := v.(type) {
	case float64:
		i := int64(t)
		return &i
	case string:
		if i, err := strconv.ParseInt(t, 10, 64); err == nil {
			return &i
		}
	}
	return nil
}
func ExpandExtensibleAttributeDefDefaultValue(ctx context.Context, defaultValue types.String, eaType types.String, diags *diag.Diagnostics) *grid.ExtensibleattributedefDefaultValue {
	if defaultValue.IsNull() || defaultValue.IsUnknown() {
		return nil
	}

	value := defaultValue.ValueString()
	if value == "" {
		return nil
	}

	// Check the type to determine if we should send as integer or string
	if !eaType.IsNull() && !eaType.IsUnknown() && eaType.ValueString() == "INTEGER" {
		// Convert string to integer for INTEGER type
		if intVal, err := strconv.ParseInt(value, 10, 32); err == nil {
			int32Val := int32(intVal)
			return &grid.ExtensibleattributedefDefaultValue{
				Int32: &int32Val,
			}
		} else {
			diags.AddError(
				"Invalid Integer Default Value",
				fmt.Sprintf("Cannot convert default_value '%s' to integer: %v", value, err),
			)
			return nil
		}
	}

	// For all other types (STRING, EMAIL, URL, DATE, ENUM), send as string
	return &grid.ExtensibleattributedefDefaultValue{
		String: &value,
	}
}

func FlattenExtensibleAttributeDefDefaultValue(ctx context.Context, from *grid.ExtensibleattributedefDefaultValue, diags *diag.Diagnostics) types.String {
	if from == nil {
		return types.StringNull()
	}

	if from.Int32 != nil {
		// Convert int32 to string for Terraform
		return types.StringValue(strconv.FormatInt(int64(*from.Int32), 10))
	}

	// Check if string value is set
	if from.String != nil {
		return types.StringValue(*from.String)
	}

	// No value set
	return types.StringNull()
}
