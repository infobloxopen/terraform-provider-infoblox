package dtc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDtcServerList(t *testing.T) {
	resourceType := "infoblox_dtc_server"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDtcServerExistsNIOS,
			Destroy: testAccCheckDtcServerDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckDtcServerExistsUDDI,
			Destroy: testAccCheckDtcServerDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dtc/dtc_server/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
