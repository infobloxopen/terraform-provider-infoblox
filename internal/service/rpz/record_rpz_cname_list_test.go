package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzCnameList(t *testing.T) {
	resourceType := "infoblox_record_rpz_cname"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzCnameExistsNIOS,
			Destroy: testAccCheckRecordRpzCnameDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "rpz/record_rpz_cname/"+backend+"_lists.tfvars", checksByBackend)
		})
	}
}
