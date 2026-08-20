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

# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: config helper 'testAccDtcLbdnAuthZones' could not be parsed (no resource block found)
case "auth_zones" {
  backend     = "nios"
  skip        = true
  skip_reason = "config helper 'testAccDtcLbdnAuthZones' could not be parsed (no resource block found)"
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

# WARNING: the extractor could not auto-record the following line(s) from
# the Go helper. Some fields may not be correctly captured — please verify
# this case manually against the original test before running:
#   %s
case "lb_method" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dtc_pool_unknown" "test_server_for_pool" {
    nios = {
      name = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
    }
  }
  resource "infoblox_dtc_server_unknown" "test_server_for_pool" {
    nios = {
      name = "{{random}}-server"
      host = "2.3.3.4"
    }
  }
  resource "infoblox_dtc_topology_unknown" "test_server_for_pool" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

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
    }
    check = {
      "nios.lb_method" = "TOPOLOGY"
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

# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: config helper 'testAccDtcLbdnPools' could not be parsed (no resource block found)
case "pools" {
  backend     = "nios"
  skip        = true
  skip_reason = "config helper 'testAccDtcLbdnPools' could not be parsed (no resource block found)"
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dtc_pool_unknown" "test_server_for_pool" {
    nios = {
      name = "{{random}}"
      lb_preferred_method = "ROUND_ROBIN"
    }
  }
  resource "infoblox_dtc_server_unknown" "test_server_for_pool" {
    nios = {
      name = "{{random}}-server"
      host = "2.3.3.4"
    }
  }
  resource "infoblox_dtc_topology_unknown" "test_server_for_pool" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "TOPOLOGY"
      topology  = "$${nios_dtc_topology.test_rules_pool.ref}"
    }
  }

  step {
    nios {
      name      = "dtc-lbdn-{{random}}"
      lb_method = "TOPOLOGY"
      topology  = "$${nios_dtc_topology.test_rules_pool.ref}"
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
