package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"regexp"
	"testing"
)

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

		internalId := res.Primary.Attributes["internal_id"]
		if internalId == "" {
			return fmt.Errorf("ID is not set")
		}

		ref, found := res.Primary.Attributes["ref"]
		if !found {
			return fmt.Errorf("'ref' attribute is not set")
		}

		connector := testAccProvider.Meta().(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(
			connector,
			"terraform_test",
			"test")
		nwObj, err := objMgr.SearchObjectByAltId("Network", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedValue == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var nw *ibclient.Network
		recJson, _ := json.Marshal(nwObj)
		err = json.Unmarshal(recJson, &nw)

		expNv := expectedValue.NetviewName
		if nw.NetviewName != expNv {
			return fmt.Errorf(
				"the value of 'network_view' field is '%s', but expected '%s'",
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
				"the object with ID '%s' has 'ext_attrs' field, but it is not expected to exist", id)
		}
		if expectedEAs != nil && nw.Ea == nil {
			return fmt.Errorf(
				"the object with ID '%s' has no 'ext_attrs' field, but it is expected to exist", id)
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

var updateNotAllowedErrorRegexp = regexp.MustCompile("changing the value of '.+' field is not allowed")

func TestAcc_resourceNetwork_ipv4(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "infoblox_ipv4_network" "foo"{
						network_view="default"
						cidr="10.10.0.0/24"
						reserve_ip = 5
						gateway = "10.10.0.250"
						comment = "10.0.0.0/24 network created"
						ext_attrs = jsonencode({
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
					}`,
				Check: validateNetwork(
					"infoblox_ipv4_network.foo",
					&ibclient.Network{
						NetviewName: "default",
						Cidr:        "10.0.0.0/24",
						Comment:     "10.0.0.0/24 network created",
						Ea: ibclient.EA{
							"Network Name": "demo-network",
							"Tenant ID":    "terraform_test_tenant",
							"Location":     "Test loc.",
							"Site":         "Test site",
						},
					},
				),
			},
			{
				// Terraform provider should be able to update the network object
				Config: `
					resource "infoblox_ipv4_network" "foo"{
						network_view="default"
						cidr="10.10.0.0/24"
						reserve_ip = 5
						gateway = "10.10.0.250"
						comment = "Updated comment"
						ext_attrs = jsonencode({
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
							"Building" = "Test Building"
						})
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_ipv4_network.foo", "comment", "Updated comment"),
					resource.TestCheckResourceAttr(
						"infoblox_ipv4_network.foo", "ext_attrs",
						`{"Building":"Test Building","Location":"Test loc.","Network Name":"demo-network","Site":"Test site","Tenant ID":"terraform_test_tenant"}`,
					),
				),
			},
			{
				Config: `
					resource "infoblox_ipv4_network" "foo"{
						network_view="default"
						cidr="10.10.0.0/24"
						reserve_ip = 6
						gateway = "10.10.0.250"
						comment = "10.0.0.0/24 network created"
						ext_attrs = jsonencode({
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
					}`,
				ExpectError: updateNotAllowedErrorRegexp,
			},
			{
				Config: `
					resource "infoblox_ipv4_network" "foo"{
						network_view="default"
						cidr="10.10.0.0/24"
						reserve_ip = 6
						gateway = "10.10.0.250"
						comment = "10.0.0.0/24 network created"
						ext_attrs = jsonencode({
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
					}`,
				// double-check that the next update (with the same changes) returns an error as well
				// (in case the field to be updated is 'computed' and the main code do not clear it to the previous state)
				ExpectError: updateNotAllowedErrorRegexp,
			},
			{
				Config: `
					resource "infoblox_ipv4_network" "foo"{
						network_view="default"
						cidr="10.10.0.0/24"
						reserve_ip = 5
						gateway = "10.10.0.251"
						comment = "10.0.0.0/24 network created"
						ext_attrs = jsonencode({
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
					}`,
				ExpectError: updateNotAllowedErrorRegexp,
			},
			{
				Config: `
					resource "infoblox_ipv4_network" "foo"{
						network_view="default"
						cidr="10.10.0.0/24"
						reserve_ip = 5
						gateway = "10.10.0.251"
						comment = "10.0.0.0/24 network created"
						ext_attrs = jsonencode({
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
					}`,
				ExpectError: updateNotAllowedErrorRegexp,
			},
		},
	})
}

// TestAcc_resourceNetwork_ipv4_ea_inheritance validates that in case of EA
// inheritance, terraform doesn't remove EAs, that set on the NIOS side.
func TestAcc_resourceNetwork_ipv4_ea_inheritance(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "infoblox_ipv4_network" "foo"{
						network_view="default"
						cidr="10.10.0.0/24"
						reserve_ip = 5
						gateway = "10.10.0.250"
						comment = "10.0.0.0/24 network created"
						ext_attrs = jsonencode({
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
					}`,
				Check: validateNetwork(
					"infoblox_ipv4_network.foo",
					&ibclient.Network{
						NetviewName: "default",
						Cidr:        "10.0.0.0/24",
						Comment:     "10.0.0.0/24 network created",
						Ea: ibclient.EA{
							"Network Name": "demo-network",
							"Tenant ID":    "terraform_test_tenant",
							"Location":     "Test loc.",
							"Site":         "Test site",
						},
					},
				),
			},
			// When extensible attributes are added by another tool,
			// terraform shouldn't remove those EAs
			{
				PreConfig: func() {
					conn := testAccProvider.Meta().(ibclient.IBConnector)

					n := &ibclient.Ipv4Network{}
					n.SetReturnFields(append(n.ReturnFields(), "extattrs"))

					qp := ibclient.NewQueryParams(
						false,
						map[string]string{
							"network":      "10.10.0.0/24",
							"network_view": "default",
						},
					)
					var res []ibclient.Ipv4Network
					err := conn.GetObject(n, "", qp, &res)
					if err != nil {
						panic(err)
					}

					res[0].NetworkView = ""
					res[0].Ea["Building"] = "Test Building"

					_, err = conn.UpdateObject(&res[0], res[0].Ref)
					if err != nil {
						panic(err)
					}
				},
				Config: `
					resource "infoblox_ipv4_network" "foo"{
						network_view="default"
						cidr="10.10.0.0/24"
						reserve_ip = 5
						gateway = "10.10.0.250"
						comment = "10.0.0.0/24 network created"
						ext_attrs = jsonencode({
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
					}`,
				Check: resource.ComposeTestCheckFunc(
					// Resource object shouldn't have Building EA, since it's omitted by provider
					resource.TestCheckResourceAttr(
						"infoblox_ipv4_network.foo", "ext_attrs",
						`{"Location":"Test loc.","Network Name":"demo-network","Site":"Test site","Tenant ID":"terraform_test_tenant"}`,
					),
					// Actual API object should have Building EA
					validateNetwork(
						"infoblox_ipv4_network.foo",
						&ibclient.Network{
							NetviewName: "default",
							Cidr:        "10.0.0.0/24",
							Comment:     "10.0.0.0/24 network created",
							Ea: ibclient.EA{
								"Network Name": "demo-network",
								"Tenant ID":    "terraform_test_tenant",
								"Location":     "Test loc.",
								"Site":         "Test site",
								"Building":     "Test Building",
							},
						},
					),
				),
			},
			// Validate that inherited EA won't be removed if some field is updated in the resource
			{
				Config: `
					resource "infoblox_ipv4_network" "foo"{
						network_view="default"
						cidr="10.10.0.0/24"
						reserve_ip = 5
						gateway = "10.10.0.250"
						comment = "Updated comment"
						ext_attrs = jsonencode({
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
					}`,
				Check: validateNetwork(
					"infoblox_ipv4_network.foo",
					&ibclient.Network{
						NetviewName: "default",
						Cidr:        "10.0.0.0/24",
						Comment:     "Updated comment",
						Ea: ibclient.EA{
							"Network Name": "demo-network",
							"Tenant ID":    "terraform_test_tenant",
							"Location":     "Test loc.",
							"Site":         "Test site",
							"Building":     "Test Building",
						},
					},
				),
			},
			// Validate that inherited EA can be updated
			{
				Config: `
					resource "infoblox_ipv4_network" "foo"{
						network_view="default"
						cidr="10.10.0.0/24"
						reserve_ip = 5
						gateway = "10.10.0.250"
						comment = "10.0.0.0/24 network created"
						ext_attrs = jsonencode({
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
							"Building" = "Test building 2"
						})
					}`,
				Check: validateNetwork(
					"infoblox_ipv4_network.foo",
					&ibclient.Network{
						NetviewName: "default",
						Cidr:        "10.0.0.0/24",
						Comment:     "10.0.0.0/24 network created",
						Ea: ibclient.EA{
							"Network Name": "demo-network",
							"Tenant ID":    "terraform_test_tenant",
							"Location":     "Test loc.",
							"Site":         "Test site",
							"Building":     "Test building 2",
						},
					},
				),
			},
			// Validate that inherited EA can be removed, if updated
			{
				Config: `
					resource "infoblox_ipv4_network" "foo"{
						network_view="default"
						cidr="10.10.0.0/24"
						reserve_ip = 5
						gateway = "10.10.0.250"
						comment = "10.0.0.0/24 network created"
						ext_attrs = jsonencode({
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
					}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"infoblox_ipv4_network.foo", "ext_attrs",
						`{"Location":"Test loc.","Network Name":"demo-network","Site":"Test site","Tenant ID":"terraform_test_tenant"}`,
					),
					func(s *terraform.State) error {
						conn := testAccProvider.Meta().(ibclient.IBConnector)

						res, found := s.RootModule().Resources["infoblox_ipv4_network.foo"]
						if !found {
							return fmt.Errorf("not found: %s", "infoblox_ipv4_network.foo")
						}

						id := res.Primary.ID
						if id == "" {
							return fmt.Errorf("ID is not set")
						}

						objMgr := ibclient.NewObjectManager(
							conn,
							"terraform_test",
							"terraform_test_tenant")
						nw, err := objMgr.GetNetworkByRef(id)
						if err != nil {
							if isNotFoundError(err) {
								return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
							}
						}

						if _, ok := nw.Ea["Building"]; ok {
							return fmt.Errorf("Building EA should've been removed, but still present in the WAPI object")
						}

						return nil
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
				Config: `
					resource "infoblox_ipv6_network" "foo"{
						network_view="default"
						cidr="2001:db8:abcd:12::/64"
						reserve_ipv6 = 10
						comment = "2001:db8:abcd:12::/64 network created"
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							"Network Name"= "demo-network"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
					}`,
				Check: validateNetwork(
					"infoblox_ipv6_network.foo",
					&ibclient.Network{
						NetviewName: "default",
						Cidr:        "2001:db8:abcd:12::/64",
						Comment:     "2001:db8:abcd:12::/64 network created",
						Ea: ibclient.EA{
							"Network Name": "demo-network",
							"Tenant ID":    "terraform_test_tenant",
							"Location":     "Test loc.",
							"Site":         "Test site",
						},
					},
				),
			},
			{
				Config: `
					resource "infoblox_ipv6_network" "foo"{
						network_view="default"
						cidr="2001:db8:abcd:12::/64"
						reserve_ipv6 = 11
						comment = "2001:db8:abcd:12::/64 network created"
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							"Network Name"= "demo-network"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
					}`,
				ExpectError: updateNotAllowedErrorRegexp,
			},
			{
				Config: `
					resource "infoblox_ipv6_network" "foo"{
						network_view="default"
						cidr="2001:db8:abcd:12::/64"
						reserve_ipv6 = 11
						comment = "2001:db8:abcd:12::/64 network created"
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							"Network Name"= "demo-network"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
					}`,
				ExpectError: updateNotAllowedErrorRegexp,
			},
		},
	})
}

var testResourceIPv4Network = `resource "infoblox_ipv4_network_container" "ipv4_network12" {
  cidr = "182.11.0.0/24"
  network_view = "default"
  comment = "small network for testing"
  ext_attrs = jsonencode({
    "Site" = "Darjeeling"
  })
}
resource "infoblox_ipv4_network" "ipv4_network13"{
	object = "networkcontainer"
    allocate_prefix_len = 26 
	comment = "network created"
	filter_params = jsonencode({
		"*Site" = "Darjeeling"
	})
	ext_attrs = jsonencode({
    	Location = "Europe"
  	})
	depends_on = [infoblox_ipv4_network_container.ipv4_network12]
}`

func TestAcc_resourceNetwork_AllocateNetworkByEA_IPV4(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testResourceIPv4Network,
				Check: resource.ComposeTestCheckFunc(
					validateNetwork("infoblox_ipv4_network.ipv4_network13",
						&ibclient.Network{
							NetviewName: "default",
							Comment:     "network created",
							Ea: ibclient.EA{
								"Location": "Europe",
							},
						},
					),
				),
			},
			{
				// Negative testcase
				Config: `
					resource "infoblox_ipv4_network" "ipv4_network14"{
					network_view="default"
					object = "networkcontainer"
					comment = "network created"
					allocate_prefix_len = 26
					filter_params = jsonencode({
						"*Site" = "Finland"
					})
					ext_attrs = jsonencode({
						Location = "Europe"
					})
				}`,
				ExpectError: regexp.MustCompile("did not return any result"),
			},
		},
	})
}
