package infoblox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

func TestAccresourceNetwork(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccresourceNetworkCreate,
				Check: resource.ComposeTestCheckFunc(
					testAccCreateNetworkExists("infoblox_network.foo", "10.10.0.0/24", "default", "demo-network"),
				),
			},
			{
				Config: testAccresourceNetworkUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCreateNetworkExists("infoblox_network.foo", "10.10.0.0/24", "default", "demo-network"),
				),
			},
		},
	})
}

func TestAccresourceNetwork_Allocate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccresourceNetworkAllocate,
				Check: resource.ComposeTestCheckFunc(
					testAccCreateNetworkExists("infoblox_network.foo0", "10.0.0.0/24", "default", "demo-network"),
					testAccCreateNetworkExists("infoblox_network.foo1", "10.0.1.0/24", "default", "demo-network"),
				),
			},
		},
	})
}

func TestAccresourceNetwork_Allocate_Fail(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccresourceNetworkAllocateFail,
				ExpectError: regexp.MustCompile(`Allocation of network block failed in network view \(default\) : Parent network container 11.11.0.0/16 not found.`),
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

func testAccCreateNetworkExists(n string, cidr string, networkViewName string, networkName string) resource.TestCheckFunc {
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

var testAccresourceNetworkCreate = `
resource "infoblox_network" "foo"{
	network_view_name="default"
	network_name="demo-network"
	cidr="10.10.0.0/24"
	tenant_id="foo"
	}`

/*
Right now no infoblox_network_container resource available
So, before run acceptance test TestAccresourceNetwork_Allocate
in default network view should be created network container 10.0.0.0/16
*/
var testAccresourceNetworkAllocate = `
resource "infoblox_network" "foo0"{
	network_view_name="default"
	network_name="demo-network"
	tenant_id="foo"
	allocate_prefix_len=24
	parent_cidr="10.0.0.0/16"
	}
resource "infoblox_network" "foo1"{
	network_view_name="default"
	network_name="demo-network"
	tenant_id="foo"
	allocate_prefix_len=24
	parent_cidr="10.0.0.0/16"
	}`

/* Network container 11.11.0.0 should NOT exists to pass this test */
var testAccresourceNetworkAllocateFail = `
resource "infoblox_network" "foo0"{
	network_view_name="default"
	network_name="demo-network"
	tenant_id="foo"
	allocate_prefix_len=24
	parent_cidr="11.11.0.0/16"
	}`

var testAccresourceNetworkUpdate = `
resource "infoblox_network" "foo"{
	network_view_name="default"
	network_name="demo-network"
	cidr="10.10.0.0/24"
	tenant_id="foo"
	}`
