package anycast_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAnycastConfigList(t *testing.T) {
	resourceType := "infoblox_anycast_config"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckAnycastConfigExistsUDDI,
			Destroy: testAccCheckAnycastConfigDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "anycast/anycast_config/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
