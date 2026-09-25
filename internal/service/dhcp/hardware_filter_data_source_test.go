package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccHardwareFilterDataSource(t *testing.T) {
	dsType := "infoblox_hardware_filter"
	resourceType := "infoblox_hardware_filter"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckHardwareFilterExistsUDDI,
			Destroy: testAccCheckHardwareFilterDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/hardware_filter/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
