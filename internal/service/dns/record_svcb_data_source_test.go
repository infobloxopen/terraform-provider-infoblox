package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordSvcbDataSource(t *testing.T) {
	dsType := "infoblox_record_svcb"
	resourceType := "infoblox_record_svcb"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckRecordSvcbExistsUDDI,
			Destroy: testAccCheckRecordSvcbDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/record_svcb/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
