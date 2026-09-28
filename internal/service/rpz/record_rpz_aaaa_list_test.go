package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzAaaaList(t *testing.T) {
	resourceType := "infoblox_record_rpz_aaaa"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzAaaaExistsNIOS,
			Destroy: testAccCheckRecordRpzAaaaDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "rpz/record_rpz_aaaa/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
