package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccHardwareFilterList(t *testing.T) {
	resourceType := "infoblox_hardware_filter"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckHardwareFilterExistsUDDI,
			Destroy: testAccCheckHardwareFilterDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dhcp/hardware_filter/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
