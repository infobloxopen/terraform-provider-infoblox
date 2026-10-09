package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzADataSource(t *testing.T) {
	dsType := "infoblox_record_rpz_a"
	resourceType := "infoblox_record_rpz_a"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzAExistsNIOS,
			Destroy: testAccCheckRecordRpzADestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "rpz/record_rpz_a/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
