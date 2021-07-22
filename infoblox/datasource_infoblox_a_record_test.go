package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceARecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceARecordsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_a_record.acctest", "dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_a_record.acctest", "zone", "test.com"),
					resource.TestCheckResourceAttr("data.infoblox_a_record.acctest", "fqdn", "test-name.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_a_record.acctest", "ip_addr", "10.0.0.20"),
				),
			},
		},
	})
}

var testAccDataSourceARecordsRead = fmt.Sprintf(`
resource "infoblox_a_record" "foo"{
	dns_view="default"
	fqdn="test-name.test.com"
	ip_addr="10.0.0.20"
}

data "infoblox_a_record" "acctest" {
	dns_view="default"
	fqdn=infoblox_a_record.foo.fqdn
	ip_addr=infoblox_a_record.foo.ip_addr
}
`)
