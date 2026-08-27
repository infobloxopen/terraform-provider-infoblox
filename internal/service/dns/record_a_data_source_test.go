package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordADataSource(t *testing.T) {
	dsType := "infoblox_record_a"
	resourceType := "infoblox_record_a"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordAExistsNIOS,
			Destroy: testAccCheckRecordADestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordAExistsUDDI,
			Destroy: testAccCheckRecordADestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/record_a/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
