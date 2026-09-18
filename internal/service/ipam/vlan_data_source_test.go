package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccVlanDataSource(t *testing.T) {
	dsType := "infoblox_vlan"
	resourceType := "infoblox_vlan"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckVlanExistsNIOS,
			Destroy: testAccCheckVlanDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "ipam/vlan/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
