package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccBulkhostnametemplateList(t *testing.T) {
	resourceType := "infoblox_bulk_hostname_template"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckBulkhostnametemplateExistsNIOS,
			Destroy: testAccCheckBulkhostnametemplateDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipam/bulk_hostname_template/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
