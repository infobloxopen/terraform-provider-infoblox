package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzCnameIpaddressdnList(t *testing.T) {
	resourceType := "infoblox_record_rpz_cname_ipaddressdn"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzCnameIpaddressdnExistsNIOS,
			Destroy: testAccCheckRecordRpzCnameIpaddressdnDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "rpz/record_rpz_cname_ipaddressdn/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
