package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/infobloxopen/infoblox-go-client"
	"testing"
)

func TestAccCreateIP(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPAllocationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCreateIP,
				Check: resource.ComposeTestCheckFunc(
					testAccIPExists(t, "createIP.foo", "", "10.0.0.1/24", "10.0.0.2", "test", "demo-network"),
				),
			},
			resource.TestStep{
				Config: testAccUpdateIP,
				Check: resource.ComposeTestCheckFunc(
					testAccIPExists(t, "UpdateIP.foo", "", "10.0.0.1/24", "10.0.0.2", "test", "demo-network"),
				),
			},
		},
	})
}

func testAccCheckIPAllocationDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "IPAssociation" {
			continue
		}
		Connector := meta.(*ibclient.Connector)
		objMgr := ibclient.NewObjectManager(Connector, "terraform_test", "test")
		recordName, _ := objMgr.GetHostRecordWithoutDNS("test-name", "demo", "10.0.0.1/24", "10.0.0.2")
		if recordName == nil {
			return fmt.Errorf("record not found")
		}

	}
	return nil
}
func testAccIPExists(t *testing.T, n string, m interface{}, cidr string, ipAddr string, networkViewName string, recordName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found:%s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID i set")
		}

		Connector := m.(*ibclient.Connector)
		objMgr := ibclient.NewObjectManager(Connector, "terraform_test", "test")

		recordName, _ := objMgr.GetHostRecordWithoutDNS(recordName, networkViewName, cidr, ipAddr)
		if recordName == nil {
			return fmt.Errorf("record not found")
		}

		return nil
	}
}

var testAccCreateIP = fmt.Sprintf(`
resource "createIP" "foo"{
	networkViewName="test"
	recordName="test-name"
	cidr="10.0.0.1/24"
	ipAddr="10.0.0.2"
	tenant_id="foo"
	}`)

var testAccUpdateIP = fmt.Sprintf(`
resource "UpdateIP" "foo"{
	networkViewName="test"
	recordName="test-name"
	cidr="10.0.0.1/24"
	ipAddr="10.0.0.2"
	tenant_id="foo"
	}`)
