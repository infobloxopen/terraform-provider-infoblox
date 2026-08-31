package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRuleset validates the Ruleset configuration.
func ValidateRuleset(ctx context.Context, data RulesetModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRulesetModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRulesetNIOSConfig(ctx, nios, resp)
	}
}

func validateRulesetNIOSConfig(ctx context.Context, m *NIOSRulesetModel, resp *resource.ValidateConfigResponse) {
}
