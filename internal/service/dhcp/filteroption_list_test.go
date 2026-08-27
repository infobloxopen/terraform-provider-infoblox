package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccFilteroptionList(t *testing.T) {
	resourceType := "infoblox_filteroption"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckFilteroptionExistsNIOS,
			Destroy: testAccCheckFilteroptionDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckFilteroptionExistsUDDI,
			Destroy: testAccCheckFilteroptionDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dhcp/filteroption/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
