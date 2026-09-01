case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      delegate_to = []
    }
    check = {
      "nios.name"                  = "{{random}}"
      "nios.delegate_to.0.address" = "2.3.4.5"
      "nios.delegate_to.0.name"    = "delegate_to_ns_group"
      "nios.comment"               = ""
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
      name        = "{{random}}"
      delegate_to = []
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      delegate_to = []
      comment     = "comment ns group"
    }
    check = {
      "nios.comment" = "comment ns group"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      delegate_to = []
      comment     = "comment ns group updated"
    }
    check = {
      "nios.comment" = "comment ns group updated"
    }
  }

}

case "delegate_to" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      delegate_to = []
    }
    check = {
      "nios.delegate_to.0.address" = "2.3.4.5"
      "nios.delegate_to.0.name"    = "delegate_to_ns_group"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      delegate_to = []
    }
    check = {
      "nios.delegate_to.0.address" = "2.3.4.6"
      "nios.delegate_to.0.name"    = "delegate_to_ns_group_update"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      delegate_to = []
      ext_attrs   = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      delegate_to = []
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
      name        = "{{random}}"
      delegate_to = []
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name        = "{{random2}}"
      delegate_to = []
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}
