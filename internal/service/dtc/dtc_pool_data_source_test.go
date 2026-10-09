package dtc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDtcPoolDataSource(t *testing.T) {
	dsType := "infoblox_dtc_pool"
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
			acctest.RunDataSourceCases(t, dsType, resourceType, "dtc/dtc_pool/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
