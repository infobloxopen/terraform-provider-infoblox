package suppressdiff

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = useStateToSuppressDiffObject{}

type useStateToSuppressDiffObject struct{}

func UseStateToSuppressDiffObject() planmodifier.Object {
	return useStateToSuppressDiffObject{}
}

func (m useStateToSuppressDiffObject) Description(_ context.Context) string {
	return "Uses the last known value for planning until the attribute is explicitly reset."
}

func (m useStateToSuppressDiffObject) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

// PlanModifyObject implements the planmodifier.Object interface.
// config has a value              -> use config (leave plan as-is)
// config null, no prior state     -> leave unknown (create: known after apply)
// config null, userset flag true  -> user is clearing it -> unknown (backend recomputes)
// config null, userset flag false -> still inherited -> carry state (quiet, converges)
func (m useStateToSuppressDiffObject) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
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
		resp.PlanValue = types.ObjectUnknown(req.StateValue.AttributeTypes(ctx))
		return
	}

	// Still inherited/unset: carry the last resolved value so the plan stays quiet
	// and converges to "No changes".
	resp.PlanValue = req.StateValue
}
