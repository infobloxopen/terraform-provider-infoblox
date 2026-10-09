package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNsgroupStubmemberList(t *testing.T) {
	resourceType := "infoblox_nsgroup_stubmember"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNsgroupStubmemberExistsNIOS,
			Destroy: testAccCheckNsgroupStubmemberDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/nsgroup_stubmember/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
