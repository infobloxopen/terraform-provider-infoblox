package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordTxtList(t *testing.T) {
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
			acctest.RunListCases(t, resourceType, "dns/record_txt/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
