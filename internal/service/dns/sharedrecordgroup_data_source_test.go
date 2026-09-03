package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSharedrecordgroupDataSource(t *testing.T) {
	dsType := "infoblox_sharedrecordgroup"
	resourceType := "infoblox_sharedrecordgroup"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSharedrecordgroupExistsNIOS,
			Destroy: testAccCheckSharedrecordgroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/sharedrecordgroup/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
