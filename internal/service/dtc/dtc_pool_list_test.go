package dtc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDtcPoolList(t *testing.T) {
	resourceType := "infoblox_dtc_pool"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDtcPoolExistsNIOS,
			Destroy: testAccCheckDtcPoolDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckDtcPoolExistsUDDI,
			Destroy: testAccCheckDtcPoolDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dtc/dtc_pool/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
