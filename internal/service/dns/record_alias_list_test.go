package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordAliasList(t *testing.T) {
	resourceType := "infoblox_record_alias"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordAliasExistsNIOS,
			Destroy: testAccCheckRecordAliasDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/record_alias/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
