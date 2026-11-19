package ipamfederation_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccReservedBlockList(t *testing.T) {
	resourceType := "infoblox_reserved_block"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckReservedBlockExistsUDDI,
			Destroy: testAccCheckReservedBlockDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipamfederation/reserved_block/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
