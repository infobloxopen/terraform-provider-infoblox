package keys_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

// Kerberos keys have no create API: /keys/kerberos exposes no POST, and keys come
// into existence by uploading a keytab to /keys/upload, which yields many keys per
// upload. The object is therefore generated as a data source only, and this test
// cannot use the generated acctest.RunDataSourceCases harness - that creates the
// managed resource and pair-checks the data source against it. Instead it reads
// keys that already exist in the backend, the same way the BloxOne provider's
// kerberos data source test does.
//
// These fixtures are the keytab uploaded to the shared test tenant. If the tenant
// is re-seeded they must be updated to match.
const (
	kerberosKeyFixturePrincipal = "DNS/ns.b1ddi.neo1.com"
	kerberosKeyFixtureDomain    = "NEO1.COM"
	kerberosKeyFixtureTagKey    = "used_for"
	kerberosKeyFixtureTagValue  = "tf_acc_test"
)

func TestAccKerberosKeyDataSource(t *testing.T) {
	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			t.Run("filters", func(t *testing.T) {
				testAccKerberosKeyDataSourceFilters(t, backend)
			})
			t.Run("tag_filters", func(t *testing.T) {
				testAccKerberosKeyDataSourceTagFilters(t, backend)
			})
			t.Run("retrieve_all", func(t *testing.T) {
				testAccKerberosKeyDataSourceRetrieveAll(t, backend)
			})
			t.Run("paging_and_limit", func(t *testing.T) {
				testAccKerberosKeyDataSourcePagingLimit(t, backend)
			})
		})
	}
}

func testAccKerberosKeyDataSourceFilters(t *testing.T, backend string) {
	dsAddr := "data.infoblox_kerberos_key.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t, backend) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config: acctest.ProviderConfigHCL(backend) + fmt.Sprintf(`
data "infoblox_kerberos_key" "test" {
  filters = {
    principal = %q
  }
}
`, kerberosKeyFixturePrincipal),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckKerberosKeyExistsUDDI(dsAddr),
				resource.TestCheckResourceAttrSet(dsAddr, "results.0.id"),
				resource.TestCheckResourceAttr(dsAddr, "results.0.uddi.principal", kerberosKeyFixturePrincipal),
				resource.TestCheckResourceAttr(dsAddr, "results.0.uddi.domain", kerberosKeyFixtureDomain),
				resource.TestCheckResourceAttrSet(dsAddr, "results.0.uddi.algorithm"),
				resource.TestCheckResourceAttrSet(dsAddr, "results.0.uddi.uploaded_at"),
				resource.TestCheckResourceAttrSet(dsAddr, "results.0.uddi.version"),
			),
		}},
	})
}

func testAccKerberosKeyDataSourceTagFilters(t *testing.T, backend string) {
	dsAddr := "data.infoblox_kerberos_key.test_tag"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t, backend) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config: acctest.ProviderConfigHCL(backend) + fmt.Sprintf(`
data "infoblox_kerberos_key" "test_tag" {
  tag_filters = {
    %s = %q
  }
}
`, kerberosKeyFixtureTagKey, kerberosKeyFixtureTagValue),
			Check: resource.ComposeTestCheckFunc(
				testAccCheckKerberosKeyExistsUDDI(dsAddr),
				resource.TestCheckResourceAttrSet(dsAddr, "results.0.id"),
				resource.TestCheckResourceAttr(dsAddr, "results.0.uddi.principal", kerberosKeyFixturePrincipal),
				resource.TestCheckResourceAttr(
					dsAddr,
					fmt.Sprintf("results.0.uddi.tags.%s", kerberosKeyFixtureTagKey),
					kerberosKeyFixtureTagValue,
				),
				resource.TestCheckResourceAttr(
					dsAddr,
					fmt.Sprintf("results.0.uddi.tags_all.%s", kerberosKeyFixtureTagKey),
					kerberosKeyFixtureTagValue,
				),
			),
		}},
	})
}

func testAccKerberosKeyDataSourceRetrieveAll(t *testing.T, backend string) {
	dsAddr := "data.infoblox_kerberos_key.test_all"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t, backend) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config: acctest.ProviderConfigHCL(backend) + `
data "infoblox_kerberos_key" "test_all" {
}
`,
			Check: resource.ComposeTestCheckFunc(
				testAccCheckKerberosKeyExistsUDDI(dsAddr),
				resource.TestCheckResourceAttrSet(dsAddr, "results.0.id"),
			),
		}},
	})
}

// testAccKerberosKeyDataSourcePagingLimit exercises the UDDI-only paging/limit
// attributes: with paging disabled the data source must return exactly one page
// capped at limit, rather than following every page.
func testAccKerberosKeyDataSourcePagingLimit(t *testing.T, backend string) {
	dsAddr := "data.infoblox_kerberos_key.test_page"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t, backend) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config: acctest.ProviderConfigHCL(backend) + `
data "infoblox_kerberos_key" "test_page" {
  paging = 0
  limit  = 2
}
`,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(dsAddr, "results.#", "2"),
				resource.TestCheckResourceAttrSet(dsAddr, "results.0.id"),
			),
		}},
	})
}

// testAccCheckKerberosKeyExistsUDDI verifies the first key the data source
// returned is genuinely readable from the backend. There is deliberately no
// Destroy counterpart: Terraform never owns these keys, so it never destroys them.
func testAccCheckKerberosKeyExistsUDDI(dataSourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[dataSourceName]
		if !ok {
			return fmt.Errorf("not found: %s", dataSourceName)
		}
		id := rs.Primary.Attributes["results.0.id"]
		if id == "" {
			return fmt.Errorf("no results returned for %s", dataSourceName)
		}
		apiRes, _, err := acctest.UDDIClient.KeysAPI.KerberosAPI.
			Read(context.Background(), id).
			Execute()
		if err != nil {
			return fmt.Errorf("failed to read KerberosKey %s: %w", id, err)
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("KerberosKey not found: %s", id)
		}
		return nil
	}
}
