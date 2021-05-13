package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

var testAccresourceIPv4AssociationCreate = fmt.Sprintf(`
resource "infoblox_ipv4_association" "foo"{
	network_view_name="%s"
	cidr="10.0.0.0/24"
	ip_addr="10.0.0.1"
	mac_addr="11:22:33:44:55:66"
	comment = "10.0.0.1 IP is associated "
	extensible_attributes = jsonencode({
		"VM Name" =  "tf-ec2-instance"
		"Tenant ID" = "terraform_test_tenant"
		Location = "Test loc."
		Site = "Test site"
		TestEA1 = ["text1","text2"]
		TestEA2 = [4,5]
	  })
	}`, testNetView)

var testAccresourceIPv4AssociationUpdate = fmt.Sprintf(`
resource "infoblox_ipv4_association" "foo"{
	network_view_name="%s"
	cidr="10.0.0.0/24"
	ip_addr="10.0.0.1"
	mac_addr="11:22:33:44:55:66"
	comment = "10.0.0.1 IP is associated updated"
	extensible_attributes = jsonencode({
		"VM Name" =  "tf-ec2-instance"
		"Tenant ID" = "terraform_test_tenant"
		Location = "Test loc. updated"
		TestEA1 = "text3"
		TestEA2 = 7
	  })
	}`, testNetView)

var testAccresourceIPv6AssociationCreate = fmt.Sprintf(`
	resource "infoblox_ipv6_association" "foo"{
		network_view_name="%s"
		cidr="2001:db8:abcd:12::/64"
		ip_addr="2001:db8:abcd:12::1"
		duid="11:22:33:44:55:66"
		comment = "2001:db8:abcd:12::1 IP is associated"
		extensible_attributes = jsonencode({
			"VM Name" =  "tf-ec2-instance-ipv6"
			"Tenant ID" = "terraform_test_tenant"
			Location = "Test loc."
			Site = "Test site"
			TestEA1 = ["text1","text2"]
			TestEA2 = [4,5]
		  })
		}`, testNetView)

var testAccresourceIPv6AssociationUpdate = fmt.Sprintf(`
	resource "infoblox_ipv6_association" "foo"{
		network_view_name="%s"
		cidr="2001:db8:abcd:12::/64"
		ip_addr="2001:db8:abcd:12::1"
		duid="11:22:33:44:55:66"
		comment = "2001:db8:abcd:12::1 IP is associated updated"
		extensible_attributes = jsonencode({
			"VM Name" =  "tf-ec2-instance-ipv6"
			"Tenant ID" = "terraform_test_tenant"
			Location = "Test loc. updated"
			TestEA1 = "text3"
			TestEA2 = 7
		  })
		}`, testNetView)

func validateIPAssociation(
	resourceName string,
	expectedValue *ibclient.FixedAddress) resource.TestCheckFunc {
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
		ipAsso, err := objMgr.GetFixedAddressByRef(id)
		if err != nil {
			if isNotFoundError(err) {
				if expectedValue == nil {
					return nil
				}
				return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
			}
		}
		expNv := expectedValue.NetviewName
		if ipAsso.NetviewName != expNv {
			return fmt.Errorf(
				"the value of 'network_view_name' field is '%s', but expected '%s'",
				ipAsso.NetviewName, expNv)
		}

		expComment := expectedValue.Comment
		if ipAsso.Comment != expComment {
			return fmt.Errorf(
				"the value of 'comment' field is '%s', but expected '%s'",
				ipAsso.Comment, expComment)
		}

		expIPv4Address := expectedValue.IPv4Address
		expIPv6Address := expectedValue.IPv6Address
		if ipAsso.IPv4Address != expIPv4Address || ipAsso.IPv6Address != expIPv6Address {
			return fmt.Errorf(
				"the value of 'IPv4Address' field is '%s', but expected '%s' or the value of 'IPv6Address' field is '%s', but expected '%s' ",
				ipAsso.IPv4Address, expIPv4Address, ipAsso.IPv6Address, expIPv6Address)
		}

		expMACAddress := expectedValue.Mac
		expDUID := expectedValue.Duid
		if ipAsso.Mac != expMACAddress || ipAsso.Duid != expDUID {
			return fmt.Errorf(
				"the value of 'IPv4Address' field is '%s', but expected '%s' or the value of 'IPv6Address' field is '%s', but expected '%s' ",
				ipAsso.IPv4Address, expIPv4Address, ipAsso.IPv6Address, expIPv6Address)
		}

		// the rest is about extensible attributes
		expectedEAs := expectedValue.Ea
		if expectedEAs == nil && ipAsso.Ea != nil {
			return fmt.Errorf(
				"the object with ID '%s' has 'extensible_attributes' field, but it is not expected to exist", id)
		}
		if expectedEAs != nil && ipAsso.Ea == nil {
			return fmt.Errorf(
				"the object with ID '%s' has no 'extensible_attributes' field, but it is expected to exist", id)
		}
		if expectedEAs == nil {
			return nil
		}

		return validateEAs(ipAsso.Ea, expectedEAs)
	}
}

func TestAcc_resourceipAssociation_ipv4(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccresourceIPv4AllocationCreate,
				Check: validateIPAssociation(
					"infoblox_ipv4_association.foo",
					&ibclient.FixedAddress{
						NetviewName: testNetView,
						Cidr:        "10.0.0.0/24",
						Comment:     "10.0.0.1 IP is associated",
						IPv4Address: "10.0.0.1",
						Mac:         "11:22:33:44:55:66",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"VM Name":   "tf-ec2-instance",
							"Location":  "Test loc.",
							"Site":      "Test site",
							"TestEA1":   []string{"text1", "text2"},
							"TestEA2":   []int{4, 5},
						},
					},
				),
			},
			{
				Config: testAccresourceIPv4AssociationUpdate,
				Check: validateIPAssociation(
					"infoblox_ipv4_association.foo",
					&ibclient.FixedAddress{
						NetviewName: testNetView,
						Cidr:        "10.0.0.0/24",
						Comment:     "10.0.0.1 IP is allocated updated",
						IPv4Address: "10.0.0.1",
						Mac:         "11:22:33:44:55:66",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"VM Name":   "tf-ec2-instance",
							"Location":  "Test loc. updated",
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

func TestAcc_resourceIPAssociation_ipv6(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccresourceIPv6AssociationCreate,
				Check: validateIPAssociation(
					"infoblox_ipv6_association.foo",
					&ibclient.FixedAddress{
						NetviewName: testNetView,
						Cidr:        "2001:db8:abcd:12::/64",
						Comment:     "2001:db8:abcd:12::1 IP is associated",
						IPv6Address: "2001:db8:abcd:12::1",
						Duid:        "11:22:33:44:55:66",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"VM Name":   "tf-ec2-instance-ipv6",
							"Location":  "Test loc.",
							"Site":      "Test site",
							"TestEA1":   []string{"text1", "text2"},
							"TestEA2":   []int{4, 5},
						},
					},
				),
			},
			{
				Config: testAccresourceIPv6AssociationUpdate,
				Check: validateIPAssociation(
					"infoblox_ipv6_association.foo",
					&ibclient.FixedAddress{
						NetviewName: testNetView,
						Cidr:        "2001:db8:abcd:12::/64",
						Comment:     "2001:db8:abcd:12::1 IP is associated updated",
						IPv6Address: "2001:db8:abcd:12::1",
						Duid:        "11:22:33:44:55:66",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"VM Name":   "tf-ec2-instance-ipv6",
							"Location":  "Test loc. updated",
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
		res, err := objMgr.GetFixedAddressByRef(rs.Primary.ID)
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
