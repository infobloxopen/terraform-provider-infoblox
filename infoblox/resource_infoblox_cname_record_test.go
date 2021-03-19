package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func TestAccResourceCNAMERecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCNAMERecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccresourceCNAMERecordCreate,
				Check: resource.ComposeTestCheckFunc(
					testAccCNAMERecordExists(t, "infoblox_cname_record.foo", "test", "test-name", "default", "a.com"),
				),
			},
			resource.TestStep{
				Config: testAccresourceCNAMERecordUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCNAMERecordExists(t, "infoblox_cname_record.foo", "test", "test-name", "default", "a.com"),
				),
			},
		},
	})
}

func testAccCheckCNAMERecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_a_record" {
			continue
		}
		Connector := meta.(*ibclient.Connector)
		objMgr := ibclient.NewObjectManager(Connector, "terraform_test", "test")
		recordName, _ := objMgr.GetCNAMERecordByRef(rs.Primary.ID)
		if recordName != nil {
			return fmt.Errorf("record not found")
		}
	}
	return nil
}
func testAccCNAMERecordExists(t *testing.T, n string, alias string, canonical string, dnsView string, zone string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found:%s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID i set")
		}
		meta := testAccProvider.Meta()
		Connector := meta.(*ibclient.Connector)
		objMgr := ibclient.NewObjectManager(Connector, "terraform_test", "test")

		recordName, _ := objMgr.GetCNAMERecordByRef(rs.Primary.ID)
		if recordName == nil {
			return fmt.Errorf("record not found")
		}

		return nil
	}
}

var testAccresourceCNAMERecordCreate = `
resource "infoblox_cname_record" "foo"{
	alias="test"
	canonical="test-name"
	dns_view="default"
	zone="a.com"
	tenant_id="foo"
	}`

var testAccresourceCNAMERecordUpdate = `
resource "infoblox_cname_record" "foo"{
	alias="test"
	canonical="test-name"
	dns_view="default"
	zone="a.com"
	tenant_id="foo"
	}`
