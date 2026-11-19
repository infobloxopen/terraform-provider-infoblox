package ipamfederation_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

// Hand-written: next-available data sources are not produced by the generator.

func testAccNextAvailableReservedBlockPrereq(realm, parentAddr, tag string) string {
	return acctest.ProviderConfigHCL("uddi") + fmt.Sprintf(`
resource "infoblox_federated_realm" "test" {
  uddi = {
    name = %q
  }
}

resource "infoblox_federated_block" "parent" {
  uddi = {
    address         = %q
    cidr            = 16
    federated_realm = infoblox_federated_realm.test.id
    tags            = { nab = %q }
  }
}
`, realm, parentAddr, tag)
}

func TestAccNextAvailableReservedBlockDataSource_ById(t *testing.T) {
	realm := acctest.RandomNameWithPrefix("nab-realm")
	addr := acctest.RandomIPv4Network()
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t, "uddi") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config: testAccNextAvailableReservedBlockPrereq(realm, addr, "x") + `
data "infoblox_next_available_reserved_block" "test" {
  id                   = infoblox_federated_block.parent.id
  cidr                 = 24
  reserved_block_count = 3
  name                 = "preview"
  comment              = "preview only"
}
`,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("data.infoblox_next_available_reserved_block.test", "results.#", "3"),
				resource.TestCheckResourceAttrPair("data.infoblox_next_available_reserved_block.test", "id", "infoblox_federated_block.parent", "id"),
			),
		}},
	})
}

func TestAccNextAvailableReservedBlockDataSource_DefaultCount(t *testing.T) {
	realm := acctest.RandomNameWithPrefix("nab-realm")
	addr := acctest.RandomIPv4Network()
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t, "uddi") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config: testAccNextAvailableReservedBlockPrereq(realm, addr, "x") + `
data "infoblox_next_available_reserved_block" "test" {
  id   = infoblox_federated_block.parent.id
  cidr = 24
}
`,
			Check: resource.TestCheckResourceAttr("data.infoblox_next_available_reserved_block.test", "results.#", "1"),
		}},
	})
}

func TestAccNextAvailableReservedBlockDataSource_ByTagFilters(t *testing.T) {
	realm := acctest.RandomNameWithPrefix("nab-realm")
	addr := acctest.RandomIPv4Network()
	tag := acctest.RandomNameWithPrefix("tag")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t, "uddi") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config: testAccNextAvailableReservedBlockPrereq(realm, addr, tag) + `
data "infoblox_next_available_reserved_block" "test" {
  depends_on           = [infoblox_federated_block.parent]
  cidr                 = 24
  reserved_block_count = 2
  tag_filters          = { nab = "` + tag + `" }
}
`,
			Check: resource.TestCheckResourceAttr("data.infoblox_next_available_reserved_block.test", "results.#", "2"),
		}},
	})
}

func TestAccNextAvailableReservedBlockDataSource_InvalidId(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t, "uddi") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config: acctest.ProviderConfigHCL("uddi") + `
data "infoblox_next_available_reserved_block" "test" {
  id   = "federation/federated_pool/abc-123"
  cidr = 24
}
`,
			ExpectError: regexp.MustCompile(`invalid resource ID specified`),
		}},
	})
}

func TestAccNextAvailableReservedBlockDataSource_NeedsIdOrTags(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t, "uddi") },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config: acctest.ProviderConfigHCL("uddi") + `
data "infoblox_next_available_reserved_block" "test" {
  cidr = 24
}
`,
			ExpectError: regexp.MustCompile(`Invalid Attribute Combination`),
		}},
	})
}
