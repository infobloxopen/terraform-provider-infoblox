package grid_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccExtensibleattributedefDataSource(t *testing.T) {
	dsType := "infoblox_extensible_attribute_def"
	resourceType := "infoblox_extensible_attribute_def"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckExtensibleattributedefExistsNIOS,
			Destroy: testAccCheckExtensibleattributedefDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "grid/extensible_attribute_def/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
