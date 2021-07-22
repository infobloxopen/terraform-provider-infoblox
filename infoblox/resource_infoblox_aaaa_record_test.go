package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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
			return fmt.Errorf("Not found: %s", resPath)
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
				"Ttl value does not match: got '%d', expected '%d'",
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
		},
	})
}
