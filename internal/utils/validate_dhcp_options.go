package utils

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// specialDHCPOptionNames lists the NIOS DHCP options that the grid always sends,
// and which therefore accept `use_option` (and an empty `value` when use_option is
// true). Declaration order is reused in the diagnostic message.
var specialDHCPOptionNames = []string{
	"routers",
	"router-templates",
	"domain-name-servers",
	"domain-name",
	"broadcast-address",
	"broadcast-address-offset",
	"dhcp-lease-time",
	"dhcp6.name-servers",
}

// Special DHCP option names that require use_option to be set
var specialDHCPOptionNums = map[int64]bool{
	3: true, 6: true, 15: true, 23: true, 28: true, 51: true,
}

var specialDHCPOptionNameSet = func() map[string]bool {
	m := make(map[string]bool, len(specialDHCPOptionNames))
	for _, n := range specialDHCPOptionNames {
		m[n] = true
	}
	return m
}()

type dhcpOption struct {
	Name        types.String `tfsdk:"name"`
	Num         types.Int64  `tfsdk:"num"`
	VendorClass types.String `tfsdk:"vendor_class"`
	Value       types.String `tfsdk:"value"`
	UseOption   types.Bool   `tfsdk:"use_option"`
}

func ValidateDHCPOptionsConfig(ctx context.Context, options types.List, basePath path.Path, diagnostics *diag.Diagnostics) {
	if options.IsNull() || options.IsUnknown() {
		return
	}

	var opts []dhcpOption
	diagnostics.Append(options.ElementsAs(ctx, &opts, false)...)
	if diagnostics.HasError() {
		return
	}

	for i, option := range opts {
		attrPath := func(a string) path.Path { return basePath.AtListIndex(i).AtName(a) }

		if option.Value.IsNull() {
			diagnostics.AddAttributeError(
				attrPath("value"),
				"Invalid configuration for DHCP Option",
				"The 'value' attribute is a required field and must be set for all DHCP Options.",
			)
		}

		var isSpecialOption bool
		var optionName string

		switch {
		case !option.Name.IsNull() && !option.Name.IsUnknown():
			optionName = option.Name.ValueString()
			isSpecialOption = specialDHCPOptionNameSet[optionName]
		case !option.Num.IsNull() && !option.Num.IsUnknown():
			optionNum := option.Num.ValueInt64()
			isSpecialOption = specialDHCPOptionNums[optionNum]
			optionName = fmt.Sprintf("with num = %d", optionNum)
		case option.Name.IsUnknown() || option.Num.IsUnknown():
			continue
		default:
			diagnostics.AddAttributeError(
				attrPath("name"),
				"Invalid configuration for DHCP Option",
				"Either the 'name' or 'num' attribute must be set for all DHCP Options. "+
					"Missing both attributes for 'option' at index "+fmt.Sprint(i)+".",
			)
			continue
		}

		if !option.Value.IsNull() && !option.Value.IsUnknown() && option.Value.ValueString() == "" {
			if !isSpecialOption {
				diagnostics.AddAttributeError(
					attrPath("value"),
					"Invalid configuration for DHCP Option",
					"The 'value' attribute cannot be set as empty for Custom DHCP Option '"+optionName+"'.",
				)
			} else if !option.UseOption.IsUnknown() && !option.UseOption.IsNull() && !option.UseOption.ValueBool() {
				diagnostics.AddAttributeError(
					attrPath("value"),
					"Invalid configuration for DHCP Option",
					"The 'value' attribute cannot be set as empty for Special DHCP Option '"+optionName+"' when 'use_option' is set to false.",
				)
			}
		}

		if !isSpecialOption && !option.UseOption.IsNull() && !option.UseOption.IsUnknown() {
			diagnostics.AddAttributeError(
				attrPath("use_option"),
				"Invalid configuration",
				fmt.Sprintf("The 'use_option' attribute should not be set for Custom DHCP Option '%s'. "+
					"It is only applicable for Special Options: %s.",
					optionName, strings.Join(specialDHCPOptionNames, ", ")),
			)
		}
	}
}
