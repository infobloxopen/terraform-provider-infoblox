package validator

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/infobloxopen/terraform-provider-unified/internal/utils"
)

var _ validator.String = timeFormatValidator{}

type timeFormatValidator struct{}

func (v timeFormatValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("Ensures the string is in the format %s", utils.NaiveDatetimeLayout)
}

func (v timeFormatValidator) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("Ensures the string is a valid timestamp in the format %s", utils.NaiveDatetimeLayout)
}

func (v timeFormatValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	_, err := time.Parse(utils.NaiveDatetimeLayout, value)
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid time format",
			fmt.Sprintf("Expected format: YYYY-MM-DDTHH:MM:SS (example: %s), Got: %s", utils.NaiveDatetimeLayout, value),
		)
	}
}

func ValidateTimeFormat() validator.String {
	return timeFormatValidator{}
}
