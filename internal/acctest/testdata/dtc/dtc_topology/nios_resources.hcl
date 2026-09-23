# Auto-generated resource acceptance-test cases for DtcTopology.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name"    = "{{random}}"
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
      name = "{{random}}"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "This is an updated comment"
    }
    check = {
      "nios.comment" = "This is an updated comment"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "rules" {
  backend     = "nios"
  skip        = false
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dtc_server" "test_server" {
    nios = {
      name = "{{random2}}"
      host = "2.2.2.2"
    }
  }
  PREREQ

  step {
    nios {
      name  = "{{random}}"
      rules = [{ dest_type = "SERVER", destination_link = infoblox_dtc_server.test_server.id }]
    }
    check = {
      "nios.rules.0.dest_type" = "SERVER"
    }
  }

  step {
    nios {
      name = "{{random}}"
      rules = [{
        dest_type        = "SERVER"
        destination_link = infoblox_dtc_server.test_server.id
        sources = [
          { source_op = "IS", source_type = "SUBNET",    source_value = "10.0.0.0/8" },
          { source_op = "IS", source_type = "CONTINENT", source_value = "Africa" },
        ]
      }]
    }
    check = {
      "nios.rules.0.dest_type"        = "SERVER"
      "nios.rules.0.sources.#"        = "2"
      "nios.rules.0.sources.0.source_value" = "10.0.0.0/8"
      "nios.rules.0.sources.1.source_value" = "Africa"
    }
  }

  # Reorder sources — no diff expected if sources are order-independent.
  step {
    nios {
      name = "{{random}}"
      rules = [{
        dest_type        = "SERVER"
        destination_link = infoblox_dtc_server.test_server.id
        sources = [
          { source_op = "IS", source_type = "CONTINENT", source_value = "Africa" },
          { source_op = "IS", source_type = "SUBNET",    source_value = "10.0.0.0/8" },
        ]
      }]
    }
    check = {
      "nios.rules.0.sources.#" = "2"
    }
  }

}

case "rules_with_pool" {
  backend     = "nios"
  skip        = false
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dtc_server" "test_server_for_pool" {
    nios = {
      name = "{{random2}}-server"
      host = "2.3.3.4"
    }
  }
  resource "infoblox_dtc_pool" "test_pool" {
    nios = {
      name = "{{random2}}"
      lb_preferred_method = "ROUND_ROBIN"
      servers = [{ server = infoblox_dtc_server.test_server_for_pool.id, ratio = 1 }]
    }
  }
  PREREQ

  step {
    nios {
      name  = "{{random}}"
      rules = [{ dest_type = "POOL", destination_link = infoblox_dtc_pool.test_pool.id }]
    }
    check = {
      "nios.rules.0.dest_type" = "POOL"
    }
  }

  step {
    nios {
      name  = "{{random}}"
      rules = [{ dest_type = "POOL", destination_link = infoblox_dtc_pool.test_pool.id }]
    }
    check = {
      "nios.rules.0.dest_type" = "POOL"
    }
  }

}
