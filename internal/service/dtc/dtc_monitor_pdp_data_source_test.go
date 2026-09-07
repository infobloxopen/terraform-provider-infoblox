package dtc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDtcMonitorPdpDataSource(t *testing.T) {
	dsType := "infoblox_dtc_monitor_pdp"
	resourceType := "infoblox_dtc_monitor_pdp"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDtcMonitorPdpExistsNIOS,
			Destroy: testAccCheckDtcMonitorPdpDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckDtcMonitorPdpExistsUDDI,
			Destroy: testAccCheckDtcMonitorPdpDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dtc/dtc_monitor_pdp/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
