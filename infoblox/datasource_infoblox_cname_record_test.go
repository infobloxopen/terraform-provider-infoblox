package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataSourceCNameRecord(t *testing.T) {
	expected_eas := map[string]string{
		"CMP Type":        "Terraform",
		"Cloud API Owned": "true",
		"Tenant ID":       "foo",
		"VM Name":         "test-name",
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceCNameRecordsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_cname_record.acctest", "dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_cname_record.acctest", "zone", "a.com"),
					resource.TestCheckResourceAttr("data.infoblox_cname_record.acctest", "fqdn", "test.a.com"),
					resource.TestCheckResourceAttr("data.infoblox_cname_record.acctest", "canonical", "test-name"),
					testCNameRecordEAs(t, "data.infoblox_cname_record.acctest", "eas", expected_eas),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testCNameRecordEAs(t *testing.T, data_source string, eas_field string, expected_EAs map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[data_source]
		if !ok {
			return fmt.Errorf("Not found:%s", data_source)
		}
		for key, expect_value := range expected_EAs {
			attr_key := fmt.Sprintf("%s.%s", eas_field, key)
			attr_value := rs.Primary.Attributes[attr_key]
			if attr_value != expect_value {
				return fmt.Errorf("Expected '%s'='%s' but found '%s'", attr_key, expect_value, attr_value)
			}
		}
		return nil
	}
}

var testAccDataSourceCNameRecordsRead = fmt.Sprintf(`
resource "infoblox_cname_record" "foo"{
	alias="test"
	canonical="test-name"
	dns_view="default"
	zone="a.com"
	tenant_id="foo"
	}

data "infoblox_cname_record" "acctest" {
	depends_on = [
		infoblox_cname_record.foo,
	]
	zone="a.com"
}
`)
