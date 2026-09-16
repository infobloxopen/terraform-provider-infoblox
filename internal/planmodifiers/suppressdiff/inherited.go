package suppressdiff

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// InheritedField describes an Optional+Computed attribute whose value is inherited
// from a parent config (e.g. a NIOS grid or zone default gated by a use_* flag).
// The server does not echo this value on POST (returns null) but does return it on
// PUT (or when additional config like nameserver is associated with it).
// After that the state carries the real value and the attribute plan modifier's
// carry-forward takes over permanently, so no further catch-up (storing grid default
// values into state via marking a field Unknown) is needed and it will be suppressed in future plans.
type InheritedField struct {
	Path         path.Path  // path to the attribute in the plan/state tree
	UnknownValue attr.Value // type-correct Unknown value (e.g. types.Int64Unknown())
}

// NOTE: This mechanism is primarily meant for fields governed by a NIOS use_* flag
// (server-inherited values not echoed on POST). For other Optional+Computed fields
// the expected behavior may differ — verify before relying on it.
//
// MarkInheritedFieldsUnknown marks each InheritedField as Unknown in the plan when all
// three conditions hold:
//  1. The resource already exists (state is not null — not a create or destroy)
//  2. A real update operation is in progress (plan differs from state in at least one field)
//  3. The field's current state value is null (server has not yet echoed the inherited value)
//  4. The field is not explicitly set in config (Terraform requires the planned value of an
//     Optional+Computed attribute to match a known config value exactly, so marking it Unknown
//     would produce a "provider produced invalid plan" error)
//
// This allows Terraform to accept whatever the server returns on the next API call,
// catching up state for inherited fields the server only populates starting from the
// first update. The null check in condition (3) is evaluated internally so callers of
// this function (per resource modify plan method) does not need to read state before calling this function.
func MarkInheritedFieldsUnknown(
	ctx context.Context,
	req resource.ModifyPlanRequest,
	resp *resource.ModifyPlanResponse,
	fields []InheritedField,
) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}

	// tftypes.Value.Equal() returns false if either side has Unknown or any field differs.
	// On an empty plan every attribute carries its state value (all known), so Equal()
	// returns true. On a real update at least one field differs -> false.
	if req.Plan.Raw.Equal(req.State.Raw) {
		return
	}

	// A real update is in progress: mark any null inherited fields Unknown so the
	// after-apply server value is accepted into state.
	for _, f := range fields {
		stateVal, err := valueAtPath(req.State.Raw, f.Path)
		if err != nil || !stateVal.IsNull() {
			continue
		}
		// The user set this field in config on this very update: the framework already
		// planned the config value and Terraform would reject an Unknown in its place.
		if configVal, err := valueAtPath(req.Config.Raw, f.Path); err == nil && configVal.IsKnown() && !configVal.IsNull() {
			continue
		}
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, f.Path, f.UnknownValue)...)
	}
}

// valueAtPath navigates a tftypes.Value tree by following each step in p, returning
// the value at the leaf. Returns the partially-navigated value on null/unknown intermediates
// so callers can check IsNull() correctly even for nested paths.
func valueAtPath(root tftypes.Value, p path.Path) (tftypes.Value, error) {
	current := root
	for _, step := range p.Steps() {
		if current.IsNull() || !current.IsKnown() {
			return current, nil
		}
		switch s := step.(type) {
		case path.PathStepAttributeName:
			raw, err := current.ApplyTerraform5AttributePathStep(tftypes.AttributeName(string(s)))
			if err != nil {
				return tftypes.Value{}, fmt.Errorf("navigating %q: %w", string(s), err)
			}
			current = raw.(tftypes.Value)
		default:
			return tftypes.Value{}, fmt.Errorf("unsupported path step type %T", step)
		}
	}
	return current, nil
}
