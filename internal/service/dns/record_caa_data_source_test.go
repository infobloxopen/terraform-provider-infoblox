package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordCaaDataSource(t *testing.T) {
	dsType := "infoblox_record_caa"
	resourceType := "infoblox_record_caa"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordCaaExistsNIOS,
			Destroy: testAccCheckRecordCaaDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordCaaExistsUDDI,
			Destroy: testAccCheckRecordCaaDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/record_caa/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
