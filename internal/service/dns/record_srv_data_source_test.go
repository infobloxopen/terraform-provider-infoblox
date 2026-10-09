package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordSrvDataSource(t *testing.T) {
	dsType := "infoblox_record_srv"
	resourceType := "infoblox_record_srv"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordSrvExistsNIOS,
			Destroy: testAccCheckRecordSrvDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordSrvExistsUDDI,
			Destroy: testAccCheckRecordSrvDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/record_srv/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
