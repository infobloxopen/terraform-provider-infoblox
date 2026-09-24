package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzPtrDataSource(t *testing.T) {
	dsType := "infoblox_record_rpz_ptr"
	resourceType := "infoblox_record_rpz_ptr"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzPtrExistsNIOS,
			Destroy: testAccCheckRecordRpzPtrDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "rpz/record_rpz_ptr/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
