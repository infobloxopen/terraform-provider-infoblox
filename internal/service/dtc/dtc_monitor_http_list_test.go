package dtc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDtcMonitorHttpList(t *testing.T) {
	resourceType := "infoblox_dtc_monitor_http"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDtcMonitorHttpExistsNIOS,
			Destroy: testAccCheckDtcMonitorHttpDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckDtcMonitorHttpExistsUDDI,
			Destroy: testAccCheckDtcMonitorHttpDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dtc/dtc_monitor_http/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
