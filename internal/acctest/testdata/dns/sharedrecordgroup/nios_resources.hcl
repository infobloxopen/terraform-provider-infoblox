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

  step {
    nios {
      name = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random}}"
    }
  }

}
