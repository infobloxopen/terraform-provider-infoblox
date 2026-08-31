package acctest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func runTest(t *testing.T, parallel bool, tc resource.TestCase) {
	if parallel {
		resource.ParallelTest(t, tc)
	} else {
		resource.Test(t, tc)
	}
}

// RunDataSourceTests runs filter, ext_attr, and tag-filter data source subtests for the given caseFilePath.
func RunDataSourceTests(t *testing.T, dsType, resourceType string, caseFilePath string, checks CheckFuncs) {
	tv, err := LoadCaseFile(caseFilePath)
	if err != nil {
		t.Fatalf("Failed to load case file: %v", err)
	}

	PreCheck(t, tv.Backend)

	t.Run("Filters", func(t *testing.T) {
		runDataSourceFilterTest(t, dsType, resourceType, tv, checks)
	})

	if tv.Backend == "nios" {
		if _, hasExtAttrs := tv.NIOS["ext_attrs"]; hasExtAttrs {
			t.Run("ExtAttrFilters", func(t *testing.T) {
				runDataSourceExtAttrFilterTest(t, dsType, resourceType, tv, checks)
			})
		}
	}

	if tv.Backend == "uddi" {
		if _, hasTags := tv.UDDI["tags"]; hasTags {
			t.Run("TagFilters", func(t *testing.T) {
				runDataSourceTagFilterTest(t, dsType, resourceType, tv, checks)
			})
		}
	}
}

func runDataSourceFilterTest(t *testing.T, dsType, resourceType string, tv *CaseConfig, checks CheckFuncs) {
	resourceAddr := resourceType + ".test"
	dsAddr := "data." + dsType + ".test"

	resourceConfig := BuildResourceHCL(resourceType, "test", tv)

	// Use ds_filter_field from case file, default to "name"
	filterField := "name"
	if tv.DSFilterField != "" {
		filterField = tv.DSFilterField
	}

	dsConfig := fmt.Sprintf(`
data %q "test" {
  filters = {
    %s = %s.test.%s
  }
  depends_on = [%s.test]
}
`, dsType, filterField, resourceType, filterField, resourceType)

	providerConfig := ProviderConfigHCL(tv.Backend)
	fullConfig := providerConfig + "\n" + tv.PrerequisitesHCL + "\n" + resourceConfig + "\n" + dsConfig

	t.Logf("Generated HCL config:\n%s", fullConfig)

	var checkFuncs []resource.TestCheckFunc
	if checks.Exists != nil {
		checkFuncs = append(checkFuncs, checks.Exists(resourceAddr))
	}

	checkFuncs = append(checkFuncs, resource.TestCheckResourceAttrSet(dsAddr, "results.0.id"))
	checkFuncs = append(checkFuncs, buildDataSourceAttrPairChecks(dsAddr, resourceAddr, tv)...)

	runTest(t, tv.Parallel, resource.TestCase{
		ProtoV6ProviderFactories: ProtoV6ProviderFactories,
		CheckDestroy:             checks.Destroy(resourceType),
		Steps: []resource.TestStep{{
			Config: fullConfig,
			Check:  resource.ComposeTestCheckFunc(checkFuncs...),
		}},
	})
}

func runDataSourceExtAttrFilterTest(t *testing.T, dsType, resourceType string, tv *CaseConfig, checks CheckFuncs) {
	resourceAddr := resourceType + ".test"
	dsAddr := "data." + dsType + ".test"

	resourceConfig := BuildResourceHCL(resourceType, "test", tv)

	dsConfig := fmt.Sprintf(`
data %q "test" {
  extattrfilters = {
    Site = %s.test.nios.ext_attrs.Site
  }
  depends_on = [%s.test]
}
`, dsType, resourceType, resourceType)

	providerConfig := ProviderConfigHCL(tv.Backend)
	fullConfig := providerConfig + "\n" + tv.PrerequisitesHCL + "\n" + resourceConfig + "\n" + dsConfig

	t.Logf("Generated HCL config:\n%s", fullConfig)

	var checkFuncs []resource.TestCheckFunc
	if checks.Exists != nil {
		checkFuncs = append(checkFuncs, checks.Exists(resourceAddr))
	}

	checkFuncs = append(checkFuncs, resource.TestCheckResourceAttrSet(dsAddr, "results.0.id"))
	checkFuncs = append(checkFuncs, buildDataSourceAttrPairChecks(dsAddr, resourceAddr, tv)...)

	runTest(t, tv.Parallel, resource.TestCase{
		ProtoV6ProviderFactories: ProtoV6ProviderFactories,
		CheckDestroy:             checks.Destroy(resourceType),
		Steps: []resource.TestStep{{
			Config: fullConfig,
			Check:  resource.ComposeTestCheckFunc(checkFuncs...),
		}},
	})
}

func runDataSourceTagFilterTest(t *testing.T, dsType, resourceType string, tv *CaseConfig, checks CheckFuncs) {
	resourceAddr := resourceType + ".test"
	dsAddr := "data." + dsType + ".test"

	resourceConfig := BuildResourceHCL(resourceType, "test", tv)

	dsConfig := fmt.Sprintf(`
data %q "test" {
  tag_filters = {
    tag1 = %s.test.uddi.tags.tag1
  }
  depends_on = [%s.test]
}
`, dsType, resourceType, resourceType)

	providerConfig := ProviderConfigHCL(tv.Backend)
	fullConfig := providerConfig + "\n" + tv.PrerequisitesHCL + "\n" + resourceConfig + "\n" + dsConfig

	t.Logf("Generated HCL config:\n%s", fullConfig)

	var checkFuncs []resource.TestCheckFunc
	if checks.Exists != nil {
		checkFuncs = append(checkFuncs, checks.Exists(resourceAddr))
	}

	checkFuncs = append(checkFuncs, resource.TestCheckResourceAttrSet(dsAddr, "results.0.id"))
	checkFuncs = append(checkFuncs, buildDataSourceAttrPairChecks(dsAddr, resourceAddr, tv)...)

	runTest(t, tv.Parallel, resource.TestCase{
		ProtoV6ProviderFactories: ProtoV6ProviderFactories,
		CheckDestroy:             checks.Destroy(resourceType),
		Steps: []resource.TestStep{{
			Config: fullConfig,
			Check:  resource.ComposeTestCheckFunc(checkFuncs...),
		}},
	})
}

// buildDataSourceAttrPairChecks returns attr-pair checks asserting the data source results match the resource.
func buildDataSourceAttrPairChecks(dsAddr, resourceAddr string, tv *CaseConfig) []resource.TestCheckFunc {
	var checks []resource.TestCheckFunc

	for k := range tv.Common {
		checks = append(checks, resource.TestCheckResourceAttrPair(
			dsAddr, "results.0."+k,
			resourceAddr, k,
		))
	}

	if tv.Backend == "nios" {
		for k := range tv.NIOS {
			checks = append(checks, resource.TestCheckResourceAttrPair(
				dsAddr, "results.0.nios."+k,
				resourceAddr, "nios."+k,
			))
		}
	}
	if tv.Backend == "uddi" {
		for k := range tv.UDDI {
			checks = append(checks, resource.TestCheckResourceAttrPair(
				dsAddr, "results.0.uddi."+k,
				resourceAddr, "uddi."+k,
			))
		}
	}

	return checks
}
