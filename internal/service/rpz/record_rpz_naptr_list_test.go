package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzNaptrList(t *testing.T) {
	resourceType := "infoblox_record_rpz_naptr"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzNaptrExistsNIOS,
			Destroy: testAccCheckRecordRpzNaptrDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "rpz/record_rpz_naptr/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
