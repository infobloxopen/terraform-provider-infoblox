package rpz_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordRpzAList(t *testing.T) {
	resourceType := "infoblox_record_rpz_a"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordRpzAExistsNIOS,
			Destroy: testAccCheckRecordRpzADestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "rpz/record_rpz_a/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
