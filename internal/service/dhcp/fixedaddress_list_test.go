package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccFixedaddressList(t *testing.T) {
	resourceType := "infoblox_fixed_address"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckFixedaddressExistsNIOS,
			Destroy: testAccCheckFixedaddressDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckFixedaddressExistsUDDI,
			Destroy: testAccCheckFixedaddressDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dhcp/fixed_address/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
