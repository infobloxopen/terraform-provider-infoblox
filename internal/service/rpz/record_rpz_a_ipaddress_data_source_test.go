package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzAIpaddressDataSource(t *testing.T) {
	dsType := "infoblox_record_rpz_a_ipaddress"
	resourceType := "infoblox_record_rpz_a_ipaddress"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzAIpaddressExistsNIOS,
			Destroy: testAccCheckRecordRpzAIpaddressDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "rpz/record_rpz_a_ipaddress/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
