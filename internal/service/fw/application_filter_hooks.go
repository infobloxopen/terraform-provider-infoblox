package fw

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateApplicationFilter validates the ApplicationFilter configuration.
func ValidateApplicationFilter(ctx context.Context, data ApplicationFilterModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIApplicationFilterModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateApplicationFilterUDDIConfig(ctx, uddi, resp)
	}
}

func validateApplicationFilterUDDIConfig(ctx context.Context, m *UDDIApplicationFilterModel, resp *resource.ValidateConfigResponse) {
	var criteria []ApplicationCriterionModel
	resp.Diagnostics.Append(m.Criteria.ElementsAs(ctx, &criteria, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	for i, c := range criteria {
		name := c.Name.ValueString()
		category := c.Category.ValueString()
		if name != "" && category != "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("uddi").AtName("criteria").AtListIndex(i),
				"Invalid Criterion",
				fmt.Sprintf("criteria[%d]: 'name' and 'category' are mutually exclusive — set one or the other, not both.", i),
			)
		}
	}
}
