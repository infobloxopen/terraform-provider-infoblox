package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

		internalId, ref, err := getAltIdFields(id)
		if err != nil {
			return err
		}
		connector := testAccProvider.Meta().(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(
			connector,
			"terraform_test",
			"terraform_test_tenant")
		ipAlloc, err := objMgr.SearchHostRecordByAltId(internalId, ref, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedValue == nil {
					return nil
				}
				return fmt.Errorf("object with reference '%s' not found, but expected to exist", ref)
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

func TestAcc_resourceIPAllocation(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIPAllocationDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "infoblox_ip_allocation" "foo3"{
					network_view="default"
					dns_view = "default"
					fqdn="testhostnameip.test.com"
					ipv6_addr="2001:db8:abcd:12::1"
					ipv4_addr="10.0.0.1"
					comment = "IPv4 and IPv6 are allocated"
					ext_attrs = jsonencode({
						"VM Name" =  "tf-ec2-instance"
						"Tenant ID" = "terraform_test_tenant"
						Location = "Test loc."
						Site = "Test site"
					  })
					}`,
				Check: validateIPAllocation(
					"infoblox_ip_allocation.foo3",
					&ibclient.HostRecord{
						NetworkView: "default",
						View:        "default",
						Name:        "testhostnameip.test.com",
						Ipv6Addr:    "2001:db8:abcd:12::1",
						Ipv4Addr:    "10.0.0.1",
						Comment:     "IPv4 and IPv6 are allocated",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"VM Name":   "tf-ec2-instance",
							"Location":  "Test loc.",
							"Site":      "Test site",
						},
					},
				),
			},
			{
				Config: `
				resource "infoblox_ip_allocation" "foo3"{
					network_view="default"
					dns_view = "default"
					fqdn="testhostnameip2.test.com"
					ipv6_addr="2001:db8:abcd:12::2"
					ipv4_addr="10.0.0.2"
					comment = "IPv4 and IPv6 are allocated"
					ext_attrs = jsonencode({
						"VM Name" =  "tf-ec2-instance"
						"Tenant ID" = "terraform_test_tenant"
						Location = "Test loc."
						Site = "Test site"
					  })
					}`,
				Check: validateIPAllocation(
					"infoblox_ip_allocation.foo3",
					&ibclient.HostRecord{
						NetworkView: "default",
						View:        "default",
						Name:        "testhostnameip2.test.com",
						Ipv6Addr:    "2001:db8:abcd:12::2",
						Ipv4Addr:    "10.0.0.2",
						Comment:     "IPv4 and IPv6 are allocated",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"VM Name":   "tf-ec2-instance",
							"Location":  "Test loc.",
							"Site":      "Test site",
						},
					},
				),
			},
			{
				Config: `
				resource "infoblox_ip_allocation" "foo3"{
					network_view="default"
					dns_view = "default"
					fqdn="testhostnameip2.test.com"
					ipv4_addr="10.0.0.2"
					comment = "IPv4 and IPv6 are allocated"
					ext_attrs = jsonencode({
						"VM Name" =  "tf-ec2-instance"
						"Tenant ID" = "terraform_test_tenant"
						Location = "Test loc."
						Site = "Test site"
					  })
					}`,
				Check: validateIPAllocation(
					"infoblox_ip_allocation.foo3",
					&ibclient.HostRecord{
						NetworkView: "default",
						View:        "default",
						Name:        "testhostnameip2.test.com",
						Ipv4Addr:    "10.0.0.2",
						Comment:     "IPv4 and IPv6 are allocated",
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

func testAccCheckIPAllocationDestroy(s *terraform.State) error {
	connector := testAccProvider.Meta().(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(
		connector,
		"terraform_test",
		"terraform_test_tenant")
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_ip_allocation" {
			continue
		}
		_, ref, err := getAltIdFields(rs.Primary.ID)
		if err != nil {
			return err
		}
		res, err := objMgr.GetHostRecordByRef(ref)
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
