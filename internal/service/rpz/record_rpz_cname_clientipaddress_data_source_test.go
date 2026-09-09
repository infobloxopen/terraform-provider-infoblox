package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzCnameClientipaddressDataSource(t *testing.T) {
	dsType := "infoblox_record_rpz_cname_clientipaddress"
	resourceType := "infoblox_record_rpz_cname_clientipaddress"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzCnameClientipaddressExistsNIOS,
			Destroy: testAccCheckRecordRpzCnameClientipaddressDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "rpz/record_rpz_cname_clientipaddress/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
