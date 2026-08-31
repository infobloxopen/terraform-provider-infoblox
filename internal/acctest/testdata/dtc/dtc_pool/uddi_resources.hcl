# Auto-generated resource acceptance-test cases for DtcPool (uddi).
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name   = "{{random}}"
      method = "round_robin"
    }
    check = {
      "uddi.name"   = "{{random}}"
      "uddi.method" = "round_robin"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    uddi {
      name   = "{{random}}"
      method = "round_robin"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      method  = "round_robin"
      comment = "pool testing"
    }
    check = {
      "uddi.comment" = "pool testing"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      method  = "round_robin"
      comment = "updated pool comment"
    }
    check = {
      "uddi.comment" = "updated pool comment"
    }
  }

}

case "disabled" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      method   = "round_robin"
      disabled = false
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      method   = "round_robin"
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

}

case "method" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name   = "{{random}}"
      method = "round_robin"
    }
    check = {
      "uddi.method" = "round_robin"
    }
  }

  step {
    uddi {
      name   = "{{random}}"
      method = "ratio"
    }
    check = {
      "uddi.method" = "ratio"
    }
  }

}

case "pool_availability" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name              = "{{random}}"
      method            = "round_robin"
      pool_availability = "any"
    }
    check = {
      "uddi.pool_availability" = "any"
    }
  }

  step {
    uddi {
      name              = "{{random}}"
      method            = "round_robin"
      pool_availability = "all"
    }
    check = {
      "uddi.pool_availability" = "all"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name   = "{{random}}"
      method = "round_robin"
      tags   = { Site = "{{random2}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random2}}"
    }
  }

  step {
    uddi {
      name   = "{{random}}"
      method = "round_robin"
      tags   = { Site = "{{random3}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random3}}"
    }
  }

}

case "ttl" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name   = "{{random}}"
      method = "round_robin"
      ttl    = 30
    }
    check = {
      "uddi.ttl" = "30"
    }
  }

  step {
    uddi {
      name   = "{{random}}"
      method = "round_robin"
      ttl    = 60
    }
    check = {
      "uddi.ttl" = "60"
    }
  }

}

case "consolidated_health_enabled" {
  backend     = "uddi"
  parallel    = true
  skip        = true
  skip_reason = "consolidated health probing needs health_checks configured, which requires a HealthCheck object with no resource to provision it"

  step {
    uddi {
      name                        = "{{random}}"
      method                      = "round_robin"
      consolidated_health_enabled = true
    }
    check = {
      "uddi.consolidated_health_enabled" = "true"
    }
  }

  step {
    uddi {
      name                        = "{{random}}"
      method                      = "round_robin"
      consolidated_health_enabled = false
    }
    check = {
      "uddi.consolidated_health_enabled" = "false"
    }
  }

}

# pool_servers_quorum only takes effect when pool_availability is quorum, so both are set
# together here.
case "pool_servers_quorum" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                = "{{random}}"
      method              = "round_robin"
      pool_availability   = "quorum"
      pool_servers_quorum = 1
    }
    check = {
      "uddi.pool_availability"   = "quorum"
      "uddi.pool_servers_quorum" = "1"
    }
  }

  step {
    uddi {
      name                = "{{random}}"
      method              = "round_robin"
      pool_availability   = "quorum"
      pool_servers_quorum = 2
    }
    check = {
      "uddi.pool_availability"   = "quorum"
      "uddi.pool_servers_quorum" = "2"
    }
  }

}

case "server_availability" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                = "{{random}}"
      method              = "round_robin"
      server_availability = "all"
    }
    check = {
      "uddi.server_availability" = "all"
    }
  }

  step {
    uddi {
      name                = "{{random}}"
      method              = "round_robin"
      server_availability = "any"
    }
    check = {
      "uddi.server_availability" = "any"
    }
  }

}

# server_health_checks_quorum only takes effect when server_availability is quorum.
case "server_health_checks_quorum" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                        = "{{random}}"
      method                      = "round_robin"
      server_availability         = "quorum"
      server_health_checks_quorum = 1
    }
    check = {
      "uddi.server_availability"         = "quorum"
      "uddi.server_health_checks_quorum" = "1"
    }
  }

  step {
    uddi {
      name                        = "{{random}}"
      method                      = "round_robin"
      server_availability         = "quorum"
      server_health_checks_quorum = 2
    }
    check = {
      "uddi.server_availability"         = "quorum"
      "uddi.server_health_checks_quorum" = "2"
    }
  }

}

case "inheritance_sources" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                = "{{random}}"
      method              = "round_robin"
      inheritance_sources = { ttl = { action = "inherit" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "inherit"
    }
  }

  step {
    uddi {
      name                = "{{random}}"
      method              = "round_robin"
      ttl                 = 45
      inheritance_sources = { ttl = { action = "override" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "override"
      "uddi.ttl"                            = "45"
    }
  }

}

case "servers" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dtc_server" "one" {
    uddi = {
      name    = "{{random2}}"
      address = "{{random_ip}}"
    }
  }
  resource "infoblox_dtc_server" "two" {
    uddi = {
      name    = "{{random3}}"
      address = "{{random_ip2}}"
    }
  }
  PREREQ

  step {
    uddi {
      name    = "{{random}}"
      method  = "ratio"
      servers = [{ server_id = infoblox_dtc_server.one.id, weight = 1 }, { server_id = infoblox_dtc_server.two.id, weight = 2 }]
    }
    check = {
      "uddi.method"           = "ratio"
      "uddi.servers.#"        = "2"
      "uddi.servers.0.name"   = "{{random2}}"
      "uddi.servers.0.weight" = "1"
      "uddi.servers.1.name"   = "{{random3}}"
      "uddi.servers.1.weight" = "2"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      method  = "ratio"
      servers = [{ server_id = infoblox_dtc_server.one.id, weight = 5 }]
    }
    check = {
      "uddi.method"           = "ratio"
      "uddi.servers.#"        = "1"
      "uddi.servers.0.name"   = "{{random2}}"
      "uddi.servers.0.weight" = "5"
    }
  }

}

