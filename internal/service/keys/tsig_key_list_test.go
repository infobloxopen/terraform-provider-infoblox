package keys_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccTsigKeyList(t *testing.T) {
	resourceType := "infoblox_tsig_key"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckTsigKeyExistsUDDI,
			Destroy: testAccCheckTsigKeyDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "keys/tsig_key/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
