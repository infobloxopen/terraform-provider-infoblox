package misc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccBfdtemplateList(t *testing.T) {
	resourceType := "infoblox_bfd_template"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckBfdtemplateExistsNIOS,
			Destroy: testAccCheckBfdtemplateDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "misc/bfd_template/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
