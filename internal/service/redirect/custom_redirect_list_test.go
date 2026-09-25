package redirect_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccCustomRedirectList(t *testing.T) {
	resourceType := "infoblox_custom_redirect"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckCustomRedirectExistsUDDI,
			Destroy: testAccCheckCustomRedirectDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "redirect/custom_redirect/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
