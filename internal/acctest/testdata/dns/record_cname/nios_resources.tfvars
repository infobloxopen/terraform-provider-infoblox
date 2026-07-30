# Auto-generated resource acceptance-test cases for RecordCname.
case "basic" {
  backend = "nios"
  parallel = true

  step {
    nios {
      canonical = "{{random}}.example.com"
      name      = "{{random2}}.example.com"
      view      = "default"
    }
    check = {
      "nios.canonical"          = "{{random}}.example.com"
      "nios.name"               = "{{random2}}.example.com"
      "nios.view"               = "default"
      "nios.creator"            = "STATIC"
      "nios.ddns_protected"     = "false"
      "nios.disable"            = "false"
      "nios.forbid_reclamation" = "false"
    }
  }

}

case "disappears" {
  backend = "nios"
  disappears = true
  expect_non_empty_plan = true
  parallel = true

  step {
    nios {
      canonical = "{{random}}.example.com"
      name      = "{{random2}}.example.com"
      view      = "default"
    }
  }

}

case "canonical" {
  backend = "nios"
  parallel = true

  step {
    nios {
      canonical = "{{random}}.example.com"
      name      = "{{random3}}.example.com"
      view      = "default"
    }
    check = {
      "nios.canonical" = "{{random}}.example.com"
    }
  }

  step {
    nios {
      canonical = "{{random2}}.example.com"
      name      = "{{random3}}.example.com"
      view      = "default"
    }
    check = {
      "nios.canonical" = "{{random2}}.example.com"
    }
  }

}

case "comment" {
  backend = "nios"
  parallel = true

  step {
    nios {
      canonical = "{{random}}.example.com"
      name      = "{{random2}}.example.com"
      view      = "default"
      comment   = "This is a new record"
    }
    check = {
      "nios.comment" = "This is a new record"
    }
  }

  step {
    nios {
      canonical = "{{random}}.example.com"
      name      = "{{random2}}.example.com"
      view      = "default"
      comment   = "This is an updated record"
    }
    check = {
      "nios.comment" = "This is an updated record"
    }
  }

}

case "creator" {
  backend = "nios"
  parallel = true

  step {
    nios {
      canonical = "{{random}}.example.com"
      name      = "{{random2}}.example.com"
      view      = "default"
      creator   = "STATIC"
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

  step {
    nios {
      canonical = "{{random}}.example.com"
      name      = "{{random2}}.example.com"
      view      = "default"
      creator   = "DYNAMIC"
    }
    check = {
      "nios.creator" = "DYNAMIC"
    }
  }

}

case "ddns_principal" {
  backend = "nios"
  parallel = true

  step {
    nios {
      canonical      = "{{random}}.example.com"
      name           = "{{random2}}.example.com"
      view           = "default"
      ddns_principal = "DDNS_PRINCIPAL_1"
      creator        = "DYNAMIC"
    }
    check = {
      "nios.ddns_principal" = "DDNS_PRINCIPAL_1"
    }
  }

  step {
    nios {
      canonical      = "{{random}}.example.com"
      name           = "{{random2}}.example.com"
      view           = "default"
      ddns_principal = "DDNS_PRINCIPAL_2"
      creator        = "DYNAMIC"
    }
    check = {
      "nios.ddns_principal" = "DDNS_PRINCIPAL_2"
    }
  }

}

case "ddns_protected" {
  backend = "nios"
  parallel = true

  step {
    nios {
      canonical      = "{{random}}.example.com"
      name           = "{{random2}}.example.com"
      view           = "default"
      ddns_protected = false
    }
    check = {
      "nios.ddns_protected" = "false"
    }
  }

  step {
    nios {
      canonical      = "{{random}}.example.com"
      name           = "{{random2}}.example.com"
      view           = "default"
      ddns_protected = true
    }
    check = {
      "nios.ddns_protected" = "true"
    }
  }

}

case "disable" {
  backend = "nios"
  parallel = true

  step {
    nios {
      canonical = "{{random}}.example.com"
      name      = "{{random2}}.example.com"
      view      = "default"
      disable   = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      canonical = "{{random}}.example.com"
      name      = "{{random2}}.example.com"
      view      = "default"
      disable   = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "extattrs" {
  backend = "nios"
  parallel = true

  step {
    nios {
      canonical = "{{random}}.example.com"
      name      = "{{random2}}.example.com"
      view      = "default"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      canonical = "{{random}}.example.com"
      name      = "{{random2}}.example.com"
      view      = "default"
      ext_attrs = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

}

case "forbid_reclamation" {
  backend = "nios"
  parallel = true

  step {
    nios {
      canonical          = "{{random}}.example.com"
      name               = "{{random2}}.example.com"
      view               = "default"
      forbid_reclamation = true
    }
    check = {
      "nios.forbid_reclamation" = "true"
    }
  }

  step {
    nios {
      canonical          = "{{random}}.example.com"
      name               = "{{random2}}.example.com"
      view               = "default"
      forbid_reclamation = false
    }
    check = {
      "nios.forbid_reclamation" = "false"
    }
  }

}

case "name" {
  backend = "nios"
  parallel = true

  step {
    nios {
      canonical = "{{random3}}.example.com"
      name      = "{{random}}.example.com"
      view      = "default"
    }
    check = {
      "nios.name" = "{{random}}.example.com"
    }
  }

  step {
    nios {
      canonical = "{{random3}}.example.com"
      name      = "{{random2}}.example.com"
      view      = "default"
    }
    check = {
      "nios.name" = "{{random2}}.example.com"
    }
  }

}

case "ttl" {
  backend = "nios"
  parallel = true

  step {
    nios {
      canonical = "{{random}}.example.com"
      name      = "{{random2}}.example.com"
      view      = "default"
      ttl       = 1000
    }
    check = {
      "nios.ttl" = "1000"
    }
  }

  step {
    nios {
      canonical = "{{random}}.example.com"
      name      = "{{random2}}.example.com"
      view      = "default"
      ttl       = 3200
    }
    check = {
      "nios.ttl" = "3200"
    }
  }

}
