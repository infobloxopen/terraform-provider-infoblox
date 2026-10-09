package validator

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = oidValidator{}

var (
	// oidSubIdentifierPattern validates that a sub-identifier contains only digits
	oidSubIdentifierPattern = regexp.MustCompile(`^\d+$`)
	// oidSubIdentifierMaxValue is the maximum value for one sub-identifier (2^32 - 1)
	oidSubIdentifierMaxValue = uint64(4294967295)
)

type oidValidator struct{}

func (v oidValidator) Description(_ context.Context) string {
	return "value must be a valid SNMP Object Identifier (OID)"
}

func (v oidValidator) MarkdownDescription(_ context.Context) string {
	return "Value must be a valid SNMP Object Identifier (OID) according to RFC3416. " +
		"OID must start with a dot, contain only digits and dots, have maximum 128 sub-identifiers, " +
		"and follow SNMP sub-identifier rules."
}

func (v oidValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	oid := req.ConfigValue.ValueString()

	// SNMP OID must start with a dot
	if !strings.HasPrefix(oid, ".") {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid OID",
			"OID must start with a dot (.)",
		)
		return
	}

	// SNMP OID must have at least one sub-identifier
	if len(oid) < 2 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid OID",
			"OID must contain at least one sub-identifier after the leading dot",
		)
		return
	}

	if strings.HasSuffix(oid, ".") {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid OID",
			"OID must not end with a dot (.)",
		)
		return
	}

	// split the OID into its sub-identifiers
	oidSubIdentifiers := strings.Split(oid, ".")

	// SNMP OID must not exceed 128 sub-identifiers
	if len(oidSubIdentifiers) > 129 { //129 because of leading dot
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid OID",
			"OID must not exceed 128 sub-identifiers",
		)
		return
	}

	// Sub-identifier can't be empty
	for i := 1; i < len(oidSubIdentifiers); i++ {
		if oidSubIdentifiers[i] == "" {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid OID Format",
				"OID sub-identifiers cannot be empty (consecutive dots are not allowed).",
			)
			return
		}
	}

	// First sub-identifier can have only the values of 0, 1, 2
	if len(oidSubIdentifiers) > 1 {
		if !oidSubIdentifierPattern.MatchString(oidSubIdentifiers[1]) {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid OID Format",
				"First sub-identifier must be a valid number (0, 1, or 2).",
			)
			return
		}

		firstSubID, err := strconv.ParseUint(oidSubIdentifiers[1], 10, 64)
		if err != nil || (firstSubID != 0 && firstSubID != 1 && firstSubID != 2) {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid OID Format",
				"First sub-identifier must be 0, 1, or 2.",
			)
			return
		}

		// Validate second sub-identifier constraints when first is 0 or 1
		if len(oidSubIdentifiers) > 2 {
			secondSubID, err := strconv.ParseUint(oidSubIdentifiers[2], 10, 64)
			if err == nil {
				if (firstSubID == 0 || firstSubID == 1) && secondSubID > 39 {
					resp.Diagnostics.AddAttributeError(
						req.Path,
						"Invalid OID Format",
						"Second sub-identifier must be between 0 and 39 when first sub-identifier is 0 or 1.",
					)
					return
				}
			}
		}

		//Validate remaining sub-identifiers (only digits and within max value)
		for i := 2; i < len(oidSubIdentifiers); i++ {
			if !oidSubIdentifierPattern.MatchString(oidSubIdentifiers[i]) {
				resp.Diagnostics.AddAttributeError(
					req.Path,
					"Invalid OID Format",
					fmt.Sprintf("Sub-identifier at index %d contains invalid characters. Only digits are allowed.", i),
				)
				return
			}

			subIDValue, err := strconv.ParseUint(oidSubIdentifiers[i], 10, 64)
			if err != nil || subIDValue > oidSubIdentifierMaxValue {
				resp.Diagnostics.AddAttributeError(
					req.Path,
					"Invalid OID Format",
					fmt.Sprintf("Sub-identifier at index %d exceeds maximum value of %d.", i, oidSubIdentifierMaxValue),
				)
				return
			}
		}

	}
}

func OIDValidator() validator.String {
	return oidValidator{}
}
