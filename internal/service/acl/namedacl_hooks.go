package acl

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateNamedacl validates the Namedacl configuration.
func ValidateNamedacl(ctx context.Context, data NamedaclModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSNamedaclModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateNamedaclNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDINamedaclModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateNamedaclUDDIConfig(ctx, uddi, resp)
	}
}

func validateNamedaclNIOSConfig(ctx context.Context, m *NIOSNamedaclModel, resp *resource.ValidateConfigResponse) {
}

func validateNamedaclUDDIConfig(ctx context.Context, m *UDDINamedaclModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenNamedaclNIOS(ctx context.Context, planned, flattened *NIOSNamedaclModel, diags *diag.Diagnostics) {
	// Preserve user-configured /32 CIDR suffixes that the NIOS API strips on read.
	if planned == nil || planned.AccessList.IsNull() || planned.AccessList.IsUnknown() {
		return
	}
	var planItems []NamedaclAccessListModel
	diags.Append(planned.AccessList.ElementsAs(ctx, &planItems, false)...)
	if diags.HasError() {
		return
	}

	var apiItems []NamedaclAccessListModel
	if flattened.AccessList.IsNull() || flattened.AccessList.IsUnknown() {
		return
	}
	diags.Append(flattened.AccessList.ElementsAs(ctx, &apiItems, false)...)
	if diags.HasError() {
		return
	}

	for i := range apiItems {
		if i >= len(planItems) {
			break
		}
		plan := planItems[i].Address
		api := apiItems[i].Address
		if !plan.IsNull() && !plan.IsUnknown() && !api.IsNull() && !api.IsUnknown() {
			if strings.HasSuffix(plan.ValueString(), "/32") && !strings.HasSuffix(api.ValueString(), "/32") {
				apiItems[i].Address = types.StringValue(api.ValueString() + "/32")
			}
		}
	}

	newList, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: NamedaclAccessListAttrTypes}, apiItems)
	diags.Append(d...)
	if !d.HasError() {
		flattened.AccessList = newList
	}
}
