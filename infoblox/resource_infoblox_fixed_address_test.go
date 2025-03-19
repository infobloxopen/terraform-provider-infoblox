package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"reflect"
	"regexp"
	"testing"
)

func testAccCheckFixedAddressDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_fixed_address" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetFixedAddressByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}

	}
	return nil
}
func testAccFixedAddressCompare(t *testing.T, resourceName string, expectedFixedAddress *ibclient.FixedAddress) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resourceName]
		if !found {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if res.Primary.ID == "" {
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
		fixedAddress, err := objMgr.SearchObjectByAltId("FixedAddress", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedFixedAddress == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		var rec *ibclient.FixedAddress
		recJson, _ := json.Marshal(fixedAddress)
		err = json.Unmarshal(recJson, &rec)
		if rec.IPv4Address != expectedFixedAddress.IPv4Address {
			return fmt.Errorf(
				"the value of 'ipv4addr' field is '%s', but expected '%s'",
				rec.IPv4Address, expectedFixedAddress.IPv4Address)
		}
		if rec.Name != expectedFixedAddress.Name {
			return fmt.Errorf("the value of 'name' field is '%s' , but expected is '%s'", rec.Name, expectedFixedAddress.Name)
		}
		if rec.Mac != expectedFixedAddress.Mac {
			return fmt.Errorf("the value of 'mac' field is '%s' , but expected is '%s'", rec.Mac, expectedFixedAddress.Mac)
		}
		if rec.MatchClient != expectedFixedAddress.MatchClient {
			return fmt.Errorf("the value of 'matchcleint' field is '%s', but expected '%s'", rec.MatchClient, expectedFixedAddress.MatchClient)
		}
		if rec.NetviewName != expectedFixedAddress.NetviewName {
			return fmt.Errorf("the value of 'network_view'field is '%s', but expected '%s'", rec.NetviewName, expectedFixedAddress.NetviewName)
		}
		if rec.Comment != expectedFixedAddress.Comment {
			return fmt.Errorf("the value of 'comment' field is '%s', but expected '%s'", rec.Comment, expectedFixedAddress.Comment)
		}
		if rec.AgentCircuitId != expectedFixedAddress.AgentCircuitId {
			return fmt.Errorf("the value of 'agent_circuit_id' field is '%s', but expected '%s'", rec.AgentCircuitId, expectedFixedAddress.AgentCircuitId)
		}
		if rec.AgentRemoteId != expectedFixedAddress.AgentRemoteId {
			return fmt.Errorf("the value of 'agent_remote_id' field is '%s', but expected '%s'", rec.AgentRemoteId, expectedFixedAddress.AgentRemoteId)
		}
		if rec.Options != nil && expectedFixedAddress.Options != nil {
			if len(rec.Options) != len(expectedFixedAddress.Options) {
				return fmt.Errorf("the length of 'options' field is '%d' but expected '%d'", len(rec.Options), len(expectedFixedAddress.Options))
			}
			if rec.ClientIdentifierPrependZero != expectedFixedAddress.ClientIdentifierPrependZero {
				return fmt.Errorf("the value of 'client_identifier_prepend_zero' field is '%t', but expected '%t'", rec.ClientIdentifierPrependZero, expectedFixedAddress.ClientIdentifierPrependZero)
			}
			if rec.DhcpClientIdentifier != expectedFixedAddress.DhcpClientIdentifier {
				return fmt.Errorf("the value of 'dhcp_client_identifier' field is '%s' , but expected '%s'", rec.DhcpClientIdentifier, expectedFixedAddress.DhcpClientIdentifier)
			}
			for i := range rec.Options {
				if !reflect.DeepEqual(rec.Options[i], expectedFixedAddress.Options[i]) {
					return fmt.Errorf("difference found at index %d: got '%v' but expected '%v'", i, rec.Options[i], expectedFixedAddress.Options[i])
				}
			}
		}
		if rec.Disable != nil && expectedFixedAddress.Disable != nil {
			if *rec.Disable != *expectedFixedAddress.Disable {
				return fmt.Errorf(
					"the value of 'disable' field is '%t', but expected '%t'",
					*rec.Disable, *expectedFixedAddress.Disable)
			}
		}
		if rec.UseOptions != nil && expectedFixedAddress.UseOptions != nil {
			if *rec.UseOptions != *expectedFixedAddress.UseOptions {
				return fmt.Errorf(
					"the value of 'use_option' field is '%t', but expected '%t'",
					*rec.UseOptions, *expectedFixedAddress.UseOptions)
			}
		}
		if rec.Cidr != expectedFixedAddress.Cidr {
			return fmt.Errorf("the value of 'network' field is '%s', but expected is '%s'", rec.Cidr, expectedFixedAddress.Cidr)
		}
		return validateEAs(rec.Ea, expectedFixedAddress.Ea)
	}
}

var regexpCreateErrorIPV4FixedAddress = regexp.MustCompile("either 'ipv4addr' or 'network' fields needs to provided to allocate a fixed address")

func TestAccResourceFixedAddress(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFixedAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "infoblox_ipv4_network" "net1" {
						cidr = "15.0.0.0/24"
					}

					resource "infoblox_ipv4_fixed_address" "fix1" {
						ipv4addr        = "15.0.0.2"
						match_client    = "CIRCUIT_ID"
						agent_circuit_id = "32"

						options {
							name         = "dhcp-lease-time"
							value        = "43200"
							vendor_class = "DHCP"
							num          = 51
							use_option   = true
						}

						depends_on = [infoblox_ipv4_network.net1]
					}`),
				Check: resource.ComposeTestCheckFunc(
					testAccFixedAddressCompare(t, "infoblox_ipv4_fixed_address.fix1", &ibclient.FixedAddress{
						IPv4Address:    "15.0.0.2",
						MatchClient:    "CIRCUIT_ID",
						AgentCircuitId: "32",
						Cidr:           "15.0.0.0/24",
						Options: []*ibclient.Dhcpoption{{
							Name:        "dhcp-lease-time",
							Value:       "43200",
							VendorClass: "DHCP",
							Num:         51,
							UseOption:   true,
						}},
						Disable:     utils.BoolPtr(false),
						NetviewName: "default",
					}),
				),
			},
			{
				//next available ip allocation
				Config: fmt.Sprintf(`
					resource "infoblox_ipv4_network" "net2" {
						cidr = "16.0.0.0/24"
					}
					resource "infoblox_ipv4_fixed_address" "fix2" {
						ipv4addr        = ""
						match_client    = "REMOTE_ID"
						agent_remote_id = "35"
						network = "16.0.0.0/24"
						options {
							name         = "dhcp-lease-time"
							value        = "43200"
							vendor_class = "DHCP"
							num          = 51
							use_option   = true
						}
						options {
							name = "routers"
   						num = "3"
   						use_option = true
   						value = "16.0.0.2"
   						vendor_class = "DHCP"
 						}
						depends_on = [infoblox_ipv4_network.net2]
					}`),
				Check: resource.ComposeTestCheckFunc(testAccFixedAddressCompare(t, "infoblox_ipv4_fixed_address.fix2", &ibclient.FixedAddress{
					IPv4Address:   "16.0.0.1",
					MatchClient:   "REMOTE_ID",
					AgentRemoteId: "35",
					Cidr:          "16.0.0.0/24",
					NetviewName:   "default",
					Options: []*ibclient.Dhcpoption{{
						Name:        "dhcp-lease-time",
						Value:       "43200",
						VendorClass: "DHCP",
						Num:         51,
						UseOption:   true,
					},
						{
							Name:        "routers",
							Num:         3,
							Value:       "16.0.0.2",
							VendorClass: "DHCP",
							UseOption:   true,
						}},
					Disable:    utils.BoolPtr(false),
					UseOptions: utils.BoolPtr(false),
				})),
			},
			{
				Config: fmt.Sprintf(`
							resource "infoblox_ipv4_network" "net3" {
									cidr = "17.0.0.0/24"
							}
							resource "infoblox_ipv4_fixed_address" "fix3"{
								ipv4addr        = "17.0.0.9"
								mac = "00:0C:24:2E:8F:2A"
							options {
								name         = "dhcp-lease-time"
								value        = "43200"
								vendor_class = "DHCP"
								num          = 51
								use_option   = true
							}
								depends_on = [infoblox_ipv4_network.net3]
							}`),
				Check: resource.ComposeTestCheckFunc(testAccFixedAddressCompare(t, "infoblox_ipv4_fixed_address.fix3", &ibclient.FixedAddress{
					IPv4Address: "17.0.0.9",
					Mac:         "00:0c:24:2e:8f:2a",
					Options: []*ibclient.Dhcpoption{{
						Name:        "dhcp-lease-time",
						Value:       "43200",
						VendorClass: "DHCP",
						Num:         51,
						UseOption:   true,
					},
					},
					Cidr:        "17.0.0.0/24",
					MatchClient: "MAC_ADDRESS",
					NetviewName: "default",
				})),
			},
			{
				//negative test case
				Config: fmt.Sprintf(`
					resource "infoblox_ipv4_fixed_address" "fix4"{
								mac = "00:0C:24:2E:8F:2A"
							options {
								name         = "dhcp-lease-time"
								value        = "43200"
								vendor_class = "DHCP"
								num          = 51
								use_option   = true
							}
					}`),
				ExpectError: regexpCreateErrorIPV4FixedAddress,
			},
			{
				Config: fmt.Sprintf(`
								resource "infoblox_ipv4_network" "net2" {
											cidr = "18.0.0.0/24"
								}
								resource "infoblox_ipv4_fixed_address" "fix5"{
  								client_identifier_prepend_zero=true
  								comment= "fixed address"
  								dhcp_client_identifier="23"
  								disable= true
  								ext_attrs = jsonencode({
    								"Site": "Blr"
  								})
 	 							match_client = "CLIENT_ID"
  								name = "fixed_address_1"
								network = "18.0.0.0/24"
  								network_view = "default"
  								options {
    								name         = "dhcp-lease-time"
    								value        = "43200"
    								vendor_class = "DHCP"
    								num          = 51
    								use_option   = true
  								}
  								options {
    								name = "routers"
    								num = "3"
    								use_option = true
    								value = "18.0.0.2"
    								vendor_class = "DHCP"
  								}
  								use_option = true
								depends_on = [infoblox_ipv4_network.net2]
								}`),
				Check: resource.ComposeTestCheckFunc(
					testAccFixedAddressCompare(t, "infoblox_ipv4_fixed_address.fix5", &ibclient.FixedAddress{
						IPv4Address:                 "18.0.0.1",
						MatchClient:                 "CLIENT_ID",
						DhcpClientIdentifier:        "23",
						ClientIdentifierPrependZero: true,
						Name:                        "fixed_address_1",
						Comment:                     "fixed address",
						Cidr:                        "18.0.0.0/24",
						Options: []*ibclient.Dhcpoption{{
							Name:        "dhcp-lease-time",
							Value:       "43200",
							VendorClass: "DHCP",
							Num:         51,
							UseOption:   true,
						},
							{
								Name:        "routers",
								Num:         3,
								Value:       "18.0.0.2",
								VendorClass: "DHCP",
								UseOption:   true,
							}},
						Disable:     utils.BoolPtr(true),
						NetviewName: "default",
						UseOptions:  utils.BoolPtr(true),
						Ea: ibclient.EA{
							"Site": "Blr",
						},
					}),
				),
			},
		},
	})
}
