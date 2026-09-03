package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNsgroupForwardingmemberDataSource(t *testing.T) {
	dsType := "infoblox_nsgroup_forwardingmember"
	resourceType := "infoblox_nsgroup_forwardingmember"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNsgroupForwardingmemberExistsNIOS,
			Destroy: testAccCheckNsgroupForwardingmemberDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/nsgroup_forwardingmember/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
