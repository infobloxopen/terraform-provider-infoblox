package infra_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccJoinTokenDataSource(t *testing.T) {
	dsType := "infoblox_join_token"
	resourceType := "infoblox_join_token"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckJoinTokenExistsUDDI,
			Destroy: testAccCheckJoinTokenDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "infra/join_token/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
