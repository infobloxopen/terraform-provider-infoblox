package infoblox

import (
	"fmt"
	"github.com/infobloxopen/infoblox-go-client/v2/utils"
	"sort"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

type v4addrsType []ibclient.HostRecordIpv4Addr

func (al v4addrsType) Len() int {
	return len(al)
}

func (al v4addrsType) Less(i, j int) bool {
	return *al[i].Ipv4Addr < *al[j].Ipv4Addr
}

func (al v4addrsType) Swap(i, j int) {
	al[i], al[j] = al[j], al[i]
}

type v6addrsType []ibclient.HostRecordIpv6Addr

func (al v6addrsType) Len() int {
	return len(al)
}

func (al v6addrsType) Less(i, j int) bool {
	return *al[i].Ipv6Addr < *al[j].Ipv6Addr
}

func (al v6addrsType) Swap(i, j int) {
	al[i], al[j] = al[j], al[i]
}

// must be used only with exp and act of the same length
func validateV4Addrs(exp, act []ibclient.HostRecordIpv4Addr) error {
	sort.Sort(v4addrsType(exp))
	sort.Sort(v4addrsType(act))

	for i, expAddr := range exp {
		actAddr := act[i]
		if actAddr.Ipv4Addr == nil {
			return fmt.Errorf("'ipv4addr' is expected to be defined but it is not")
		}
		if *expAddr.Ipv4Addr != *actAddr.Ipv4Addr {
			return fmt.Errorf(
				"expected IPv4 address '%s' does not equal to the actual one '%s'",
				*expAddr.Ipv4Addr, *actAddr.Ipv4Addr)
		}
	}

	return nil
}

func validateV6Addrs(exp, act []ibclient.HostRecordIpv6Addr) error {
	sort.Sort(v6addrsType(exp))
	sort.Sort(v6addrsType(act))

	for i, expAddr := range exp {
		actAddr := act[i]
		if actAddr.Ipv6Addr == nil {
			return fmt.Errorf("'ipv6addr' is expected to be defined but it is not")
		}
		if *expAddr.Ipv6Addr != *actAddr.Ipv6Addr {
			return fmt.Errorf(
				"expected IPv6 address '%s' does not equal to the actual one '%s'",
				*expAddr.Ipv6Addr, *actAddr.Ipv6Addr)
		}
	}

	return nil
}

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

		ref, found := res.Primary.Attributes["ref"]
		if !found {
			return fmt.Errorf("'ref' attribute is not set")
		}

		internalId := newInternalResourceIdFromString(id)
		if internalId == nil {
			return fmt.Errorf("resource ID '%s' has an invalid format", id)
		}
		connector := testAccProvider.Meta().(ibclient.IBConnector)
		objMgr := ibclient.NewObjectManager(
			connector,
			"terraform_test",
			"terraform_test_tenant")
		ipAlloc, err := objMgr.SearchHostRecordByAltId(internalId.String(), ref, eaNameForInternalId)
		if err != nil {
			if isNotFoundError(err) {
				if expectedValue == nil {
					return nil
				}
				return fmt.Errorf("object with Terraform ID '%s' not found, but expected to exist", internalId)
			}
		}
		expNv := expectedValue.NetworkView
		if ipAlloc.NetworkView != expNv {
			return fmt.Errorf(
				"the value of 'network_view' field is '%s', but expected '%s'",
				ipAlloc.NetworkView, expNv)
		}

		if ipAlloc.View == nil {
			return fmt.Errorf("'dns_view' is expected to be defined but it is not")
		}
		expDnsView := expectedValue.View
		if *ipAlloc.View != *expDnsView {
			return fmt.Errorf(
				"the value of 'dns_view' field is '%s', but expected '%s'",
				*ipAlloc.View, *expDnsView)
		}

		if ipAlloc.EnableDns == nil {
			return fmt.Errorf("'enable_dns' is expected to be defined but it is not")
		}
		expEnableDns := expectedValue.EnableDns
		if *ipAlloc.EnableDns != *expEnableDns {
			return fmt.Errorf(
				"the value of 'enable_dns' field is '%t', but expected '%t'",
				*ipAlloc.EnableDns, *expEnableDns)
		}

		if ipAlloc.UseTtl != nil {
			if expectedValue.UseTtl == nil {
				return fmt.Errorf("'use_ttl' is expected to be undefined but it is not")
			}
			if *ipAlloc.UseTtl != *expectedValue.UseTtl {
				return fmt.Errorf(
					"'use_ttl' does not match: got '%t', expected '%t'",
					*ipAlloc.UseTtl, *expectedValue.UseTtl)
			}
			if *ipAlloc.UseTtl {
				if *ipAlloc.Ttl != *expectedValue.Ttl {
					return fmt.Errorf(
						"'TTL' usage does not match: got '%d', expected '%d'",
						ipAlloc.Ttl, expectedValue.Ttl)
				}
			}
		}

		if ipAlloc.Name == nil {
			return fmt.Errorf("'fqdn' is expected not to be nil")
		}
		if expectedValue.Name == nil {
			panic("'fqdn' is expected not to be nil")
		}
		if *ipAlloc.Name != *expectedValue.Name {
			return fmt.Errorf(
				"'fqdn' does not match: got '%s', expected '%s'",
				*ipAlloc.Name,
				*expectedValue.Name)
		}

		expComment := expectedValue.Comment
		if ipAlloc.Comment != nil {
			if expComment == nil {
				return fmt.Errorf("'comment' is expected to be undefined but it is not")
			}
			if *ipAlloc.Comment != *expComment {
				return fmt.Errorf(
					"'comment' does not match: got '%s', expected '%s'",
					*ipAlloc.Comment, *expComment)
			}
		} else if expectedValue.Comment != nil {
			return fmt.Errorf("'comment' is expected to be defined but it is not")
		}

		expV4Addrs := expectedValue.Ipv4Addrs
		actualV4Addrs := ipAlloc.Ipv4Addrs
		if (expV4Addrs == nil) != (actualV4Addrs == nil) {
			return fmt.Errorf("one of the expected and actual IPv4 address lists is 'nil' while the other one is not")
		}
		if expV4Addrs != nil {
			if len(expV4Addrs) != len(actualV4Addrs) {
				return fmt.Errorf("expected and actual IPv4 address lists are not of equal length")
			}
			if err = validateV4Addrs(expV4Addrs, actualV4Addrs); err != nil {
				return err
			}
		}

		expV6Addrs := expectedValue.Ipv6Addrs
		actualV6Addrs := ipAlloc.Ipv6Addrs
		if (expV6Addrs == nil) != (actualV6Addrs == nil) {
			return fmt.Errorf("one of the expected and actual IPv6 address lists is 'nil' while the other one is not")
		}
		if expV6Addrs != nil {
			if len(expV6Addrs) != len(actualV6Addrs) {
				return fmt.Errorf("expected and actual IPv6 address lists are not of equal length")
			}
			if err = validateV6Addrs(expV6Addrs, actualV6Addrs); err != nil {
				return err
			}
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
						View:        utils.StringPtr("default"),
						EnableDns:   utils.BoolPtr(true),
						Name:        utils.StringPtr("testhostnameip.test.com"),
						Ipv6Addrs:   []ibclient.HostRecordIpv6Addr{*ibclient.NewHostRecordIpv6Addr("2001:db8:abcd:12::1", "", false, "")},
						Ipv4Addrs:   []ibclient.HostRecordIpv4Addr{*ibclient.NewHostRecordIpv4Addr("10.0.0.1", "", false, "")},
						UseTtl:      utils.BoolPtr(false),
						Comment:     utils.StringPtr("IPv4 and IPv6 are allocated"),
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
					ttl = 10
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
						View:        utils.StringPtr("default"),
						EnableDns:   utils.BoolPtr(true),
						Name:        utils.StringPtr("testhostnameip2.test.com"),
						Ipv6Addrs:   []ibclient.HostRecordIpv6Addr{*ibclient.NewHostRecordIpv6Addr("2001:db8:abcd:12::2", "", false, "")},
						Ipv4Addrs:   []ibclient.HostRecordIpv4Addr{*ibclient.NewHostRecordIpv4Addr("10.0.0.2", "", false, "")},
						UseTtl:      utils.BoolPtr(true),
						Ttl:         utils.Uint32Ptr(10),
						Comment:     utils.StringPtr("IPv4 and IPv6 are allocated"),
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"VM Name":   "tf-ec2-instance",
							"Location":  "Test loc.",
							"Site":      "Test site",
						},
					},
				),
			},
			// Validate that ipv6_addr can be removed
			{
				Config: `
				resource "infoblox_ip_allocation" "foo3"{
					network_view="default"
					dns_view = "default"
					fqdn="testhostnameip2.test.com"
					ipv4_addr="10.0.0.2"
					comment = "IPv4 is allocated"
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
						View:        utils.StringPtr("default"),
						EnableDns:   utils.BoolPtr(true),
						Name:        utils.StringPtr("testhostnameip2.test.com"),
						Ipv4Addrs:   []ibclient.HostRecordIpv4Addr{*ibclient.NewHostRecordIpv4Addr("10.0.0.2", "", false, "")},
						UseTtl:      utils.BoolPtr(false),
						Comment:     utils.StringPtr("IPv4 is allocated"),
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
				resource "infoblox_ipv4_network" "net1" {
					cidr = "10.0.0.0/24"
				}
				resource "infoblox_ip_allocation" "foo3"{
					network_view="default"
					enable_dns = "false"
					fqdn="testhostnameip3"
					ipv4_addr="10.0.0.2"
					comment = "DNS disabled"
					ext_attrs = jsonencode({
						"VM Name" =  "tf-ec2-instance"
						"Tenant ID" = "terraform_test_tenant"
						Location = "Test loc."
						Site = "Test site"
					  })
					depends_on = [infoblox_ipv4_network.net1]
				}`,
				Check: validateIPAllocation(
					"infoblox_ip_allocation.foo3",
					&ibclient.HostRecord{
						NetworkView: "default",
						View:        utils.StringPtr(" "),
						EnableDns:   utils.BoolPtr(false),
						Name:        utils.StringPtr("testhostnameip3"),
						Ipv4Addrs:   []ibclient.HostRecordIpv4Addr{*ibclient.NewHostRecordIpv4Addr("10.0.0.2", "", false, "")},
						UseTtl:      utils.BoolPtr(false),
						Comment:     utils.StringPtr("DNS disabled"),
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
				resource "infoblox_ipv4_network" "net1" {
					cidr = "10.0.0.0/24"
				}
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
						View:        utils.StringPtr("default"),
						EnableDns:   utils.BoolPtr(true),
						Name:        utils.StringPtr("testhostnameip2.test.com"),
						Ipv4Addrs:   []ibclient.HostRecordIpv4Addr{*ibclient.NewHostRecordIpv4Addr("10.0.0.2", "", false, "")},
						UseTtl:      utils.BoolPtr(false),
						Comment:     utils.StringPtr("IPv4 and IPv6 are allocated"),
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
		ref, found := rs.Primary.Attributes["ref"]
		if !found {
			return fmt.Errorf("resource with ID '%s' has no NIOS object reference", rs.Primary.ID)
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
