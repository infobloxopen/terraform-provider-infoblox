# Auto-generated resource acceptance-test cases for Nsgroup.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.name"                   = "{{random}}"
      "nios.comment"                = ""
      "nios.grid_primary.#"         = "1"
      "nios.grid_primary.0.name"    = "{{grid_master_hostname}}"
      "nios.is_grid_default"        = "false"
      "nios.is_multimaster"         = "false"
      "nios.use_external_primary"   = "false"
      "nios.grid_primary.0.stealth" = "false"
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
      name         = "{{random}}"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
      comment      = "This is a test comment"
    }
    check = {
      "nios.comment" = "This is a test comment"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
      comment      = "This is an updated test comment"
    }
    check = {
      "nios.comment" = "This is an updated test comment"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
      ext_attrs    = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
      ext_attrs    = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "external_primaries" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      external_primaries   = [{ name = "external.primary.1", address = "2.3.4.5" }]
      grid_secondaries     = [{ name = "{{grid_master_hostname}}" }]
      use_external_primary = true
    }
    check = {
      "nios.external_primaries.0.name"    = "external.primary.1"
      "nios.external_primaries.0.address" = "2.3.4.5"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      external_primaries   = [{ name = "external.primary.2", address = "20.1.12.23" }]
      grid_secondaries     = [{ name = "{{grid_master_hostname}}" }]
      use_external_primary = true
    }
    check = {
      "nios.external_primaries.0.name"    = "external.primary.2"
      "nios.external_primaries.0.address" = "20.1.12.23"
    }
  }

}

case "external_primaries_tsig" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      external_primaries   = [{ name = "external.primary.1", address = "2.3.4.5", tsig_key_alg = "HMAC-SHA256", tsig_key = "X4oRe92t54I+T98NdQpV2w==", use_tsig_key_name = true, tsig_key_name = "{{random2}}" }]
      grid_secondaries     = [{ name = "{{grid_master_hostname}}" }]
      use_external_primary = true
    }
    check = {
      "nios.external_primaries.0.tsig_key_alg"      = "HMAC-SHA256"
      "nios.external_primaries.0.tsig_key"          = "X4oRe92t54I+T98NdQpV2w=="
      "nios.external_primaries.0.use_tsig_key_name" = "true"
      "nios.external_primaries.0.tsig_key_name"     = "{{random2}}"
    }
  }

}

case "external_secondaries" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      grid_primary         = [{ name = "{{grid_master_hostname}}" }]
      external_secondaries = [{ name = "external.secondary.1", address = "2.3.3.3" }]
    }
    check = {
      "nios.external_secondaries.0.name"    = "external.secondary.1"
      "nios.external_secondaries.0.address" = "2.3.3.3"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      grid_primary         = [{ name = "{{grid_master_hostname}}" }]
      external_secondaries = [{ name = "external.secondary.2", address = "20.3.32.3" }]
    }
    check = {
      "nios.external_secondaries.0.name"    = "external.secondary.2"
      "nios.external_secondaries.0.address" = "20.3.32.3"
    }
  }

}

case "grid_primary" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.grid_primary.0.name" = "{{grid_master_hostname}}"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      grid_primary = [{ name = "{{grid_member_hostname}}" }]
    }
    check = {
      "nios.grid_primary.0.name" = "{{grid_member_hostname}}"
    }
  }

}

case "grid_secondaries" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      grid_secondaries     = [{ name = "{{grid_master_hostname}}" }]
      external_primaries   = [{ name = "external.primaries.example.com", address = "2.3.3.4" }]
      use_external_primary = true
    }
    check = {
      "nios.grid_secondaries.0.name" = "{{grid_master_hostname}}"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      grid_secondaries     = [{ name = "{{grid_member_hostname}}" }]
      external_primaries   = [{ name = "external.primaries.example.com", address = "2.3.3.4" }]
      use_external_primary = true
    }
    check = {
      "nios.grid_secondaries.0.name" = "{{grid_member_hostname}}"
    }
  }

}

case "is_grid_default" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name            = "{{random}}"
      grid_primary    = [{ name = "{{grid_master_hostname}}" }]
      is_grid_default = true
    }
    check = {
      "nios.is_grid_default" = "true"
    }
  }

  step {
    nios {
      name            = "{{random}}"
      grid_primary    = [{ name = "{{grid_master_hostname}}" }]
      is_grid_default = false
    }
    check = {
      "nios.is_grid_default" = "false"
    }
  }

}

case "is_multimaster" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name           = "{{random}}"
      grid_primary   = [{ name = "{{grid_master_hostname}}" }, { name = "{{grid_member_hostname}}" }]
      is_multimaster = true
    }
    check = {
      "nios.is_multimaster" = "true"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name         = "{{random2}}"
      grid_primary = [{ name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "use_external_primary" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      grid_primary         = [{ name = "{{grid_member_hostname}}" }]
      use_external_primary = false
    }
    check = {
      "nios.use_external_primary" = "false"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      grid_secondaries     = [{ name = "{{grid_master_hostname}}" }]
      external_primaries   = [{ name = "external.primary.1", address = "2.3.4.5" }]
      use_external_primary = true
    }
    check = {
      "nios.use_external_primary" = "true"
    }
  }

}
