package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccVlanviewDataSource(t *testing.T) {
	dsType := "infoblox_vlan_view"
	resourceType := "infoblox_vlan_view"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckVlanviewExistsNIOS,
			Destroy: testAccCheckVlanviewDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "ipam/vlan_view/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
