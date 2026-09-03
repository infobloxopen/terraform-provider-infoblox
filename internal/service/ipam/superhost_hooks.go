package ipam

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateSuperhost validates the Superhost configuration.
func ValidateSuperhost(ctx context.Context, data SuperhostModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSSuperhostModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateSuperhostNIOSConfig(ctx, nios, resp)
	}
}

func validateSuperhostNIOSConfig(ctx context.Context, m *NIOSSuperhostModel, resp *resource.ValidateConfigResponse) {
	// NIOS stores a host record under dns_associated_objects and leaves this attribute
	// empty, which fails the apply.
	if m.DhcpAssociatedObjects.IsNull() || m.DhcpAssociatedObjects.IsUnknown() {
		return
	}

	var refs []types.String
	resp.Diagnostics.Append(m.DhcpAssociatedObjects.ElementsAs(ctx, &refs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	for i, ref := range refs {
		// A ref from another resource is unknown here; PostExpandSuperhostNIOS catches those.
		if ref.IsNull() || ref.IsUnknown() {
			continue
		}
		if strings.HasPrefix(ref.ValueString(), "record:host") {
			resp.Diagnostics.AddAttributeError(
				path.Root("nios").AtName("dhcp_associated_objects").AtListIndex(i),
				"Invalid DHCP Associated Object",
				"Host record can only be associated with DNS Associated Objects, not DHCP Associated Objects.",
			)
		}
	}
}

// PostExpandSuperhostNIOS repeats the host record check on create and update, for refs that
// were still unknown when the config was validated.
func PostExpandSuperhostNIOS(ctx context.Context, ext *coremodel.NIOSSuperhostExt, diags *diag.Diagnostics) *coremodel.NIOSSuperhostExt {
	if ext == nil {
		return nil
	}

	for _, ref := range ext.DhcpAssociatedObjects {
		if strings.HasPrefix(ref, "record:host") {
			diags.AddError(
				"Invalid DHCP Associated Object",
				"Host record can only be associated with DNS Associated Objects, not DHCP Associated Objects.",
			)
			return ext
		}
	}

	return ext
}
