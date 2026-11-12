package migration_test

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	tfresource "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

// Each test creates objects with a released provider, then expects an empty first plan with this one.
// Run without dev_overrides, so the old release is downloaded.

// ttl, quorum and delegated_ttl stay unset: v2 sends them as 0.
func TestAccV2ToV3Upgrade(t *testing.T) {
	name := acctest.RandomNameWithPrefix("mig")
	zone := name + ".test"

	tfresource.ParallelTest(t, tfresource.TestCase{
		PreCheck:               func() { preCheck(t) },
		TerraformVersionChecks: movedBlockVersion(),
		Steps: []tfresource.TestStep{
			v2Step("2.12.0", v2Config(name, zone)),
			v3Step(v3Config(name, zone)),
		},
	})
}

// State before v2.6.0 has no internal_id: the move fails until a newer v2 apply adds it.
func TestAccV2ToV3UpgradeFromPreV26(t *testing.T) {
	name := acctest.RandomNameWithPrefix("mig")
	zone := name + ".test"
	v2 := v2Provider() + fmt.Sprintf(`
resource "infoblox_zone_auth" "zone" {
  fqdn = %[1]q
}

resource "infoblox_a_record" "a" {
  fqdn      = "a.${infoblox_zone_auth.zone.fqdn}"
  ip_addr   = "10.99.1.20"
  ext_attrs = jsonencode({ Site = "Blr" })
}
`, zone)
	v3 := fmt.Sprintf(`
resource "infoblox_zone_auth" "zone" {
  nios = {
    fqdn = %[1]q
  }
}

moved {
  from = infoblox_a_record.a
  to   = infoblox_record_a.a
}

resource "infoblox_record_a" "a" {
  nios = {
    name      = "a.${infoblox_zone_auth.zone.nios.fqdn}"
    ipv4addr  = "10.99.1.20"
    ext_attrs = { Site = "Blr" }
  }
}
`, zone)

	failing := v3Step(v3)
	failing.ConfigPlanChecks = tfresource.ConfigPlanChecks{}
	failing.ExpectError = regexp.MustCompile(`Missing Internal ID`)

	tfresource.ParallelTest(t, tfresource.TestCase{
		PreCheck:               func() { preCheck(t) },
		TerraformVersionChecks: movedBlockVersion(),
		Steps: []tfresource.TestStep{
			v2Step("2.5.0", v2),
			failing,
			// The fix the error suggests.
			v2Step("2.13.0", v2),
			v3Step(v3),
		},
	})
}

// 3.0.0-beta state is also version 0 but already v3-shaped, so it must be kept as is.
func TestAccBetaToCurrent(t *testing.T) {
	zone := acctest.RandomNameWithPrefix("mig") + ".test"
	config := fmt.Sprintf(`
resource "infoblox_zone_auth" "zone" {
  nios = {
    fqdn = %[1]q
  }
}

resource "infoblox_record_a" "a" {
  nios = {
    name      = "a.${infoblox_zone_auth.zone.nios.fqdn}"
    ipv4addr  = "10.99.1.30"
    ext_attrs = { Site = "Blr" }
  }
}
`, zone)

	tfresource.ParallelTest(t, tfresource.TestCase{
		PreCheck: func() { preCheck(t) },
		Steps: []tfresource.TestStep{
			{
				ExternalProviders: map[string]tfresource.ExternalProvider{
					"infoblox": {Source: "infobloxopen/infoblox", VersionConstraint: "3.0.0-beta"},
				},
				Config: v3Provider() + config,
			},
			v3Step(config),
		},
	})
}

func movedBlockVersion() []tfversion.TerraformVersionCheck {
	// Cross-type `moved` blocks need Terraform 1.8.
	return []tfversion.TerraformVersionCheck{tfversion.SkipBelow(tfversion.Version1_8_0)}
}

func v2Step(version, config string) tfresource.TestStep {
	return tfresource.TestStep{
		ExternalProviders: map[string]tfresource.ExternalProvider{
			"infoblox": {Source: "infobloxopen/infoblox", VersionConstraint: version},
		},
		Config: config,
	}
}

func v3Step(config string) tfresource.TestStep {
	return tfresource.TestStep{
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Config:                   v3Provider() + config,
		ConfigPlanChecks: tfresource.ConfigPlanChecks{
			PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
		},
	}
}

// v2 has no Portal passthrough, so these tests always use the Grid directly.
func preCheck(t *testing.T) {
	for _, k := range []string{"NIOS_HOST_URL", "NIOS_USERNAME", "NIOS_PASSWORD"} {
		if os.Getenv(k) == "" {
			t.Fatalf("%s must be set for migration acceptance tests", k)
		}
	}
}

func v3Provider() string {
	return fmt.Sprintf(`
provider "infoblox" {
  nios = {
    host_url = %q
    username = %q
    password = %q
  }
}
`, os.Getenv("NIOS_HOST_URL"), os.Getenv("NIOS_USERNAME"), os.Getenv("NIOS_PASSWORD"))
}

// v2 takes a host name, not a URL.
func v2Provider() string {
	host := strings.TrimPrefix(strings.TrimPrefix(os.Getenv("NIOS_HOST_URL"), "https://"), "http://")
	return fmt.Sprintf(`
provider "infoblox" {
  server   = %q
  username = %q
  password = %q
}
`, host, os.Getenv("NIOS_USERNAME"), os.Getenv("NIOS_PASSWORD"))
}

func v2Config(name, zone string) string {
	return v2Provider() + fmt.Sprintf(`
resource "infoblox_zone_auth" "zone" {
  fqdn = %[2]q
}

resource "infoblox_a_record" "a" {
  fqdn      = "a.${infoblox_zone_auth.zone.fqdn}"
  ip_addr   = "10.99.1.10"
  ext_attrs = jsonencode({ Site = "Blr" })
}

resource "infoblox_ns_record" "ns" {
  name       = infoblox_zone_auth.zone.fqdn
  nameserver = "ns2.example.org"
  addresses {
    address = "10.99.1.13"
  }
}

resource "infoblox_zone_delegated" "deleg" {
  fqdn = "sub.${infoblox_zone_auth.zone.fqdn}"
  delegate_to {
    name    = "ns1.example.org"
    address = "10.99.1.11"
  }
}

resource "infoblox_dtc_server" "srv" {
  name = "%[1]s-srv"
  host = "10.99.1.12"
}

resource "infoblox_dtc_pool" "pool" {
  name                = "%[1]s-pool"
  lb_preferred_method = "ROUND_ROBIN"
  servers {
    server = infoblox_dtc_server.srv.name
    ratio  = 1
  }
}

resource "infoblox_dtc_lbdn" "lbdn" {
  name      = "%[1]s-lbdn"
  lb_method = "ROUND_ROBIN"
  types     = ["A"]
  pools {
    pool  = infoblox_dtc_pool.pool.name
    ratio = 1
  }
}

resource "infoblox_ipv4_network" "net" {
  cidr = "10.97.77.0/24"
}

resource "infoblox_ip_allocation" "host" {
  fqdn       = "host.${infoblox_zone_auth.zone.fqdn}"
  ipv4_addr  = "10.97.77.10"
  enable_dns = true
  dns_view   = "default"
  depends_on = [infoblox_ipv4_network.net]
}

resource "infoblox_ip_association" "host" {
  internal_id = infoblox_ip_allocation.host.internal_id
  mac_addr    = "00:11:22:33:44:77"
  enable_dhcp = true
}
`, name, zone)
}

func v3Config(name, zone string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "zone" {
  nios = {
    fqdn = %[2]q
  }
}

moved {
  from = infoblox_a_record.a
  to   = infoblox_record_a.a
}

resource "infoblox_record_a" "a" {
  nios = {
    name      = "a.${infoblox_zone_auth.zone.nios.fqdn}"
    ipv4addr  = "10.99.1.10"
    ext_attrs = { Site = "Blr" }
  }
}

moved {
  from = infoblox_ns_record.ns
  to   = infoblox_record_ns.ns
}

resource "infoblox_record_ns" "ns" {
  nios = {
    name       = infoblox_zone_auth.zone.nios.fqdn
    nameserver = "ns2.example.org"
    addresses  = [{ address = "10.99.1.13" }]
  }
}

resource "infoblox_zone_delegated" "deleg" {
  nios = {
    fqdn        = "sub.${infoblox_zone_auth.zone.nios.fqdn}"
    delegate_to = [{ name = "ns1.example.org", address = "10.99.1.11" }]
  }
}

resource "infoblox_dtc_server" "srv" {
  nios = {
    name = "%[1]s-srv"
    host = "10.99.1.12"
  }
}

resource "infoblox_dtc_pool" "pool" {
  nios = {
    name                = "%[1]s-pool"
    lb_preferred_method = "ROUND_ROBIN"
    servers             = [{ server = infoblox_dtc_server.srv.id, ratio = 1 }]
  }
}

resource "infoblox_dtc_lbdn" "lbdn" {
  nios = {
    name      = "%[1]s-lbdn"
    lb_method = "ROUND_ROBIN"
    types     = ["A"]
    pools     = [{ pool = infoblox_dtc_pool.pool.id, ratio = 1 }]
  }
}

moved {
  from = infoblox_ipv4_network.net
  to   = infoblox_network.net
}

resource "infoblox_network" "net" {
  nios = {
    network = "10.97.77.0/24"
  }
}

moved {
  from = infoblox_ip_allocation.host
  to   = infoblox_record_host.host
}

resource "infoblox_record_host" "host" {
  nios = {
    name      = "host.${infoblox_zone_auth.zone.nios.fqdn}"
    view      = "default"
    ipv4addrs = [{ ipv4addr = "10.97.77.10" }]
  }
  depends_on = [infoblox_network.net]
}

resource "infoblox_ip_association" "host" {
  nios = {
    ref                = infoblox_record_host.host.id
    mac                = "00:11:22:33:44:77"
    configure_for_dhcp = true
  }
}
`, name, zone)
}
