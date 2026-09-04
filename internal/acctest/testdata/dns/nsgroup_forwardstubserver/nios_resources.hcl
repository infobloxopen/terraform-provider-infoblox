# Auto-generated resource acceptance-test cases for NsgroupForwardstubserver.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name             = "{{random}}"
      external_servers = [{ name = "example.com", address = "2.3.3.4" }]
    }
    check = {
      "nios.name"                       = "{{random}}"
      "nios.external_servers.#"         = "1"
      "nios.external_servers.0.name"    = "example.com"
      "nios.external_servers.0.address" = "2.3.3.4"
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
      name             = "{{random}}"
      external_servers = [{ name = "example.com", address = "2.3.3.4" }]
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name             = "{{random}}"
      external_servers = [{ name = "example.com", address = "2.3.3.4" }]
      comment          = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name             = "{{random}}"
      external_servers = [{ name = "example.com", address = "2.3.3.4" }]
      comment          = "This is an updated comment"
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
      name             = "{{random}}"
      external_servers = [{ name = "example.com", address = "2.3.3.4" }]
      ext_attrs        = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name             = "{{random}}"
      external_servers = [{ name = "example.com", address = "2.3.3.4" }]
      ext_attrs        = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "external_servers" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name             = "{{random}}"
      external_servers = [{ name = "example.com", address = "2.3.3.4" }]
    }
    check = {
      "nios.external_servers.0.name"    = "example.com"
      "nios.external_servers.0.address" = "2.3.3.4"
    }
  }

  step {
    nios {
      name             = "{{random}}"
      external_servers = [{ name = "example1.com", address = "2.3.4.4" }]
    }
    check = {
      "nios.external_servers.0.name"    = "example1.com"
      "nios.external_servers.0.address" = "2.3.4.4"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name             = "{{random}}"
      external_servers = [{ name = "example.com", address = "2.3.3.4" }]
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name             = "{{random2}}"
      external_servers = [{ name = "example.com", address = "2.3.3.4" }]
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}
