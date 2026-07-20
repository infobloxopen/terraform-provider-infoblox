package validator

import (
	"context"
	"fmt"
	"strings"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type noLeadingTrailingWhitespaceUnlessQuoted struct{}

func NoLeadingTrailingWhitespaceUnlessQuoted() validator.String {
	return noLeadingTrailingWhitespaceUnlessQuoted{}
}

func (v noLeadingTrailingWhitespaceUnlessQuoted) Description(ctx context.Context) string {
	return `Disallow leading/trailing whitespace unless the value is quoted (e.g., "\"    hello\""). "\"\"" is allowed and treated as empty.`
}
func (v noLeadingTrailingWhitespaceUnlessQuoted) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v noLeadingTrailingWhitespaceUnlessQuoted) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// Ignore unknown/null—other validators or defaults may handle these.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	str := req.ConfigValue.ValueString()

	// Special-case: "\"\"" is allowed and is treated as empty by normalization logic.
	if str == "\"\"" {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid value for text",
			fmt.Sprintf(`Invalid value for text: %q: empty string with quotes is not allowed.`, str))
		return
	}

	isQuoted := strings.HasPrefix(str, "\"") && strings.HasSuffix(str, "\"")

	// Leading/trailing whitespace detection (Unicode).
	hasLeading := len(str) > 0 && unicode.IsSpace(rune(str[0]))
	hasTrailing := len(str) > 0 && unicode.IsSpace(rune(str[len(str)-1]))

	if (hasLeading || hasTrailing) && !isQuoted {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid value for text",
			fmt.Sprintf(
				`Invalid value for text: %q: leading or trailing whitespace is not allowed.
To enter leading, trailing or embedded spaces in the text, add quotes (" ") around the text to preserve the spaces.`,
				str,
			),
		)
	}
}
