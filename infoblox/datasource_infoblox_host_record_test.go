package infoblox

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"regexp"
	"testing"
)

func TestAccDataSourceHostRecordRead_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceHostRecordReadConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.infoblox_host_record.test", "id"),
					resource.TestCheckResourceAttr("data.infoblox_host_record.test", "results.0.fqdn", "testhostnameip1.test.com"),
				),
			},
		},
	})
}

func TestAccDataSourceHostRecordRead_noResult(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceHostRecordReadConfig_noResult(),
				ExpectError: regexp.MustCompile("failed getting Host-record: not found"),
			},
		},
	})
}

func testAccDataSourceHostRecordReadConfig_basic() string {
	return `
resource "infoblox_zone_auth" "zone" {
  fqdn = "test.com"
}
resource "infoblox_ip_allocation" "foo1"{
  network_view="default"
  dns_view = "default"
  fqdn="testhostnameip1.test.com"
  ipv6_addr="2001:db8:abcd:12::1"
  ipv4_addr="10.0.0.1"
  ttl = 10
  comment = "IPv4 and IPv6 are allocated"
  ext_attrs = jsonencode({
    Site = "Test site"
    })
  depends_on = [infoblox_zone_auth.zone]
	}

data "infoblox_host_record" "test" {
  filters = {
    name = infoblox_ip_allocation.foo1.fqdn
  }
  depends_on = [infoblox_ip_allocation.foo1]
}
`
}

func testAccDataSourceHostRecordReadConfig_noResult() string {
	return `
data "infoblox_host_record" "test" {
  filters = {
    name = "nonexistent.example.com"
  }
}
`
}

func TestAccDataSourceHostRecordRead_noIPv4Address(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceHostRecordReadConfig_noIPv4Address(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.infoblox_host_record.test", "id"),
					resource.TestCheckResourceAttr("data.infoblox_host_record.test", "results.0.ipv4_addr", ""),
					resource.TestCheckResourceAttr("data.infoblox_host_record.test", "results.0.ipv6_addr", "2001:db8:abcd:12::1"),
				),
			},
		},
	})
}

func testAccDataSourceHostRecordReadConfig_noIPv4Address() string {
	return `
resource "infoblox_zone_auth" "zone" {
  fqdn = "test.com"
}
resource "infoblox_ip_allocation" "foo1"{
  network_view="default"
  dns_view = "default"
  fqdn="testhostnameip1.test.com"
  ipv6_addr="2001:db8:abcd:12::1"
  ttl = 10
  comment = "IPv6 allocated"
  ext_attrs = jsonencode({
    Site = "Test site"
    })
  depends_on = [infoblox_zone_auth.zone]
 }

data "infoblox_host_record" "test" {
  filters = {
    name = infoblox_ip_allocation.foo1.fqdn
  }
  depends_on = [infoblox_ip_allocation.foo1]
}
`
}
