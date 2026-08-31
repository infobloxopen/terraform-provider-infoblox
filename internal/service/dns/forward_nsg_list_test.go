package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccForwardNsgList(t *testing.T) {
	resourceType := "infoblox_forward_nsg"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckForwardNsgExistsUDDI,
			Destroy: testAccCheckForwardNsgDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/forward_nsg/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
