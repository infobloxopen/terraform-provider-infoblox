package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

var testAccresourceIPv4NetworkCreate = fmt.Sprintf(`
resource "infoblox_ipv4_network" "foo"{
	network_view_name="default"
	cidr="10.10.0.0/24"
	comment = "10.0.0.0/24 network created"
	extensible_attributes = jsonencode({
		"Network Name"= "demo-network"
		"Tenant ID" = "terraform_test_tenant"
		Location = "Test loc."
		Site = "Test site"
		TestEA1 = ["text1","text2"]
		TestEA2 = [4,5]
	  })
	}`)

var testAccresourceIPv4NetworkUpdate = fmt.Sprintf(`
resource "infoblox_ipv4_network" "foo"{
	network_view_name="default"
	cidr="10.10.0.0/24"
	comment = "10.0.0.0/24 network updated"
	extensible_attributes = jsonencode({
		"Network Name"= "demo-network"
		"Tenant ID" = "terraform_test_tenant"
		Location = "Test loc. 2"
		Site = "Test site"
		TestEA1 = "text3"
		TestEA2 = 7
	  })
	}`)

var testAccresourceIPv6NetworkCreate = fmt.Sprintf(`
	resource "infoblox_ipv6_network" "foo"{
		network_view_name="default"
		cidr="2001:db8:abcd:12::/64"
		comment = "2001:db8:abcd:12::/64 network created"
		extensible_attributes = jsonencode({
			"Tenant ID" = "terraform_test_tenant"
			"Network Name"= "demo-network"
			Location = "Test loc."
			Site = "Test site"
			TestEA1 = ["text1","text2"]
			TestEA2 = [4,5]
		})
	}`)

var testAccresourceIPv6NetworkUpdate = fmt.Sprintf(`
	resource "infoblox_ipv6_network" "foo"{
		network_view_name="default"
		cidr="2001:db8:abcd:12::/64"
		comment = "2001:db8:abcd:12::/64 network updated"
		extensible_attributes = jsonencode({
			"Tenant ID" = "terraform_test_tenant"
			"Network Name"= "demo-network"
			Location = "Test loc. 2"
			Site = "Test site"
			TestEA1 = ["text3"]
			TestEA2 = 7
		})
	}`)

func validateNetwork(
	resourceName string,
	expectedValue *ibclient.Network) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resourceName]
		if !found {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id := res.Primary.ID
		if id == "" {
			return fmt.Errorf("ID is not set")
		}

		connector := testAccProvider.Meta().(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(
			connector,
			"terraform_test",
			"terraform_test_tenant")
		nw, err := objMgr.GetNetworkByRef(id)
		if err != nil {
			if isNotFoundError(err) {
				if expectedValue == nil {
					return nil
				}
				return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
			}
		}
		expNv := expectedValue.NetviewName
		if nw.NetviewName != expNv {
			return fmt.Errorf(
				"the value of 'network_view_name' field is '%s', but expected '%s'",
				nw.NetviewName, expNv)
		}

		expComment := expectedValue.Comment
		if nw.Comment != expComment {
			return fmt.Errorf(
				"the value of 'comment' field is '%s', but expected '%s'",
				nw.Comment, expComment)
		}

		// the rest is about extensible attributes
		expectedEAs := expectedValue.Ea
		if expectedEAs == nil && nw.Ea != nil {
			return fmt.Errorf(
				"the object with ID '%s' has 'extensible_attributes' field, but it is not expected to exist", id)
		}
		if expectedEAs != nil && nw.Ea == nil {
			return fmt.Errorf(
				"the object with ID '%s' has no 'extensible_attributes' field, but it is expected to exist", id)
		}
		if expectedEAs == nil {
			return nil
		}

		return validateEAs(nw.Ea, expectedEAs)
	}
}

func testAccCheckNetworkDestroy(s *terraform.State) error {
	connector := testAccProvider.Meta().(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(
		connector,
		"terraform_test",
		"terraform_test_tenant")
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_network" && rs.Type != "infoblox_ipv6_network" {
			continue
		}
		res, err := objMgr.GetNetworkByRef(rs.Primary.ID)
		if err != nil {
			if isNotFoundError(err) {
				continue
			}
			return err
		}
		if res != nil {
			return fmt.Errorf("object with ID '%s' remains", rs.Primary.ID)
		}
	}
	return nil
}

func TestAcc_resourceNetwork_ipv4(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccresourceIPv4NetworkCreate,
				Check: validateNetwork(
					"infoblox_ipv4_network.foo",
					&ibclient.Network{
						Cidr:    "10.0.0.0/24",
						Comment: "10.0.0.0/24 network created",
						Ea: ibclient.EA{
							"Network Name": "demo-network",
							"Tenant ID":    "terraform_test_tenant",
							"Location":     "Test loc.",
							"Site":         "Test site",
							"TestEA1":      []string{"text1", "text2"},
							"TestEA2":      []int{4, 5},
						},
					},
				),
			},
			{
				Config: testAccresourceIPv4NetworkUpdate,
				Check: validateNetwork(
					"infoblox_ipv4_network.foo",
					&ibclient.Network{
						Cidr:    "10.0.0.0/24",
						Comment: "10.0.0.0/24 network updated",
						Ea: ibclient.EA{
							"Network Name": "demo-network",
							"Tenant ID":    "terraform_test_tenant",
							"Location":     "Test loc. 2",
							// lists which contain ony one element are reduced by NIOS to a single-value element
							"TestEA1": "text3",
							"TestEA2": 7,
						},
					},
				),
			},
		},
	})
}

func TestAcc_resourceNetwork_ipv6(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccresourceIPv6NetworkCreate,
				Check: validateNetwork(
					"infoblox_ipv6_network.foo",
					&ibclient.Network{
						Cidr:    "2001:db8:abcd:12::/64",
						Comment: "2001:db8:abcd:12::/64 network created",
						Ea: ibclient.EA{
							"Network Name": "demo-network",
							"Tenant ID":    "terraform_test_tenant",
							"Location":     "Test loc.",
							"Site":         "Test site",
							"TestEA1":      []string{"text1", "text2"},
							"TestEA2":      []int{4, 5},
						},
					},
				),
			},
			{
				Config: testAccresourceIPv6NetworkUpdate,
				Check: validateNetwork(
					"infoblox_ipv6_network.foo",
					&ibclient.Network{
						Cidr:    "2001:db8:abcd:12::/64",
						Comment: "2001:db8:abcd:12::/64 network updated",
						Ea: ibclient.EA{
							"Network Name": "demo-network",
							"Tenant ID":    "terraform_test_tenant",
							"Location":     "Test loc. 2",
							// lists which contain ony one element are reduced by NIOS to a single-value element
							"TestEA1": "text3",
							"TestEA2": 7,
						},
					},
				),
			},
		},
	})
}
