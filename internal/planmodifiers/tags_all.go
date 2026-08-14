package planmodifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TagsAllUseStateForUnknown returns a plan modifier for tags_all computed maps
// that preserves the prior state value when the sibling tags field has not changed.
// This prevents perpetual diffs on no-op plans while still allowing tags_all to
// reflect actual API values when tags are updated.
func TagsAllUseStateForUnknown(tagsPath path.Path) planmodifier.Map {
	return tagsAllUseStateForUnknown{tagsPath: tagsPath}
}

type tagsAllUseStateForUnknown struct {
	tagsPath path.Path
}

func (m tagsAllUseStateForUnknown) Description(_ context.Context) string {
	return "Preserve prior state for tags_all when tags are unchanged."
}

func (m tagsAllUseStateForUnknown) MarkdownDescription(_ context.Context) string {
	return "Preserve prior state for tags_all when tags are unchanged."
}

func (m tagsAllUseStateForUnknown) PlanModifyMap(ctx context.Context, req planmodifier.MapRequest, resp *planmodifier.MapResponse) {
	if !req.PlanValue.IsUnknown() {
		return
	}

	// Skip during Create — the resource has no id in state yet.
	var stateId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &stateId)
	if stateId.IsNull() || stateId.IsUnknown() {
		return
	}

	var planTags, stateTags types.Map
	req.Plan.GetAttribute(ctx, m.tagsPath, &planTags)
	req.State.GetAttribute(ctx, m.tagsPath, &stateTags)

	if planTags.Equal(stateTags) {
		resp.PlanValue = req.StateValue
	}
}
