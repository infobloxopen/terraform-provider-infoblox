package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func TestAccResourceIPAllocation(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPAllocationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccresourceIPAllocationCreate,
				Check: resource.ComposeTestCheckFunc(
					testAccIPExists("infoblox_ip_allocation.foo", "10.0.0.1/24", "10.0.0.1", "default", "demo-network"),
				),
			},
			{
				Config: testAccresourceIPAllocationUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccIPExists("infoblox_ip_allocation.foo", "10.0.0.1/24", "10.0.0.1", "default", "demo-network"),
				),
			},
		},
	})
}

func testAccCheckIPAllocationDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_ip_allocation" {
			continue
		}
		Connector := meta.(*ibclient.Connector)
		objMgr := ibclient.NewObjectManager(Connector, "terraform_test", "test")
		recordName, _ := objMgr.GetFixedAddress("default", "10.0.0.0/24", "10.0.0.2", "")
		if recordName == nil {
			return fmt.Errorf("record not found")
		}
	}
	return nil
}
func testAccIPExists(n string, cidr string, ipAddr string, networkViewName string, recordName string) resource.TestCheckFunc {
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

		recordName, _ := objMgr.GetFixedAddress(networkViewName, cidr, ipAddr, "")
		if recordName != nil {
			return fmt.Errorf("record not found")
		}

		return nil
	}
}

var testAccresourceIPAllocationCreate = `
resource "infoblox_ip_allocation" "foo"{
	network_view_name="default"
	vm_name="test-name"
	cidr="10.0.0.0/24"
	ip_addr="10.0.0.1"
	tenant_id="foo"
	}`

var testAccresourceIPAllocationUpdate = `
resource "infoblox_ip_allocation" "foo"{
	network_view_name="default"
	vm_name="test-name"
	cidr="10.0.0.0/24"
	ip_addr="10.0.0.1"
	tenant_id="foo"
	}`
