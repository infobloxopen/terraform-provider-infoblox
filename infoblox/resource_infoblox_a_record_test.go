package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func TestAccResourceARecord(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckARecordDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccresourceARecordCreate,
				Check: resource.ComposeTestCheckFunc(
					testAccARecordExists(t, "infoblox_a_record.foo", "10.0.0.0/24", "10.0.0.2", "test", "demo-network", "default", "a.com"),
				),
			},
			resource.TestStep{
				Config: testAccresourceARecordAllocate,
				Check: resource.ComposeTestCheckFunc(
					testAccARecordExists(t, "infoblox_a_record.foo1", "10.0.0.0/24", "10.0.0.1", "test", "demo-network", "default", "a.com"),
					testAccARecordExists(t, "infoblox_a_record.foo2", "10.0.0.0/24", "10.0.0.2", "test", "demo-network", "default", "a.com"),
				),
			},
			resource.TestStep{
				Config: testAccresourceARecordUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccARecordExists(t, "infoblox_a_record.foo", "10.0.0.0/24", "10.0.0.2", "test", "demo-network", "default", "a.com"),
				),
			},
		},
	})
}

func testAccCheckARecordDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_a_record" {
			continue
		}
		Connector := meta.(*ibclient.Connector)
		objMgr := ibclient.NewObjectManager(Connector, "terraform_test", "test")
		recordName, _ := objMgr.GetARecordByRef(rs.Primary.ID)
		if recordName != nil {
			return fmt.Errorf("record not found")
		}

	}
	return nil
}
func testAccARecordExists(t *testing.T, n string, cidr string, ipAddr string, networkViewName string, recordName string, dnsView string, zone string) resource.TestCheckFunc {
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

		recordName, _ := objMgr.GetARecordByRef(rs.Primary.ID)
		if recordName == nil {
			return fmt.Errorf("record not found")
		}

		return nil
	}
}

var testAccresourceARecordCreate = fmt.Sprintf(`
resource "infoblox_a_record" "foo"{
	vm_name="test-name"
	zone="a.com"
	ip_addr="10.0.0.2"
	tenant_id="foo"
	}`)

var testAccresourceARecordAllocate = fmt.Sprintf(`
resource "infoblox_a_record" "foo1"{
	vm_name="test-name"
	zone="a.com"
	ip_addr=""
	cidr="10.0.0.0/24"
	tenant_id="foo"
	}
resource "infoblox_a_record" "foo2"{
	vm_name="test-name"
	zone="a.com"
	ip_addr=""
	cidr="10.0.0.0/24"
	tenant_id="foo"
	}`)

var testAccresourceARecordUpdate = fmt.Sprintf(`
resource "infoblox_a_record" "foo"{
	vm_name="test-name"
	dns_view="default"
	zone="a.com"
	cidr="10.0.0.0/24"
	ip_addr="10.0.0.2"
	tenant_id="foo"
	}`)
