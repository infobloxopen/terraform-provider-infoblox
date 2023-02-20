package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTXTRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceTXTRecordsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_txt_record.acctest", "dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.acctest", "zone", "test.com"),
					resource.TestCheckResourceAttr("data.infoblox_txt_record.acctest", "fqdn", "test-name.test.com"),
				),
			},
		},
	})
}

var testAccDataSourceTXTRecordsRead = fmt.Sprintf(`
resource "infoblox_txt_record" "foo"{
	dns_view="default"
	fqdn="test-name.test.com"
}

data "infoblox_txt_record" "acctest"{
	dns_view="default"
	fqdn=infoblox_txt_record.foo.fqdn
}
`)
