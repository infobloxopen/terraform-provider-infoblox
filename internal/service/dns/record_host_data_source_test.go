package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordHostDataSource(t *testing.T) {
	dsType := "infoblox_record_host"
	resourceType := "infoblox_record_host"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordHostExistsNIOS,
			Destroy: testAccCheckRecordHostDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/record_host/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
