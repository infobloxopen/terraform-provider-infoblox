package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNsgroupForwardstubserverDataSource(t *testing.T) {
	dsType := "infoblox_nsgroup_forwardstubserver"
	resourceType := "infoblox_nsgroup_forwardstubserver"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNsgroupForwardstubserverExistsNIOS,
			Destroy: testAccCheckNsgroupForwardstubserverDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/nsgroup_forwardstubserver/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
