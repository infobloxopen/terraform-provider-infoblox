package fw_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNamedListList(t *testing.T) {
	resourceType := "infoblox_named_list"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckNamedListExistsUDDI,
			Destroy: testAccCheckNamedListDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "fw/named_list/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
