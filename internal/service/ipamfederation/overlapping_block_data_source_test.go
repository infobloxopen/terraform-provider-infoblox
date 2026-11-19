package ipamfederation_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccOverlappingBlockDataSource(t *testing.T) {
	dsType := "infoblox_overlapping_block"
	resourceType := "infoblox_overlapping_block"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckOverlappingBlockExistsUDDI,
			Destroy: testAccCheckOverlappingBlockDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "ipamfederation/overlapping_block/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
