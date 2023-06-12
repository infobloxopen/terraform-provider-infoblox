package infoblox

import (
	"fmt"
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

func testAccAAAARecordCompare(t *testing.T, resPath string, expectedRec *ibclient.RecordAAAA) resource.TestCheckFunc {
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

		if expectedRec.Ipv6Addr == "" {
			expectedRec.Ipv6Addr = res.Primary.Attributes["ipv6_addr"]
			if expectedRec.Ipv6Addr == "" {
				return fmt.Errorf(
					"the value of 'ipv6_addr' field is empty, but expected some value")
			}
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

var changeNotAllowedExp = regexp.MustCompile("when '.+' exists '.+' value is not allowed to update")

var bothChangesNotAllowedExp = regexp.MustCompile("both '.+' and '.+' values are not allowed to update at once")

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
						dns_view = "default"
						comment = "test comment 1"
						ext_attrs = jsonencode({
							"Tenant ID"="terraform_test_tenant"
							"Location"="Test loc"
							"Site"="Test site"
							"TestEA1"=["text1","text2"]
						})
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccAAAARecordCompare(t, "infoblox_aaaa_record.foo", &ibclient.RecordAAAA{
						Ipv6Addr: "2000::1",
						Name:     "name1.test.com",
						View:     "default",
						Ttl:      0,
						Comment:  "test comment 1",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc",
							"Site":      "Test site",
							"TestEA1":   []string{"text1", "text2"},
						},
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_aaaa_record" "foo2"{
						fqdn = "name3.test.com"
						ipv6_addr = "2000::3"
						ttl = 155
						dns_view = "default"
						comment = "test comment 2"
						ext_attrs = jsonencode({
							"Tenant ID"="terraform_test_tenant"
							"Location"="Test loc"
						})
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccAAAARecordCompare(t, "infoblox_aaaa_record.foo2", &ibclient.RecordAAAA{
						Ipv6Addr: "2000::3",
						Name:     "name3.test.com",
						View:     "default",
						Ttl:      155,
						Comment:  "test comment 2",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc",
						},
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_aaaa_record" "foo2"{
						fqdn = "name3.test.com"
						ipv6_addr = "2000::3"
						cidr = "2001::/64"
						dns_view = "default"
						comment = "test comment 2"
					}`),
				ExpectError: changeNotAllowedExp,
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_aaaa_record" "foo2"{
						fqdn = "name3.test.com"
						ipv6_addr = "2000::5"
						cidr = "2001::/64"
						dns_view = "default"
						comment = "test comment 2"
					}`),
				ExpectError: bothChangesNotAllowedExp,
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_ipv6_network" "v6n1" {
						cidr = "2001::/64"
					}
					resource "infoblox_aaaa_record" "foo3"{
						fqdn = "name4.test.com"
						cidr = infoblox_ipv6_network.v6n1.cidr
						dns_view = "default"
						comment = "test comment 3"
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccAAAARecordCompare(t, "infoblox_aaaa_record.foo3", &ibclient.RecordAAAA{
						Name:    "name4.test.com",
						View:    "default",
						Comment: "test comment 3",
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_aaaa_record" "foo3"{
						fqdn = "name4.test.com"
						cidr = "2001::/64"
						ipv6_addr = "2000::7"
						dns_view = "default"
						comment = "test comment 2"
					}`),
				ExpectError: changeNotAllowedExp,
			},
		},
	})
}
