package types

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ basetypes.StringValuableWithSemanticEquals = DUIDValue{}
var _ basetypes.StringTypable = DUIDType{}

// DUIDType is a custom type for DHCP Unique Identifiers with semantic equality
type DUIDType struct {
	basetypes.StringType
}

func (t DUIDType) String() string {
	return "DUIDType"
}

func (t DUIDType) ValueType(ctx context.Context) attr.Value {
	return DUIDValue{}
}

func (t DUIDType) Equal(o attr.Type) bool {
	other, ok := o.(DUIDType)
	if !ok {
		return false
	}
	return t.StringType.Equal(other.StringType)
}

func (t DUIDType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return DUIDValue{
		StringValue: in,
	}, nil
}

func (t DUIDType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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

// DUIDValue is the value type with semantic equality
type DUIDValue struct {
	basetypes.StringValue
}

func (v DUIDValue) Type(ctx context.Context) attr.Type {
	return DUIDType{}
}

func (v DUIDValue) Equal(o attr.Value) bool {
	other, ok := o.(DUIDValue)
	if !ok {
		return false
	}
	return v.StringValue.Equal(other.StringValue)
}

func (v DUIDValue) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(DUIDValue)
	if !ok {
		return false, diags
	}

	if v.IsNull() && newValue.IsNull() {
		return true, diags
	}
	if v.IsNull() || newValue.IsNull() {
		return false, diags
	}
	if v.IsUnknown() && newValue.IsUnknown() {
		return true, diags
	}
	if v.IsUnknown() || newValue.IsUnknown() {
		return false, diags
	}

	currentDUID := v.ValueString()
	newDUID := newValue.ValueString()

	if currentDUID == "" && newDUID == "" {
		return true, diags
	}
	if currentDUID == "" || newDUID == "" {
		return false, diags
	}

	// Normalize both values
	currentNormalized, err1 := normalizeDUID(currentDUID)
	newNormalized, err2 := normalizeDUID(newDUID)

	if err1 != nil || err2 != nil {
		return false, diags
	}

	// Compare normalized values
	return currentNormalized == newNormalized, diags
}

func NewDUIDValue(value string) DUIDValue {
	return DUIDValue{StringValue: basetypes.NewStringValue(value)}
}

func NewDUIDNull() DUIDValue {
	return DUIDValue{
		StringValue: basetypes.NewStringNull(),
	}
}

// normalizeDUID normalizes DUID to standard lowercase colon format
func normalizeDUID(address string) (string, error) {
	if address == "" {
		return "", nil
	}

	address = strings.TrimSpace(address)
	if address == "" {
		return "", nil
	}

	separatorPattern := regexp.MustCompile(`[.:;\-\s]+`)

	tokens := separatorPattern.Split(address, -1)

	var sb strings.Builder
	for _, token := range tokens {
		normalizedGroup := normalizeGroup(token)
		if normalizedGroup == "" {
			return "", fmt.Errorf("invalid group: %s", token)
		}
		sb.WriteString(normalizedGroup)
	}

	duid := sb.String()

	duidPattern := regexp.MustCompile(`^[0-9a-fA-F]{4,260}$`)
	if !duidPattern.MatchString(duid) {
		return "", fmt.Errorf("invalid DUID format: must be 4-260 hex characters")
	}

	if len(duid)%2 != 0 {
		return "", fmt.Errorf("invalid DUID: must have even number of hex characters")
	}

	return formatDUIDWithColons(strings.ToLower(duid)), nil
}

// normalizeGroup validates and pads a group to 2 characters
func normalizeGroup(group string) string {
	if group == "" {
		return ""
	}

	hexPattern := regexp.MustCompile(`^[0-9a-fA-F]+$`)
	if !hexPattern.MatchString(group) {
		return ""
	}

	if len(group) >= 2 {
		return group
	}

	return "0" + group
}

// formatDUIDWithColons converts continuous hex string to colon-separated pairs
func formatDUIDWithColons(duid string) string {
	var result strings.Builder
	for i := 0; i < len(duid); i += 2 {
		if i > 0 {
			result.WriteString(":")
		}
		result.WriteString(duid[i : i+2])
	}
	return result.String()
}
