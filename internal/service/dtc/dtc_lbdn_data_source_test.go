package dtc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDtcLbdnDataSource(t *testing.T) {
	dsType := "infoblox_dtc_lbdn"
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
			acctest.RunDataSourceCases(t, dsType, resourceType, "dtc/dtc_lbdn/"+backend+"_datasources.tfvars", checksByBackend)
		})
	}
}
