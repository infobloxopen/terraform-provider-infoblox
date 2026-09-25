package fw_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNetworkListDataSource(t *testing.T) {
	dsType := "infoblox_network_list"
	resourceType := "infoblox_network_list"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckNetworkListExistsUDDI,
			Destroy: testAccCheckNetworkListDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "fw/network_list/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
