package immutable

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.Object = immutableObject{}

// immutableObject validates that the provided object is not mutated after resource creation.
type immutableObject struct{}

func (m immutableObject) Description(ctx context.Context) string {
	return "Ensures this attribute cannot be changed after resource creation"
}

func (m immutableObject) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m immutableObject) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
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

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Immutable Attribute Changed",
		fmt.Sprintf("The attribute cannot be changed after creation. "+
			"To change this value, the resource must be destroyed and recreated."),
	)
}

// ImmutableObject returns a plan modifier that ensures the given object attribute cannot be changed after creation.
func ImmutableObject() planmodifier.Object {
	return immutableObject{}
}
