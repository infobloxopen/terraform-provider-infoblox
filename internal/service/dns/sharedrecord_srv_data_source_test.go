package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSharedrecordSrvDataSource(t *testing.T) {
	dsType := "infoblox_sharedrecord_srv"
	resourceType := "infoblox_sharedrecord_srv"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSharedrecordSrvExistsNIOS,
			Destroy: testAccCheckSharedrecordSrvDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/sharedrecord_srv/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
