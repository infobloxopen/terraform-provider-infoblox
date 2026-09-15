package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzAaaaDataSource(t *testing.T) {
	dsType := "infoblox_record_rpz_aaaa"
	resourceType := "infoblox_record_rpz_aaaa"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzAaaaExistsNIOS,
			Destroy: testAccCheckRecordRpzAaaaDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "rpz/record_rpz_aaaa/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
