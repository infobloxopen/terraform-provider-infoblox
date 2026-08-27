package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRecordAliasDataSource(t *testing.T) {
	dsType := "infoblox_record_alias"
	resourceType := "infoblox_record_alias"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRecordAliasExistsNIOS,
			Destroy: testAccCheckRecordAliasDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/record_alias/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
