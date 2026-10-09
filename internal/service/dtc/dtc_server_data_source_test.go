package dtc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDtcServerDataSource(t *testing.T) {
	dsType := "infoblox_dtc_server"
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
			acctest.RunDataSourceCases(t, dsType, resourceType, "dtc/dtc_server/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
