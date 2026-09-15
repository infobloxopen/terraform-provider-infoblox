package anycast_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAnycastHostList(t *testing.T) {
	resourceType := "infoblox_anycast_host"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckAnycastHostExistsUDDI,
			Destroy: testAccCheckAnycastHostDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "anycast/anycast_host/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
