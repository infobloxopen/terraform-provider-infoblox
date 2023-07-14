package infoblox

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
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
						ext_attrs = {
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						  }
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
						ext_attrs = {
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
							"Building" = "Test Building"
						  }
						}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_ipv4_network.foo", "comment", "Updated comment"),
					resource.TestCheckResourceAttr("infoblox_ipv4_network.foo", "ext_attrs.%", "5"),
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
						ext_attrs = {
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						  }
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
						ext_attrs = {
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						  }
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
						ext_attrs = {
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						  }
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
						ext_attrs = {
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						  }
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
						ext_attrs = {
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						  }
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
				// When extensible attributes are added by another tool,
				// terraform shouldn't remove those EAs
				PreConfig: func() {
					getEnvDefault := func(key, def string) string {
						if value, ok := os.LookupEnv(key); ok {
							return value
						}
						return def
					}

					hostConfig := ibclient.HostConfig{
						Host:    os.Getenv("INFOBLOX_SERVER"),
						Port:    getEnvDefault("PORT", "443"),
						Version: getEnvDefault("WAPI_VERSION", "2.7"),
					}

					authConfig := ibclient.AuthConfig{
						Username: os.Getenv("INFOBLOX_USERNAME"),
						Password: os.Getenv("INFOBLOX_PASSWORD"),
					}

					sslVerify, err := strconv.ParseBool(getEnvDefault("SSLMODE", "false"))
					if err != nil {
						panic(err)
					}

					secondsStr := getEnvDefault("CONNECT_TIMEOUT", "60")
					seconds, err := strconv.Atoi(secondsStr)
					if err != nil {
						panic(err)
					}

					poolConnectionsStr := getEnvDefault("POOL_CONNECTIONS", "10")
					poolConnections, err := strconv.Atoi(poolConnectionsStr)
					if err != nil {
						panic(err)
					}

					transportConfig := ibclient.TransportConfig{
						SslVerify:           sslVerify,
						HttpRequestTimeout:  time.Duration(seconds),
						HttpPoolConnections: poolConnections,
					}

					requestBuilder := &ibclient.WapiRequestBuilder{}
					requestor := &ibclient.WapiHttpRequestor{}

					conn, err := ibclient.NewConnector(hostConfig, authConfig, transportConfig, requestBuilder, requestor)
					if err != nil {
						panic(err)
					}

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
					err = conn.GetObject(n, "", qp, &res)
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
						ext_attrs = {
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
						  }
						}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("infoblox_ipv4_network.foo", "ext_attrs.%", "5"),
					resource.TestCheckResourceAttr("infoblox_ipv4_network.foo", "ext_attrs.Building", "Test Building"),
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
						ext_attrs = {
							"Network Name"= "demo-network"
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc."
							"Site" = "Test site"
							"Building" = "Test building 2"
						  }
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
						ext_attrs = {
							"Tenant ID" = "terraform_test_tenant"
							"Network Name"= "demo-network"
							"Location" = "Test loc."
							"Site" = "Test site"
						}
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
						ext_attrs = {
							"Tenant ID" = "terraform_test_tenant"
							"Network Name"= "demo-network"
							"Location" = "Test loc."
							"Site" = "Test site"
						}
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
						ext_attrs = {
							"Tenant ID" = "terraform_test_tenant"
							"Network Name"= "demo-network"
							"Location" = "Test loc."
							"Site" = "Test site"
						}
					}`,
				ExpectError: updateNotAllowedErrorRegexp,
			},
		},
	})
}
