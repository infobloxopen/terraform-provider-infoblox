package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccOptionGroupList(t *testing.T) {
	resourceType := "infoblox_option_group"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckOptionGroupExistsUDDI,
			Destroy: testAccCheckOptionGroupDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dhcp/option_group/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
