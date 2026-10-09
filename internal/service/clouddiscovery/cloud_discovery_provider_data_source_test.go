package clouddiscovery_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccCloudDiscoveryProviderDataSource(t *testing.T) {
	dsType := "infoblox_cloud_discovery_provider"
	resourceType := "infoblox_cloud_discovery_provider"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckCloudDiscoveryProviderExistsUDDI,
			Destroy: testAccCheckCloudDiscoveryProviderDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "clouddiscovery/cloud_discovery_provider/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
