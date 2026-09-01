package keys_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccTsigKeyDataSource(t *testing.T) {
	dsType := "infoblox_tsig_key"
	resourceType := "infoblox_tsig_key"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckTsigKeyExistsUDDI,
			Destroy: testAccCheckTsigKeyDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "keys/tsig_key/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
