package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSharedrecordTxtList(t *testing.T) {
	resourceType := "infoblox_sharedrecord_txt"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSharedrecordTxtExistsNIOS,
			Destroy: testAccCheckSharedrecordTxtDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/sharedrecord_txt/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
