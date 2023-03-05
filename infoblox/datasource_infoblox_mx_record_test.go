package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceMXRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceMXRecordsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_mx_record.acctest", "dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.acctest", "fqdn", "test-name.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.acctest", "mail_exchanger", "mx-test.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_mx_record.acctest", "priority", 25),
				),
			},
		},
	})
}

var testAccDataSourceMXRecordsRead = fmt.Sprintf(`
resource "infoblox_mx_record" "foo" {
	dns_view="default"
	fqdn="test-name.test.com"
	mail_exchanger="mxt-test.test.com"
	priority=25
}

data "infoblox_mx_record" "acctest" {
	dns_view="default"
	fqdn=infoblox_mx_record.foo.fqdn
}
`)
