package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzCnameIpaddressdnDataSource(t *testing.T) {
	dsType := "infoblox_record_rpz_cname_ipaddressdn"
	resourceType := "infoblox_record_rpz_cname_ipaddressdn"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzCnameIpaddressdnExistsNIOS,
			Destroy: testAccCheckRecordRpzCnameIpaddressdnDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "rpz/record_rpz_cname_ipaddressdn/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
