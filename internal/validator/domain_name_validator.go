package validator

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

const (
	labelMaxLength = 63
	dot            = "."
)

var (
	escapedPattern = regexp.MustCompile(`\\(\d{3}|.)`)
	labelsPattern  = regexp.MustCompile(`(.*?)?([^\\])\.`)
)

// DomainNameOption defines a functional option for domainNameValidator
type DomainNameOption func(*domainNameValidator)

// domainNameValidator validates if the provided value is a valid domain name.
type domainNameValidator struct {
	isMultiLabel        bool
	allowNullOrEmpty    bool
	allowRootZone       bool
	allowTrailingDot    bool
	checkPrintableChars bool
}

// IsValidDomainName creates a new domain name validator with default options
// (multi-label enabled, null/empty not allowed, root zone allowed, trailing dot not allowed, printable chars not checked)
func IsValidDomainName(opts ...DomainNameOption) validator.String {
	v := domainNameValidator{
		isMultiLabel:        true,
		allowRootZone:       true,
		allowNullOrEmpty:    false,
		allowTrailingDot:    false,
		checkPrintableChars: false,
	}

	for _, opt := range opts {
		opt(&v)
	}

	return v
}

func (v domainNameValidator) Description(ctx context.Context) string {
	return "value must be a valid domain name format"
}

func (v domainNameValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v domainNameValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	// Check for root zone value
	if v.allowRootZone && value == dot {
		return
	}

	// Check for null or empty value
	if v.allowNullOrEmpty && strings.TrimSpace(value) == "" {
		return
	}

	validatable := value
	// Handle trailing dot
	if v.allowTrailingDot {
		if len(value) > 1 && strings.HasSuffix(value, dot) {
			validatable = value[:len(value)-1]
		}
	}

	// Perform domain name validation
	if err := isDomainName(validatable, v.isMultiLabel, v.checkPrintableChars); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Domain Name Format",
			err.Error(),
		)
	}
}

// isDomainName validates if a string is a valid domain name
func isDomainName(s string, allowMultiLabel bool, checkPrintableChars bool) error {
	// Check for leading or trailing dots
	if s == "" || strings.HasPrefix(s, dot) || strings.HasSuffix(s, dot) {
		return fmt.Errorf("domain name cannot be empty or have leading/trailing dots")
	}

	// Check for leading or trailing whitespace
	if strings.TrimSpace(s) != s {
		return fmt.Errorf("domain name cannot have leading or trailing whitespace")
	}

	// Check for uppercase characters
	if s != strings.ToLower(s) {
		return fmt.Errorf("domain name must not contain uppercase characters")
	}

	// Check printable characters (ASCII 32-126)
	if checkPrintableChars {
		for _, ch := range s {
			if ch < 32 || ch > 126 {
				return fmt.Errorf("domain name contains non-printable characters (must be ASCII 32-126)")
			}
		}
	}

	// Check total length
	length := findLabelLength(s)
	if length < 1 {
		return fmt.Errorf("domain name has invalid escape sequences or is empty")
	}
	if length > 255 {
		return fmt.Errorf("domain name exceeds maximum length of 255 characters")
	}

	// Split into labels and validate each
	matches := labelsPattern.FindAllStringSubmatch(s, -1)
	var last string

	for _, match := range matches {
		if len(match) < 3 {
			continue
		}

		label := match[0]
		if strings.HasPrefix(label, dot) {
			return fmt.Errorf("domain name contains empty label (consecutive dots)")
		}

		// Reconstruct the label from captured groups
		labelPart := match[1] + match[2]

		labelLen := findLabelLength(labelPart)
		if labelLen > labelMaxLength {
			return fmt.Errorf("domain name contains a label exceeding maximum length of 63 characters")
		}
		if labelLen < 1 {
			return fmt.Errorf("domain name contains an empty label")
		}

		// Find the position after this match
		matchEnd := strings.Index(s, label) + len(label)
		if matchEnd < len(s) {
			last = s[matchEnd:]
		}
	}

	// Check if multi-label is required
	if last != "" && (!allowMultiLabel || strings.HasPrefix(last, dot)) {
		if !allowMultiLabel {
			return fmt.Errorf("domain name contains multiple labels but only single label is allowed")
		}
		return fmt.Errorf("domain name contains empty label (consecutive dots)")
	}

	// Validate the last label (or the entire string if it's a single label)
	lastLabel := last
	if lastLabel == "" {
		lastLabel = s
	}
	lastLen := findLabelLength(lastLabel)
	if lastLen > labelMaxLength {
		return fmt.Errorf("domain name contains a label exceeding maximum length of 63 characters")
	}
	if lastLen < 1 {
		return fmt.Errorf("domain name contains an empty label")
	}

	return nil
}

// findLabelLength calculates the length of a label according to the rules:
// [\.], [\\], [\"], and [\nnn] count as one character
// Returns -1 if label is invalid, positive integer with label length otherwise
func findLabelLength(s string) int {
	if s == "" {
		return -1
	}

	matches := escapedPattern.FindAllStringSubmatch(s, -1)
	skipChars := 0

	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		group := match[1]

		// Check if it's a digit escape sequence (\nnn)
		if len(group) == 3 && unicode.IsDigit(rune(group[0])) {
			num, err := strconv.Atoi(group)
			if err != nil || num > 255 {
				return -1
			}
			skipChars += 3 // [\nnn] counts as one char, skip 3 extra
		} else if group != "\"" && group != "." && group != "\\" {
			// Only \", \., and \\ are allowed as single char escapes
			return -1
		} else {
			skipChars += 1 // bypass backslash
		}
	}

	return len(s) - skipChars
}

// WithAllowNullOrEmpty returns a DomainNameOption that allows null or empty string values.
func WithAllowNullOrEmpty() DomainNameOption {
	return func(v *domainNameValidator) {
		v.allowNullOrEmpty = true
	}
}
