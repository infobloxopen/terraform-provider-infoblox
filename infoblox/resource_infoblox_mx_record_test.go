package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func testAccCheckMXRecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_mx_record" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetMXRecordByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}

	}
	return nil
}

func testAccMXRecordCompare(t *testing.T, resPath string, expectedRec *ibclient.RecordMX) resource.TestCheckFunc {
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

		rec, _ := objMgr.GetMXRecordByRef(res.Primary.ID)
		if rec == nil {
			return fmt.Errorf("record not found")
		}

		if rec.Fqdn != expectedRec.Fqdn {
			return fmt.Errorf(
				"'fqdn' does not match: got '%s', expected '%s'",
				rec.Fqdn,
				expectedRec.Fqdn)
		}

		if rec.View != expectedRec.View {
			return fmt.Errorf(
				"'dns_view' does not match: got '%s', expected '%s'",
				rec.View, expectedRec.View)
		}

		if rec.MX != expectedRec.MX {
			return fmt.Errorf(
				"'mail_exchanger' does not match: got '%s', expected '%s'",
				rec.MX, expectedRec.MX)
		}

		if rec.Priority != expectedRec.Priority {
			return fmt.Errorf(
				"'priority' does not match: got '%d', expected '%d'",
				rec.Priority, expectedRec.Priority)
		}

		if rec.UseTtl != expectedRec.UseTtl {
			return fmt.Errorf(
				"TTL usage does not match: got '%t', expected '%t'",
				rec.UseTtl, expectedRec.UseTtl)
		}
		if rec.UseTtl {
			if rec.Ttl != expectedRec.Ttl {
				return fmt.Errorf(
					"'Ttl' usage does not match: got '%d', expected '%d'",
					rec.Ttl, expectedRec.Ttl)
			}
		}
		if rec.Comment != expectedRec.Comment {
			return fmt.Errorf(
				"'comment' does not match: got '%s', expected '%s'",
				rec.Comment, expectedRec.Comment)
		}
		return validateEAs(rec.Ea, expectedRec.Ea)
	}
}

func TestAccResourceMXRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMXRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_mx_record" "foo"{
						fqdn = "name1.test.com"
						mail_exchanger = "sample.mx1.com"
						preference = 25
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccMXRecordCompare(t, "infoblox_mx_record.foo", &ibclient.RecordMX{
						Fqdn:     "name1.test.com",
						MX:       "sample.mx1.com",
						Priority: 25,
						View:     "default",
						Ttl:      0,
						UseTtl:   false,
						Comment:  "",
						Ea:       nil,
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_mx_record" "foo1"{
						fqdn = "name2.test.com"
						mail_exchanger = "sample.mx2.com"
						preference = 30
						ttl = 300
						dns_view = "nondefault_view"
						comment = "test comment 1"
						extattrs = jsonencode({
							"Location" = "Los Angeles"
							"Site" = "HQ"
						}) 
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccMXRecordCompare(t, "infoblox_mx_record.foo1", &ibclient.RecordMX{
						Fqdn:     "name2.test.com",
						MX:       "sample.mx1.com",
						Priority: 30,
						View:     "nondefault_view",
						Ttl:      300,
						UseTtl:   true,
						Comment:  "test comment 1",
						Ea: ibclient.EA{
							"Location": "Los Angeles",
							"Site":     "HQ",
						},
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_mx_record" "foo2"{
						fqdn = "name3.test.com"
						mail_exchanger = "sample.mx3.com"
						preference = 35
						ttl = 150
						dns_view = "nondefault_view"
						comment = "test comment 2" 
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccMXRecordCompare(t, "infoblox_mx_record.foo2", &ibclient.RecordMX{
						Fqdn:     "name3.test.com",
						MX:       "sample.mx3.com",
						Priority: 35,
						View:     "nondefault_view",
						Ttl:      150,
						UseTtl:   true,
						Comment:  "test comment 2",
					}),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "infoblox_a_record" "foo3"{
						fqdn = "name3.test.com"
						dns_view = "nondefault_view"
						mail_exchanger = "sample.mx3.com"
						preference = 20
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccMXRecordCompare(t, "infoblox_mx_record.foo3", &ibclient.RecordMX{
						Fqdn:     "name3.test.com",
						View:     "nondefault_view",
						MX:       "sample.mx3.com",
						Priority: 20,
					}),
				),
			},
		},
	})
}
