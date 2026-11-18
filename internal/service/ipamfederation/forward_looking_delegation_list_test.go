package ipamfederation_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccForwardLookingDelegationList(t *testing.T) {
	resourceType := "infoblox_forward_looking_delegation"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckForwardLookingDelegationExistsUDDI,
			Destroy: testAccCheckForwardLookingDelegationDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipamfederation/forward_looking_delegation/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
