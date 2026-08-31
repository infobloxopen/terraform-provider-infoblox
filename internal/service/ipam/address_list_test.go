package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAddressList(t *testing.T) {
	resourceType := "infoblox_address"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckAddressExistsUDDI,
			Destroy: testAccCheckAddressDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipam/address/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
