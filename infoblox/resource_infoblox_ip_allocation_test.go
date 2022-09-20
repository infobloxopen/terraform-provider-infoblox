package infoblox

import (
	"fmt"
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
	return al[i].Ipv4Addr < al[j].Ipv4Addr
}

func (al v4addrsType) Swap(i, j int) {
	al[i], al[j] = al[j], al[i]
}

type v6addrsType []ibclient.HostRecordIpv6Addr

func (al v6addrsType) Len() int {
	return len(al)
}

func (al v6addrsType) Less(i, j int) bool {
	return al[i].Ipv6Addr < al[j].Ipv6Addr
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
		if expAddr.Ipv4Addr != actAddr.Ipv4Addr {
			return fmt.Errorf(
				"expected IPv4 address '%s' does not equal to the actual one '%s'",
				expAddr.Ipv4Addr, actAddr.Ipv4Addr)
		}
	}

	return nil
}

func validateV6Addrs(exp, act []ibclient.HostRecordIpv6Addr) error {
	sort.Sort(v6addrsType(exp))
	sort.Sort(v6addrsType(act))

	for i, expAddr := range exp {
		actAddr := act[i]
		if expAddr.Ipv6Addr != actAddr.Ipv6Addr {
			return fmt.Errorf(
				"expected IPv6 address '%s' does not equal to the actual one '%s'",
				expAddr.Ipv6Addr, actAddr.Ipv6Addr)
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

		internalId, ref := getAltIdFields(id)
		if internalId == nil || ref == "" {
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
				return fmt.Errorf("object with reference '%s' not found, but expected to exist", ref)
			}
		}
		expNv := expectedValue.NetworkView
		if ipAlloc.NetworkView != expNv {
			return fmt.Errorf(
				"the value of 'network_view' field is '%s', but expected '%s'",
				ipAlloc.NetworkView, expNv)
		}

		expDnsView := expectedValue.View
		if ipAlloc.View != expDnsView {
			return fmt.Errorf(
				"the value of 'dns_view' field is '%s', but expected '%s'",
				ipAlloc.View, expDnsView)
		}

		expEnableDns := expectedValue.EnableDns
		if ipAlloc.EnableDns != expEnableDns {
			return fmt.Errorf(
				"the value of 'enable_dns' field is '%t', but expected '%t'",
				ipAlloc.EnableDns, expEnableDns)
		}

		expUseTtl := expectedValue.UseTtl
		if ipAlloc.UseTtl != expUseTtl {
			return fmt.Errorf(
				"the value of 'use_ttl' field is '%t', but expected '%t'",
				ipAlloc.UseTtl, expUseTtl)
		}

		expTtl := expectedValue.Ttl
		if ipAlloc.Ttl != expTtl {
			return fmt.Errorf(
				"the value of 'ttl' field is '%d', but expected '%d'",
				ipAlloc.Ttl, expTtl)
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

		expV4Addrs := expectedValue.Ipv4Addrs
		actualV4Addrs := ipAlloc.Ipv4Addrs
		if expV4Addrs == nil && actualV4Addrs != nil || expV4Addrs != nil && actualV4Addrs == nil {
			return fmt.Errorf("one of expected IPv4 address list and actual IPv4 address list is 'nil' while the other one is not")
		}
		if expV4Addrs != nil {
			if len(expV4Addrs) != len(actualV4Addrs) {
				return fmt.Errorf("expected IPv4 address list and actual IPv4 address list are not of equal length")
			}
			if err = validateV4Addrs(expV4Addrs, actualV4Addrs); err != nil {
				return err
			}
		}

		expV6Addrs := expectedValue.Ipv6Addrs
		actualV6Addrs := ipAlloc.Ipv6Addrs
		if expV6Addrs == nil && actualV6Addrs != nil || expV6Addrs != nil && actualV6Addrs == nil {
			return fmt.Errorf("one of expected IPv6 address list and actual IPv6 address list is 'nil' while the other one is not")
		}
		if expV6Addrs != nil {
			if len(expV6Addrs) != len(actualV6Addrs) {
				return fmt.Errorf("expected IPv6 address list and actual IPv6 address list are not of equal length")
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
						View:        "default",
						EnableDns:   true,
						Name:        "testhostnameip.test.com",
						Ipv6Addrs:   []ibclient.HostRecordIpv6Addr{*ibclient.NewHostRecordIpv6Addr("2001:db8:abcd:12::1", "", false, "")},
						Ipv4Addrs:   []ibclient.HostRecordIpv4Addr{*ibclient.NewHostRecordIpv4Addr("10.0.0.1", "", false, "")},
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
						View:        "default",
						EnableDns:   true,
						Name:        "testhostnameip2.test.com",
						Ipv6Addrs:   []ibclient.HostRecordIpv6Addr{*ibclient.NewHostRecordIpv6Addr("2001:db8:abcd:12::2", "", false, "")},
						Ipv4Addrs:   []ibclient.HostRecordIpv4Addr{*ibclient.NewHostRecordIpv4Addr("10.0.0.2", "", false, "")},
						UseTtl:      true,
						Ttl:         10,
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
						EnableDns:   true,
						Name:        "testhostnameip2.test.com",
						Ipv4Addrs:   []ibclient.HostRecordIpv4Addr{*ibclient.NewHostRecordIpv4Addr("10.0.0.2", "", false, "")},
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
						View:        " ",
						EnableDns:   false,
						Name:        "testhostnameip3",
						Ipv4Addrs:   []ibclient.HostRecordIpv4Addr{*ibclient.NewHostRecordIpv4Addr("10.0.0.2", "", false, "")},
						Comment:     "DNS disabled",
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
		_, ref := getAltIdFields(rs.Primary.ID)
		if ref == "" {
			return fmt.Errorf("resource ID '%s' has an invalid format", rs.Primary.ID)
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
