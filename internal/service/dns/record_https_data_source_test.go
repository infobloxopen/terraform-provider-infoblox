package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordHttpsDataSource(t *testing.T) {
	dsType := "infoblox_record_https"
	resourceType := "infoblox_record_https"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckRecordHttpsExistsUDDI,
			Destroy: testAccCheckRecordHttpsDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/record_https/"+backend+"_datasources.tfvars", checksByBackend)
		})
	}
}
