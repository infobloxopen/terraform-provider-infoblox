package ipamfederation_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccFederatedRealmList(t *testing.T) {
	resourceType := "infoblox_federated_realm"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckFederatedRealmExistsUDDI,
			Destroy: testAccCheckFederatedRealmDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipamfederation/federated_realm/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
