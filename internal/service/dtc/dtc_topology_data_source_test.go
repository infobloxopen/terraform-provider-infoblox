package dtc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDtcTopologyDataSource(t *testing.T) {
	dsType := "infoblox_dtc_topology"
	resourceType := "infoblox_dtc_topology"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDtcTopologyExistsNIOS,
			Destroy: testAccCheckDtcTopologyDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dtc/dtc_topology/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
