package anycast

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/anycast"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateAnycastHost validates the AnycastHost configuration.
func ValidateAnycastHost(ctx context.Context, data AnycastHostModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIAnycastHostModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateAnycastHostUDDIConfig(ctx, uddi, resp)
	}
}

func validateAnycastHostUDDIConfig(ctx context.Context, m *UDDIAnycastHostModel, resp *resource.ValidateConfigResponse) {
}

func (r *AnycastHostResource) lookupAnycastHost(ctx context.Context, data *AnycastHostModel, obj *coremodel.AnycastHost, diags *diag.Diagnostics) {
	results, _, _, err := r.lookupService.List(ctx, &core.ListOptions{
		InternalFilters: map[string]string{
			"legacy_id": fmt.Sprintf("%d", *obj.Id),
		},
	})
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to look up AnycastHost: %s", err))
		return
	}
	if len(results) != 1 {
		diags.AddError("Client Error", fmt.Sprintf("Expected exactly one host with legacy_id %d, found %d", *obj.Id, len(results)))
		return
	}

	host := results[0]
	if host.UDDI != nil {
		if obj.UDDI == nil {
			obj.UDDI = &coremodel.UDDIAnycastHostExt{}
		}
		obj.UDDI.Name = &host.UDDI.DisplayName
	}
}
