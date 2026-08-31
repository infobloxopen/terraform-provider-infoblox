case "basic" {
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
      "nios.name"    = "{{random}}"
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "This is an updated comment"
    }
    check = {
      "nios.name"    = "{{random}}"
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
      "nios.name"            = "{{random}}"
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.name"            = "{{random}}"
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
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dtc_server" "test_server" {
    nios = {
      name = "{{random2}}-server"
      host = "2.2.2.2"
    }
  }
  PREREQ

  step {
    depends_on = [infoblox_dtc_server.test_server]
    nios {
      name  = "{{random}}"
      rules = [
        {
          dest_type        = "SERVER"
          destination_link = "$${infoblox_dtc_server.test_server.id}"
        }
      ]
    }
    check = {
      "nios.name"               = "{{random}}"
      "nios.rules.0.dest_type"  = "SERVER"
    }
  }

  step {
    depends_on = [infoblox_dtc_server.test_server]
    nios {
      name  = "{{random}}"
      rules = [
        {
          dest_type        = "SERVER"
          destination_link = "$${infoblox_dtc_server.test_server.id}"
          return_type      = "REGULAR"
          sources = [
            {
              source_type  = "COUNTRY"
              source_op    = "IS"
              source_value = "US"
            }
          ]
        }
      ]
    }
    check = {
      "nios.name"                        = "{{random}}"
      "nios.rules.0.dest_type"           = "SERVER"
      "nios.rules.0.return_type"         = "REGULAR"
      "nios.rules.0.sources.0.source_type"  = "COUNTRY"
      "nios.rules.0.sources.0.source_op"    = "IS"
      "nios.rules.0.sources.0.source_value" = "US"
    }
  }

}

case "rules_with_pool" {
  backend     = "nios"
  skip        = true
  skip_reason = "infoblox_dtc_pool not yet implemented — activate once dtc_pool is generated"
  parallel    = true
}
