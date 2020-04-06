package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/infobloxopen/infoblox-go-client"
	"testing"
)

func TestAccresourceIPAssociation(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRecordHostDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccresourceIPAssociationCreate,
				Check: resource.ComposeTestCheckFunc(
					testAccRecordHostExists(t, "infoblox_ip_association.foo", "10.0.0.0/24", "10.0.0.2", "default", "demo-network"),
				),
			},
			resource.TestStep{
				Config: testAccresourceIPAssociationUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccRecordHostExists(t, "infoblox_ip_association.foo", "10.0.0.0/24", "10.0.0.22", "default", "demo-network"),
				),
			},
		},
	})
}

func testAccCheckRecordHostDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_ip_association" {
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
func testAccRecordHostExists(t *testing.T, n string, cidr string, ipAddr string, networkViewName string, recordName string) resource.TestCheckFunc {
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

		recordName, _ := objMgr.GetFixedAddress("default", "10.0.0.0/24", "10.0.0.2", "")
		if recordName == nil {
			return fmt.Errorf("record not found")
		}

		return nil
	}
}

var testAccresourceIPAssociationCreate = fmt.Sprintf(`
resource "infoblox_ip_association" "foo"{
	network_view_name="default"
	vm_name="test-name"
	mac_addr="11:22:33:44:55:66"
	cidr="10.0.0.0/24"
	ip_addr="10.0.0.2"
	tenant_id="foo"
	}`)

var testAccresourceIPAssociationUpdate = fmt.Sprintf(`
resource "infoblox_ip_association" "foo"{
	network_view_name="default"
	vm_name="test-name"
	cidr="10.0.0.0/24"
	ip_addr="10.0.0.2"
	mac_addr="12:22:33:44:55:66"
	tenant_id="foo"
	}`)
