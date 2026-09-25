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
  resource "infoblox_dtc_server" "test_server2" {
    nios = {
      name = "{{random3}}"
      host = "3.3.3.3"
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
          { source_op = "IS", source_type = "COUNTRY",   source_value = "United States" },
        ]
      }]
    }
    check = {
      "nios.rules.0.dest_type"              = "SERVER"
      "nios.rules.0.sources.#"              = "3"
      "nios.rules.0.sources.0.source_value" = "10.0.0.0/8"
      "nios.rules.0.sources.1.source_value" = "Africa"
      "nios.rules.0.sources.2.source_value" = "United States"
    }
  }

  step {
    nios {
      name = "{{random}}"
      rules = [{
        dest_type        = "SERVER"
        destination_link = infoblox_dtc_server.test_server.id
        sources = [
          { source_op = "IS", source_type = "CONTINENT", source_value = "Africa" },
          { source_op = "IS", source_type = "COUNTRY",   source_value = "United States" },
          { source_op = "IS", source_type = "SUBNET",    source_value = "10.0.0.0/8" },
        ]
      }]
    }
    check = {
      "nios.rules.0.dest_type"              = "SERVER"
      "nios.rules.0.sources.#"              = "3"
      "nios.rules.0.sources.0.source_value" = "Africa"
      "nios.rules.0.sources.1.source_value" = "United States"
      "nios.rules.0.sources.2.source_value" = "10.0.0.0/8"
    }
  }

  step {
    nios {
      name = "{{random}}"
      rules = [
        {
          dest_type        = "SERVER"
          destination_link = infoblox_dtc_server.test_server.id
          sources          = [{ source_op = "IS", source_type = "SUBNET", source_value = "10.0.0.0/8" }]
        },
        {
          dest_type        = "SERVER"
          destination_link = infoblox_dtc_server.test_server2.id
          sources          = [{ source_op = "IS", source_type = "SUBNET", source_value = "192.168.0.0/16" }]
        },
      ]
    }
    check = {
      "nios.rules.#"           = "2"
      "nios.rules.0.dest_type" = "SERVER"
      "nios.rules.1.dest_type" = "SERVER"
    }
  }

  step {
    nios {
      name = "{{random}}"
      rules = [{
        dest_type        = "SERVER"
        destination_link = infoblox_dtc_server.test_server.id
        sources = [
          { source_op = "IS",     source_type = "CONTINENT", source_value = "Europe" },
          { source_op = "IS_NOT", source_type = "COUNTRY",   source_value = "Russia" },
          { source_op = "IS",     source_type = "SUBNET",    source_value = "10.0.0.0/8" },
        ]
      }]
    }
    check = {
      "nios.rules.0.sources.#"             = "3"
      "nios.rules.0.sources.0.source_op"   = "IS"
      "nios.rules.0.sources.0.source_type" = "CONTINENT"
      "nios.rules.0.sources.1.source_op"   = "IS_NOT"
      "nios.rules.0.sources.1.source_type" = "COUNTRY"
      "nios.rules.0.sources.2.source_op"   = "IS"
      "nios.rules.0.sources.2.source_type" = "SUBNET"
    }
  }

  step {
    nios {
      name = "{{random}}"
      rules = [{
        dest_type        = "SERVER"
        destination_link = infoblox_dtc_server.test_server.id
        sources = [
          { source_op = "IS", source_type = "COUNTRY",   source_value = "United States" },
          { source_op = "IS", source_type = "CONTINENT", source_value = "North America" },
          { source_op = "IS", source_type = "CITY",      source_value = "New York" },
        ]
      }]
    }
    check = {
      "nios.rules.0.sources.0.source_value" = "United States"
      "nios.rules.0.sources.1.source_value" = "North America"
      "nios.rules.0.sources.2.source_value" = "New York"
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
  resource "infoblox_dtc_server" "test_server2_for_pool" {
    nios = {
      name = "{{random3}}-server2"
      host = "3.4.4.5"
    }
  }
  resource "infoblox_dtc_pool" "test_pool2" {
    nios = {
      name = "{{random3}}-pool2"
      lb_preferred_method = "ROUND_ROBIN"
      servers = [{ server = infoblox_dtc_server.test_server2_for_pool.id, ratio = 1 }]
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
      rules = [
        {
          dest_type        = "POOL"
          destination_link = infoblox_dtc_pool.test_pool.id
          sources          = [{ source_op = "IS", source_type = "SUBNET", source_value = "10.0.0.0/8" }]
        },
        {
          dest_type        = "POOL"
          destination_link = infoblox_dtc_pool.test_pool2.id
          sources          = [{ source_op = "IS", source_type = "SUBNET", source_value = "192.168.0.0/16" }]
        },
      ]
    }
    check = {
      "nios.rules.#"           = "2"
      "nios.rules.0.dest_type" = "POOL"
      "nios.rules.1.dest_type" = "POOL"
    }
  }

}

case "rules_return_type" {
  backend     = "nios"
  skip        = false
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dtc_server" "test_server_rt" {
    nios = {
      name = "{{random2}}-server-rt"
      host = "4.4.4.4"
    }
  }
  resource "infoblox_dtc_pool" "test_pool_rt" {
    nios = {
      name = "{{random2}}-pool-rt"
      lb_preferred_method = "ROUND_ROBIN"
      servers = [{ server = infoblox_dtc_server.test_server_rt.id, ratio = 1 }]
    }
  }
  PREREQ

  step {
    nios {
      name = "{{random}}"
      rules = [{
        dest_type        = "POOL"
        destination_link = infoblox_dtc_pool.test_pool_rt.id
        return_type      = "REGULAR"
        sources          = [{ source_op = "IS", source_type = "SUBNET", source_value = "10.0.0.0/8" }]
      }]
    }
    check = {
      "nios.rules.#"                        = "1"
      "nios.rules.0.return_type"            = "REGULAR"
      "nios.rules.0.sources.0.source_value" = "10.0.0.0/8"
    }
  }

  step {
    nios {
      name = "{{random}}"
      rules = [
        {
          dest_type        = "POOL"
          destination_link = infoblox_dtc_pool.test_pool_rt.id
          return_type      = "REGULAR"
          sources          = [{ source_op = "IS", source_type = "SUBNET", source_value = "10.0.0.0/8" }]
        },
        {
          dest_type   = "POOL"
          return_type = "NOERR"
          sources     = [{ source_op = "IS", source_type = "SUBNET", source_value = "192.168.0.0/16" }]
        },
      ]
    }
    check = {
      "nios.rules.#"                         = "2"
      "nios.rules.0.return_type"             = "REGULAR"
      "nios.rules.0.sources.#"              = "1"
      "nios.rules.0.sources.0.source_value" = "10.0.0.0/8"
      "nios.rules.1.return_type"             = "NOERR"
      "nios.rules.1.sources.#"              = "1"
      "nios.rules.1.sources.0.source_value" = "192.168.0.0/16"
    }
  }

  step {
    nios {
      name = "{{random}}"
      rules = [
        {
          dest_type        = "POOL"
          destination_link = infoblox_dtc_pool.test_pool_rt.id
          return_type      = "REGULAR"
          sources          = [{ source_op = "IS", source_type = "SUBNET", source_value = "10.0.0.0/8" }]
        },
        {
          dest_type   = "POOL"
          return_type = "NXDOMAIN"
          sources     = [{ source_op = "IS", source_type = "SUBNET", source_value = "192.168.0.0/16" }]
        },
      ]
    }
    check = {
      "nios.rules.#"                         = "2"
      "nios.rules.0.return_type"             = "REGULAR"
      "nios.rules.0.sources.0.source_value"  = "10.0.0.0/8"
      "nios.rules.1.return_type"             = "NXDOMAIN"
      "nios.rules.1.sources.#"              = "1"
      "nios.rules.1.sources.0.source_value"  = "192.168.0.0/16"
    }
  }

}

