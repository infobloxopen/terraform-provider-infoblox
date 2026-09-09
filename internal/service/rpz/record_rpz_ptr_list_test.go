package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzPtrList(t *testing.T) {
	resourceType := "infoblox_record_rpz_ptr"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzPtrExistsNIOS,
			Destroy: testAccCheckRecordRpzPtrDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "rpz/record_rpz_ptr/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
