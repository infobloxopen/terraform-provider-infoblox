package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSuperhostDataSource(t *testing.T) {
	dsType := "infoblox_superhost"
	resourceType := "infoblox_superhost"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSuperhostExistsNIOS,
			Destroy: testAccCheckSuperhostDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "ipam/superhost/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
