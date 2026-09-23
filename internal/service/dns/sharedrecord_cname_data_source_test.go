package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSharedrecordCnameDataSource(t *testing.T) {
	dsType := "infoblox_sharedrecord_cname"
	resourceType := "infoblox_sharedrecord_cname"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSharedrecordCnameExistsNIOS,
			Destroy: testAccCheckSharedrecordCnameDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/sharedrecord_cname/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
