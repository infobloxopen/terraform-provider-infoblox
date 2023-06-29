package infoblox

import (
	"fmt"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"net"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func testAccCheckARecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_a_record" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetARecordByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}

	}
	return nil
}

func testAccARecordCompare(
	t *testing.T,
	resPath string,
	expectedRec *ibclient.RecordA,
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

		rec, _ := objMgr.GetARecordByRef(res.Primary.ID)
		if rec == nil {
			return fmt.Errorf("record not found")
		}

		if rec.Name != expectedRec.Name {
			return fmt.Errorf(
				"'fqdn' does not match: got '%s', expected '%s'",
				*rec.Name,
				*expectedRec.Name)
		}
		if notExpectedIpAddr != "" && notExpectedIpAddr == *rec.Ipv4Addr {
			return fmt.Errorf(
				"'ip_addr' field has value '%s' but that is not expected to happen",
				notExpectedIpAddr)
		}
		if expectedCidr != "" {
			_, parsedCidr, err := net.ParseCIDR(expectedCidr)
			if err != nil {
				panic(fmt.Sprintf("cannot parse CIDR '%s': %s", expectedCidr, err))
			}

			if !parsedCidr.Contains(net.ParseIP(*rec.Ipv4Addr)) {
				return fmt.Errorf(
					"IP address '%s' does not belong to the expected CIDR '%s'",
					*rec.Ipv4Addr, expectedCidr)
			}
		}
		if *expectedRec.Ipv4Addr == "" {
			expectedRec.Ipv4Addr = utils.StringPtr(res.Primary.Attributes["ip_addr"])
		}
		if rec.Ipv4Addr != expectedRec.Ipv4Addr {
			return fmt.Errorf(
				"'ipv4address' does not match: got '%s', expected '%s'",
				*rec.Ipv4Addr, *expectedRec.Ipv4Addr)
		}
		if rec.View != expectedRec.View {
			return fmt.Errorf(
				"'dns_view' does not match: got '%s', expected '%s'",
				rec.View, expectedRec.View)
		}
		if rec.UseTtl != expectedRec.UseTtl {
			return fmt.Errorf(
				"TTL usage does not match: got '%t', expected '%t'",
				*rec.UseTtl, *expectedRec.UseTtl)
		}
		if *rec.UseTtl {
			if *rec.Ttl != *expectedRec.Ttl {
				return fmt.Errorf(
					"'Ttl' usage does not match: got '%d', expected '%d'",
					rec.Ttl, expectedRec.Ttl)
			}
		}
		if rec.Comment != expectedRec.Comment {
			return fmt.Errorf(
				"'comment' does not match: got '%s', expected '%s'",
				*rec.Comment, *expectedRec.Comment)
		}

		return validateEAs(rec.Ea, expectedRec.Ea)
	}
}

var (
	regexpRequiredMissingIPv4    = regexp.MustCompile("either of 'ip_addr' and 'cidr' values is required")
	regexpCidrIpAddrConflictIPv4 = regexp.MustCompile("only one of 'ip_addr' and 'cidr' values is allowed to be defined")

	regexpNetviewUpdateNotAllowed = regexp.MustCompile("changing the value of 'network_view' field is not allowed")
	regexpDnsviewUpdateNotAllowed = regexp.MustCompile("changing the value of 'dns_view' field is not allowed")
)

func TestAccResourceARecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckARecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_a_record" "foo"{
						fqdn = "name1.test.com"
						ip_addr = "10.0.0.2"
						cidr = "10.20.30.0/24"
                        network_view = "default"
					}`),
				ExpectError: regexpCidrIpAddrConflictIPv4,
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_a_record" "foo"{
						fqdn = "name1.test.com"
					}`),
				ExpectError: regexpRequiredMissingIPv4,
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_a_record" "foo"{
						fqdn = "name1.test.com"
						ip_addr = "10.0.0.2"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccARecordCompare(t, "infoblox_a_record.foo", &ibclient.RecordA{
						Ipv4Addr: utils.StringPtr("10.0.0.2"),
						Name:     utils.StringPtr("name1.test.com"),
						View:     "default",
						Ttl:      utils.Uint32Ptr(0),
						UseTtl:   utils.BoolPtr(false),
						Comment:  utils.StringPtr(""),
						Ea:       nil,
					}, "", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_a_record" "foo2"{
						fqdn = "name2.test.com"
						ip_addr = "192.168.31.31"
						ttl = 10
						dns_view = "nondefault_view"
						comment = "test comment 1"
						ext_attrs = jsonencode({
						  "Location" = "New York"
						  "Site" = "HQ"
						})
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccARecordCompare(t, "infoblox_a_record.foo2", &ibclient.RecordA{
						Ipv4Addr: utils.StringPtr("192.168.31.31"),
						Name:     utils.StringPtr("name2.test.com"),
						View:     "nondefault_view",
						Ttl:      utils.Uint32Ptr(10),
						UseTtl:   utils.BoolPtr(true),
						Comment:  utils.StringPtr("test comment 1"),
						Ea: ibclient.EA{
							"Location": "New York",
							"Site":     "HQ",
						},
					}, "", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_a_record" "foo2"{
						fqdn = "name3.test.com"
						ip_addr = "10.10.0.1"
						ttl = 155
						dns_view = "nondefault_view"
						comment = "test comment 2"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccARecordCompare(t, "infoblox_a_record.foo2", &ibclient.RecordA{
						Ipv4Addr: utils.StringPtr("10.10.0.1"),
						Name:     utils.StringPtr("name3.test.com"),
						View:     "nondefault_view",
						Ttl:      utils.Uint32Ptr(155),
						UseTtl:   utils.BoolPtr(true),
						Comment:  utils.StringPtr("test comment 2"),
					}, "", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_a_record" "foo2"{
						fqdn = "name3.test.com"
						ip_addr = "10.10.0.1"
						dns_view = "nondefault_view"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccARecordCompare(t, "infoblox_a_record.foo2", &ibclient.RecordA{
						Ipv4Addr: utils.StringPtr("10.10.0.1"),
						Name:     utils.StringPtr("name3.test.com"),
						View:     "nondefault_view",
						UseTtl:   utils.BoolPtr(false),
					}, "", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
                    resource "infoblox_ipv4_network" "net1" {
                        cidr = "10.20.30.0/24"
                        network_view = "default"
                    }
					resource "infoblox_a_record" "foo2"{
						fqdn = "name3.test.com"
                        cidr = infoblox_ipv4_network.net1.cidr
                        network_view = infoblox_ipv4_network.net1.network_view
						dns_view = "nondefault_view"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccARecordCompare(t, "infoblox_a_record.foo2", &ibclient.RecordA{
						Name:   utils.StringPtr("name3.test.com"),
						View:   "nondefault_view",
						UseTtl: utils.BoolPtr(false),
					}, "10.10.0.1", "10.20.30.0/24"),
				),
			},
			{
				Config: fmt.Sprintf(`
                    resource "infoblox_ipv4_network" "net2" {
                        cidr = "10.20.33.0/24"
                        network_view = "default"
                    }
					resource "infoblox_a_record" "foo2"{
						fqdn = "name3.test.com"
                        cidr = infoblox_ipv4_network.net2.cidr
                        network_view = infoblox_ipv4_network.net2.network_view
						dns_view = "nondefault_view"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccARecordCompare(t, "infoblox_a_record.foo2", &ibclient.RecordA{
						Name:   utils.StringPtr("name3.test.com"),
						View:   "nondefault_view",
						UseTtl: utils.BoolPtr(false),
					}, "", "10.20.33.0/24"),
				),
			},
			{
				Config: fmt.Sprintf(`
                    resource "infoblox_ipv4_network" "net3" {
                        cidr = "10.20.34.0/24"
                        network_view = "nondefault_netview"
                    }
					resource "infoblox_a_record" "foo2"{
						fqdn = "name3.test.com"
                        cidr = infoblox_ipv4_network.net3.cidr
                        network_view = infoblox_ipv4_network.net3.network_view
						dns_view = "nondefault_view"
					}`),
				ExpectError: regexpNetviewUpdateNotAllowed,
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_a_record" "foo2"{
						fqdn = "name3.test.com"
						ip_addr = "10.10.0.2"
						dns_view = "nondefault_view"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccARecordCompare(t, "infoblox_a_record.foo2", &ibclient.RecordA{
						Ipv4Addr: utils.StringPtr("10.10.0.2"),
						Name:     utils.StringPtr("name3.test.com"),
						View:     "nondefault_view",
						UseTtl:   utils.BoolPtr(false),
					}, "", ""),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_a_record" "foo2"{
						fqdn = "name3.test.com"
						ip_addr = "10.10.0.2"
						dns_view = "default"
					}`),
				ExpectError: regexpDnsviewUpdateNotAllowed,
			},
		},
	})
}
