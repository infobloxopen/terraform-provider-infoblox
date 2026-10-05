package dns

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDnsHost validates the DnsHost configuration.
func ValidateDnsHost(ctx context.Context, data DnsHostModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIDnsHostModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDnsHostUDDIConfig(ctx, uddi, resp)
	}
}

func validateDnsHostUDDIConfig(ctx context.Context, m *UDDIDnsHostModel, resp *resource.ValidateConfigResponse) {
}

func (r *DnsHostResource) preDeleteDnsHost(ctx context.Context, data *DnsHostModel, diags *diag.Diagnostics) {
	obj := data.Expand(ctx, diags, false)
	if diags.HasError() {
		return
	}

	if obj.UDDI == nil {
		obj.UDDI = &coremodel.UDDIDnsHostExt{}
	}

	obj.UDDI.Server = nil
	_, _, err := r.service.Update(ctx, data.Id.ValueString(), obj, &core.Options{
		ReturnFields: DnsHostReturnFields,
	})
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to detach Server before Host deletion: %s", err))
	}

}
