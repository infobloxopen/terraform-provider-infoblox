package fw_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAccessCodeList(t *testing.T) {
	resourceType := "infoblox_access_code"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckAccessCodeExistsUDDI,
			Destroy: testAccCheckAccessCodeDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "fw/access_code/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
