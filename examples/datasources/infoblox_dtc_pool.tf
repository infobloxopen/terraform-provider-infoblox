resource "infoblox_dtc_pool" "dtc_pool"{
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
  availability = "QUORUM"
  quorum = 2
  ttl = 120
  ext_attrs = jsonencode({
    "Site" = "Blr"
  })
  consolidated_monitors{
    monitor_name = "http"
    monitor_type = "http"
    members = ["infoblox.localdomain"]
    availability= "ALL"
    full_health_communication= true
  }
  disble = true
}


data "infoblox_dtc_pool" "testPool_read" {
  filters = {
    name = "pool-test.com"
  }
  depends_on=[infoblox_dtc_pool.dtc_pool]
}

output "dtc_rec_res" {
  value = data.infoblox_dtc_pool.testPool_read
}

data "infoblox_dtc_pool" "dtcPool_ea" {
  filters = {
    "*Site" = "Blr"
  }
}

// throws matching pools with EA, if any
output "dtcPool_out" {
  value = data.infoblox_dtc_pool.dtcPool_ea
}