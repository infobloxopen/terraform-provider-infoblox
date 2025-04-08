//Pool creation with minimal parameters
resource "infoblox_dtc_pool" "test_pool1" {
  name                = "Pool63"
  lb_preferred_method = "ROUND_ROBIN"
}

//Pool creation with maximal parameters
//parameters for DTC pool when preferred load balancing is TOPOLOGY and alternate load balancing is DYNAMIC_RATIO
resource "infoblox_dtc_pool" "pool" {
  name                  = "terraform_pool.com"
  comment               = "testing pool terraform"
  lb_alternate_method   = "DYNAMIC_RATIO"
  lb_preferred_method   = "TOPOLOGY"
  lb_preferred_topology = "topology_ruleset"
  availability          = "QUORUM"
  quorum                = 2
  ttl                   = 120
  disable               = true

  lb_dynamic_ratio_alternate = jsonencode({
    "monitor_name"          = "snmp"
    "monitor_type"          = "snmp"
    "method"                = "MONITOR"
    "monitor_metric"        = ".1.2"
    "monitor_weighing"      = "PRIORITY"
    "invert_monitor_metric" = true
  })

  servers {
    server = "server.com"
    ratio  = 3
  }
  servers {
    server = "terraform_server.com"
    ratio  = 3
  }
  servers {
    server = "terraform_server1.com"
    ratio  = 4
  }

  monitors {
    monitor_name = "http"
    monitor_type = "http"
  }
  monitors {
    monitor_name = "snmp"
    monitor_type = "snmp"
  }

  consolidated_monitors {
    monitor_name              = "http"
    monitor_type              = "http"
    members                   = ["infoblox.localdomain"]
    availability              = "ALL"
    full_health_communication = true
  }

  ext_attrs = jsonencode({
    "Site" = "Blr"
  })
}

//parameters for DTC pool when preferred load balancing method is DYNAMIC_RATIO
resource "infoblox_dtc_pool" "test_pool3" {
  name                = "Pool64"
  lb_preferred_method = "DYNAMIC_RATIO"
  lb_dynamic_ratio_preferred = jsonencode({
    "monitor_name"          = "snmp"
    "monitor_type"          = "snmp"
    "method"                = "MONITOR"
    "monitor_metric"        = ".1.2"
    "monitor_weighing"      = "PRIORITY"
    "invert_monitor_metric" = true
  })

  monitors {
    monitor_name = "snmp"
    monitor_type = "snmp"
  }
}
