# Auto-generated resource acceptance-test cases for DtcLbdn.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "ROUND_ROBIN"
    }
    check = {
      "nios.lb_method"                  = "ROUND_ROBIN"
      "nios.name"                       = "dtc-lbdn-{{random}}"
      "nios.auto_consolidated_monitors" = "false"
      "nios.disable"                    = "false"
      "nios.persistence"                = "0"
      "nios.priority"                   = "1"
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
      name      = "dtc-lbdn-{{random}}"
      lb_method = "ROUND_ROBIN"
    }
  }

}

case "auth_zones" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dtc_server" "test_server1" {
    nios = {
      name = "server-{{random}}"
      host = "{{random_ip}}"
    }
  }
  resource "infoblox_zone_auth" "test_zone1" {
    nios = {
      fqdn         = "{{random}}.test.com"
      view         = "default"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
    }
  }
  resource "infoblox_zone_auth" "test_zone2" {
    nios = {
      fqdn         = "{{random2}}.record_test.com"
      view         = "default"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
    }
  }
  resource "infoblox_zone_auth" "test_zone3" {
    nios = {
      fqdn         = "{{random3}}.test.com"
      view         = "default"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
    }
  }
  resource "infoblox_dtc_pool" "test_pool1" {
    nios = {
      name                = "pool-{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      servers             = [{ server = infoblox_dtc_server.test_server1.id, ratio = 1 }]
    }
  }
  PREREQ

  step {
    nios {
      name       = "dtc-lbdn-{{random}}"
      lb_method  = "SOURCE_IP_HASH"
      auth_zones = ["$${infoblox_zone_auth.test_zone1.id}", "$${infoblox_zone_auth.test_zone2.id}"]
      pools      = [{ pool = "$${infoblox_dtc_pool.test_pool1.id}", ratio = 2 }]
      patterns   = ["*.test.com", "*.record_test.com"]
      disable    = true
    }
    check = {
      "nios.auth_zones.#" = "2"
    }
  }

  step {
    nios {
      name       = "dtc-lbdn-{{random}}"
      lb_method  = "SOURCE_IP_HASH"
      auth_zones = ["$${infoblox_zone_auth.test_zone3.id}"]
      pools      = [{ pool = "$${infoblox_dtc_pool.test_pool1.id}", ratio = 2 }]
      patterns   = ["*.test.com", "*.record_test.com"]
      disable    = true
    }
    check = {
      "nios.auth_zones.#" = "1"
    }
  }

}

case "auto_consolidated_monitors" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                       = "dtc-lbdn-{{random}}"
      lb_method                  = "RATIO"
      auto_consolidated_monitors = true
    }
    check = {
      "nios.auto_consolidated_monitors" = "true"
    }
  }

  step {
    nios {
      name                       = "dtc-lbdn-{{random}}"
      lb_method                  = "RATIO"
      auto_consolidated_monitors = false
    }
    check = {
      "nios.auto_consolidated_monitors" = "false"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "GLOBAL_AVAILABILITY"
      comment   = "resource-comment"
    }
    check = {
      "nios.comment" = "resource-comment"
    }
  }

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "GLOBAL_AVAILABILITY"
      comment   = "resource-comment-update"
    }
    check = {
      "nios.comment" = "resource-comment-update"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "RATIO"
      disable   = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "RATIO"
      disable   = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "ROUND_ROBIN"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "ROUND_ROBIN"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "lb_method" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "GLOBAL_AVAILABILITY"
    }
    check = {
      "nios.lb_method" = "GLOBAL_AVAILABILITY"
    }
  }

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "RATIO"
    }
    check = {
      "nios.lb_method" = "RATIO"
    }
  }

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "ROUND_ROBIN"
    }
    check = {
      "nios.lb_method" = "ROUND_ROBIN"
    }
  }

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "SOURCE_IP_HASH"
    }
    check = {
      "nios.lb_method" = "SOURCE_IP_HASH"
    }
  }

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "TOPOLOGY"
      topology  = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGMwMV9zaW5nbGVfcnVsZQ:tc01_single_rule"
    }
    check = {
      "nios.lb_method" = "TOPOLOGY"
      "nios.topology"  = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGMwMV9zaW5nbGVfcnVsZQ:tc01_single_rule"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "SOURCE_IP_HASH"
    }
    check = {
      "nios.name" = "dtc-lbdn-{{random}}"
    }
  }

  step {
    nios {
      name      = "dtc-lbdn-{{random2}}"
      lb_method = "SOURCE_IP_HASH"
    }
    check = {
      "nios.name" = "dtc-lbdn-{{random2}}"
    }
  }

}

case "patterns" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "SOURCE_IP_HASH"
      patterns  = ["*.test.com", "*.info.com"]
    }
    check = {
      "nios.patterns.#" = "2"
      "nios.patterns.0" = "*.test.com"
      "nios.patterns.1" = "*.info.com"
    }
  }

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "SOURCE_IP_HASH"
      patterns  = ["*.test123.com", "*.info*.com"]
    }
    check = {
      "nios.patterns.#" = "2"
      "nios.patterns.0" = "*.test123.com"
      "nios.patterns.1" = "*.info*.com"
    }
  }

}

case "persistence" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "dtc-lbdn-{{random}}"
      lb_method   = "ROUND_ROBIN"
      persistence = 3
    }
    check = {
      "nios.persistence" = "3"
    }
  }

  step {
    nios {
      name        = "dtc-lbdn-{{random}}"
      lb_method   = "ROUND_ROBIN"
      persistence = 8
    }
    check = {
      "nios.persistence" = "8"
    }
  }

}

case "pools" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dtc_server" "test_server1" {
    nios = {
      name = "server-{{random}}"
      host = "{{random_ip}}"
    }
  }
  resource "infoblox_dtc_pool" "test_pool1" {
    nios = {
      name                = "pool1-{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      servers             = [{ server = infoblox_dtc_server.test_server1.id, ratio = 1 }]
    }
  }
  resource "infoblox_dtc_pool" "test_pool2" {
    nios = {
      name                = "pool2-{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
      servers             = [{ server = infoblox_dtc_server.test_server1.id, ratio = 1 }]
    }
  }
  PREREQ

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "ROUND_ROBIN"
      pools     = [{ pool = "$${infoblox_dtc_pool.test_pool1.id}", ratio = 2 }]
    }
    check = {
      "nios.pools.#"       = "1"
      "nios.pools.0.ratio" = "2"
    }
  }

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "ROUND_ROBIN"
      pools     = [{ pool = "$${infoblox_dtc_pool.test_pool2.id}", ratio = 2 }]
    }
    check = {
      "nios.pools.#"       = "1"
      "nios.pools.0.ratio" = "2"
    }
  }

}

case "priority" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "RATIO"
      priority  = 1
    }
    check = {
      "nios.priority" = "1"
    }
  }

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "RATIO"
      priority  = 3
    }
    check = {
      "nios.priority" = "3"
    }
  }

}

case "topology" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "TOPOLOGY"
      topology  = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGMwMV9zaW5nbGVfcnVsZQ:tc01_single_rule"
    }
    check = {
      "nios.topology" = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGMwMV9zaW5nbGVfcnVsZQ:tc01_single_rule"
    }
  }

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "TOPOLOGY"
      topology  = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGMwMl9kZXN0X3Bvb2w:tc02_dest_pool"
    }
    check = {
      "nios.topology" = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdGMwMl9kZXN0X3Bvb2w:tc02_dest_pool"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "GLOBAL_AVAILABILITY"
      ttl       = 260
    }
    check = {
      "nios.ttl" = "260"
    }
  }

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "GLOBAL_AVAILABILITY"
      ttl       = 480
    }
    check = {
      "nios.ttl" = "480"
    }
  }

}

case "types" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "ROUND_ROBIN"
      types     = ["A", "AAAA", "CNAME"]
    }
    check = {
      "nios.types.#" = "3"
    }
  }

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "ROUND_ROBIN"
      types     = ["A", "AAAA", "CNAME", "SRV"]
    }
    check = {
      "nios.types.#" = "4"
    }
  }

}
