package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

func validateRecordPTR(
	resourceName string,
	expectedValue *ibclient.RecordPTR) resource.TestCheckFunc {
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
		recPtr, err := objMgr.GetPTRRecordByRef(id)
		if err != nil {
			if isNotFoundError(err) {
				if expectedValue == nil {
					return nil
				}
				return fmt.Errorf("object with ID '%s' not found, but expected to exist", id)
			}
		}
		expPtrdName := expectedValue.PtrdName
		if recPtr.PtrdName != expPtrdName {
			return fmt.Errorf(
				"the value of 'ptrdname' field is '%s', but expected '%s'",
				recPtr.PtrdName, expPtrdName)
		}

		expComment := expectedValue.Comment
		if recPtr.Comment != expComment {
			return fmt.Errorf(
				"the value of 'comment' field is '%s', but expected '%s'",
				recPtr.Comment, expComment)
		}

		expName := expectedValue.Name
		if recPtr.Name != expName {
			return fmt.Errorf(
				"the value of 'name' field is '%s', but expected '%s'",
				recPtr.Name, expName)
		}

		expUseTtl := expectedValue.UseTtl
		if recPtr.UseTtl != expUseTtl {
			return fmt.Errorf(
				"the value of 'use_ttl' field is '%t', but expected '%t'",
				recPtr.UseTtl, expUseTtl)
		}
		if expUseTtl {
			expTtl := expectedValue.Ttl
			if recPtr.Ttl != expTtl {
				return fmt.Errorf(
					"the value of 'ttl' field is '%d', but expected '%d'",
					recPtr.Ttl, expTtl)
			}
		}

		expView := expectedValue.View
		if recPtr.View != expView {
			return fmt.Errorf(
				"the value of 'view' field is '%s', but expected '%s'",
				recPtr.View, expView)
		}

		expZone := expectedValue.Zone
		if recPtr.Zone != expZone {
			return fmt.Errorf(
				"the value of 'zone' field is '%s', but expected '%s'",
				recPtr.Zone, expZone)
		}

		expIpv4Addr := expectedValue.Ipv4Addr
		if recPtr.Ipv4Addr != expIpv4Addr {
			return fmt.Errorf(
				"the value of 'ipv4addr' field is '%s', but expected '%s'",
				recPtr.Ipv4Addr, expIpv4Addr)
		}

		expIpv6Addr := expectedValue.Ipv6Addr
		if recPtr.Ipv6Addr != expIpv6Addr {
			return fmt.Errorf(
				"the value of 'ipv6addr' field is '%s', but expected '%s'",
				recPtr.Ipv6Addr, expIpv6Addr)
		}

		// the rest is about extensible attributes
		expectedEAs := expectedValue.Ea
		if expectedEAs == nil && recPtr.Ea != nil {
			return fmt.Errorf(
				"the object with ID '%s' has 'ext_attrs' field, but it is not expected to exist", id)
		}
		if expectedEAs != nil && recPtr.Ea == nil {
			return fmt.Errorf(
				"the object with ID '%s' has no 'ext_attrs' field, but it is expected to exist", id)
		}
		if expectedEAs == nil {
			return nil
		}

		return validateEAs(recPtr.Ea, expectedEAs)
	}
}

func testAccCheckRecordPTRDestroy(s *terraform.State) error {
	connector := testAccProvider.Meta().(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(
		connector,
		"terraform_test",
		"terraform_test_tenant")
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "infoblox_ptr_record" {
			continue
		}
		res, err := objMgr.GetPTRRecordByRef(rs.Primary.ID)
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

func TestAcc_resourceRecordPTR(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRecordPTRDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "infoblox_ptr_record" "foo"{
						dns_view="default"
						ptrdname="testptrdname.test.com"
						record_name="testname.test.com"
						comment="PTR record created in forward mapping zone"
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc"
							"Site" = "Test site"
							"TestEA1"=["text1","text2"]
						  })
					}`,
				Check: validateRecordPTR(
					"infoblox_ptr_record.foo",
					&ibclient.RecordPTR{
						View:     "default",
						PtrdName: "testptrdname.test.com",
						Name:     "testname.test.com",
						Zone:     "test.com",
						Comment:  "PTR record created in forward mapping zone",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc",
							"Site":      "Test site",
							"TestEA1":   []string{"text1", "text2"},
						},
					},
				),
			},
			{
				Config: `
					resource "infoblox_ptr_record" "foo"{
						dns_view="default"
						ptrdname="testptrdname2.test.com"
						record_name="testname2.test.com"
						comment="PTR record created in forward mapping zone"
						ext_attrs = jsonencode({
							"Tenant ID" = "terraform_test_tenant"
							"Location" = "Test loc2"
							"Site" = "Test site"
							"TestEA1" = ["text1","text2"]
						  })
					}`,
				Check: validateRecordPTR(
					"infoblox_ptr_record.foo",
					&ibclient.RecordPTR{
						View:     "default",
						PtrdName: "testptrdname2.test.com",
						Name:     "testname2.test.com",
						Zone:     "test.com",
						Comment:  "PTR record created in forward mapping zone",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc2",
							"Site":      "Test site",
							"TestEA1":   []string{"text1", "text2"},
						},
					},
				),
			},
			{
				Config: `
					resource "infoblox_ptr_record" "foo2"{
						network_view="default"
						dns_view="default"
						ptrdname="testptrdname2.test.com"
						ip_addr = "10.0.0.2"
						comment="PTR record created in reverse mapping zone with IP"
						ext_attrs=jsonencode({
							"Tenant ID"="terraform_test_tenant"
							"Location"="Test loc."
							"Site"="Test site"
							"TestEA1"=["text1","text2"]
						  })
					}`,
				Check: validateRecordPTR(
					"infoblox_ptr_record.foo2",
					&ibclient.RecordPTR{
						View:     "default",
						PtrdName: "testptrdname2.test.com",
						Ipv4Addr: "10.0.0.2",
						Name:     "2.0.0.10.in-addr.arpa",
						Zone:     "0.0.10.in-addr.arpa",
						Comment:  "PTR record created in reverse mapping zone with IP",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc.",
							"Site":      "Test site",
							"TestEA1":   []string{"text1", "text2"},
						},
					},
				),
			},
			{
				Config: `
					resource "infoblox_ptr_record" "foo2"{
						network_view = "default"
						dns_view="default"
						ptrdname="testptrdname3.test.com"
						ip_addr = "10.0.0.3"
						comment="PTR record created in reverse mapping zone with IP"
						ext_attrs = jsonencode({
							"Tenant ID"="terraform_test_tenant"
							"Location"="Test loc."
							"Site"="Test site2"
							"TestEA1"=["text1","text2"]
						  })
					}`,
				Check: validateRecordPTR(
					"infoblox_ptr_record.foo2",
					&ibclient.RecordPTR{
						View:     "default",
						PtrdName: "testptrdname3.test.com",
						Ipv4Addr: "10.0.0.3",
						Name:     "3.0.0.10.in-addr.arpa",
						Zone:     "0.0.10.in-addr.arpa",
						Comment:  "PTR record created in reverse mapping zone with IP",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc.",
							"Site":      "Test site2",
							"TestEA1":   []string{"text1", "text2"},
						},
					},
				),
			},
			{
				// TODO: implement a requirement of lower-case FQDNs
				Config: `
					resource "infoblox_ptr_record" "foo2"{
						ptrdname="testPtrdName3.test.com"
						record_name = "4.0.0.10.in-addr.arpa"
						comment="PTR record created in reverse mapping zone with IP"
						ext_attrs = jsonencode({
							"Tenant ID"="terraform_test_tenant"
							"Location"="Test loc."
							"Site"="Test site2"
							"TestEA1"=["text1","text2"]
						  })
					}`,
				Check: validateRecordPTR(
					"infoblox_ptr_record.foo2",
					&ibclient.RecordPTR{
						View:     "default",
						PtrdName: "testPtrdName3.test.com",
						Ipv4Addr: "10.0.0.4",
						Name:     "4.0.0.10.in-addr.arpa",
						Zone:     "0.0.10.in-addr.arpa",
						Comment:  "PTR record created in reverse mapping zone with IP",
						Ea: ibclient.EA{
							"Tenant ID": "terraform_test_tenant",
							"Location":  "Test loc.",
							"Site":      "Test site2",
							"TestEA1":   []string{"text1", "text2"},
						},
					},
				),
			},
		},
	})
}
