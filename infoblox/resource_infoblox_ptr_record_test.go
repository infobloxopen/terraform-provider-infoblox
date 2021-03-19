package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func TestAccResourcePTRRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPTRRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccresourcePTRRecordCreate,
				Check: resource.ComposeTestCheckFunc(
					testAccPTRRecordExists("infoblox_ptr_record.foo", "10.0.0.0/24", "10.0.0.2", "test", "demo-network", "default", "a.com"),
				),
			},
			{
				Config: testAccresourcePTRRecordAllocate,
				Check: resource.ComposeTestCheckFunc(
					testAccPTRRecordExists("infoblox_ptr_record.foo1", "10.0.0.0/24", "10.0.0.1", "test", "demo-network", "default", "a.com"),
					testAccPTRRecordExists("infoblox_ptr_record.foo2", "10.0.0.0/24", "10.0.0.2", "test", "demo-network", "default", "a.com"),
				),
			},
			{
				Config: testAccresourcePTRRecordUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccPTRRecordExists("infoblox_ptr_record.foo", "10.0.0.0/24", "10.0.0.2", "test", "demo-network", "default", "a.com"),
				),
			},
		},
	})
}

func testAccCheckPTRRecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_a_record" {
			continue
		}
		Connector := meta.(*ibclient.Connector)
		objMgr := ibclient.NewObjectManager(Connector, "terraform_test", "test")
		recordName, _ := objMgr.GetPTRRecordByRef(rs.Primary.ID)
		if recordName != nil {
			return fmt.Errorf("record not found")
		}
	}
	return nil
}
func testAccPTRRecordExists(n string, cidr string, ipAddr string, networkViewName string, recordName string, dnsView string, zone string) resource.TestCheckFunc {
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

		recordName, _ := objMgr.GetPTRRecordByRef(rs.Primary.ID)
		if recordName == nil {
			return fmt.Errorf("record not found")
		}

		return nil
	}
}

var testAccresourcePTRRecordCreate = `
resource "infoblox_ptr_record" "foo"{
	vm_name="test-name"
	zone="a.com"
	ip_addr="10.0.0.2"
	tenant_id="foo"
	}`

var testAccresourcePTRRecordAllocate = `
resource "infoblox_ptr_record" "foo1"{
	vm_name="test-name"
	zone="a.com"
	ip_addr=""
	cidr="10.0.0.0/24"
	tenant_id="foo"
	}
resource "infoblox_ptr_record" "foo2"{
	vm_name="test-name"
	zone="a.com"
	ip_addr=""
	cidr="10.0.0.0/24"
	tenant_id="foo"
	}`

var testAccresourcePTRRecordUpdate = `
resource "infoblox_ptr_record" "foo"{
	vm_name="test-name"
	dns_view="default"
	zone="a.com"
	cidr="10.0.0.0/24"
	ip_addr="10.0.0.2"
	tenant_id="foo"
	}`
