package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzTxtList(t *testing.T) {
	resourceType := "infoblox_record_rpz_txt"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzTxtExistsNIOS,
			Destroy: testAccCheckRecordRpzTxtDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "rpz/record_rpz_txt/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
