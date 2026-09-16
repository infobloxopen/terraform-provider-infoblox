package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNsgroupForwardingmemberList(t *testing.T) {
	resourceType := "infoblox_nsgroup_forwardingmember"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNsgroupForwardingmemberExistsNIOS,
			Destroy: testAccCheckNsgroupForwardingmemberDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/nsgroup_forwardingmember/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
