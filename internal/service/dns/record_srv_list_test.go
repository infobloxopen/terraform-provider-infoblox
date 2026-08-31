package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordSrvList(t *testing.T) {
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
			acctest.RunListCases(t, resourceType, "dns/record_srv/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
