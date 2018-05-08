package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/infobloxopen/infoblox-go-client"
	"testing"
)

func TestAccCreateRecordHost(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRecordHostDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCreateRecordHost,
				Check: resource.ComposeTestCheckFunc(
					testAccRecordHostExists(t, "createRecordHost.foo", "", "10.0.0.1/24", "10.0.0.2", "test", "demo-network"),
				),
			},
			resource.TestStep{
				Config: testAccUpdateRecordHost,
				Check: resource.ComposeTestCheckFunc(
					testAccRecordHostExists(t, "UpdateRecordHost.foo", "", "10.0.0.1/24", "10.0.0.2", "test", "demo-network"),
				),
			},
		},
	})
}

func testAccCheckRecordHostDestroy(s *terraform.State) error {
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
func testAccRecordHostExists(t *testing.T, n string, m interface{}, cidr string, ipAddr string, networkViewName string, recordName string) resource.TestCheckFunc {
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

var testAccCreateRecordHost = fmt.Sprintf(`
resource "createRecordHost" "foo"{
	networkViewName="test"
	recordName="test-name"
	cidr="10.0.0.1/24"
	ipAddr="10.0.0.2"
	tenant_id="foo"
	}`)

var testAccUpdateRecordHost = fmt.Sprintf(`
resource "UpdateRecordHost" "foo"{
	networkViewName="test"
	recordName="test-name"
	cidr="10.0.0.1/24"
	ipAddr="10.0.0.2"
	tenant_id="foo"
	}`)
