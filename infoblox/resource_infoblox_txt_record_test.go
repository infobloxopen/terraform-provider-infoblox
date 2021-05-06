package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func TestAccResourceTXTRecord(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTXTRecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccresourceTXTRecordCreate,
				Check: resource.ComposeTestCheckFunc(
					testAccTXTRecordExists(t, "infoblox_txt_record.foo", "test-name", "a.com", 3600, "default"),
				),
			},
			resource.TestStep{
				Config: testAccresourceTXTRecordAllocate,
				Check: resource.ComposeTestCheckFunc(
					testAccTXTRecordExists(t, "infoblox_txt_record.foo1", "test-name", "a.com", 3600, "default"),
					testAccTXTRecordExists(t, "infoblox_txt_record.foo2", "test-name", "a.com", 3600, "default"),
				),
			},
			resource.TestStep{
				Config: testAccresourceTXTRecordUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccTXTRecordExists(t, "infoblox_txt_record.foo", "test-name", "a.com", 3600, "default"),
				),
			},
		},
	})
}

func testAccCheckTXTRecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_txt_record" {
			continue
		}
		Connector := meta.(*ibclient.Connector)
		objMgr := ibclient.NewObjectManager(Connector, "terraform_test", "test")
		recordName, _ := objMgr.GetTXTRecordByRef(rs.Primary.ID)
		if recordName != nil {
			return fmt.Errorf("record not found")
		}

	}
	return nil
}
func testAccTXTRecordExists(t *testing.T, n string, recordName string, text string, ttl int, dnsView string) resource.TestCheckFunc {
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

		recordName, _ := objMgr.GetTXTRecordByRef(rs.Primary.ID)
		if recordName == nil {
			return fmt.Errorf("record not found")
		}

		return nil
	}
}

var testAccresourceTXTRecordCreate = fmt.Sprintf(`
resource "infoblox_txt_record" "foo"{
	name="test-name"
	text="a.com"
	ttl=3600
	dns_view="default"
	tenant_id="foo"
	}`)

var testAccresourceTXTRecordAllocate = fmt.Sprintf(`
resource "infoblox_txt_record" "foo1"{
	name="test-name"
	text="a.com"
	ttl=3600
	dns_view="default"
	tenant_id="foo"
	}
resource "infoblox_txt_record" "foo2"{
	name="test-name"
	text="a.com"
	ttl=3600
	dns_view="default"
	tenant_id="foo"
	}`)

var testAccresourceTXTRecordUpdate = fmt.Sprintf(`
resource "infoblox_txt_record" "foo"{
	name="test-name"
	text="a.com"
	ttl=3600
	dns_view="default"
	tenant_id="foo"
	}`)
