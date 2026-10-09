package infra_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccJoinTokenList(t *testing.T) {
	resourceType := "infoblox_join_token"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckJoinTokenExistsUDDI,
			Destroy: testAccCheckJoinTokenDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "infra/join_token/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
