package ipamfederation_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccReservedBlockDataSource(t *testing.T) {
	dsType := "infoblox_reserved_block"
	resourceType := "infoblox_reserved_block"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckReservedBlockExistsUDDI,
			Destroy: testAccCheckReservedBlockDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "ipamfederation/reserved_block/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
