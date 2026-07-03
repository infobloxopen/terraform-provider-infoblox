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
	// Require backend-specific block (can be empty)
	if backend == core.BackendNIOS && niosBlock.IsNull() {
		diags.AddError(
			"Missing Required Block",
			"The 'nios' block is required when using the NIOS backend. Use 'nios = {}' if no attributes needed.",
		)
	}
	if backend == core.BackendUDDI && uddiBlock.IsNull() {
		diags.AddError(
			"Missing Required Block",
			"The 'uddi' block is required when using the UDDI backend. Use 'uddi = {}' if no attributes needed.",
		)
	}

	// Disallow wrong backend block
	if backend == core.BackendNIOS && !uddiBlock.IsNull() {
		diags.AddError(
			"Invalid Configuration",
			"The 'uddi' block is not allowed when using the NIOS backend.",
		)
	}
	if backend == core.BackendUDDI && !niosBlock.IsNull() {
		diags.AddError(
			"Invalid Configuration",
			"The 'nios' block is not allowed when using the UDDI backend.",
		)
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
func ValidateDataSourceFilters(backend core.BackendType, extAttrFilters, tagFilters types.Map, maxResults, paging types.Int32, diags *diag.Diagnostics) {
	// ext_attr_filters is NIOS only
	if !extAttrFilters.IsNull() && backend == core.BackendUDDI {
		AddBackendFieldError(diags, "ext_attr_filters", "NIOS")
	}

	// max_results is NIOS only
	if !maxResults.IsNull() && backend == core.BackendUDDI {
		AddBackendFieldError(diags, "max_results", "NIOS")
	}

	// paging is NIOS only
	if !paging.IsNull() && backend == core.BackendUDDI {
		AddBackendFieldError(diags, "paging", "NIOS")
	}

	// tag_filters is UDDI only
	if !tagFilters.IsNull() && backend == core.BackendNIOS {
		AddBackendFieldError(diags, "tag_filters", "UDDI")
	}
}

// ValidateListFilters validates backend-specific filter fields for lists.
func ValidateListFilters(backend core.BackendType, extAttrFilters, tagFilters types.Map, diags *diag.Diagnostics) {
	// ext_attr_filters is NIOS only
	if !extAttrFilters.IsNull() && backend == core.BackendUDDI {
		AddBackendFieldError(diags, "ext_attr_filters", "NIOS")
	}

	// tag_filters is UDDI only
	if !tagFilters.IsNull() && backend == core.BackendNIOS {
		AddBackendFieldError(diags, "tag_filters", "UDDI")
	}
}
