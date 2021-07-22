package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func validateIPAllocation(
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
		ipAlloc, err := objMgr.GetHostRecordByRef(id)
		if err != nil {
			if isNotFoundError(err) {
				if expectedValue == nil {
					return nil
				}
				return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
			}
		}
		expNv := expectedValue.NetworkView
		if ipAlloc.NetworkView != expNv {
			return fmt.Errorf(
				"the value of 'network_view' field is '%s', but expected '%s'",
				ipAlloc.NetworkView, expNv)
		}

		expFqdn := expectedValue.Name
		if ipAlloc.Name != expFqdn {
			return fmt.Errorf(
				"the value of 'fqdn' field is '%s', but expected '%s'",
				ipAlloc.Name, expFqdn)
		}

		expComment := expectedValue.Comment
		if ipAlloc.Comment != expComment {
			return fmt.Errorf(
				"the value of 'comment' field is '%s', but expected '%s'",
				ipAlloc.Comment, expComment)
		}

		// the rest is about extensible attributes
		expectedEAs := expectedValue.Ea
		if expectedEAs == nil && ipAlloc.Ea != nil {
			return fmt.Errorf(
				"the object with ID '%s' has 'ext_attrs' field, but it is not expected to exist", id)
		}
		if expectedEAs != nil && ipAlloc.Ea == nil {
			return fmt.Errorf(
				"the object with ID '%s' has no 'ext_attrs' field, but it is expected to exist", id)
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
				Config: fmt.Sprintf(`
				resource "infoblox_ipv4_allocation" "foo"{
					network_view="default"
					dns_view = "default"
					fqdn="testhostname.test.com"
					cidr="10.0.0.0/24"
					ip_addr="10.0.0.1"
					enable_dns = "true"
					comment = "10.0.0.1 IP is allocated"
					ext_attrs = jsonencode({
						"Tenant ID" = "terraform_test_tenant"
						"VM Name" =  "tf-ec2-instance"
						"Location" = "Test loc."
						"Site" = "Test site"
					  })
					}`),
				Check: validateIPAllocation(
					"infoblox_ipv4_allocation.foo",
					&ibclient.HostRecord{
						NetworkView: "default",
						View:        "default",
						Name:        "testhostname.test.com",
						Ipv4Addr:    "10.0.0.1",
						Comment:     "10.0.0.1 IP is allocated",
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

func TestAcc_resourceIPAllocation_ipv6(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPAllocationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "infoblox_ipv6_allocation" "foo2"{
					network_view="default"
					dns_view = "default"
					fqdn="testhostnameipv6.test.com"
					ip_addr="2001:db8:abcd:12::1"
					duid="11:22:33:44:55:66"
					comment = "2001:db8:abcd:12::1 IP is allocated"
					ext_attrs = jsonencode({
						"VM Name" =  "tf-ec2-instance-ipv6"
						"Tenant ID" = "terraform_test_tenant"
						Location = "Test loc."
						Site = "Test site"
					  })
					}`),
				Check: validateIPAllocation(
					"infoblox_ipv6_allocation.foo2",
					&ibclient.HostRecord{
						NetworkView: "default",
						View:        "default",
						Name:        "testhostnameipv6.test.com",
						Ipv6Addr:    "2001:db8:abcd:12::1",
						Comment:     "2001:db8:abcd:12::1 IP is allocated",
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
