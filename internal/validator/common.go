package validator

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
)

// ValidateBackendBlocks validates that the correct backend-specific block is present
// and the wrong backend block is not specified.
func ValidateBackendBlocks(backend core.BackendType, niosBlock, uddiBlock types.Object, diags *diag.Diagnostics) {
	switch backend {
	case core.BackendNIOS:
		if !uddiBlock.IsNull() {
			diags.AddError(
				"Invalid Configuration",
				"The 'uddi' block is not allowed when using the NIOS backend or NIOS via the Infoblox Portal. Use the 'nios' block here, or to manage UDDI objects use a provider configured with a 'uddi' block and 'enable_nios_passthru = false'.",
			)
			return
		}
		if niosBlock.IsNull() {
			diags.AddError(
				"Missing Required Block",
				"The 'nios' block is required when using the NIOS backend or NIOS via the Infoblox Portal. Add a 'nios' block here, or to manage UDDI objects use a provider configured with a 'uddi' block and 'enable_nios_passthru = false'.",
			)
		}
	case core.BackendUDDI:
		if !niosBlock.IsNull() {
			diags.AddError(
				"Invalid Configuration",
				"The 'nios' block is not allowed when using the UDDI backend. Use the 'uddi' block here, or to manage NIOS objects use a provider configured with a 'nios' block, or a 'uddi' block with 'enable_nios_passthru = true'.",
			)
			return
		}
		if uddiBlock.IsNull() {
			diags.AddError(
				"Missing Required Block",
				"The 'uddi' block is required when using the UDDI backend. Add a 'uddi' block here, or to manage NIOS objects use a provider configured with a 'nios' block, or a 'uddi' block with 'enable_nios_passthru = true'.",
			)
		}
	}
}

// AddBackendFieldError adds an error for a field that is only valid for a specific backend
func AddBackendFieldError(diags *diag.Diagnostics, fieldName, requiredBackend string) {
	diags.AddError(
		"Invalid Configuration",
		fmt.Sprintf("The '%s' field is only applicable for the %s backend.", fieldName, requiredBackend),
	)
}

// ValidateDataSourceFilters validates backend-specific filter fields for datasources.
func ValidateDataSourceFilters(backend core.BackendType, extAttrFilters, tagFilters types.Map, maxResults, limit types.Int32, diags *diag.Diagnostics) {
	// ext_attr_filters is NIOS only
	if !extAttrFilters.IsNull() && backend == core.BackendUDDI {
		AddBackendFieldError(diags, "ext_attr_filters", "NIOS")
	}

	// max_results is NIOS only
	if !maxResults.IsNull() && backend == core.BackendUDDI {
		AddBackendFieldError(diags, "max_results", "NIOS")
	}

	// limit is UDDI only
	if !limit.IsNull() && backend == core.BackendNIOS {
		AddBackendFieldError(diags, "limit", "UDDI")
	}

	// tag_filters is UDDI only
	if !tagFilters.IsNull() && backend == core.BackendNIOS {
		AddBackendFieldError(diags, "tag_filters", "UDDI")
	}
}

// ValidateListFilters validates backend-specific filter fields for lists.
func ValidateListFilters(backend core.BackendType, extAttrFilters, tagFilters types.Map, diags *diag.Diagnostics) {
	// ext_attr_filters is NIOS only
	if !extAttrFilters.IsNull() && !extAttrFilters.IsUnknown() && backend == core.BackendUDDI {
		AddBackendFieldError(diags, "ext_attr_filters", "NIOS")
	}

	// tag_filters is UDDI only
	if !tagFilters.IsNull() && !tagFilters.IsUnknown() && backend == core.BackendNIOS {
		AddBackendFieldError(diags, "tag_filters", "UDDI")
	}
}
