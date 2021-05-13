package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/infobloxopen/infoblox-go-client"
	"testing"
)

var testAccresourceNetworkView = fmt.Sprintf(`
	resource "infoblox_network_view" "foo"{
		network_view_name="test1"
		tenant_id="foo"
	}`)

func TestAccresourceNetworkView(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccresourceNetworkView,
				Check: resource.ComposeTestCheckFunc(
					testAccCreateNetworkViewExists(t, "infoblox_network_view.foo", "test"),
				),
			},
		},
	})
}

func testAccCreateNetworkViewExists(t *testing.T, n string, networkViewName string) resource.TestCheckFunc {
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

		netview, _ := objMgr.GetNetworkView(networkViewName)
		if netview != nil {
			return fmt.Errorf("Network View not found")
		}
		return nil
	}
}
