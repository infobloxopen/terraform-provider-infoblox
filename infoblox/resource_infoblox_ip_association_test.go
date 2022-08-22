package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func validateIPAssociationIpV4Addr(a, e *ibclient.HostRecordIpv4Addr) error {
	if e == nil {
		if a != nil {
			return fmt.Errorf("IPv4 address at a host record is expected to be empty")
		}
		return nil
	}

	if a == nil {
		return fmt.Errorf("IPv4 address at a host record is expected to be non-empty")
	}

	if a.Ipv4Addr != e.Ipv4Addr || a.EnableDhcp != e.EnableDhcp || a.Mac != e.Mac {
		return fmt.Errorf(
			"IPv4 address at a host record is not the same as expected;"+
				" actual: '%+v'; expected: '%+v'",
			a, e)
	}

	return nil
}

func validateIPAssociationIpV6Addr(a, e *ibclient.HostRecordIpv6Addr) error {
	if e == nil {
		if a != nil {
			return fmt.Errorf("IPv6 address at a host record is expected to be empty")
		}
		return nil
	}

	if a == nil {
		return fmt.Errorf("IPv6 address at a host record is expected to be non-empty")
	}

	if a.Ipv6Addr != e.Ipv6Addr || a.EnableDhcp != e.EnableDhcp || a.Duid != e.Duid {
		return fmt.Errorf(
			"IPv6 address at a host record is not the same as expected;"+
				" actual: '%+v'; expected: '%+v'",
			a, e)
	}

	return nil
}

func validateIPAssociation(
	resourceName string,
	expectedValue *ibclient.HostRecord) resource.TestCheckFunc {
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

		internalId := res.Primary.Attributes["internal_id"]
		if internalId == "" {
			return fmt.Errorf("internal ID is not set")
		}

		hostRec, err := objMgr.SearchHostRecordByAltId(internalId, "", eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedValue == nil {
					return nil
				}
				return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
			}
		}

		var (
			internalIdEaVal string
			recIpV4Addr     *ibclient.HostRecordIpv4Addr
			recIpV6Addr     *ibclient.HostRecordIpv6Addr
		)
		if hostRec.Ea != nil {
			if tempVal, found := hostRec.Ea[eaNameForInternalId]; found {
				if tempStrVal, ok := tempVal.(string); ok {
					internalIdEaVal = tempStrVal
				}
			}

		}
		if internalIdEaVal == "" {
			return fmt.Errorf("internal ID EA not set")
		}

		if internalIdEaVal != internalId {
			return fmt.Errorf(
				"internal ID in EA (%s differs from expected one (%s)",
				internalIdEaVal, internalId)
		}

		if len(hostRec.Ipv4Addrs) > 0 {
			recIpV4Addr = &hostRec.Ipv4Addrs[0]
		}
		if len(hostRec.Ipv6Addrs) > 0 {
			recIpV6Addr = &hostRec.Ipv6Addrs[0]
		}

		var (
			expIpv4Addr *ibclient.HostRecordIpv4Addr
			expIpv6Addr *ibclient.HostRecordIpv6Addr
		)

		if len(expectedValue.Ipv4Addrs) > 0 {
			expIpv4Addr = &expectedValue.Ipv4Addrs[0]
		}
		err = validateIPAssociationIpV4Addr(recIpV4Addr, expIpv4Addr)
		if err != nil {
			return err
		}

		if len(expectedValue.Ipv6Addrs) > 0 {
			expIpv6Addr = &expectedValue.Ipv6Addrs[0]
		}
		err = validateIPAssociationIpV6Addr(recIpV6Addr, expIpv6Addr)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIPAssociationDestroy(s *terraform.State) error {
	connector := testAccProvider.Meta().(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(
		connector,
		"terraform_test",
		"terraform_test_tenant")
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_ip_association" {
			continue
		}
		internalId := rs.Primary.Attributes["internal_id"]
		if internalId == "" {
			return fmt.Errorf("internal ID is not set")
		}

		hostRec, err := objMgr.SearchHostRecordByAltId(internalId, "", eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				continue
			}
			return err
		}
		if hostRec != nil {
			return fmt.Errorf("object with ID '%s' remains", internalId)
		}
	}
	return nil
}

func TestAcc_resourceipAssociation(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
					resource infoblox_ipv4_network "net1" {
						cidr = "10.0.0.0/24"
					}
					resource "infoblox_ip_allocation" "foo" {
						network_view="default"
						fqdn="testhostname.test.com"
						ipv4_addr="10.0.0.12"
						enable_dns = "true"
						comment = "10.0.0.12 IP is allocated"
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							"VM Name" =  "tf-ec2-instance"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
						depends_on = [infoblox_ipv4_network.net1]
					}

					resource "infoblox_ip_association" "foo" {
					  internal_id = infoblox_ip_allocation.foo.internal_id
					  enable_dhcp = true
					  mac_addr = "11:22:33:44:55:66"
					}`,
				Check: validateIPAssociation(
					"infoblox_ip_association.foo",
					ibclient.NewHostRecord(
						"default",
						"testhostname.test.com",
						"", "",
						[]ibclient.HostRecordIpv4Addr{
							*ibclient.NewHostRecordIpv4Addr(
								"10.0.0.12",
								"11:22:33:44:55:66",
								true, "")},
						[]ibclient.HostRecordIpv6Addr{},
						nil,
						true, "default",
						"test.com", "",
						false, 0,
						"10.0.0.12 IP is allocated",
						[]string{})),
			},
			{
				Config: `
					resource infoblox_ipv4_network "net1" {
						cidr = "10.0.0.0/24"
					}
					resource infoblox_ipv6_network "net2" {
						cidr = "2001::/56"
					}
					resource "infoblox_ip_allocation" "foo" {
						network_view="default"
						fqdn="testhostname.test.com"
						ipv4_addr="10.0.0.12"
						ipv6_addr="2001::10"
						enable_dns = "false"
						comment = "10.0.0.12 IP is allocated"
						ttl=0
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							"VM Name" =  "tf-ec2-instance"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
						depends_on = [infoblox_ipv4_network.net1, infoblox_ipv6_network.net2]
					}
		
					resource "infoblox_ip_association" "foo" {
					  enable_dhcp = true
					  mac_addr = "11:22:33:44:55:66"
					  duid = "22:44:66"
					  internal_id = infoblox_ip_allocation.foo.internal_id
					}`,
				Check: validateIPAssociation(
					"infoblox_ip_association.foo",
					ibclient.NewHostRecord(
						"default",
						"testhostname.test.com",
						"", "",
						[]ibclient.HostRecordIpv4Addr{
							*ibclient.NewHostRecordIpv4Addr(
								"10.0.0.12",
								"11:22:33:44:55:66",
								true, "")},
						[]ibclient.HostRecordIpv6Addr{
							*ibclient.NewHostRecordIpv6Addr(
								"2001::10",
								"22:44:66",
								true, "")},
						nil,
						false, "default",
						"", "",
						true, 0,
						"10.0.0.12 IP is allocated",
						[]string{})),
			},
			{
				Config: `
					resource infoblox_ipv4_network "net1" {
						cidr = "10.0.0.0/24"
					}
					resource infoblox_ipv6_network "net2" {
						cidr = "2001::/56"
					}
					resource "infoblox_ip_allocation" "foo" {
						network_view="default"
						fqdn="testhostname.test.com"
						ipv4_addr="10.0.0.12"
						ipv6_addr="2001::10"
						enable_dns = "false"
						comment = "10.0.0.12 IP is allocated"
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							"VM Name" =  "tf-ec2-instance"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
						depends_on = [infoblox_ipv4_network.net1, infoblox_ipv6_network.net2]
					}
		
					resource "infoblox_ip_association" "foo" {
					  enable_dhcp = true
					  mac_addr = "11:22:33:44:55:66"
					  duid = "22:44:66"
					  internal_id = infoblox_ip_allocation.foo.internal_id
					}`,
				Check: validateIPAssociation(
					"infoblox_ip_association.foo",
					ibclient.NewHostRecord(
						"default",
						"testhostname.test.com",
						"", "",
						[]ibclient.HostRecordIpv4Addr{
							*ibclient.NewHostRecordIpv4Addr(
								"10.0.0.12",
								"11:22:33:44:55:66",
								true, "")},
						[]ibclient.HostRecordIpv6Addr{
							*ibclient.NewHostRecordIpv6Addr(
								"2001::10",
								"22:44:66",
								true, "")},
						nil,
						false, "default",
						"", "",
						false, 0,
						"10.0.0.12 IP is allocated",
						[]string{})),
			},
			{
				Config: `
					resource infoblox_ipv4_network "net1" {
						cidr = "10.0.0.0/24"
					}
					resource infoblox_ipv6_network "net2" {
						cidr = "2001::/56"
					}
					resource "infoblox_ip_allocation" "foo" {
						network_view="default"
						fqdn="testhostname.test.com"
						ipv4_addr="10.0.0.12"
						ipv6_addr="2001::10"
						enable_dns = "false"
						comment = "10.0.0.12 IP is allocated"
						ttl=10
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							"VM Name" =  "tf-ec2-instance"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
						depends_on = [infoblox_ipv4_network.net1, infoblox_ipv6_network.net2]
					}
		
					resource "infoblox_ip_association" "foo" {
					  enable_dhcp = true
					  mac_addr = "11:22:33:44:55:66"
					  duid = "22:44:66"
					  internal_id = infoblox_ip_allocation.foo.internal_id
					}`,
				Check: validateIPAssociation(
					"infoblox_ip_association.foo",
					ibclient.NewHostRecord(
						"default",
						"testhostname.test.com",
						"", "",
						[]ibclient.HostRecordIpv4Addr{
							*ibclient.NewHostRecordIpv4Addr(
								"10.0.0.12",
								"11:22:33:44:55:66",
								true, "")},
						[]ibclient.HostRecordIpv6Addr{
							*ibclient.NewHostRecordIpv6Addr(
								"2001::10",
								"22:44:66",
								true, "")},
						nil,
						false, "default",
						"", "",
						true, 10,
						"10.0.0.12 IP is allocated",
						[]string{})),
			},
			{
				Config: `
					resource infoblox_ipv4_network "net1" {
						cidr = "10.0.0.0/24"
					}
					resource infoblox_ipv6_network "net2" {
						cidr = "2001::/56"
					}
					resource "infoblox_ip_allocation" "foo" {
						network_view="default"
						fqdn="testhostname.test.com"
						ipv4_addr="10.0.0.12"
						ipv6_addr="2001::10"
						enable_dns = "false"
						comment = "10.0.0.12 IP is allocated"
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							"VM Name" =  "tf-ec2-instance"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
						depends_on = [infoblox_ipv4_network.net1, infoblox_ipv6_network.net2]
					}
		
					resource "infoblox_ip_association" "foo" {
					  enable_dhcp = false
					  mac_addr = "11:22:33:44:55:66"
					  duid = "22:44:66"
					  internal_id = infoblox_ip_allocation.foo.internal_id
					}`,
				Check: validateIPAssociation(
					"infoblox_ip_association.foo",
					ibclient.NewHostRecord(
						"default",
						"testhostname.test.com",
						"", "",
						[]ibclient.HostRecordIpv4Addr{
							*ibclient.NewHostRecordIpv4Addr(
								"10.0.0.12",
								"11:22:33:44:55:66",
								false, "")},
						[]ibclient.HostRecordIpv6Addr{
							*ibclient.NewHostRecordIpv6Addr(
								"2001::10",
								"22:44:66",
								false, "")},
						nil,
						false, "default",
						"", "",
						false, 0,
						"10.0.0.12 IP is allocated",
						[]string{})),
			},
			{
				Config: `
					resource infoblox_ipv4_network "net1" {
						cidr = "10.0.0.0/24"
					}
					resource infoblox_ipv6_network "net2" {
						cidr = "2001::/56"
					}
					resource "infoblox_ip_allocation" "foo" {
						network_view="default"
						fqdn="testhostname.test.com"
						ipv4_addr="10.0.0.12"
						ipv6_addr="2001::10"
						enable_dns = "false"
						comment = "10.0.0.12 IP is allocated"
						ttl=10
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							"VM Name" =  "tf-ec2-instance"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
						depends_on = [infoblox_ipv4_network.net1, infoblox_ipv6_network.net2]
					}
		
					resource "infoblox_ip_association" "foo" {
					  enable_dhcp = false
					  mac_addr = "11:22:33:44:55:66"
					  duid = "22:44:66"
					  internal_id = infoblox_ip_allocation.foo.internal_id
					}`,
				Check: validateIPAssociation(
					"infoblox_ip_association.foo",
					ibclient.NewHostRecord(
						"default",
						"testhostname.test.com",
						"", "",
						[]ibclient.HostRecordIpv4Addr{
							*ibclient.NewHostRecordIpv4Addr(
								"10.0.0.12",
								"11:22:33:44:55:66",
								false, "")},
						[]ibclient.HostRecordIpv6Addr{
							*ibclient.NewHostRecordIpv6Addr(
								"2001::10",
								"22:44:66",
								false, "")},
						nil,
						false, "default",
						"", "",
						true, 10,
						"10.0.0.12 IP is allocated",
						[]string{})),
			},
			{
				Config: `
					resource infoblox_ipv4_network "net1" {
						cidr = "10.0.0.0/24"
					}
					resource infoblox_ipv6_network "net2" {
						cidr = "2001::/56"
					}
					resource "infoblox_ip_allocation" "foo" {
						network_view="default"
						fqdn="testhostname.test.com"
						ipv4_addr="10.0.0.12"
						ipv6_addr="2001::10"
						enable_dns = "false"
						comment = "10.0.0.12 IP is allocated"
						ttl=10
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							"VM Name" =  "tf-ec2-instance"
							"Location" = "Test loc."
							"Site" = "Test site"
						})
						depends_on = [infoblox_ipv4_network.net1, infoblox_ipv6_network.net2]
					}
		
					resource "infoblox_ip_association" "foo" {
					  enable_dhcp = true
					  mac_addr = "11:22:33:44:55:66"
					  duid = "22:44:66"
					  internal_id = infoblox_ip_allocation.foo.internal_id
					}`,
				Check: validateIPAssociation(
					"infoblox_ip_association.foo",
					ibclient.NewHostRecord(
						"default",
						"testhostname.test.com",
						"", "",
						[]ibclient.HostRecordIpv4Addr{
							*ibclient.NewHostRecordIpv4Addr(
								"10.0.0.12",
								"11:22:33:44:55:66",
								true, "")},
						[]ibclient.HostRecordIpv6Addr{
							*ibclient.NewHostRecordIpv6Addr(
								"2001::10",
								"22:44:66",
								true, "")},
						nil,
						false, "default",
						"", "",
						true, 10,
						"10.0.0.12 IP is allocated",
						[]string{})),
			},
		},
	})
}
