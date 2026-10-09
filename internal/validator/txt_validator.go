package validator

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = txtValidator{}

// txtValidator validates TXT record text field constraints.
type txtValidator struct {
	checkIfEmpty bool
}

func (v txtValidator) Description(ctx context.Context) string {
	return "value must not exceed 512 bytes total, with each substring limited to 255 bytes"
}

func (v txtValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v txtValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	// Parse substrings respecting quotes and escapes
	substrings, err := parseTXTSubstrings(value)
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			err.Error(),
			value,
		)
		return
	}

	// Check per-substring and total byte length
	total := 0
	isEmpty := true
	for _, substr := range substrings {
		byteLen := len(substr)

		if byteLen > 0 {
			isEmpty = false
		}

		if byteLen > 255 {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Text Value",
				fmt.Sprintf("substring exceeds 255 bytes: %q", value),
			)
			return
		}
		total += byteLen
	}

	// Check for empty text even with quotes
	if v.checkIfEmpty && isEmpty {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Text Value",
			fmt.Sprintf("Text cannot be empty: %q", value),
		)
		return
	}

	if total > 512 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Text Value",
			fmt.Sprintf("total text exceeds 512 bytes: %q", value),
		)
		return
	}
}

// parseTXTSubstrings parses TXT value into substrings.
// Quotes preserve/aren't counted, outside spaces delimit. Backslash escapes inside quotes. Errors on unbalanced quotes.
func parseTXTSubstrings(s string) ([]string, error) {
	var substrings []string
	var current strings.Builder
	inQuotes := false

	for i := 0; i < len(s); i++ {
		ch := s[i]

		// Handle quote characters
		if ch == '"' {
			if inQuotes {
				// Closing quote - save the substring
				substrings = append(substrings, current.String())
				current.Reset()
			}
			inQuotes = !inQuotes
			continue
		}

		// Handle escapes inside quotes
		if ch == '\\' && inQuotes {
			i++
			if i >= len(s) {
				return nil, fmt.Errorf("dangling escape in quoted substring")
			}
			current.WriteByte(s[i])
			continue
		}

		// Handle spaces outside quotes - split substrings
		if ch == ' ' && !inQuotes {
			if current.Len() > 0 {
				substrings = append(substrings, current.String())
				current.Reset()
			}
			continue
		}

		// Regular character
		current.WriteByte(ch)
	}

	// Check for unbalanced quotes
	if inQuotes {
		return nil, fmt.Errorf("unbalanced quotes in text")
	}

	// Add final substring if any
	if current.Len() > 0 {
		substrings = append(substrings, current.String())
	}

	return substrings, nil
}

// ValidateTXT returns a validator for TXT record text fields.
func ValidateTXT(checkIfEmpty bool) validator.String {
	return txtValidator{checkIfEmpty: checkIfEmpty}
}
