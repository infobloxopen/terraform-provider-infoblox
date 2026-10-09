package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordSvcbList(t *testing.T) {
	resourceType := "infoblox_record_svcb"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckRecordSvcbExistsUDDI,
			Destroy: testAccCheckRecordSvcbDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/record_svcb/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
