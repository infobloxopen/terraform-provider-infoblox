package infoblox

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"regexp"
	"testing"
)

var testResourceIpv4SharedNetwork = `resource "infoblox_ipv4_shared_network" "record1" {
    name = "shared-network11"
    comment = "test ipv4 shared network record"
    networks = ["31.12.3.0/24","31.13.3.0/24"]
	network_view = "default"
	disable = false
	ext_attrs = jsonencode({
    	"Site" = "Tokyo"
  	})
	use_options = true
	options {
    	name = "domain-name-servers"
    	value = "11.22.33.44"
		vendor_class = "DHCP"
		num = 6
		use_option = true
  	}
	options {
    	name = "dhcp-lease-time"
    	value = "43200"
		vendor_class = "DHCP"
		num = 51
		use_option = true
  	}
	depends_on = [infoblox_ipv4_network.ipv4_network1, infoblox_ipv4_network.ipv4_network2]
}

resource "infoblox_ipv4_network" "ipv4_network1" {
  cidr = "31.12.3.0/24"
}

resource "infoblox_ipv4_network" "ipv4_network2" {
  cidr = "31.13.3.0/24"
}`

var testResourceIpv4SharedNetwork2 = `resource "infoblox_ipv4_shared_network" "record2" {
    name = "shared-network12"
    networks = ["29.11.3.0/24"]
	options {
    	name = "dhcp-lease-time"
    	value = "89200"
		vendor_class = "DHCP"
		num = 51
		use_option = true
  	}
	depends_on = [infoblox_ipv4_network.ipv4_network3]
}

resource "infoblox_ipv4_network" "ipv4_network3" {
  cidr = "29.11.3.0/24"
}`

var testResourceIpv4SharedNetwork3 = `resource "infoblox_ipv4_shared_network" "record3" {
    name = "shared-network3"
	networks = ["30.1.12.0/24"]
	
}`

func testIpv4SharedNetworkDestroy(s *terraform.State) error {
	meta := testAccProvider.Meta()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "resource_ipv4_shared_network" {
			continue
		}
		connector := meta.(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(connector, "terraform_test", "test")
		rec, _ := objMgr.GetIpv4SharedNetworkByRef(rs.Primary.ID)
		if rec != nil {
			return fmt.Errorf("record not found")
		}
	}
	return nil
}

func testIpv4SharedNetworkCompare(t *testing.T, resourceName string, ipv4SharedNetwork *ibclient.SharedNetwork) resource.TestCheckFunc {
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
		zf, err := objMgr.SearchObjectByAltId("SharedNetwork", ref, internalId, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if ipv4SharedNetwork == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		// Assertion of object type and error handling
		var rec *ibclient.SharedNetwork
		recJson, _ := json.Marshal(zf)
		err = json.Unmarshal(recJson, &rec)

		if rec.Name == nil && ipv4SharedNetwork.Name != nil {
			if *rec.Name != *ipv4SharedNetwork.Name {
				return fmt.Errorf(
					"the value of 'name' field is '%s', but expected '%s'",
					*rec.Name, *ipv4SharedNetwork.Name)
			}
		}
		if rec.Comment != nil && ipv4SharedNetwork.Comment != nil {
			if *rec.Comment != *ipv4SharedNetwork.Comment {
				return fmt.Errorf(
					"the value of 'comment' field is '%s', but expected '%s'",
					*rec.Comment, *ipv4SharedNetwork.Comment)
			}
		}
		if rec.Disable != nil && ipv4SharedNetwork.Disable != nil {
			if *rec.Disable != *ipv4SharedNetwork.Disable {
				return fmt.Errorf(
					"the value of 'disable' field is '%t', but expected '%t'",
					*rec.Disable, *ipv4SharedNetwork.Disable)
			}
		}
		if rec.NetworkView != ipv4SharedNetwork.NetworkView {
			return fmt.Errorf("the value of 'network_view' field is '%s', but expected '%s'",
				rec.NetworkView, ipv4SharedNetwork.NetworkView)
		}
		if rec.UseOptions != nil && ipv4SharedNetwork.UseOptions != nil {
			if *rec.UseOptions != *ipv4SharedNetwork.UseOptions {
				return fmt.Errorf(
					"the value of 'use_options' field is '%t', but expected '%t'",
					*rec.UseOptions, *ipv4SharedNetwork.UseOptions)
			}
		}
		if rec.Options != nil && ipv4SharedNetwork.Options != nil {
			if !compareDHCPOptions(rec.Options, ipv4SharedNetwork.Options) {
				return fmt.Errorf("the value of 'options' field is '%v', but expected '%v'", rec.Options, ipv4SharedNetwork.Options)
			}
		}
		if rec.Networks != nil && ipv4SharedNetwork.Networks != nil {
			if !comapreIpv4SharedNetworks(rec.Networks, ipv4SharedNetwork.Networks) {
				return fmt.Errorf("the value of 'networks' field is '%v', but expected '%v'", rec.Networks, ipv4SharedNetwork.Networks)
			}
		}
		return validateEAs(rec.Ea, ipv4SharedNetwork.Ea)
	}
}

func comapreIpv4SharedNetworks(networks1 []*ibclient.Ipv4Network, networks2 []*ibclient.Ipv4Network) bool {
	if len(networks1) != len(networks2) {
		return false
	}
	for i, _ := range networks1 {
		if networks1[i].Network != nil && networks2[i].Network != nil {
			if *networks1[i].Network != *networks2[i].Network {
				return false
			}
		}
	}
	return true
}

func compareDHCPOptions(options1, options2 []*ibclient.Dhcpoption) bool {
	if len(options1) != len(options2) {
		return false
	}
	for i := range options1 {
		if options1[i].Name != options2[i].Name || options1[i].Value != options2[i].Value {
			return false
		}
	}
	return true
}

func TestAccResourceipv4SharedNetwork(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testIpv4SharedNetworkDestroy,
		Steps: []resource.TestStep{
			// maximum params
			{
				Config: testResourceIpv4SharedNetwork,
				Check: testIpv4SharedNetworkCompare(t, "infoblox_ipv4_shared_network.record1", &ibclient.SharedNetwork{
					Name:        utils.StringPtr("shared-network11"),
					NetworkView: defaultNetView,
					Networks: []*ibclient.Ipv4Network{
						{Network: utils.StringPtr("31.12.3.0/24")},
						{Network: utils.StringPtr("31.13.3.0/24")},
					},
					Disable:    utils.BoolPtr(false),
					Ea:         map[string]interface{}{"Site": "Tokyo"},
					Comment:    utils.StringPtr("test ipv4 shared network record"),
					UseOptions: utils.BoolPtr(true),
					Options: []*ibclient.Dhcpoption{
						{
							Name:        "domain-name-servers",
							Num:         6,
							UseOption:   true,
							Value:       "11.22.33.44",
							VendorClass: "DHCP",
						},
						{
							Name:        "dhcp-lease-time",
							Num:         51,
							UseOption:   true,
							Value:       "43200",
							VendorClass: "DHCP",
						},
					},
				}),
			},
			// minimum params
			{
				Config: testResourceIpv4SharedNetwork2,
				Check: testIpv4SharedNetworkCompare(t, "infoblox_ipv4_shared_network.record2", &ibclient.SharedNetwork{
					Name: utils.StringPtr("shared-network12"),
					Networks: []*ibclient.Ipv4Network{
						{Network: utils.StringPtr("29.11.3.0/24")},
					},
					NetworkView: defaultNetView,
					Options: []*ibclient.Dhcpoption{
						{
							Name:        "dhcp-lease-time",
							Num:         51,
							UseOption:   true,
							Value:       "89200",
							VendorClass: "DHCP",
						},
					},
				}),
			},
			// negative test case
			{
				Config:      testResourceIpv4SharedNetwork3,
				ExpectError: regexp.MustCompile("No network objects were matched by {'network': '30.1.12.0/24', 'network_view': 'default'}"),
			},
		},
	})
}
