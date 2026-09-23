package dtc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDtcMonitorSnmpDataSource(t *testing.T) {
	dsType := "infoblox_dtc_monitor_snmp"
	resourceType := "infoblox_dtc_monitor_snmp"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDtcMonitorSnmpExistsNIOS,
			Destroy: testAccCheckDtcMonitorSnmpDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckDtcMonitorSnmpExistsUDDI,
			Destroy: testAccCheckDtcMonitorSnmpDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dtc/dtc_monitor_snmp/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
