package clouddiscovery

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/clouddiscovery"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateCloudDiscoveryProvider validates the CloudDiscoveryProvider configuration.
func ValidateCloudDiscoveryProvider(ctx context.Context, data CloudDiscoveryProviderModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDICloudDiscoveryProviderModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateCloudDiscoveryProviderUDDIConfig(ctx, uddi, resp)
	}
}

func validateCloudDiscoveryProviderUDDIConfig(ctx context.Context, m *UDDICloudDiscoveryProviderModel, resp *resource.ValidateConfigResponse) {
}

func (r *CloudDiscoveryProviderResource) preDeleteCloudDiscoveryProvider(ctx context.Context, data *CloudDiscoveryProviderModel, diags *diag.Diagnostics) {

	obj := data.Expand(ctx, diags, false)
	if diags.HasError() {
		return
	}

	if obj.UDDI == nil {
		obj.UDDI = &coremodel.UDDICloudDiscoveryProviderExt{}
	}
	disabled := "disabled"
	obj.UDDI.DesiredState = &disabled

	_, _, err := r.service.Update(ctx, data.Id.ValueString(), obj, &core.Options{
		ReturnFields: CloudDiscoveryProviderReturnFields,
	})
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to disable Providers before deletion: %s", err))
	}
}
