package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSharednetworkDataSource(t *testing.T) {
	dsType := "infoblox_sharednetwork"
	resourceType := "infoblox_sharednetwork"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSharednetworkExistsNIOS,
			Destroy: testAccCheckSharednetworkDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/sharednetwork/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
