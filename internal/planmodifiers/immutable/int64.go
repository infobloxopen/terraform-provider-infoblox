package immutable

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.Int64 = immutableInt64Modifier{}

// immutableInt64Modifier validates that the provided int64 is not mutated after resource creation.
type immutableInt64Modifier struct{}

func (m immutableInt64Modifier) Description(ctx context.Context) string {
	return "Ensures this attribute cannot be changed after resource creation"
}

func (m immutableInt64Modifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m immutableInt64Modifier) PlanModifyInt64(ctx context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
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

	stateVal := req.StateValue.ValueInt64()
	planVal := req.PlanValue.ValueInt64()

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Immutable Attribute Changed",
		fmt.Sprintf("The attribute cannot be changed after creation. "+
			"Existing value: %d, Planned value: %d. "+
			"To change this value, the resource must be destroyed and recreated.",
			stateVal,
			planVal,
		),
	)
}

// ImmutableInt64 returns a plan modifier that ensures the given int64 attribute cannot be changed after creation.
func ImmutableInt64() planmodifier.Int64 {
	return immutableInt64Modifier{}
}
