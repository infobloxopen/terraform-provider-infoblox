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

var _ basetypes.StringValuableWithSemanticEquals = MACAddressValue{}
var _ basetypes.StringTypable = MACAddressType{}

// MACAddressType is a custom type for MAC addresses with semantic equality
type MACAddressType struct {
	basetypes.StringType
}

func (t MACAddressType) String() string {
	return "MACAddressType"
}

func (t MACAddressType) ValueType(ctx context.Context) attr.Value {
	return MACAddressValue{}
}

func (t MACAddressType) Equal(o attr.Type) bool {
	other, ok := o.(MACAddressType)
	if !ok {
		return false
	}
	return t.StringType.Equal(other.StringType)
}

func (t MACAddressType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return MACAddressValue{
		StringValue: in,
	}, nil
}

func (t MACAddressType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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

// MACAddressValue is the value type with semantic equality
type MACAddressValue struct {
	basetypes.StringValue
}

func (v MACAddressValue) Type(ctx context.Context) attr.Type {
	return MACAddressType{}
}

func (v MACAddressValue) Equal(o attr.Value) bool {
	other, ok := o.(MACAddressValue)
	if !ok {
		return false
	}
	return v.StringValue.Equal(other.StringValue)
}

func (v MACAddressValue) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(MACAddressValue)
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

	currentMAC := v.ValueString()
	newMAC := newValue.ValueString()

	if currentMAC == "" && newMAC == "" {
		return true, diags
	}
	if currentMAC == "" || newMAC == "" {
		return false, diags
	}

	// Normalize both values
	currentNormalized, err1 := normalizeMAC(currentMAC)
	newNormalized, err2 := normalizeMAC(newMAC)

	if err1 != nil || err2 != nil {
		return false, diags
	}

	// Compare normalized values
	return currentNormalized == newNormalized, diags
}

func NewMACAddressValue(value string) MACAddressValue {
	return MACAddressValue{StringValue: basetypes.NewStringValue(value)}
}

func NewMACAddressNull() MACAddressValue {
	return MACAddressValue{
		StringValue: basetypes.NewStringNull(),
	}
}

// normalizeMAC normalizes MAC address to standard lowercase colon format
func normalizeMAC(address string) (string, error) {
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
		if token == "" {
			return "", fmt.Errorf("empty token found")
		}
		sb.WriteString(token)
	}

	mac := sb.String()

	macPattern := regexp.MustCompile(`^[0-9a-fA-F]{1,12}$`)
	if !macPattern.MatchString(mac) {
		return "", fmt.Errorf("invalid MAC format")
	}

	if len(mac) == 12 {
		// Converts AABBCCDDEEFF to aa:bb:cc:dd:ee:ff (lowercase)
		mac = strings.ToLower(mac)
		return fmt.Sprintf("%s:%s:%s:%s:%s:%s",
			mac[0:2], mac[2:4], mac[4:6], mac[6:8], mac[8:10], mac[10:12]), nil
	}

	// Check for AABB.CCDD.EEFF
	ciscoPattern := regexp.MustCompile(`^[0-9a-fA-F]{1,4}\.[0-9a-fA-F]{1,4}\.[0-9a-fA-F]{1,4}$`)
	if ciscoPattern.MatchString(address) {
		return normalizeCiscoFormat(address)
	}

	// Check for AABBCC-DDEEFF
	hyphenatedPattern := regexp.MustCompile(`^[0-9a-fA-F]{1,6}-[0-9a-fA-F]{1,6}$`)
	if hyphenatedPattern.MatchString(address) {
		return normalizeHyphenatedFormat(address)
	}

	// Standard format AA:BB:CC:DD:EE:FF
	return normalizeStandardFormat(tokens)
}

func normalizeCiscoFormat(address string) (string, error) {
	parts := strings.Split(address, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid Cisco format")
	}

	var result []string
	for _, part := range parts {
		if len(part) == 0 || len(part) > 4 {
			return "", fmt.Errorf("invalid group length")
		}

		// Pad with leading zeros to make it 4 chars
		for len(part) < 4 {
			part = "0" + part
		}

		// Validate hex characters
		if !regexp.MustCompile(`^[0-9a-fA-F]+$`).MatchString(part) {
			return "", fmt.Errorf("invalid hex characters")
		}

		// Convert to lowercase and split into 2-char groups
		part = strings.ToLower(part)
		result = append(result, part[0:2], part[2:4])
	}

	return strings.Join(result, ":"), nil
}

func normalizeHyphenatedFormat(address string) (string, error) {
	parts := strings.Split(address, "-")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid hyphenated format")
	}

	var result []string
	for _, part := range parts {
		if len(part) == 0 || len(part) > 6 {
			return "", fmt.Errorf("invalid group length")
		}

		// Pad with leading zeros to make it 6 chars
		for len(part) < 6 {
			part = "0" + part
		}

		// Validate hex characters
		if !regexp.MustCompile(`^[0-9a-fA-F]+$`).MatchString(part) {
			return "", fmt.Errorf("invalid hex characters")
		}

		// Convert to lowercase and split into 2-char groups
		part = strings.ToLower(part)
		for i := 0; i < 6; i += 2 {
			result = append(result, part[i:i+2])
		}
	}

	return strings.Join(result, ":"), nil
}

func normalizeStandardFormat(tokens []string) (string, error) {
	if len(tokens) != 6 {
		return "", fmt.Errorf("MAC address must have 6 groups")
	}

	var result []string
	for _, token := range tokens {
		if len(token) == 0 || len(token) > 2 {
			return "", fmt.Errorf("invalid group length")
		}

		// Validate hex characters
		if !regexp.MustCompile(`^[0-9a-fA-F]+$`).MatchString(token) {
			return "", fmt.Errorf("invalid hex characters")
		}

		// Pad with leading zero if needed and convert to lowercase
		if len(token) == 1 {
			token = "0" + token
		}

		result = append(result, strings.ToLower(token))
	}

	return strings.Join(result, ":"), nil
}
