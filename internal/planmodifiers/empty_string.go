package planmodifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func UseEmptyStringForNull() planmodifier.String {
	return useEmptyStringForNull{}
}

// useEmptyStringForNull implements the plan modifier.
type useEmptyStringForNull struct{}

// Description returns a human-readable description of the plan modifier.
func (m useEmptyStringForNull) Description(_ context.Context) string {
	return "Use an empty string for null values in the plan."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m useEmptyStringForNull) MarkdownDescription(_ context.Context) string {
	return "Use an empty string for null values in the plan."
}

// PlanModifyString implements the plan modification logic.
func (m useEmptyStringForNull) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Do nothing if there is a known planned value.
	if !req.PlanValue.IsUnknown() {
		return
	}

	// If the config value is unknown or null, use an empty string.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		resp.PlanValue = types.StringValue("")
	}
}
