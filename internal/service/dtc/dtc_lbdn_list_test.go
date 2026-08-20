package dtc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDtcLbdnList(t *testing.T) {
	resourceType := "infoblox_dtc_lbdn"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDtcLbdnExistsNIOS,
			Destroy: testAccCheckDtcLbdnDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckDtcLbdnExistsUDDI,
			Destroy: testAccCheckDtcLbdnDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dtc/dtc_lbdn/"+backend+"_lists.tfvars", checksByBackend)
		})
	}
}
