package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordPtrDataSource(t *testing.T) {
	dsType := "infoblox_record_ptr"
	resourceType := "infoblox_record_ptr"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordPtrExistsNIOS,
			Destroy: testAccCheckRecordPtrDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordPtrExistsUDDI,
			Destroy: testAccCheckRecordPtrDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/record_ptr/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
