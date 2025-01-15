package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

var testAccDataSourceDtcPool = fmt.Sprintf(`resource "infoblox_dtc_pool" "pool11"{
	name = "testPool_read1"
	lb_preferred_method  = "ROUND_ROBIN"
}
data "infoblox_dtc_pool" "testPool_read1" {	
	filters = {
	    name = infoblox_dtc_pool.pool11.name
    }
    depends_on = [infoblox_dtc_pool.pool11]
}`)

func TestAccDataSourceDtcPool(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDtcPool,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read1", "results.0.name", "testPool_read1"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read1", "results.0.lb_preferred_method", "ROUND_ROBIN"),
				),
			}},
	})
}

func TestAccDataSourceDtcPoolSearchByEA(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
						resource "infoblox_dtc_pool" "pool12"{
									name="pool-test.com"
									comment="pool creation"
									lb_preferred_method="TOPOLOGY"
									lb_preferred_topology="topology_ruleset1"
									monitors{
											monitor_name = "snmp"
     										monitor_type="snmp"
									}
									monitors{
											monitor_name = "http"
     										monitor_type="http"
									}
									lb_alternate_method="DYNAMIC_RATIO"
									lb_dynamic_ratio_alternate = jsonencode({
    										"monitor_name"="snmp"
    										"monitor_type"="snmp"
    										"method"="MONITOR"
   	 										"monitor_metric"=".1.2"
    										"monitor_weighing"="PRIORITY"
    										"invert_monitor_metric"=true
										})
									servers{
    										server = "dummy-server.com"
    										ratio=3
  									}
									servers{
    										server = "server-test.com"
    										ratio=3
  									}
  									servers{
   	 										server = "server-test1.com"
    										ratio= 4
									}
									auto_consolidated_monitors=true
									availability = "QUORUM"
									quorum = 2
									ttl = 120
									ext_attrs = jsonencode({
    								"Site" = "Blr"
  									})
						}
						data "infoblox_dtc_pool" "testPool_read" {	
						filters = {
	   			 			name = infoblox_dtc_pool.pool12.name
    						}
						depends_on=[infoblox_dtc_pool.pool12]
						}
						resource "infoblox_dtc_server" "server11"{
											name="server-test.com"
											host="2.3.4.5"
						}
						resource "infoblox_dtc_server" "server12"{
											name="server-test1.com"
											host="2.3.4.6"
						}`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.#", "1"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.name", "pool-test.com"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.comment", "pool creation"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.lb_preferred_method", "TOPOLOGY"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.monitors.0.monitor_name", "snmp"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.monitors.0.monitor_type", "snmp"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.lb_dynamic_ratio_alternate", "{\"invert_monitor_metric\":true,\"method\":\"MONITOR\",\"monitor_metric\":\".1.2\",\"monitor_name\":\"snmp\",\"monitor_type\":\"snmp\",\"monitor_weighing\":\"PRIORITY\"}"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.servers.0.server", "dummy-server.com"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.servers.0.ratio", "3"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.servers.1.server", "server-test.com"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.servers.1.ratio", "3"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.servers.2.ratio", "4"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.servers.2.server", "server-test1.com"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.lb_preferred_topology", "topology_ruleset1"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.lb_alternate_method", "DYNAMIC_RATIO"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.auto_consolidated_monitors", "true"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.consolidated_monitors.0.monitor_name", "snmp"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.consolidated_monitors.0.monitor_type", "snmp"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.consolidated_monitors.0.availability", "ALL"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.consolidated_monitors.0.full_health_communication", "true"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.availability", "QUORUM"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.quorum", "2"),
					resource.TestCheckResourceAttr("data.infoblox_dtc_pool.testPool_read", "results.0.ttl", "120"),
					resource.TestCheckResourceAttrPair("data.infoblox_dtc_pool.testPool_read", "results.0.ext_attrs.Site", "infoblox_dtc_pool.pool12", "ext_attrs.Site"),
				),
			}},
	})
}
