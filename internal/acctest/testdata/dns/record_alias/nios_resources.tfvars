# Auto-generated resource acceptance-test cases for RecordAlias.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
    }
    check = {
      "nios.name"        = "{{random}}.example.com"
      "nios.target_name" = "server.example.com"
      "nios.target_type" = "A"
      "nios.view"        = "default"
      "nios.creator"     = "STATIC"
      "nios.disable"     = "false"
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
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
      comment     = "This is a sample comment."
    }
    check = {
      "nios.comment" = "This is a sample comment."
    }
  }

  step {
    nios {
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
      comment     = "This is an updated comment."
    }
    check = {
      "nios.comment" = "This is an updated comment."
    }
  }

}

case "creator" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
      creator     = "STATIC"
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
      disable     = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
      disable     = true
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
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
      ext_attrs   = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
      ext_attrs   = { Site = "{{random3}}" }
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
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
    }
    check = {
      "nios.name" = "{{random}}.example.com"
    }
  }

  step {
    nios {
      name        = "{{random2}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
    }
    check = {
      "nios.name" = "{{random2}}.example.com"
    }
  }

}

case "target_name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
    }
    check = {
      "nios.target_name" = "server.example.com"
    }
  }

  step {
    nios {
      name        = "{{random}}.example.com"
      target_name = "updated-server.example.com"
      target_type = "A"
      view        = "default"
    }
    check = {
      "nios.target_name" = "updated-server.example.com"
    }
  }

}

case "target_type" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
    }
    check = {
      "nios.target_type" = "A"
    }
  }

  step {
    nios {
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "AAAA"
      view        = "default"
    }
    check = {
      "nios.target_type" = "AAAA"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
      ttl         = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name        = "{{random}}.example.com"
      target_name = "server.example.com"
      target_type = "A"
      view        = "default"
      ttl         = 0
    }
    check = {
      "nios.ttl" = "0"
    }
  }

}

case "view" {
  backend  = "nios"
  parallel = true

  step {
    prerequisites_hcl = <<-PREREQ
    resource "infoblox_view" "test" {
      nios = {
        name = "{{random3}}"
      }
    }
    resource "infoblox_zone_auth" "test" {
      nios = {
        fqdn = "{{random}}.com"
        view = infoblox_view.test.nios.name
      }
    }
    PREREQ
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = "{{random3}}"
    }
    check = {
      "nios.view" = "{{random3}}"
    }
  }

  step {
    prerequisites_hcl = <<-PREREQ
    resource "infoblox_view" "test" {
      nios = {
        name = "{{random4}}"
      }
    }
    resource "infoblox_zone_auth" "test" {
      nios = {
        fqdn = "{{random}}.com"
        view = infoblox_view.test.nios.name
      }
    }
    PREREQ
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = "{{random4}}"
    }
    check = {
      "nios.view" = "{{random4}}"
    }
  }

}
