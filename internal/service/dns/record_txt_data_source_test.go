package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordTxtDataSource(t *testing.T) {
	dsType := "infoblox_record_txt"
	resourceType := "infoblox_record_txt"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordTxtExistsNIOS,
			Destroy: testAccCheckRecordTxtDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRecordTxtExistsUDDI,
			Destroy: testAccCheckRecordTxtDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/record_txt/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
