package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccVlanrangeDataSource(t *testing.T) {
	dsType := "infoblox_vlanrange"
	resourceType := "infoblox_vlanrange"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckVlanrangeExistsNIOS,
			Destroy: testAccCheckVlanrangeDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "ipam/vlanrange/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
