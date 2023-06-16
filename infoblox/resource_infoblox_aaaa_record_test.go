package infoblox

import (
	"fmt"
	"net"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func testAccCheckAAAARecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_aaaa_record" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetAAAARecordByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}

	}
	return nil
}

func testAccAAAARecordCompare(
	t *testing.T,
	resPath string,
	expectedRec *ibclient.RecordAAAA,
	notExpectedIpAddr string,
	expectedCidr string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found: %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID is not set")
		}
		meta := testAccProvider.Meta()
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")

		rec, _ := objMgr.GetAAAARecordByRef(res.Primary.ID)
		if rec == nil {
			return fmt.Errorf("record not found")
		}

		if rec.Name != expectedRec.Name {
			return fmt.Errorf(
				"'fqdn' does not match: got '%s', expected '%s'",
				rec.Name,
				expectedRec.Name)
		}
		if notExpectedIpAddr != "" && notExpectedIpAddr == rec.Ipv6Addr {
			return fmt.Errorf(
				"'ipv6_addr' field has value '%s' but that is not expected to happen",
				notExpectedIpAddr)
		}
		if expectedCidr != "" {
			_, parsedCidr, err := net.ParseCIDR(expectedCidr)
			if err != nil {
				panic(fmt.Sprintf("cannot parse CIDR '%s': %s", expectedCidr, err))
			}

			if !parsedCidr.Contains(net.ParseIP(rec.Ipv6Addr)) {
				return fmt.Errorf(
					"IP address '%s' does not belong to the expected CIDR '%s'",
					rec.Ipv6Addr, expectedCidr)
			}
		}
		if expectedRec.Ipv6Addr == "" {
			expectedRec.Ipv6Addr = res.Primary.Attributes["ipv6_addr"]
		}
		if rec.Ipv6Addr != expectedRec.Ipv6Addr {
			return fmt.Errorf(
				"'ipv6address' does not match: got '%s', expected '%s'",
				rec.Ipv6Addr, expectedRec.Ipv6Addr)
		}
		if rec.View != expectedRec.View {
			return fmt.Errorf(
				"'dns_view' does not match: got '%s', expected '%s'",
				rec.View, expectedRec.View)
		}
		if rec.Ttl != expectedRec.Ttl {
			return fmt.Errorf(
				"TTL value does not match: got '%d', expected '%d'",
				rec.Ttl, expectedRec.Ttl)
		}
		if rec.Comment != expectedRec.Comment {
			return fmt.Errorf(
				"'comment' does not match: got '%s', expected '%s'",
				rec.Comment, expectedRec.Comment)
		}
		return validateEAs(rec.Ea, expectedRec.Ea)
	}
}

var (
	regexpRequiredMissingIPv6    = regexp.MustCompile("either of 'ipv6_addr' and 'cidr' values is required")
	regexpCidrIpAddrConflictIPv6 = regexp.MustCompile("only one of 'ipv6_addr' and 'cidr' values is allowed to be defined")
)

func TestAccResourceAAAARecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAAAARecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_aaaa_record" "foo"{
						fqdn = "name1.test.com"
						ipv6_addr = "2000::1"
						cidr = "2000:1fde::/96"
                        network_view = "default"
					}`),
				ExpectError: regexpCidrIpAddrConflictIPv6,
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_aaaa_record" "foo"{
						fqdn = "name1.test.com"
					}`),
				ExpectError: regexpRequiredMissingIPv6,
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_aaaa_record" "foo"{
						fqdn = "name1.test.com"
						ipv6_addr = "2000::1"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccAAAARecordCompare(t, "infoblox_aaaa_record.foo", &ibclient.RecordAAAA{
						Ipv6Addr: "2000::1",
						Name:     "name1.test.com",
						View:     "default",
						Ttl:      0,
						UseTtl:   false,
						Comment:  "",
						Ea:       nil,
					}, "", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_aaaa_record" "foo2"{
						fqdn = "name2.test.com"
						ipv6_addr = "2002::10"
						ttl = 10
						dns_view = "nondefault_view"
						comment = "test comment 1"
						ext_attrs = jsonencode({
						  "Location" = "New York"
						  "Site" = "HQ"
						})
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccAAAARecordCompare(t, "infoblox_aaaa_record.foo2", &ibclient.RecordAAAA{
						Ipv6Addr: "2002::10",
						Name:     "name2.test.com",
						View:     "nondefault_view",
						Ttl:      10,
						UseTtl:   true,
						Comment:  "test comment 1",
						Ea: ibclient.EA{
							"Location": "New York",
							"Site":     "HQ",
						},
					}, "", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_aaaa_record" "foo2"{
						fqdn = "name3.test.com"
						ipv6_addr = "2000::1"
						ttl = 155
						dns_view = "nondefault_view"
						comment = "test comment 2"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccAAAARecordCompare(t, "infoblox_aaaa_record.foo2", &ibclient.RecordAAAA{
						Ipv6Addr: "2000::1",
						Name:     "name3.test.com",
						View:     "nondefault_view",
						Ttl:      155,
						UseTtl:   true,
						Comment:  "test comment 2",
					}, "", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_aaaa_record" "foo2"{
						fqdn = "name3.test.com"
						ipv6_addr = "2000::1"
						dns_view = "nondefault_view"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccAAAARecordCompare(t, "infoblox_aaaa_record.foo2", &ibclient.RecordAAAA{
						Ipv6Addr: "2000::1",
						Name:     "name3.test.com",
						View:     "nondefault_view",
						UseTtl:   false,
					}, "", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
                    resource "infoblox_ipv6_network" "net1" {
                        cidr = "2000:1fde::/96"
                        network_view = "default"
                    }
					resource "infoblox_aaaa_record" "foo2"{
						fqdn = "name3.test.com"
                        cidr = infoblox_ipv6_network.net1.cidr
                        network_view = infoblox_ipv6_network.net1.network_view
						dns_view = "nondefault_view"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccAAAARecordCompare(t, "infoblox_aaaa_record.foo2", &ibclient.RecordAAAA{
						Name:   "name3.test.com",
						View:   "nondefault_view",
						UseTtl: false,
					}, "2000::1", "2000:1fde::/96"),
				),
			},
			{
				Config: fmt.Sprintf(`
                    resource "infoblox_ipv6_network" "net2" {
                        cidr = "2000:1fcc::/96"
                        network_view = "default"
                    }
					resource "infoblox_aaaa_record" "foo2"{
						fqdn = "name3.test.com"
                        cidr = infoblox_ipv6_network.net2.cidr
                        network_view = infoblox_ipv6_network.net2.network_view
						dns_view = "nondefault_view"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccAAAARecordCompare(t, "infoblox_aaaa_record.foo2", &ibclient.RecordAAAA{
						Name:   "name3.test.com",
						View:   "nondefault_view",
						UseTtl: false,
					}, "", "2000:1fcc::/96"),
				),
			},
			{
				Config: fmt.Sprintf(`
                    resource "infoblox_ipv6_network" "net3" {
                        cidr = "2000:1fcd::/96"
                        network_view = "nondefault_netview"
                    }
					resource "infoblox_aaaa_record" "foo2"{
						fqdn = "name3.test.com"
                        cidr = infoblox_ipv6_network.net3.cidr
                        network_view = infoblox_ipv6_network.net3.network_view
						dns_view = "nondefault_view"
					}`),
				ExpectError: regexpNetviewUpdateNotAllowed,
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_aaaa_record" "foo2"{
						fqdn = "name3.test.com"
						ipv6_addr = "2000::2"
						dns_view = "nondefault_view"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccAAAARecordCompare(t, "infoblox_aaaa_record.foo2", &ibclient.RecordAAAA{
						Ipv6Addr: "2000::2",
						Name:     "name3.test.com",
						View:     "nondefault_view",
						UseTtl:   false,
					}, "", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_aaaa_record" "foo2"{
						fqdn = "name3.test.com"
						ipv6_addr = "2000::2"
						dns_view = "default"
					}`),
				ExpectError: regexpDnsviewUpdateNotAllowed,
			},
		},
	})
}
