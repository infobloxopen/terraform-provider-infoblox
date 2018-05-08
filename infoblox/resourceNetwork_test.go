package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/infobloxopen/infoblox-go-client"
	"testing"
)

func TestAccCreateNetwork(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCreateNetwork,
				Check: resource.ComposeTestCheckFunc(
					testAccCreateNetworkExists(t, "createNetwork.foo", "", "10.0.0.1/24", "test", "demo-network"),
				),
			},
			resource.TestStep{
				Config: testAccUpdateNetwork,
				Check: resource.ComposeTestCheckFunc(
					testAccCreateNetworkExists(t, "createNetwork.foo", "", "10.0.0.1/24", "test", "demo-network"),
				),
			},
		},
	})
}

func testAccCheckNetworkDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "Network" {
			continue
		}
		Connector := meta.(*ibclient.Connector)
		objMgr := ibclient.NewObjectManager(Connector, "terraform_test", "test")
		networkName, _ := objMgr.GetNetwork("demo-network", "10.0.0.1/24", nil)
		if networkName == nil {
			return fmt.Errorf("Network not found")
		}

	}
	return nil
}

func testAccCreateNetworkExists(t *testing.T, n string, m interface{}, cidr string, networkViewName string, networkName string) resource.TestCheckFunc {
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

		networkName, _ := objMgr.GetNetwork(networkName, cidr, nil)
		if networkName == nil {
			return fmt.Errorf("Network not found")
		}
		return nil
	}
}

var testAccCreateNetwork = fmt.Sprintf(`
resource "createNetwork" "foo"{
	networkViewName="test"
	networkName="demo-network"
	cidr="10.0.0.1/24"
	tenant_id="foo"
	}`)

var testAccUpdateNetwork = fmt.Sprintf(`
resource "UpdateNetwork" "foo"{
	networkViewName="test"
	networkName="demo-network"
	cidr="10.0.0.1/24"
	tenant_id="foo"
	}`)
