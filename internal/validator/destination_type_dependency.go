package validator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ validator.List = destinationTypeDependencyValidator{}
)

// destinationTypeDependencyValidator validates that if a destination of type 'DNS' is present, there must also be a destination of type 'IPAM/DHCP'.
// This is required for destinations field in Cloud Discovery Config.

type destinationTypeDependencyValidator struct{}

func (v destinationTypeDependencyValidator) Description(ctx context.Context) string {
	return "ensures that if a destination of type 'DNS' is present, there must also be a destination of type 'IPAM/DHCP'"
}

func (v destinationTypeDependencyValidator) MarkdownDescription(ctx context.Context) string {
	return "ensures that if a destination of type `DNS` is present, there must also be a destination of type `IPAM/DHCP`"
}

func (v destinationTypeDependencyValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	hasDNS := false
	hasIPAM := false
	hasACCOUNTS := false

	for _, item := range req.ConfigValue.Elements() {
		if destination, ok := item.(types.Object); ok {
			destinationMap := destination.Attributes()
			if destType, ok := destinationMap["destination_type"]; ok {
				if destType.Equal(types.StringValue("DNS")) {
					hasDNS = true
				}
				if destType.Equal(types.StringValue("IPAM/DHCP")) {
					hasIPAM = true
				}
				if destType.Equal(types.StringValue("ACCOUNTS")) {
					hasACCOUNTS = true
				}
			}
		}
	}

	if (hasDNS && !hasIPAM) || (hasACCOUNTS && !hasIPAM) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Missing Required Destination Type",
			"When a destination of type 'DNS' or 'ACCOUNTS' is provided, a destination of type 'IPAM/DHCP' must also be included.",
		)
	}
}

func DestinationTypeDependency() validator.List {
	return destinationTypeDependencyValidator{}
}
