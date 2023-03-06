package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSRVRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceSRVRecordsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_srv_record.acctest", "dns_view", "nondefault_view"),
					resource.TestCheckResourceAttr("data.infoblox_srv_record.acctest", "name", "_http._udp.test1.com"),
				),
			},
		},
	})
}

var testAccDataSourceSRVRecordsRead = fmt.Sprintf(`
resource "infoblox_srv_record" "foo" {
	dns_view="nondefault_view"
	name="_http._udp.test1.com"
	priority=50
	weight=40
	port=8080
	target="sub.test.com"
}

data "infoblox_srv_record" "acctest" {
	dns_view=infoblox_srv_record.foo.dns_view
	name=infoblox_srv_record.foo.name
}
`)
