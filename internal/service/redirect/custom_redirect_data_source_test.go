package redirect_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccCustomRedirectDataSource(t *testing.T) {
	dsType := "infoblox_custom_redirect"
	resourceType := "infoblox_custom_redirect"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckCustomRedirectExistsUDDI,
			Destroy: testAccCheckCustomRedirectDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "redirect/custom_redirect/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
