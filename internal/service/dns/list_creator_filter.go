package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// defaultStaticCreatorFilter limits NIOS NS and PTR list results to STATIC records unless the user
// filters on creator themselves. NIOS also returns SYSTEM records it generates (zone NS records,
// auto-created PTRs), which Terraform cannot manage, so Terraform Search would otherwise emit
// configuration for objects that fail on apply.
func defaultStaticCreatorFilter(backend core.BackendType, filters map[string]string) map[string]string {
	if backend != core.BackendNIOS {
		return filters
	}
	if _, ok := filters["creator"]; ok {
		return filters
	}
	if _, ok := filters["nios.creator"]; ok {
		return filters
	}
	if filters == nil {
		filters = make(map[string]string, 1)
	}
	filters["creator"] = "STATIC"
	return filters
}
