package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzCnameClientipaddressdnList(t *testing.T) {
	resourceType := "infoblox_record_rpz_cname_clientipaddressdn"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzCnameClientipaddressdnExistsNIOS,
			Destroy: testAccCheckRecordRpzCnameClientipaddressdnDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "rpz/record_rpz_cname_clientipaddressdn/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
