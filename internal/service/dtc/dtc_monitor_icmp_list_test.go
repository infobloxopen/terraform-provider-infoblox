package dtc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDtcMonitorIcmpList(t *testing.T) {
	resourceType := "infoblox_dtc_monitor_icmp"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDtcMonitorIcmpExistsNIOS,
			Destroy: testAccCheckDtcMonitorIcmpDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckDtcMonitorIcmpExistsUDDI,
			Destroy: testAccCheckDtcMonitorIcmpDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dtc/dtc_monitor_icmp/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
