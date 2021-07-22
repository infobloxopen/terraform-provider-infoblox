package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

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
		ipAsso, err := objMgr.GetHostRecordByRef(id)
		if err != nil {
			if isNotFoundError(err) {
				if expectedValue == nil {
					return nil
				}
				return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
			}
		}
		expNv := expectedValue.NetworkView
		if ipAsso.NetworkView != expNv {
			return fmt.Errorf(
				"the value of 'network_view' field is '%s', but expected '%s'",
				ipAsso.NetworkView, expNv)
		}

		expFqdn := expectedValue.Name
		if ipAsso.Name != expFqdn {
			return fmt.Errorf(
				"the value of 'fqdn' field is '%s', but expected '%s'",
				ipAsso.Name, expFqdn)
		}

		expComment := expectedValue.Comment
		if ipAsso.Comment != expComment {
			return fmt.Errorf(
				"the value of 'comment' field is '%s', but expected '%s'",
				ipAsso.Comment, expComment)
		}

		/*
			expIPv4Address := expectedValue.IPv4Addr
			expIPv6Address := expectedValue.IPv6Addr
			if ipAsso.IPv4Addr != expIPv4Address || ipAsso.IPv6Addr != expIPv6Address {
				return fmt.Errorf(
					"the value of 'IPv4Address' field is '%s', but expected '%s' or the value of 'IPv6Address' field is '%s', but expected '%s' ",
					ipAsso.IPv4Addr, expIPv4Address, ipAsso.IPv6Addr, expIPv6Address)
			}

			expMACAddress := expectedValue.Ipv4Addrs[0].Mac
			expDUID := expectedValue.Ipv6Addrs[0].Duid
			if ipAsso.Ipv4Addrs[0].Mac != expMACAddress || ipAsso.Ipv6Addrs[0].Duid != expDUID {
				return fmt.Errorf(
					"the value of 'IPv4Address' field is '%s', but expected '%s' or the value of 'IPv6Address' field is '%s', but expected '%s' ",
					ipAsso.IPv4Addr, expIPv4Address, ipAsso.IPv6Addr, expIPv6Address)
			}
		*/

		// the rest is about extensible attributes
		expectedEAs := expectedValue.Ea
		if expectedEAs == nil && ipAsso.Ea != nil {
			return fmt.Errorf(
				"the object with ID '%s' has 'ext_attrs' field, but it is not expected to exist", id)
		}
		if expectedEAs != nil && ipAsso.Ea == nil {
			return fmt.Errorf(
				"the object with ID '%s' has no 'ext_attrs' field, but it is expected to exist", id)
		}
		if expectedEAs == nil {
			return nil
		}

		return validateEAs(ipAsso.Ea, expectedEAs)
	}
}

func testAccCheckIPAssociationDestroy(s *terraform.State) error {
	connector := testAccProvider.Meta().(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(
		connector,
		"terraform_test",
		"terraform_test_tenant")
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_ipv4_association" && rs.Type != "infoblox_ipv6_association" {
			continue
		}
		res, err := objMgr.GetHostRecordByRef(rs.Primary.ID)
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

func TestAcc_resourceipAssociation_ipv4(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "infoblox_ipv4_allocation" "foo"{
					network_view="default"
					fqdn="testhostname.test.com"
					ip_addr="10.0.0.12"
					enable_dns = "true"
					comment = "10.0.0.12 IP is allocated"
					ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						"VM Name" =  "tf-ec2-instance"
						"Location" = "Test loc."
						"Site" = "Test site"
					})
				}
				
				resource "infoblox_ipv4_association" "foo"{
					network_view="default"
					dns_view = "default"
					fqdn=infoblox_ipv4_allocation.foo.fqdn
					ip_addr=infoblox_ipv4_allocation.foo.ip_addr
				    mac_addr = "11:22:33:44:55:66"
					enable_dns = "true"
					comment = "10.0.0.12 IP is associated"
					ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						"VM Name" =  "tf-ec2-instance"
						"Location" = "Test loc."
						"Site" = "Test site"
					})
				}	`),
				Check: validateIPAssociation(
					"infoblox_ipv4_association.foo",
					&ibclient.HostRecord{
						NetworkView: "default",
						View:        "default",
						Name:        "testhostname.test.com",
						Ipv4Addr:    "10.0.0.12",
						Comment:     "10.0.0.12 IP is associated",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"VM Name":   "tf-ec2-instance",
							"Location":  "Test loc.",
							"Site":      "Test site",
						},
					},
				),
			},
		},
	})
}

func TestAcc_resourceIPAssociation_ipv6(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "infoblox_ipv6_allocation" "ipv6_allocation" {
					network_view= "default"
					fqdn="testhostnameipv6.test.com"
					ip_addr = "2001:db8:abcd:12::10"
					duid = "00:00:00:00:00:00:00:10"
					comment = "tf IPv6 allocation"
					ext_attrs = jsonencode({
					  "Tenant ID" = "tf-plugin"
					  "Network Name" = "ipv6-tf-network"
					  "VM Name" =  "tf-ec2-instance-ipv6"
					  "Location" = "Test loc."
					  "Site" = "Test site"
					})
				  }

				resource "infoblox_ipv6_association" "foo2"{
					network_view="default"
					fqdn=infoblox_ipv6_allocation.ipv6_allocation.fqdn
					ip_addr=infoblox_ipv6_allocation.ipv6_allocation.ip_addr
					duid="11:22:33:44:55:66"
					comment = "2001:db8:abcd:12::10 IP is associated"
					ext_attrs = jsonencode({
						"VM Name" =  "tf-ec2-instance-ipv6"
						"Tenant ID" = "terraform_test_tenant"
						"Location" = "Test loc."
						"Site" = "Test site"
					  })
					}`),
				Check: validateIPAssociation(
					"infoblox_ipv6_association.foo2",
					&ibclient.HostRecord{
						NetworkView: "default",
						View:        "default",
						Name:        "testhostnameipv6.test.com",
						Ipv6Addr:    "2001:db8:abcd:12::10",
						Comment:     "2001:db8:abcd:12::10 IP is associated",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"VM Name":   "tf-ec2-instance-ipv6",
							"Location":  "Test loc.",
							"Site":      "Test site",
						},
					},
				),
			},
		},
	})
}
