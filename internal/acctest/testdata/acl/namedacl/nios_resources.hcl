# Auto-generated resource acceptance-test cases for Namedacl.
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

case "access_list" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      access_list = [{ struct = "addressac", address = "10.0.0.5", permission = "DENY" }, { struct = "addressac", address = "10.0.2.1/32", permission = "DENY" }]
    }
    check = {
      "nios.access_list.#"            = "2"
      "nios.access_list.0.address"    = "10.0.0.5"
      "nios.access_list.1.address"    = "10.0.2.1/32"
      "nios.access_list.0.permission" = "DENY"
      "nios.access_list.1.permission" = "DENY"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      access_list = [{ struct = "tsigac", tsig_key = "X4oRe92t54I+T98NdQpV2w==", tsig_key_name = "example-tsig-key", tsig_key_alg = "HMAC-SHA256" }]
    }
    check = {
      "nios.access_list.#"               = "1"
      "nios.access_list.0.tsig_key"      = "X4oRe92t54I+T98NdQpV2w=="
      "nios.access_list.0.tsig_key_name" = "example-tsig-key"
      "nios.access_list.0.tsig_key_alg"  = "HMAC-SHA256"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a new named acl"
    }
    check = {
      "nios.comment" = "This is a new named acl"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a updated named acl"
    }
    check = {
      "nios.comment" = "This is a updated named acl"
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
