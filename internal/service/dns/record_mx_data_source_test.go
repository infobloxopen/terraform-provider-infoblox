package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordMxDataSource(t *testing.T) {
	dsType := "infoblox_record_mx"
	resourceType := "infoblox_record_mx"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordMxExistsNIOS,
			Destroy: testAccCheckRecordMxDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordMxExistsUDDI,
			Destroy: testAccCheckRecordMxDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/record_mx/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
