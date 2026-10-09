package misc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRulesetDataSource(t *testing.T) {
	dsType := "infoblox_ruleset"
	resourceType := "infoblox_ruleset"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRulesetExistsNIOS,
			Destroy: testAccCheckRulesetDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "misc/ruleset/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
