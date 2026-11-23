// Package inherited provides plan modifiers for NIOS fields governed by a use_*
// flag (for example ttl and use_ttl).
//
// The provider sets the use flag to true only when one of its governed fields is
// set in config. When none is set, the flag is false and NIOS uses the inherited
// value, so the provider stores the governed fields as null. These modifiers plan
// null for that case, so that removing an override shows in the plan as
// "<value> -> null" instead of keeping the last value from state.
package inherited

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const description = "Plans null when neither this field nor any other field governed by the same use flag is set in config, because NIOS then uses the inherited value."

// unset reports whether the use flag will be false: the field and all its
// sibling fields (the other fields governed by the same use flag) are null in config.
func unset(ctx context.Context, config tfsdk.Config, configValue attr.Value, siblings []path.Path, diags *diag.Diagnostics) bool {
	if !configValue.IsNull() {
		return false
	}
	for _, p := range siblings {
		var v attr.Value
		diags.Append(config.GetAttribute(ctx, p, &v)...)
		if diags.HasError() || (v != nil && !v.IsNull()) {
			return false
		}
	}
	return true
}

// NullWhenUnsetInt64 returns a plan modifier for an Int64 field governed by a use flag.
// siblings are the paths of the other fields governed by the same flag.
func NullWhenUnsetInt64(siblings ...path.Path) planmodifier.Int64 {
	return nullWhenUnsetInt64{siblings: siblings}
}

type nullWhenUnsetInt64 struct{ siblings []path.Path }

func (m nullWhenUnsetInt64) Description(_ context.Context) string { return description }

func (m nullWhenUnsetInt64) MarkdownDescription(_ context.Context) string { return description }

func (m nullWhenUnsetInt64) PlanModifyInt64(ctx context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	if unset(ctx, req.Config, req.ConfigValue, m.siblings, &resp.Diagnostics) {
		resp.PlanValue = types.Int64Null()
	}
}

// NullWhenUnsetBool returns a plan modifier for a Bool field governed by a use flag.
// siblings are the paths of the other fields governed by the same flag.
func NullWhenUnsetBool(siblings ...path.Path) planmodifier.Bool {
	return nullWhenUnsetBool{siblings: siblings}
}

type nullWhenUnsetBool struct{ siblings []path.Path }

func (m nullWhenUnsetBool) Description(_ context.Context) string { return description }

func (m nullWhenUnsetBool) MarkdownDescription(_ context.Context) string { return description }

func (m nullWhenUnsetBool) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if unset(ctx, req.Config, req.ConfigValue, m.siblings, &resp.Diagnostics) {
		resp.PlanValue = types.BoolNull()
	}
}

// NullWhenUnsetObject returns a plan modifier for an Object field governed by a use flag.
// siblings are the paths of the other fields governed by the same flag.
func NullWhenUnsetObject(siblings ...path.Path) planmodifier.Object {
	return nullWhenUnsetObject{siblings: siblings}
}

type nullWhenUnsetObject struct{ siblings []path.Path }

func (m nullWhenUnsetObject) Description(_ context.Context) string { return description }

func (m nullWhenUnsetObject) MarkdownDescription(_ context.Context) string { return description }

func (m nullWhenUnsetObject) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if unset(ctx, req.Config, req.ConfigValue, m.siblings, &resp.Diagnostics) {
		resp.PlanValue = types.ObjectNull(req.PlanValue.AttributeTypes(ctx))
	}
}
