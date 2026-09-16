package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordCnameDataSource(t *testing.T) {
	dsType := "infoblox_record_cname"
	resourceType := "infoblox_record_cname"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordCnameExistsNIOS,
			Destroy: testAccCheckRecordCnameDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordCnameExistsUDDI,
			Destroy: testAccCheckRecordCnameDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/record_cname/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
