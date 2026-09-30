package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzCnameClientipaddressList(t *testing.T) {
	resourceType := "infoblox_record_rpz_cname_clientipaddress"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzCnameClientipaddressExistsNIOS,
			Destroy: testAccCheckRecordRpzCnameClientipaddressDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "rpz/record_rpz_cname_clientipaddress/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
