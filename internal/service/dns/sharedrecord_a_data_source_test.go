package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSharedrecordADataSource(t *testing.T) {
	dsType := "infoblox_sharedrecord_a"
	resourceType := "infoblox_sharedrecord_a"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSharedrecordAExistsNIOS,
			Destroy: testAccCheckSharedrecordADestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/sharedrecord_a/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
