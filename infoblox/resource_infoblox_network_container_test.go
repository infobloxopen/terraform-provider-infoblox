package infoblox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func validateNetworkContainer(
	resourceName string,
	expectedValue *ibclient.NetworkContainer) resource.TestCheckFunc {
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

		nc, err := objMgr.GetNetworkContainerByRef(id)
		if err != nil {
			if isNotFoundError(err) {
				if expectedValue == nil {
					return nil
				}
				return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
			}
		}

		expNv := expectedValue.NetviewName
		if nc.NetviewName != expNv {
			return fmt.Errorf(
				"the value of 'network_view' field is '%s', but expected '%s'",
				nc.NetviewName, expNv)
		}

		expComment := expectedValue.Comment
		if nc.Comment != expComment {
			return fmt.Errorf(
				"the value of 'comment' field is '%s', but expected '%s'",
				nc.Comment, expComment)
		}

		expCidr := expectedValue.Cidr
		//cidr is not passed in nextavailable network container test case
		if expCidr == "" {
			expCidr = res.Primary.Attributes["cidr"]
			if expCidr == "" {
				return fmt.Errorf(
					"the value of 'cidr' field is empty, but expected some value")
			}
		}

		if nc.Cidr != expCidr {
			return fmt.Errorf(
				"the value of 'cidr' field is '%s', but expected '%s'",
				nc.Cidr, expCidr)
		}

		// the rest is about extensible attributes
		expectedEAs := expectedValue.Ea
		if expectedEAs == nil && nc.Ea != nil {
			return fmt.Errorf(
				"the object with ID '%s' has 'ext_attrs' field, but it is not expected to exist", id)
		}
		if expectedEAs != nil && nc.Ea == nil {
			return fmt.Errorf(
				"the object with ID '%s' has no 'ext_attrs' field, but it is expected to exist", id)
		}
		if expectedEAs == nil {
			return nil
		}

		return validateEAs(nc.Ea, expectedEAs)
	}
}

var updateNotAllowedRegexp = regexp.MustCompile("changing the value of '.+' field is not allowed")

func TestAcc_resourceNetworkContainer_ipv4(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkContainerDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "infoblox_ipv4_network_container" "nc_1" {
					  network_view = "default"
					  cidr = "10.0.0.0/16"
					  comment = "10.0.0.0/16 network container"
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						Location = "Test location"
						Site = "Test site"
					  })
					}`,
				Check: validateNetworkContainer(
					"infoblox_ipv4_network_container.nc_1",
					&ibclient.NetworkContainer{
						NetviewName: "default",
						Cidr:        "10.0.0.0/16",
						Comment:     "10.0.0.0/16 network container",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test location",
							"Site":      "Test site",
						},
					},
				),
			},
			{
				Config: `
					resource "infoblox_ipv4_network_container" "nc_1" {
					  network_view = "default"
					  cidr = "10.0.0.0/16"
					  comment = "new 10.0.0.0/16 network container"
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						Location = "Test loc. 2"
					  })
					}`,
				Check: validateNetworkContainer(
					"infoblox_ipv4_network_container.nc_1",
					&ibclient.NetworkContainer{
						NetviewName: "default",
						Cidr:        "10.0.0.0/16",
						Comment:     "new 10.0.0.0/16 network container",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc. 2",
						},
					},
				),
			},
			{
				Config: `
					resource "infoblox_ipv4_network_container" "nc_1" {
					  network_view = "default"
					  cidr = "10.0.0.0/16"
					  comment = ""
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
					  })
					}`,
				Check: validateNetworkContainer(
					"infoblox_ipv4_network_container.nc_1",
					&ibclient.NetworkContainer{
						NetviewName: "default",
						Cidr:        "10.0.0.0/16",
						Comment:     "",
						Ea:          ibclient.EA{},
					},
				),
			},
			{
				Config: `
					resource "infoblox_ipv4_network_container" "nc_2" {
					  network_view = "default"
					  cidr = "25.0.0.0/24"
					  comment = "25.0.0.0/24 network container"
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						Location = "Test location"
						Site = "Test site"
					  })
					}
					
					resource "infoblox_ipv4_network_container" "nc_3" {
					  network_view = infoblox_ipv4_network_container.nc_2.network_view
					  parent_cidr = infoblox_ipv4_network_container.nc_2.cidr
					  allocate_prefix_len = 26
					  comment = "dynamic creation of network container"
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						Site = "Test site"
					  })
					}`,
				Check: validateNetworkContainer(
					"infoblox_ipv4_network_container.nc_3",
					&ibclient.NetworkContainer{
						NetviewName: "default",
						Comment:     "dynamic creation of network container",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Site":      "Test site",
						},
					},
				),
			},
			{
				Config: `
					resource "infoblox_ipv4_network_container" "nc_2" {
					  network_view = "default"
					  cidr = "25.0.0.0/24"
					  comment = "25.0.0.0/24 network container"
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						Location = "Test location"
						Site = "Test site"
					  })
					}
					
					resource "infoblox_ipv4_network_container" "nc_3" {
					  network_view = infoblox_ipv4_network_container.nc_2.network_view
					  parent_cidr = infoblox_ipv4_network_container.nc_2.cidr
					  allocate_prefix_len = 26
					  comment = ""
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
					  })
					}`,
				Check: validateNetworkContainer(
					"infoblox_ipv4_network_container.nc_3",
					&ibclient.NetworkContainer{
						NetviewName: "default",
						Comment:     "",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
						},
					},
				),
			},
			{
				Config: `
					resource "infoblox_ipv4_network_container" "nc_2" {
					  network_view = "default"
					  cidr = "25.0.0.0/24"
					  comment = "25.0.0.0/24 network container"
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						Location = "Test location"
						Site = "Test site"
					  })
					}
					
					resource "infoblox_ipv4_network_container" "nc_3" {
					  network_view = infoblox_ipv4_network_container.nc_2.network_view
					  parent_cidr = infoblox_ipv4_network_container.nc_2.cidr
					  allocate_prefix_len = 28
					  comment = ""
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
					  })
					}`,
				ExpectError: updateNotAllowedRegexp,
			},
			{
				Config: `
					resource "infoblox_network_view" "nv1" {
						name = "non_default_view_54397156"
					}
					resource "infoblox_ipv4_network_container" "nc_4" {
					  network_view = infoblox_network_view.nv1.name
					  cidr = "26.0.0.0/24"
					}
					resource "infoblox_ipv4_network_container" "nc_5" {
					  network_view = infoblox_ipv4_network_container.nc_4.network_view
					  parent_cidr = infoblox_ipv4_network_container.nc_4.cidr
					  allocate_prefix_len = 26
					  comment = "comment 1"
					  ext_attrs = jsonencode({
						"Site" = "North Pole"
					  })
					}`,
				Check: validateNetworkContainer(
					"infoblox_ipv4_network_container.nc_5",
					&ibclient.NetworkContainer{
						NetviewName: "non_default_view_54397156",
						Comment:     "comment 1",
						Ea: ibclient.EA{
							"Site": "North Pole",
						},
					},
				),
			},
			{
				Config: `
					resource "infoblox_network_view" "nv1" {
						name = "non_default_view_54397156"
					}
					resource "infoblox_ipv4_network_container" "nc_4" {
					  network_view = infoblox_network_view.nv1.name
					  cidr = "26.0.0.0/24"
					}
					resource "infoblox_ipv4_network_container" "nc_5" {
					  network_view = infoblox_ipv4_network_container.nc_4.network_view
					  parent_cidr = infoblox_ipv4_network_container.nc_4.cidr
					  allocate_prefix_len = 26
					  comment = ""
					}`,
				Check: validateNetworkContainer(
					"infoblox_ipv4_network_container.nc_5",
					&ibclient.NetworkContainer{
						NetviewName: "non_default_view_54397156",
						Comment:     "",
						Ea:          ibclient.EA{},
					},
				),
			},
		},
	})
}

func TestAcc_resourceNetworkContainer_ipv6(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkContainerDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "infoblox_ipv6_network_container" "nc_1" {
					  network_view = "default"
					  cidr = "fc00::/56"
					  comment = "fc00::/56 network container"
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						Location = "Test loc."
						Site = "Test site"
					  })
					}`,
				Check: validateNetworkContainer(
					"infoblox_ipv6_network_container.nc_1",
					&ibclient.NetworkContainer{
						NetviewName: "default",
						Cidr:        "fc00::/56",
						Comment:     "fc00::/56 network container",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc.",
							"Site":      "Test site",
						},
					},
				),
			},
			{
				Config: `
					resource "infoblox_ipv6_network_container" "nc_1" {
					  network_view = "default"
					  cidr = "fc00::/56"
					  comment = "new comment for fc00::/56 network container"
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						Location = "Test loc. 2"
					  })
					}`,
				Check: validateNetworkContainer(
					"infoblox_ipv6_network_container.nc_1",
					&ibclient.NetworkContainer{
						NetviewName: "default",
						Cidr:        "fc00::/56",
						Comment:     "new comment for fc00::/56 network container",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc. 2",
						},
					},
				),
			},
			{
				Config: `
					resource "infoblox_ipv6_network_container" "nc6_2" {
					  network_view = "default"
					  cidr = "fc01::/56"
					  comment = "fc01::/56 network container"
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						Site = "Test site"
					  })
					}
					
					resource "infoblox_ipv6_network_container" "nc6_3" {
					  network_view = infoblox_ipv6_network_container.nc6_2.network_view
					  parent_cidr = infoblox_ipv6_network_container.nc6_2.cidr
					  allocate_prefix_len = 58
					  comment = "fc01::/58 dynamic network container"
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						Site = "Test site"
					  })
					}`,
				Check: validateNetworkContainer(
					"infoblox_ipv6_network_container.nc6_3",
					&ibclient.NetworkContainer{
						NetviewName: "default",
						Comment:     "fc01::/58 dynamic network container",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Site":      "Test site",
						},
					},
				),
			},
			{
				Config: `
					resource "infoblox_ipv6_network_container" "nc6_2" {
					  network_view = "default"
					  cidr = "fc01::/56"
					  comment = "fc01::/56 network container"
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						Site = "Test site"
					  })
					}
					
					resource "infoblox_ipv6_network_container" "nc6_3" {
					  network_view = infoblox_ipv6_network_container.nc6_2.network_view
					  parent_cidr = infoblox_ipv6_network_container.nc6_2.cidr
					  allocate_prefix_len = 58
					  comment = "dynamic network container testing"
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						Site = "Test site"
					  })
					}`,
				Check: validateNetworkContainer(
					"infoblox_ipv6_network_container.nc6_3",
					&ibclient.NetworkContainer{
						NetviewName: "default",
						Comment:     "dynamic network container testing",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Site":      "Test site",
						},
					},
				),
			},
			{
				Config: `
					resource "infoblox_ipv6_network_container" "nc6_2" {
					  network_view = "default"
					  cidr = "fc01::/56"
					  comment = "fc01::/56 network container"
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						Site = "Test site"
					  })
					}
					
					resource "infoblox_ipv6_network_container" "nc6_3" {
					  network_view = infoblox_ipv6_network_container.nc6_2.network_view
					  parent_cidr = infoblox_ipv6_network_container.nc6_2.cidr
					  allocate_prefix_len = 59
					  comment = "dynamic network container testing"
					  ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						Site = "Test site"
					  })
					}`,
				ExpectError: updateNotAllowedRegexp,
			},
			{
				Config: `
					resource "infoblox_network_view" "nv1" {
						name = "non_default_view_54397156"
					}
					resource "infoblox_ipv6_network_container" "nc6_4" {
					  network_view = infoblox_network_view.nv1.name
					  cidr = "fc02::/56"
					}
					resource "infoblox_ipv6_network_container" "nc6_5" {
					  network_view = infoblox_ipv6_network_container.nc6_4.network_view
					  parent_cidr = infoblox_ipv6_network_container.nc6_4.cidr
					  allocate_prefix_len = 58
					  comment = "dynamic IPv6 network container testing"
					  ext_attrs = jsonencode({
						"Site" = "Antarctiсa"
					  })
					}`,
				Check: validateNetworkContainer(
					"infoblox_ipv6_network_container.nc6_5",
					&ibclient.NetworkContainer{
						NetviewName: "non_default_view_54397156",
						Comment:     "dynamic IPv6 network container testing",
						Ea: ibclient.EA{
							"Site": "Antarctiсa",
						},
					},
				),
			},
			{
				Config: `
					resource "infoblox_network_view" "nv1" {
						name = "non_default_view_54397156"
					}
					resource "infoblox_ipv6_network_container" "nc6_4" {
					  network_view = infoblox_network_view.nv1.name
					  cidr = "fc02::/56"
					}
					resource "infoblox_ipv6_network_container" "nc6_5" {
					  network_view = infoblox_ipv6_network_container.nc6_4.network_view
					  parent_cidr = infoblox_ipv6_network_container.nc6_4.cidr
					  allocate_prefix_len = 58
					  ext_attrs = jsonencode({})
					}`,
				Check: validateNetworkContainer(
					"infoblox_ipv6_network_container.nc6_5",
					&ibclient.NetworkContainer{
						NetviewName: "non_default_view_54397156",
						Comment:     "",
						Ea:          ibclient.EA{},
					},
				),
			},
		},
	})
}

func testAccCheckNetworkContainerDestroy(s *terraform.State) error {
	connector := testAccProvider.Meta().(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(
		connector,
		"terraform_test",
		"terraform_test_tenant")
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_ipv4_network_container" && rs.Type != "infoblox_ipv6_network_container" {
			continue
		}
		res, err := objMgr.GetNetworkContainerByRef(rs.Primary.ID)
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
