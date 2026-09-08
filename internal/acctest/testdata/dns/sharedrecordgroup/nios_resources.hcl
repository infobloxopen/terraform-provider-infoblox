# Auto-generated resource acceptance-test cases for Sharedrecordgroup.
#
# TODO: The zone_associations case references the auth zones "tf-srg-zone-1.com"
#       and "tf-srg-zone-2.com", which must already exist on the grid. They are
#       not created as Terraform prerequisites on purpose: associating a shared
#       record group mutates the parent zone in NIOS, so a managed zone_auth
#       resource drifts after every apply and the step fails on a non-empty
#       refresh plan.
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

  step {
    nios {
      name = "{{random}}"
      zone_associations = [{
        fqdn = "tf-srg-zone-1.com"
        view = "default"
      }]
    }
    check = {
      "nios.zone_associations.#"      = "1"
      "nios.zone_associations.0.fqdn" = "tf-srg-zone-1.com"
      "nios.zone_associations.0.view" = "default"
    }
  }

  step {
    nios {
      name = "{{random}}"
      zone_associations = [{
        fqdn = "tf-srg-zone-2.com"
        view = "default"
      }]
    }
    check = {
      "nios.zone_associations.#"      = "1"
      "nios.zone_associations.0.fqdn" = "tf-srg-zone-2.com"
      "nios.zone_associations.0.view" = "default"
    }
  }

  // NIOS refuses to delete a shared record group while a zone still references
  // it, so the association must be cleared before the test tears down.
  step {
    nios {
      name = "{{random}}"
    }
  }

}
