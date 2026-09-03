package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzAaaaIpaddressList(t *testing.T) {
	resourceType := "infoblox_record_rpz_aaaa_ipaddress"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzAaaaIpaddressExistsNIOS,
			Destroy: testAccCheckRecordRpzAaaaIpaddressDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "rpz/record_rpz_aaaa_ipaddress/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
