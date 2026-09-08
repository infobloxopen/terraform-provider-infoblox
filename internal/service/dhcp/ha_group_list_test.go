package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccHaGroupList(t *testing.T) {
	resourceType := "infoblox_ha_group"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckHaGroupExistsUDDI,
			Destroy: testAccCheckHaGroupDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dhcp/ha_group/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
