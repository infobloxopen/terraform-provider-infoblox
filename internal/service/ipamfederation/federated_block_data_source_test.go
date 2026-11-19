package ipamfederation_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccFederatedBlockDataSource(t *testing.T) {
	dsType := "infoblox_federated_block"
	resourceType := "infoblox_federated_block"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckFederatedBlockExistsUDDI,
			Destroy: testAccCheckFederatedBlockDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "ipamfederation/federated_block/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
