package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func TestAccresourceIPv6Network(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPv6NetworkDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: TestAccresourceIPv6NetworkCreate,
				Check: resource.ComposeTestCheckFunc(
					testAccCreateIPv6NetworkExists(t, "infoblox_ipv6network", "2600:1f18:a0a:f700::/64", "default", "demo-network"),
				),
			},
			resource.TestStep{
				Config: TestAccresourceIPv6NetworkUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCreateIPv6NetworkExists(t, "infoblox_ipv6network", "2600:1f18:a0a:f700::/64", "default", "demo-network"),
				),
			},
		},
	})
}

func testAccCreateIPv6NetworkExists(t *testing.T, n string, cidr string, networkViewName string, networkName string) resource.TestCheckFunc {
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

func testAccCheckIPv6NetworkDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infobloxIPv6Network" {
			continue
		}
		Connector := meta.(*ibclient.Connector)
		objMgr := ibclient.NewObjectManager(Connector, "terraform_test", "test")
		networkName, _ := objMgr.GetIPv6Network("demo-network", "2600:1f18:a0a:f700::/64", nil)
		if networkName != nil {
			return fmt.Errorf("Network not found")
		}
	}
	return nil
}

var TestAccresourceIPv6NetworkCreate = fmt.Sprintf(`
resource "infoblox_ipv6network" "foo"{
	network_view_name = "default"
	network_name = "demo-network"
	cidr = "2600:1f18:a0a:f700::/64"
	tenant_id = "foo"
	}`)

var TestAccresourceIPv6NetworkUpdate = fmt.Sprintf(`
resource "infoblox_ipv6network" "foo"{
	network_view_name = "default"
	network_name = "demo-network"
	cidr = "2600:1f18:a0a:f700::/64"
	tenant_id = "foo" 
	}`)
