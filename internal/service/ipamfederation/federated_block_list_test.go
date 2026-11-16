package ipamfederation_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccFederatedBlockList(t *testing.T) {
	resourceType := "infoblox_federated_block"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckFederatedBlockExistsUDDI,
			Destroy: testAccCheckFederatedBlockDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipamfederation/federated_block/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
