# Auto-generated resource acceptance-test cases for Sharedrecordgroup.
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
      comment = "shared record group comment"
    }
    check = {
      "nios.comment" = "shared record group comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "shared record group comment updated"
    }
    check = {
      "nios.comment" = "shared record group comment updated"
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

case "record_name_policy" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name               = "{{random}}"
      record_name_policy = "Allow Underscore"
    }
    check = {
      "nios.record_name_policy" = "Allow Underscore"
    }
  }

  step {
    nios {
      name               = "{{random}}"
      record_name_policy = "Allow Any"
    }
    check = {
      "nios.record_name_policy" = "Allow Any"
    }
  }

}

case "zone_associations" {
  backend  = "nios"
  parallel = true
  // Associating a shared record group mutates the parent zone in NIOS, so the
  // zone_auth prerequisite legitimately drifts after each apply.
  expect_non_empty_plan = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "srg_zone_1" {
    nios = {
      fqdn = "{{random2}}.com"
    }
  }
  resource "infoblox_zone_auth" "srg_zone_2" {
    nios = {
      fqdn = "{{random3}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name = "{{random}}"
      zone_associations = [{
        fqdn = infoblox_zone_auth.srg_zone_1.nios.fqdn
        view = infoblox_zone_auth.srg_zone_1.nios.view
      }]
    }
    check = {
      "nios.zone_associations.#"      = "1"
      "nios.zone_associations.0.fqdn" = "{{random2}}.com"
      "nios.zone_associations.0.view" = "default"
    }
  }

  // Unset the association so the group can be re-associated and, finally, destroyed.
  step {
    nios {
      name = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random}}"
      zone_associations = [{
        fqdn = infoblox_zone_auth.srg_zone_2.nios.fqdn
        view = infoblox_zone_auth.srg_zone_2.nios.view
      }]
    }
    check = {
      "nios.zone_associations.#"      = "1"
      "nios.zone_associations.0.fqdn" = "{{random3}}.com"
      "nios.zone_associations.0.view" = "default"
    }
  }

  step {
    nios {
      name = "{{random}}"
    }
  }

}
