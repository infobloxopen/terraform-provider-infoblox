package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordNsDataSource(t *testing.T) {
	dsType := "infoblox_record_ns"
	resourceType := "infoblox_record_ns"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordNsExistsNIOS,
			Destroy: testAccCheckRecordNsDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordNsExistsUDDI,
			Destroy: testAccCheckRecordNsDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/record_ns/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
