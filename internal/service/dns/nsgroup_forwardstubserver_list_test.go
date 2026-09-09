package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNsgroupForwardstubserverList(t *testing.T) {
	resourceType := "infoblox_nsgroup_forwardstubserver"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNsgroupForwardstubserverExistsNIOS,
			Destroy: testAccCheckNsgroupForwardstubserverDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/nsgroup_forwardstubserver/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
