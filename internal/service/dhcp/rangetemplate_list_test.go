package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRangetemplateList(t *testing.T) {
	resourceType := "infoblox_range_template"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRangetemplateExistsNIOS,
			Destroy: testAccCheckRangetemplateDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dhcp/range_template/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
