package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzCnameDataSource(t *testing.T) {
	dsType := "infoblox_record_rpz_cname"
	resourceType := "infoblox_record_rpz_cname"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzCnameExistsNIOS,
			Destroy: testAccCheckRecordRpzCnameDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "rpz/record_rpz_cname/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
