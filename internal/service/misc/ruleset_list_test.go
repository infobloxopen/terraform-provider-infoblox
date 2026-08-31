package misc_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRulesetList(t *testing.T) {
	resourceType := "infoblox_ruleset"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRulesetExistsNIOS,
			Destroy: testAccCheckRulesetDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "misc/ruleset/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
