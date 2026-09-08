package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccHaGroupDataSource(t *testing.T) {
	dsType := "infoblox_ha_group"
	resourceType := "infoblox_ha_group"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckHaGroupExistsUDDI,
			Destroy: testAccCheckHaGroupDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/ha_group/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
