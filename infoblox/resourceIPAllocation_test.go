package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/infobloxopen/infoblox-go-client"
	"testing"
)

func TestAccResourceIPAllocation(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPAllocationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccresourceIPAllocationCreate,
				Check: resource.ComposeTestCheckFunc(
					testAccIPExists(t, "infoblox_ip_allocation.foo", "10.0.0.1/24", "10.0.0.2", "test", "demo-network"),
				),
			},
			resource.TestStep{
				Config: testAccresourceIPAllocationUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccIPExists(t, "infoblox_ip_allocation.foo", "10.0.0.1/24", "10.0.0.2", "test", "demo-network"),
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
		recordName, _ := objMgr.GetHostRecord("test-name", "demo", "10.0.0.0/24", "10.0.0.2")
		if recordName == nil {
			return fmt.Errorf("record not found")
		}

	}
	return nil
}
func testAccIPExists(t *testing.T, n string, cidr string, ipAddr string, networkViewName string, recordName string) resource.TestCheckFunc {
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

		recordName, _ := objMgr.GetHostRecord(recordName, networkViewName, cidr, ipAddr)
		if recordName != nil {
			return fmt.Errorf("record not found")
		}

		return nil
	}
}

var testAccresourceIPAllocationCreate = fmt.Sprintf(`
resource "infoblox_ip_allocation" "foo"{
	network_view_name="test"
	vm_name="test-name"
	cidr="10.0.0.0/24"
	ip_addr="10.0.0.2"
	tenant_id="foo"
	}`)

var testAccresourceIPAllocationUpdate = fmt.Sprintf(`
resource "infoblox_ip_allocation" "foo"{
	network_view_name="test"
	vm_name="test-name"
	cidr="10.0.0.0/24"
	ip_addr="10.0.0.2"
	tenant_id="foo"
	}`)
