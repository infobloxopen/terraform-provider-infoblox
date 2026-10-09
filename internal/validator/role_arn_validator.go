package validator

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = awsRoleArnValidator{}

type awsRoleArnValidator struct{}

var (
	awsRoleArnRegex = regexp.MustCompile(`^arn:aws:iam::[0-9]{12}:role/[a-zA-Z0-9+=,.@_-]+$`)
)

func (v awsRoleArnValidator) Description(ctx context.Context) string {
	return "Role ARN should be of the format arn:aws:iam::123456789012:role/Role-name"
}

func (v awsRoleArnValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v awsRoleArnValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if !awsRoleArnRegex.MatchString(value) {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"Invalid Role ARN. Role ARN should be of the format arn:aws:iam::123456789012:role/Role-name",
			value,
		))
	}
}

// IsValidAwsRoleArn returns a validator.String that validates AWS Role ARN format.
func IsValidAwsRoleArn() validator.String {
	return awsRoleArnValidator{}
}
