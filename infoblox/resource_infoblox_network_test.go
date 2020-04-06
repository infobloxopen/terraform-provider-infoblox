package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/infobloxopen/infoblox-go-client"
	"testing"
)

func TestAccresourceNetwork(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccresourceNetworkCreate,
				Check: resource.ComposeTestCheckFunc(
					testAccCreateNetworkExists(t, "infoblox_network.foo", "10.10.0.0/24", "default", "demo-network"),
				),
			},
			resource.TestStep{
				Config: testAccresourceNetworkUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCreateNetworkExists(t, "infoblox_network.foo", "10.10.0.0/24", "default", "demo-network"),
				),
			},
		},
	})
}

func testAccCheckNetworkDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_network" {
			continue
		}
		Connector := meta.(*ibclient.Connector)
		objMgr := ibclient.NewObjectManager(Connector, "terraform_test", "test")
		networkName, _ := objMgr.GetNetwork("demo-network", "10.10.0.0/24", nil)
		if networkName != nil {
			return fmt.Errorf("Network not found")
		}

	}
	return nil
}

func testAccCreateNetworkExists(t *testing.T, n string, cidr string, networkViewName string, networkName string) resource.TestCheckFunc {
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

		networkName, _ := objMgr.GetNetwork(networkName, cidr, nil)
		if networkName != nil {
			return fmt.Errorf("Network not found")
		}
		return nil
	}
}

var testAccresourceNetworkCreate = fmt.Sprintf(`
resource "infoblox_network" "foo"{
	network_view_name="default"
	network_name="demo-network"
	cidr="10.10.0.0/24"
	tenant_id="foo"
	}`)

var testAccresourceNetworkUpdate = fmt.Sprintf(`
resource "infoblox_network" "foo"{
	network_view_name="default"
	network_name="demo-network"
	cidr="10.10.0.0/24"
	tenant_id="foo"
	}`)
