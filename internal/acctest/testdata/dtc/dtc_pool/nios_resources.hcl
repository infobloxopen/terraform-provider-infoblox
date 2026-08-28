# Auto-generated resource acceptance-test cases for DtcPool.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
    }
    check = {
      "nios.name"                                             = "{{random}}"
      "nios.lb_preferred_method"                              = "ROUND_ROBIN"
      "nios.auto_consolidated_monitors"                       = "false"
      "nios.availability"                                     = "ALL"
      "nios.disable"                                          = "false"
      "nios.lb_alternate_method"                              = "NONE"
      "nios.lb_dynamic_ratio_preferred.method"                = "MONITOR"
      "nios.lb_dynamic_ratio_preferred.invert_monitor_metric" = "false"
      "nios.lb_dynamic_ratio_preferred.monitor_weighing"      = "RATIO"
      "nios.lb_dynamic_ratio_alternate.method"                = "MONITOR"
      "nios.lb_dynamic_ratio_alternate.invert_monitor_metric" = "false"
      "nios.lb_dynamic_ratio_alternate.monitor_weighing"      = "RATIO"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
    }
  }

}

case "auto_consolidated_monitors" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                       = "{{random}}"
      lb_preferred_method        = "ROUND_ROBIN"
      auto_consolidated_monitors = true
    }
    check = {
      "nios.auto_consolidated_monitors" = "true"
    }
  }

  step {
    nios {
      name                       = "{{random}}"
      lb_preferred_method        = "ROUND_ROBIN"
      auto_consolidated_monitors = false
    }
    check = {
      "nios.auto_consolidated_monitors" = "false"
    }
  }

}

case "availability" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      availability        = "ANY"
    }
    check = {
      "nios.availability" = "ANY"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      availability        = "ALL"
    }
    check = {
      "nios.availability" = "ALL"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      comment             = "pool testing"
    }
    check = {
      "nios.comment" = "pool testing"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      comment             = "updated pool comment"
    }
    check = {
      "nios.comment" = "updated pool comment"
    }
  }

}

case "consolidated_monitors" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                  = "{{random}}"
      lb_preferred_method   = "ROUND_ROBIN"
      disable = true
      monitors              = ["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"]
      consolidated_monitors = [{ monitor = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", availability = "ANY", full_health_communication = false, members = ["infoblox.172_28_82_8"] }]
    }
    check = {
      "nios.consolidated_monitors.0.monitor"                   = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"
      "nios.consolidated_monitors.0.availability"              = "ANY"
      "nios.consolidated_monitors.0.full_health_communication" = "false"
      "nios.consolidated_monitors.0.members.0"                 = "infoblox.172_28_82_8"
    }
  }

  step {
    nios {
      name                  = "{{random}}"
      lb_preferred_method   = "ROUND_ROBIN"
      disable = true 
      monitors              = ["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"]
      consolidated_monitors = [{ monitor = "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp", availability = "ALL", full_health_communication = false, members = ["infoblox.172_28_82_8"] }]
    }
    check = {
      "nios.consolidated_monitors.0.monitor"                   = "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"
      "nios.consolidated_monitors.0.availability"              = "ALL"
      "nios.consolidated_monitors.0.full_health_communication" = "false"
      "nios.consolidated_monitors.0.members.0"                 = "infoblox.172_28_82_8"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      disable             = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      disable             = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      ext_attrs           = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      ext_attrs           = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

# servers stays hardcoded: a test-created server would also need registering in the
# pre-provisioned topology, and infoblox_dtc_topology is not implemented yet.
case "lb_alternate_method" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                  = "{{random}}"
      lb_preferred_method   = "TOPOLOGY"
      lb_preferred_topology = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGVycmFmb3JtX3RvcG9sb2d5X3Rlc3Q:terraform_topology_test"
      servers               = [{ server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Rfc2VydmVyLmNvbQ:test_server.com", ratio = 1 }, { server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Rfc2VydmVyMi5jb20:test_server2.com", ratio = 2 }]
      lb_alternate_method   = "ALL_AVAILABLE"
    }
    check = {
      "nios.lb_preferred_method" = "TOPOLOGY"
      "nios.lb_alternate_method" = "ALL_AVAILABLE"
    }
  }

  step {
    nios {
      name                  = "{{random}}"
      lb_preferred_method   = "TOPOLOGY"
      lb_preferred_topology = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGVycmFmb3JtX3RvcG9sb2d5X3Rlc3Q:terraform_topology_test"
      servers               = [{ server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Rfc2VydmVyLmNvbQ:test_server.com", ratio = 1 }, { server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Rfc2VydmVyMi5jb20:test_server2.com", ratio = 2 }]
      lb_alternate_method   = "GLOBAL_AVAILABILITY"
    }
    check = {
      "nios.lb_preferred_method" = "TOPOLOGY"
      "nios.lb_alternate_method" = "GLOBAL_AVAILABILITY"
    }
  }

}

# servers stays hardcoded: a test-created server would also need registering in the
# pre-provisioned topology, and infoblox_dtc_topology is not implemented yet.
case "lb_alternate_topology" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                  = "{{random}}"
      lb_preferred_method   = "TOPOLOGY"
      lb_preferred_topology = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGVycmFmb3JtX3RvcG9sb2d5X3Rlc3Q:terraform_topology_test"
      lb_alternate_method   = "TOPOLOGY"
      lb_alternate_topology = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGVycmFmb3JtX3RvcG9sb2d5X3Rlc3Qy:terraform_topology_test2"
      servers               = [{ server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Rfc2VydmVyLmNvbQ:test_server.com", ratio = 1 }, { server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Rfc2VydmVyMi5jb20:test_server2.com", ratio = 2 }]
    }
    check = {
      "nios.lb_preferred_method"   = "TOPOLOGY"
      "nios.lb_preferred_topology" = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGVycmFmb3JtX3RvcG9sb2d5X3Rlc3Q:terraform_topology_test"
      "nios.lb_alternate_method"   = "TOPOLOGY"
      "nios.lb_alternate_topology" = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGVycmFmb3JtX3RvcG9sb2d5X3Rlc3Qy:terraform_topology_test2"
    }
  }

  step {
    nios {
      name                  = "{{random}}"
      lb_preferred_method   = "TOPOLOGY"
      lb_preferred_topology = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGVycmFmb3JtX3RvcG9sb2d5X3Rlc3Qy:terraform_topology_test2"
      lb_alternate_method   = "TOPOLOGY"
      lb_alternate_topology = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGVycmFmb3JtX3RvcG9sb2d5X3Rlc3Q:terraform_topology_test"
      servers               = [{ server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Rfc2VydmVyLmNvbQ:test_server.com", ratio = 1 }, { server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Rfc2VydmVyMi5jb20:test_server2.com", ratio = 2 }]
    }
    check = {
      "nios.lb_preferred_method"   = "TOPOLOGY"
      "nios.lb_preferred_topology" = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGVycmFmb3JtX3RvcG9sb2d5X3Rlc3Qy:terraform_topology_test2"
      "nios.lb_alternate_method"   = "TOPOLOGY"
      "nios.lb_alternate_topology" = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGVycmFmb3JtX3RvcG9sb2d5X3Rlc3Q:terraform_topology_test"
    }
  }

}

# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: config helper 'testAccDtcPoolLbDynamicRatioAlternate' could not be parsed (no resource block found)
case "lb_dynamic_ratio_alternate" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                       = "{{random}}"
      lb_preferred_method        = "ROUND_ROBIN"
      lb_alternate_method        = "DYNAMIC_RATIO"
      monitors                   = ["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"]
      lb_dynamic_ratio_alternate = { method = "ROUND_TRIP_DELAY", monitor = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", monitor_metric = ".0", monitor_weighing = "RATIO", invert_monitor_metric = false }
    }
    check = {
      "nios.lb_alternate_method"                              = "DYNAMIC_RATIO"
      "nios.lb_dynamic_ratio_alternate.method"                = "ROUND_TRIP_DELAY"
      "nios.lb_dynamic_ratio_alternate.monitor_metric"        = ".0"
      "nios.lb_dynamic_ratio_alternate.monitor_weighing"      = "RATIO"
      "nios.lb_dynamic_ratio_alternate.invert_monitor_metric" = "false"
    }
  }

  step {
    nios {
      name                       = "{{random}}"
      lb_preferred_method        = "ROUND_ROBIN"
      lb_alternate_method        = "DYNAMIC_RATIO"
      monitors                   = ["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"]
      lb_dynamic_ratio_alternate = { method = "MONITOR", monitor = "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp", monitor_metric = ".2", monitor_weighing = "RATIO", invert_monitor_metric = false }
    }
    check = {
      "nios.lb_alternate_method"                              = "DYNAMIC_RATIO"
      "nios.lb_dynamic_ratio_alternate.method"                = "MONITOR"
      "nios.lb_dynamic_ratio_alternate.monitor_metric"        = ".2"
      "nios.lb_dynamic_ratio_alternate.monitor_weighing"      = "RATIO"
      "nios.lb_dynamic_ratio_alternate.invert_monitor_metric" = "false"
    }
  }

}

case "lb_dynamic_ratio_preferred" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                       = "{{random}}"
      lb_preferred_method        = "DYNAMIC_RATIO"
      monitors                   = ["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"]
      lb_dynamic_ratio_preferred = { method = "ROUND_TRIP_DELAY", monitor = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", monitor_metric = ".0", monitor_weighing = "RATIO", invert_monitor_metric = false }
    }
    check = {
      "nios.lb_preferred_method"                              = "DYNAMIC_RATIO"
      "nios.lb_dynamic_ratio_preferred.method"                = "ROUND_TRIP_DELAY"
      "nios.lb_dynamic_ratio_preferred.monitor_metric"        = ".0"
      "nios.lb_dynamic_ratio_preferred.monitor_weighing"      = "RATIO"
      "nios.lb_dynamic_ratio_preferred.invert_monitor_metric" = "false"
    }
  }

  step {
    nios {
      name                       = "{{random}}"
      lb_preferred_method        = "DYNAMIC_RATIO"
      monitors                   = ["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"]
      lb_dynamic_ratio_preferred = { method = "ROUND_TRIP_DELAY", monitor = "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp", monitor_metric = ".2", monitor_weighing = "RATIO", invert_monitor_metric = true }
    }
    check = {
      "nios.lb_preferred_method"                              = "DYNAMIC_RATIO"
      "nios.lb_dynamic_ratio_preferred.method"                = "ROUND_TRIP_DELAY"
      "nios.lb_dynamic_ratio_preferred.monitor_metric"        = ".2"
      "nios.lb_dynamic_ratio_preferred.monitor_weighing"      = "RATIO"
      "nios.lb_dynamic_ratio_preferred.invert_monitor_metric" = "true"
    }
  }

}

case "lb_preferred_method" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      lb_preferred_method = "ROUND_ROBIN"
      name                = "{{random}}"
    }
    check = {
      "nios.lb_preferred_method" = "ROUND_ROBIN"
    }
  }

  step {
    nios {
      lb_preferred_method = "ALL_AVAILABLE"
      name                = "{{random}}"
    }
    check = {
      "nios.lb_preferred_method" = "ALL_AVAILABLE"
    }
  }

}

case "lb_preferred_method_source_ip_hash" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      lb_preferred_method = "ROUND_ROBIN"
      name                = "{{random}}"
    }
    check = {
      "nios.lb_preferred_method" = "ROUND_ROBIN"
    }
  }

  step {
    nios {
      lb_preferred_method = "SOURCE_IP_HASH"
      name                = "{{random}}"
    }
    check = {
      "nios.lb_preferred_method" = "SOURCE_IP_HASH"
    }
  }

}

case "lb_preferred_method_ratio" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      lb_preferred_method = "ROUND_ROBIN"
      name                = "{{random}}"
    }
    check = {
      "nios.lb_preferred_method" = "ROUND_ROBIN"
    }
  }

  step {
    nios {
      lb_preferred_method = "RATIO"
      name                = "{{random}}"
    }
    check = {
      "nios.lb_preferred_method" = "RATIO"
    }
  }

}

# servers stays hardcoded: a test-created server would also need registering in the
# pre-provisioned topology, and infoblox_dtc_topology is not implemented yet.
case "lb_preferred_topology" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                  = "{{random}}"
      lb_preferred_method   = "TOPOLOGY"
      lb_preferred_topology = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGVycmFmb3JtX3RvcG9sb2d5X3Rlc3Q:terraform_topology_test"
      servers             = [{ server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Rfc2VydmVyLmNvbQ:test_server.com", ratio = 1 }, { server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Rfc2VydmVyMi5jb20:test_server2.com", ratio = 2 }]
    }
    check = {
      "nios.lb_preferred_method"   = "TOPOLOGY"
      "nios.lb_preferred_topology" = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGVycmFmb3JtX3RvcG9sb2d5X3Rlc3Q:terraform_topology_test"
    }
  }

  step {
    nios {
      name                  = "{{random}}"
      lb_preferred_method   = "TOPOLOGY"
      lb_preferred_topology = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGVycmFmb3JtX3RvcG9sb2d5X3Rlc3Qy:terraform_topology_test2"
      servers             = [{ server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Rfc2VydmVyLmNvbQ:test_server.com", ratio = 1 }, { server = "dtc:server/ZG5zLmlkbnNfc2VydmVyJHRlc3Rfc2VydmVyMi5jb20:test_server2.com", ratio = 2 }]
    }
    check = {
      "nios.lb_preferred_method"   = "TOPOLOGY"
      "nios.lb_preferred_topology" = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGVycmFmb3JtX3RvcG9sb2d5X3Rlc3Qy:terraform_topology_test2"
    }
  }

}

case "monitors" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      monitors = ["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http","dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"]
    }
    check = {
      "nios.monitors.#" = "2"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      monitors = ["dtc:monitor:pdp/ZG5zLmlkbnNfbW9uaXRvcl9wZHAkcGRw:pdp"]
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name                = "{{random2}}"
      lb_preferred_method = "ROUND_ROBIN"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "quorum" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      availability        = "QUORUM"
      quorum              = 1
      monitors            = ["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"]
    }
    check = {
      "nios.availability" = "QUORUM"
      "nios.quorum"       = "1"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      availability        = "QUORUM"
      quorum              = 2
      monitors            = ["dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http", "dtc:monitor:snmp/ZG5zLmlkbnNfbW9uaXRvcl9zbm1wJHNubXA:snmp"]
    }
    check = {
      "nios.availability" = "QUORUM"
      "nios.quorum"       = "2"
    }
  }

}

# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: config helper 'testAccDtcPoolServers' could not be parsed (no resource block found)
case "servers" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dtc_server" "one" {
    nios = {
      name = "{{random2}}"
      host = "{{random_ip}}"
    }
  }
  resource "infoblox_dtc_server" "two" {
    nios = {
      name = "{{random3}}"
      host = "{{random_ip2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      servers             = [{ server = infoblox_dtc_server.one.id, ratio = 1 }, { server = infoblox_dtc_server.two.id, ratio = 2 }]
    }
    check = {
      "nios.servers.#"       = "2"
      "nios.servers.0.ratio" = "1"
      "nios.servers.1.ratio" = "2"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      servers             = [{ server = infoblox_dtc_server.one.id, ratio = 1 }]
    }
    check = {
      "nios.servers.0.ratio" = "1"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      ttl                 = 24
    }
  }

  step {
    nios {
      name                = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      ttl                 = 48
    }
  }

}
