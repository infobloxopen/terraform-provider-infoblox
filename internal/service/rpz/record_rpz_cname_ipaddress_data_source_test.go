package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzCnameIpaddressDataSource(t *testing.T) {
	dsType := "infoblox_record_rpz_cname_ipaddress"
	resourceType := "infoblox_record_rpz_cname_ipaddress"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzCnameIpaddressExistsNIOS,
			Destroy: testAccCheckRecordRpzCnameIpaddressDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "rpz/record_rpz_cname_ipaddress/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
