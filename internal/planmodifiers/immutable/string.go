package immutable

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.String = immutableStringModifier{}

// immutableStringModifier validates that the provided string is not mutated after resource creation.
type immutableStringModifier struct{}

func (m immutableStringModifier) Description(ctx context.Context) string {
	return "Ensures this attribute cannot be changed after resource creation"
}

func (m immutableStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m immutableStringModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}

	if req.PlanValue.IsNull() {
		return
	}

	if req.StateValue.IsUnknown() || req.PlanValue.IsUnknown() {
		return
	}

	if req.StateValue.Equal(req.PlanValue) {
		return
	}

	stateVal := req.StateValue.ValueString()
	planVal := req.PlanValue.ValueString()

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Immutable Attribute Changed",
		fmt.Sprintf("The attribute cannot be changed after creation. "+
			"Existing value: %q, Planned value: %q. "+
			"To change this value, the resource must be destroyed and recreated.",
			stateVal,
			planVal,
		),
	)
}

// ImmutableString returns a plan modifier that ensures the given string attribute cannot be changed after creation.
func ImmutableString() planmodifier.String {
	return immutableStringModifier{}
}

var _ planmodifier.String = immutableIfValueStringModifier{}

// immutableIfValueStringModifier prevents changes to/from a specific value.
type immutableIfValueStringModifier struct {
	value string
}

func (m immutableIfValueStringModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("Ensures this attribute cannot be changed to or from %q after resource creation", m.value)
}

func (m immutableIfValueStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m immutableIfValueStringModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}

	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}

	stateVal := req.StateValue.ValueString()
	planVal := req.PlanValue.ValueString()

	if stateVal == m.value || planVal == m.value {
		if stateVal != planVal {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Immutable Attribute Changed",
				fmt.Sprintf("The attribute cannot be changed to or from %q. "+
					"Existing value: %q, Planned value: %q. "+
					"To change this value, the resource must be destroyed and recreated.",
					m.value,
					stateVal,
					planVal,
				),
			)
		}
	}
}

// ImmutableIfValue returns a plan modifier that prevents changes to or from a specific value.
func ImmutableIfValue(value string) planmodifier.String {
	return immutableIfValueStringModifier{value: value}
}
