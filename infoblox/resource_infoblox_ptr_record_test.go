package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
)

var testAccresourceRecordPTRCreate = fmt.Sprintf(`
resource "infoblox_ptr_record" "foo"{
    dns_view="default"
	ptrdname="testPtrdName.test.com"
	record_name="testName.test.com"
	comment="PTR record created in forward mapping zone"
	ext_attrs = jsonencode({
		"Tenant ID" = "terraform_test_tenant"
		"Location" = "Test loc"
		"Site" = "Test site"
		"TestEA1"=["text1","text2"]
	  })
}`)

var testAccresourceRecordPTRCreate_2 = fmt.Sprintf(`
resource "infoblox_ptr_record" "foo2"{
	network_view="default"
    dns_view="default"
	ptrdname="testPtrdName2.test.com"
	ip_addr = "10.0.0.2"
	comment="PTR record created in reverse mapping zone with IP"
	ext_attrs=jsonencode({
		"Tenant ID"="terraform_test_tenant"
		"Location"="Test loc."
		"Site"="Test site"
		"TestEA1"=["text1","text2"]
	  })
}`)

var testAccresourceRecordPTRUpdate = fmt.Sprintf(`
resource "infoblox_ptr_record" "foo"{
	dns_view="default"
	ptrdname="testPtrdName.test.com"
	record_name="testName.test.com"
	comment="PTR record created in forward mapping zone"
	ext_attrs = jsonencode({
		"Tenant ID" = "terraform_test_tenant"
		"Location" = "Test loc"
		"Site" = "Test site"
		"TestEA1" = ["text1","text2"]
	  })
}`)

var testAccresourceRecordPTRUpdate_2 = fmt.Sprintf(`
resource "infoblox_ptr_record" "foo2"{
	network_view = "default"
	dns_view="default"
	ptrdname="testPtrdName2.test.com"
	ip_addr = "10.0.0.2"
	comment="PTR record created in reverse mapping zone with IP"
	ext_attrs = jsonencode({
		"Tenant ID"="terraform_test_tenant"
		"Location"="Test loc."
		"Site"="Test site"
		"TestEA1"=["text1","text2"]
	  })
}`)

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
				Config: testAccresourceRecordPTRCreate,
				Check: validateRecordPTR(
					"infoblox_ptr_record.foo",
					&ibclient.RecordPTR{
						View:     "default",
						PtrdName: "testPtrdName.test.com",
						Name:     "testName",
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
				Config: testAccresourceRecordPTRUpdate,
				Check: validateRecordPTR(
					"infoblox_ptr_record.foo",
					&ibclient.RecordPTR{
						View:     "default",
						PtrdName: "testPtrdName.test.com",
						Name:     "testName",
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
				Config: testAccresourceRecordPTRCreate_2,
				Check: validateRecordPTR(
					"infoblox_ptr_record.foo2",
					&ibclient.RecordPTR{
						View:     "default",
						PtrdName: "testPtrdName2.test.com",
						Ipv4Addr: "10.0.0.2",
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
				Config: testAccresourceRecordPTRUpdate_2,
				Check: validateRecordPTR(
					"infoblox_ptr_record.foo2",
					&ibclient.RecordPTR{
						View:     "default",
						PtrdName: "testPtrdName2.test.com",
						Ipv4Addr: "10.0.0.2",
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
		},
	})
}
