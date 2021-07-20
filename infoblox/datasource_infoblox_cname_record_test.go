package infoblox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceCNameRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceCNameRecordsRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_cname_record.acctest", "dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_cname_record.acctest", "zone", "test.com"),
					resource.TestCheckResourceAttr("data.infoblox_cname_record.acctest", "alias", "test.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_cname_record.acctest", "canonical", "test-name.test.com"),
				),
			},
		},
	})
}

var testAccDataSourceCNameRecordsRead = fmt.Sprintf(`
resource "infoblox_cname_record" "foo"{
	dns_view="default"
	
	alias="test.test.com"
	canonical="test-name.test.com"	
  }

data "infoblox_cname_record" "acctest" {
	dns_view="default"
	
	alias= infoblox_cname_record.foo.alias
	canonical=infoblox_cname_record.foo.canonical
  }
`)
