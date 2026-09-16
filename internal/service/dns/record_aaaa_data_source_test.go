package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordAaaaDataSource(t *testing.T) {
	dsType := "infoblox_record_aaaa"
	resourceType := "infoblox_record_aaaa"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordAaaaExistsNIOS,
			Destroy: testAccCheckRecordAaaaDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordAaaaExistsUDDI,
			Destroy: testAccCheckRecordAaaaDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/record_aaaa/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
