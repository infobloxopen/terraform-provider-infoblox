package suppressdiff

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Int64 = useStateToSuppressDiffInt64{}

type useStateToSuppressDiffInt64 struct{}

func UseStateToSuppressDiffInt64() planmodifier.Int64 {
	return useStateToSuppressDiffInt64{}
}

func (m useStateToSuppressDiffInt64) Description(_ context.Context) string {
	return "Uses the last known value for planning until the attribute is explicitly reset."
}

func (m useStateToSuppressDiffInt64) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

// userSetKey is the private-state key for the attribute indicating if the user set the value,
// derived from the attribute path so each field is distinct within resource
func userSetKey(p path.Path) string {
	return "userset::" + p.String()
}

// PlanModifyInt64 implements the planmodifier.Int64 interface for useStateToSuppressDiffInt64
// config has a value                          -> use config (leave plan as-is)
// config null, no prior state (create)        -> leave as-is (null on first apply)
// config null, userset flag true              -> user clearing it -> unknown (backend recomputes)
// config null, state non-null                 -> carry state forward (suppress server-side mutations)
// config null, state null, prior state exists -> unknown (server didn't echo value on create; accept result)
func (m useStateToSuppressDiffInt64) PlanModifyInt64(ctx context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	key := userSetKey(req.Path)

	// Determine if the user explicitly set a value in the prior plan (or cleared it)
	wasSet := false
	if req.Private != nil {
		if v, diags := req.Private.GetKey(ctx, key); !diags.HasError() && string(v) == "true" {
			wasSet = true
		}
	}

	// Set the private-state flag for this attribute so we can detect if the user clears it in a next plan
	if resp.Private != nil {
		flag := "false"
		if !req.ConfigValue.IsNull() {
			flag = "true"
		}
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, key, []byte(flag))...)
	}

	// User set an explicit value: the framework already put it in the plan
	if !req.ConfigValue.IsNull() {
		return
	}

	// Detect a real create via the whole prior state, not this field: a field can
	// stay null after create (no inheritance/echo), and treating that as "create"
	// would keep it unknown forever (perpetual diff).
	if req.State.Raw.IsNull() {
		return
	}

	if wasSet {
		// The user is clearing a previously-set value: hand control back to the backend to recompute
		resp.PlanValue = types.Int64Unknown()
		return
	}

	// State has a known non-null value: carry it forward to suppress server-side mutations.
	if !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
		return
	}

	// State exists but this field is null — the server didn't echo the inherited value on
	// create. Mark as unknown so the post-apply server value
	// (e.g. the grid-default TTL) is accepted into state. After that one update the field
	// will be non-null and the steady-state carry-forward above takes over permanently.
	resp.PlanValue = types.Int64Unknown()
}
