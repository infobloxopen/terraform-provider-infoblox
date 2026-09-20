package infra

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateInfraHost validates the InfraHost configuration.
func ValidateInfraHost(ctx context.Context, data InfraHostModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIInfraHostModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateInfraHostUDDIConfig(ctx, uddi, resp)
	}
}

func validateInfraHostUDDIConfig(_ context.Context, m *UDDIInfraHostModel, resp *resource.ValidateConfigResponse) {
	if m.SerialNumber.IsNull() || m.SerialNumber.IsUnknown() {
		return
	}

	// tags cannot be validated when unknown (e.g. computed from another resource).
	if m.Tags.IsUnknown() {
		return
	}

	tagsPath := path.Root("uddi").AtName("tags")

	if m.Tags.IsNull() {
		resp.Diagnostics.AddAttributeError(tagsPath,
			"Missing Required Tag",
			`When "serial_number" is set, "tags" must contain the key "host/serial_number" with the same value.`,
		)
		return
	}

	tagVal, exists := m.Tags.Elements()["host/serial_number"]
	if !exists {
		resp.Diagnostics.AddAttributeError(tagsPath,
			"Missing Required Tag",
			`When "serial_number" is set, "tags" must contain the key "host/serial_number" with the same value as serial number.`,
		)
		return
	}

	tagStr, ok := tagVal.(types.String)
	if !ok || tagStr.IsNull() || tagStr.IsUnknown() {
		return
	}

	if tagStr.ValueString() != m.SerialNumber.ValueString() {
		resp.Diagnostics.AddAttributeError(
			tagsPath.AtMapKey("host/serial_number"),
			"Tag Value Mismatch",
			fmt.Sprintf(`The "tags" key "host/serial_number" must match "serial_number" (got %q, want %q).`,
				tagStr.ValueString(), m.SerialNumber.ValueString()),
		)
	}
}
