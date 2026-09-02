package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSharedrecordAaaaDataSource(t *testing.T) {
	dsType := "infoblox_sharedrecord_aaaa"
	resourceType := "infoblox_sharedrecord_aaaa"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSharedrecordAaaaExistsNIOS,
			Destroy: testAccCheckSharedrecordAaaaDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/sharedrecord_aaaa/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
