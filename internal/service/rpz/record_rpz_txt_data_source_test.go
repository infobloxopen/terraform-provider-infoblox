package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzTxtDataSource(t *testing.T) {
	dsType := "infoblox_record_rpz_txt"
	resourceType := "infoblox_record_rpz_txt"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzTxtExistsNIOS,
			Destroy: testAccCheckRecordRpzTxtDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "rpz/record_rpz_txt/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
