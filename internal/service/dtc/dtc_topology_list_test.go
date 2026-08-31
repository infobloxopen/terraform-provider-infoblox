package dtc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDtcTopologyList(t *testing.T) {
	resourceType := "infoblox_dtc_topology"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDtcTopologyExistsNIOS,
			Destroy: testAccCheckDtcTopologyDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dtc/dtc_topology/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
