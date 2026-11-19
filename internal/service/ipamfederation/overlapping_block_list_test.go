package ipamfederation_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccOverlappingBlockList(t *testing.T) {
	resourceType := "infoblox_overlapping_block"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckOverlappingBlockExistsUDDI,
			Destroy: testAccCheckOverlappingBlockDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipamfederation/overlapping_block/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
