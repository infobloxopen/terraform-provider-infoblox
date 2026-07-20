package flex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func DeriveUseFlag(ctx context.Context, config tfsdk.Config, diags *diag.Diagnostics, valuePaths ...path.Path) *bool {
	val := false
	for _, p := range valuePaths {
		var v attr.Value
		diags.Append(config.GetAttribute(ctx, p, &v)...)
		if diags.HasError() {
			return &val
		}
		if v != nil && !v.IsNull() && !v.IsUnknown() {
			val = true
		}
	}
	return &val
}
