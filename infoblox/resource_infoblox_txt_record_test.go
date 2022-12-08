package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func testAccCheckTXTRecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_txt_record" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetTXTRecordByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}
	}
	return nil
}

func testAccTXTRecordCompare(t *testing.T, resPath string, expectedRec *ibclient.RecordTXT) resource.TestCheckFunc {
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

		rec, _ := objMgr.GetTXTRecordByRef(res.Primary.ID)
		if rec == nil {
			return fmt.Errorf("record not found")
		}

		if rec.Name != expectedRec.Name {
			return fmt.Errorf(
				"'name' does not match: got '%s', expected '%s'",
				rec.Name,
				expectedRec.Name)
		}
		if rec.Text != expectedRec.Text {
			return fmt.Errorf(
				"'text' does not match: got '%s', expected '%s'",
				rec.Text, expectedRec.Text)
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

func TestAccResourceTXTRecord(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTXTRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
						resource "infoblox_txt_record" "foo"{
							name="name1.test.com"
							text="test"
							dns_view="default"
							ttl = 0
							comment="test comment 1"
							ext_attrs = jsonencode({
								"Tenant ID" = "terraform_test_tenant"
								"Location" = "Test loc"
								"Site" = "Test site"
								"TestEA1"=["text1","text2"]
							})
							}`),
				Check: resource.ComposeTestCheckFunc(
					testAccTXTRecordCompare(t, "infoblox_txt_record.foo", &ibclient.RecordTXT{
						Name:    "name1.test.com",
						Text:    "test",
						View:    "default",
						Ttl:     0,
						Comment: "test comment 1",
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
						resource "infoblox_txt_record" "foo2"{
							name="name2.test.com"
							text="test"
							dns_view="default"
							ttl = 3600
							comment="test comment 1"
							ext_attrs = jsonencode({
								"Tenant ID" = "terraform_test_tenant"
								"Location" = "Test loc"
								"Site" = "Test site"
								"TestEA1"=["text1","text2"]
							})
							}`),
				Check: resource.ComposeTestCheckFunc(
					testAccTXTRecordCompare(t, "infoblox_txt_record.foo2", &ibclient.RecordTXT{
						Name:    "name2.test.com",
						Text:    "test",
						View:    "default",
						Ttl:     3600,
						Comment: "test comment 1",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc",
							"Site":      "Test site",
							"TestEA1":   []string{"text1", "text2"},
						},
					}),
				),
			},
		},
	})
}
