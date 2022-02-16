package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceAAARecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceARecordsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_aaaa_record.acctest", "dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_aaaa_record.acctest", "zone", "test.com"),
					resource.TestCheckResourceAttr("data.infoblox_aaaa_record.acctest", "fqdn", "test-name.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_aaaa_record.acctest", "ipv6_addr", "2000::20"),
				),
			},
		},
	})
}

var testAccDataSourceAAAARecordsRead = fmt.Sprintf(`
resource "infoblox_aaaa_record" "foo"{
	dns_view="default"
	fqdn="test-name.test.com"
	ipv6_addr="2000::20"
}

data "infoblox_a_record" "acctest" {
	dns_view="default"
	fqdn=infoblox_aaaa_record.foo.fqdn
	ipv6_addr=infoblox_aaaa_record.foo.ipv6_addr
}
`)
