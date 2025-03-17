package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceNSRecord(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceNSRecordRead,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_ns_record.acctest", "results.0.dns_view", "default"),
					resource.TestCheckResourceAttr("data.infoblox_ns_record.acctest", "results.0.name", "test.com"),
					resource.TestCheckResourceAttr("data.infoblox_ns_record.acctest", "results.0.nameserver", "name24.test.com"),
					resource.TestCheckResourceAttr("data.infoblox_ns_record.acctest", "results.0.addresses.0.address", "2.3.2.5"),
					resource.TestCheckResourceAttr("data.infoblox_ns_record.acctest", "results.0.addresses.0.auto_create_ptr", "true"),
					resource.TestCheckResourceAttr("data.infoblox_ns_record.acctest", "results.0.addresses.1.address", "2.3.23.5"),
					resource.TestCheckResourceAttr("data.infoblox_ns_record.acctest", "results.0.addresses.1.auto_create_ptr", "false"),
					resource.TestCheckResourceAttr("data.infoblox_ns_record.acctest", "results.0.addresses.2.address", "2.3.1.5"),
					resource.TestCheckResourceAttr("data.infoblox_ns_record.acctest", "results.0.addresses.2.auto_create_ptr", "true"),
					resource.TestCheckResourceAttr("data.infoblox_ns_record.acctest", "results.0.policy", "Allow Underscore"),
					resource.TestCheckResourceAttr("data.infoblox_ns_record.acctest", "results.0.creator", "STATIC"),
					resource.TestCheckResourceAttr("data.infoblox_ns_record.acctest", "results.0.dns_name", "test.com"),
					resource.TestCheckResourceAttr("data.infoblox_ns_record.acctest", "results.0.ms_delegation_name", ""),
					resource.TestCheckResourceAttr("data.infoblox_ns_record.acctest", "results.0.zone", "test.com"),
				),
			},
		},
	})
}

var testAccDataSourceNSRecordRead = fmt.Sprintf(`
resource "infoblox_ns_record" "foo"{
	name = "test.com"
  nameserver = "name24.test.com"
    addresses{
     address = "2.3.2.5"
      auto_create_ptr=true
    }
  addresses{
     address = "2.3.23.5"
     auto_create_ptr=false
   }
   addresses{
     address = "2.3.1.5"
    auto_create_ptr=true
   }
}

data "infoblox_ns_record" "acctest" {
	filters = {
		view="default"
		name=infoblox_ns_record.foo.name
		nameserver=infoblox_ns_record.foo.nameserver
	}
}
`)
