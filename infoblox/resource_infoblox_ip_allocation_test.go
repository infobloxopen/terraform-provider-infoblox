package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client"
)

var testAccresourceIPv4AllocationCreate = fmt.Sprintf(`
resource "infoblox_ipv4_allocation" "foo"{
	network_view_name="%s"
	host_name="testHostName"
	cidr="10.0.0.0/24"
	ip_addr="10.0.0.1"
	comment = "10.0.0.1 IP is allocated"
	extensible_attributes = jsonencode({
		"VM Name" =  "tf-ec2-instance"
		"Tenant ID" = "terraform_test_tenant"
		Location = "Test loc."
		Site = "Test site"
		TestEA1 = ["text1","text2"]
		TestEA2 = [4,5]
	  })
	}`, testNetView)

var testAccresourceIPv4AllocationUpdate = fmt.Sprintf(`
resource "infoblox_ipv4_allocation" "foo"{
	network_view_name="%s"
	host_name="testHostName"
	cidr="10.0.0.0/24"
	ip_addr="10.0.0.1"
	comment = "10.0.0.1 IP is allocated updated"
	extensible_attributes = jsonencode({
		"VM Name" =  "tf-ec2-instance"
		"Tenant ID" = "terraform_test_tenant"
		Location = "Test loc. updated"
		TestEA1 = "text3"
		TestEA2 = 7
	  })
	}`, testNetView)

var testAccresourceIPv6AllocationCreate = fmt.Sprintf(`
	resource "infoblox_ipv6_allocation" "foo"{
		network_view_name="%s"
		cidr="2001:db8:abcd:12::/64"
		ip_addr="2001:db8:abcd:12::1"
		duid="11:22:33:44:55:66"
		comment = "2001:db8:abcd:12::1 IP is allocated"
		extensible_attributes = jsonencode({
			"VM Name" =  "tf-ec2-instance-ipv6"
			"Tenant ID" = "terraform_test_tenant"
			Location = "Test loc."
			Site = "Test site"
			TestEA1 = ["text1","text2"]
			TestEA2 = [4,5]
		  })
		}`, testNetView)

var testAccresourceIPv6AllocationUpdate = fmt.Sprintf(`
	resource "infoblox_ipv6_allocation" "foo"{
		network_view_name="%s"
		cidr="2001:db8:abcd:12::/64"
		ip_addr="2001:db8:abcd:12::1"
		duid="11:22:33:44:55:66"
		comment = "2001:db8:abcd:12::1 IP is allocated updated"
		extensible_attributes = jsonencode({
			"VM Name" =  "tf-ec2-instance-ipv6"
			"Tenant ID" = "terraform_test_tenant"
			Location = "Test loc. updated"
			TestEA1 = "text3"
			TestEA2 = 7
		  })
		}`, testNetView)

func validateIPAllocation(
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
		ipAlloc, err := objMgr.GetFixedAddressByRef(id)
		if err != nil {
			if isNotFoundError(err) {
				if expectedValue == nil {
					return nil
				}
				return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
			}
		}
		expNv := expectedValue.NetviewName
		if ipAlloc.NetviewName != expNv {
			return fmt.Errorf(
				"the value of 'network_view_name' field is '%s', but expected '%s'",
				ipAlloc.NetviewName, expNv)
		}

		expComment := expectedValue.Comment
		if ipAlloc.Comment != expComment {
			return fmt.Errorf(
				"the value of 'comment' field is '%s', but expected '%s'",
				ipAlloc.Comment, expComment)
		}

		expIPv4Address := expectedValue.IPv4Address
		expIPv6Address := expectedValue.IPv6Address
		if ipAlloc.IPv4Address != expIPv4Address || ipAlloc.IPv6Address != expIPv6Address {
			return fmt.Errorf(
				"the value of 'IPv4Address' field is '%s', but expected '%s'or 'IPv6Address' field is '%s', but expected %s",
				ipAlloc.IPv4Address, expIPv4Address, ipAlloc.IPv6Address, expIPv6Address)
		}

		// the rest is about extensible attributes
		expectedEAs := expectedValue.Ea
		if expectedEAs == nil && ipAlloc.Ea != nil {
			return fmt.Errorf(
				"the object with ID '%s' has 'extensible_attributes' field, but it is not expected to exist", id)
		}
		if expectedEAs != nil && ipAlloc.Ea == nil {
			return fmt.Errorf(
				"the object with ID '%s' has no 'extensible_attributes' field, but it is expected to exist", id)
		}
		if expectedEAs == nil {
			return nil
		}

		return validateEAs(ipAlloc.Ea, expectedEAs)
	}
}

func TestAcc_resourceIPAllocation_ipv4(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPAllocationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccresourceIPv4AllocationCreate,
				Check: validateIPAllocation(
					"infoblox_ipv4_allocation.foo",
					&ibclient.FixedAddress{
						NetviewName: testNetView,
						Cidr:        "10.0.0.0/24",
						Comment:     "10.0.0.1 IP is allocated",
						IPv4Address: "10.0.0.1",
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
				Config: testAccresourceIPv4AllocationUpdate,
				Check: validateIPAllocation(
					"infoblox_ipv4_allocation.foo",
					&ibclient.FixedAddress{
						NetviewName: testNetView,
						Cidr:        "10.0.0.0/24",
						Comment:     "10.0.0.1 IP is allocated updated",
						IPv4Address: "10.0.0.1",
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

func TestAcc_resourceIPAllocation_ipv6(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPAllocationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccresourceIPv6AllocationCreate,
				Check: validateIPAllocation(
					"infoblox_ipv6_allocation.foo",
					&ibclient.FixedAddress{
						NetviewName: testNetView,
						Cidr:        "2001:db8:abcd:12::/64",
						Comment:     "2001:db8:abcd:12::1 IP is allocated",
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
				Config: testAccresourceIPv6AllocationUpdate,
				Check: validateIPAllocation(
					"infoblox_ipv6_allocation.foo",
					&ibclient.FixedAddress{
						NetviewName: testNetView,
						Cidr:        "2001:db8:abcd:12::/64",
						Comment:     "2001:db8:abcd:12::1 IP is allocated updated",
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

func testAccCheckIPAllocationDestroy(s *terraform.State) error {
	connector := testAccProvider.Meta().(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(
		connector,
		"terraform_test",
		"terraform_test_tenant")
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_ipv4_allocation" && rs.Type != "infoblox_ipv6_allocation" {
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
