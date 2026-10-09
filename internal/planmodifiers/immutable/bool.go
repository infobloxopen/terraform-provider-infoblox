package immutable

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.Bool = immutableBoolModifier{}

// immutableBoolModifier validates that the provided bool is not mutated after resource creation.
type immutableBoolModifier struct{}

func (m immutableBoolModifier) Description(ctx context.Context) string {
	return "Ensures this attribute cannot be changed after resource creation"
}

func (m immutableBoolModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m immutableBoolModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
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

	stateVal := req.StateValue.ValueBool()
	planVal := req.PlanValue.ValueBool()

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Immutable Attribute Changed",
		fmt.Sprintf("The attribute cannot be changed after creation. "+
			"Existing value: %t, Planned value: %t. "+
			"To change this value, the resource must be destroyed and recreated.",
			stateVal,
			planVal,
		),
	)
}

// ImmutableBool returns a plan modifier that ensures the given bool attribute cannot be changed after creation.
func ImmutableBool() planmodifier.Bool {
	return immutableBoolModifier{}
}
