package ipamfederation_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccFederatedRealmDataSource(t *testing.T) {
	dsType := "infoblox_federated_realm"
	resourceType := "infoblox_federated_realm"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckFederatedRealmExistsUDDI,
			Destroy: testAccCheckFederatedRealmDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "ipamfederation/federated_realm/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
