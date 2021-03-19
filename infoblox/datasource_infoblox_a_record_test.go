package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataSourceARecord(t *testing.T) {
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
				Config: testAccDataSourceARecordsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_a_record.acctest", "dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_a_record.acctest", "zone", "a.com"),
					resource.TestCheckResourceAttr("data.infoblox_a_record.acctest", "fqdn", "test-name.a.com"),
					resource.TestCheckResourceAttr("data.infoblox_a_record.acctest", "ip_addr", "10.0.0.2"),
					testARecordEAs(t, "data.infoblox_a_record.acctest", "eas", expected_eas),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testARecordEAs(t *testing.T, data_source string, eas_field string, expected_EAs map[string]string) resource.TestCheckFunc {
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

var testAccDataSourceARecordsRead = `
resource "infoblox_a_record" "foo"{
	network_view_name="test"
	vm_name="test-name"
	dns_view="default"
	zone="a.com"
	cidr="10.0.0.0/24"
	ip_addr="10.0.0.2"
	tenant_id="foo"
}

data "infoblox_a_record" "acctest" {
	depends_on = [
		infoblox_a_record.foo,
	]
	zone="a.com"
}
`
