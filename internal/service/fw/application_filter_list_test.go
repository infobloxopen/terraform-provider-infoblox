package fw_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccApplicationFilterList(t *testing.T) {
	resourceType := "infoblox_application_filter"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckApplicationFilterExistsUDDI,
			Destroy: testAccCheckApplicationFilterDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "fw/application_filter/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
