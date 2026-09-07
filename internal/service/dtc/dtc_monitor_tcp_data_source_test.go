package dtc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDtcMonitorTcpDataSource(t *testing.T) {
	dsType := "infoblox_dtc_monitor_tcp"
	resourceType := "infoblox_dtc_monitor_tcp"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDtcMonitorTcpExistsNIOS,
			Destroy: testAccCheckDtcMonitorTcpDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckDtcMonitorTcpExistsUDDI,
			Destroy: testAccCheckDtcMonitorTcpDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dtc/dtc_monitor_tcp/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
