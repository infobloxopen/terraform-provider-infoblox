package anycast

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateAnycastConfig validates the AnycastConfig configuration.
func ValidateAnycastConfig(ctx context.Context, data AnycastConfigModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIAnycastConfigModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateAnycastConfigUDDIConfig(ctx, uddi, resp)
	}
}

func validateAnycastConfigUDDIConfig(ctx context.Context, m *UDDIAnycastConfigModel, resp *resource.ValidateConfigResponse) {
	if !m.OnpremHosts.IsNull() && !m.OnpremHosts.IsUnknown() {
		var hosts []OnpremHostRefModel
		resp.Diagnostics.Append(m.OnpremHosts.ElementsAs(ctx, &hosts, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		for i, host := range hosts {
			if host.Id.IsUnknown() || host.Name.IsUnknown() {
				continue
			}
			if host.Id.IsNull() {
				resp.Diagnostics.AddError(
					"Missing Required Field",
					fmt.Sprintf("onprem_hosts[%d]: 'id' is required when 'onprem_hosts' is configured", i),
				)
			}
			if host.Name.IsNull() || host.Name.ValueString() == "" {
				resp.Diagnostics.AddError(
					"Missing Required Field",
					fmt.Sprintf("onprem_hosts[%d]: 'name' is required when 'onprem_hosts' is configured", i),
				)
			}
		}
	}
}
