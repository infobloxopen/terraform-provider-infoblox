package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzNaptrDataSource(t *testing.T) {
	dsType := "infoblox_record_rpz_naptr"
	resourceType := "infoblox_record_rpz_naptr"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzNaptrExistsNIOS,
			Destroy: testAccCheckRecordRpzNaptrDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "rpz/record_rpz_naptr/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
