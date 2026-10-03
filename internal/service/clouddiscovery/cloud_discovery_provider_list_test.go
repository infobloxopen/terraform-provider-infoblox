package clouddiscovery_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccCloudDiscoveryProviderList(t *testing.T) {
	resourceType := "infoblox_cloud_discovery_provider"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckCloudDiscoveryProviderExistsUDDI,
			Destroy: testAccCheckCloudDiscoveryProviderDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "clouddiscovery/cloud_discovery_provider/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
