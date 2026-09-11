package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzCnameIpaddressList(t *testing.T) {
	resourceType := "infoblox_record_rpz_cname_ipaddress"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzCnameIpaddressExistsNIOS,
			Destroy: testAccCheckRecordRpzCnameIpaddressDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "rpz/record_rpz_cname_ipaddress/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
