package planmodifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func UseDefaultAclForNull() planmodifier.List {
	return useDefaultAclForNull{}
}

// useDefaultAclForNull implements the plan modifier.
type useDefaultAclForNull struct{}

// Description returns a human-readable description of the plan modifier.
func (m useDefaultAclForNull) Description(_ context.Context) string {
	return "Sets the default value for the acl attribute in the plan if the value is null."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m useDefaultAclForNull) MarkdownDescription(_ context.Context) string {
	return "Sets the default value for the acl attribute in the plan if the value is null."
}

// PlanModifyList implements the plan modification logic.
func (m useDefaultAclForNull) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	// Do nothing if there is no state value.
	if req.StateValue.IsNull() {
		return
	}

	// Only default when the user left the value unset.
	if !req.ConfigValue.IsUnknown() && !req.ConfigValue.IsNull() {
		return
	}

	objType, ok := req.StateValue.ElementType(ctx).(types.ObjectType)
	if !ok {
		return
	}
	tsigType, _ := objType.AttributeTypes()["tsig_key"].(types.ObjectType)

	obj, diags := types.ObjectValue(objType.AttributeTypes(), map[string]attr.Value{
		"access":   types.StringValue("allow"),
		"acl":      types.StringNull(),
		"address":  types.StringValue(""),
		"element":  types.StringValue("any"),
		"tsig_key": types.ObjectNull(tsigType.AttributeTypes()),
	})
	if resp.Diagnostics = append(resp.Diagnostics, diags...); resp.Diagnostics.HasError() {
		return
	}

	resp.PlanValue, diags = types.ListValue(objType, []attr.Value{obj})
	if resp.Diagnostics = append(resp.Diagnostics, diags...); resp.Diagnostics.HasError() {
		return
	}
}
