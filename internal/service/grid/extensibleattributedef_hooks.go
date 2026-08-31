package grid

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateExtensibleattributedef validates the Extensibleattributedef configuration.
func ValidateExtensibleattributedef(ctx context.Context, data ExtensibleattributedefModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSExtensibleattributedefModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateExtensibleattributedefNIOSConfig(ctx, nios, resp)
	}
}

func validateExtensibleattributedefNIOSConfig(ctx context.Context, m *NIOSExtensibleattributedefModel, resp *resource.ValidateConfigResponse) {
	if m.Type.IsUnknown() || m.Type.IsNull() {
		return
	}
	typeValue := m.Type.ValueString()

	// min and max are only applicable for INTEGER type (per WAPI docs)
	if (!m.Min.IsNull() || !m.Max.IsNull()) && typeValue != "INTEGER" {
		resp.Diagnostics.AddError(
			"Invalid Min/Max Configuration",
			fmt.Sprintf("'min' and 'max' are only valid for INTEGER type, but type is %q.", typeValue),
		)
	}

	// default_value for INTEGER must be a parseable 32-bit integer
	if !m.DefaultValue.IsNull() && !m.DefaultValue.IsUnknown() && typeValue == "INTEGER" {
		if _, err := strconv.ParseInt(m.DefaultValue.ValueString(), 10, 32); err != nil {
			resp.Diagnostics.AddError(
				"Invalid Integer Default Value",
				fmt.Sprintf("'default_value' %q is not a valid integer.", m.DefaultValue.ValueString()),
			)
		}
	}
}
