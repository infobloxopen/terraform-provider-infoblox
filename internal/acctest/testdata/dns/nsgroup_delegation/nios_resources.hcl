case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      delegate_to = [
        {
          address = "2.3.3.4"
          name    = "delegate_to_ns_group"
        }
      ]
    }
    check = {
      "nios.name"                  = "{{random}}"
      "nios.delegate_to.0.address" = "2.3.3.4"
      "nios.delegate_to.0.name"    = "delegate_to_ns_group"
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
      delegate_to = [
        {
          address = "2.3.3.4"
          name    = "delegate_to_ns_group"
        }
      ]
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      delegate_to = [
        {
          address = "2.3.3.4"
          name    = "delegate_to_ns_group"
        }
      ]
      comment = "comment ns group"
    }
    check = {
      "nios.comment" = "comment ns group"
    }
  }

  step {
    nios {
      name = "{{random}}"
      delegate_to = [
        {
          address = "2.3.3.4"
          name    = "delegate_to_ns_group"
        }
      ]
      comment = "comment ns group updated"
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
      name = "{{random}}"
      delegate_to = [
        {
          address = "2.3.3.4"
          name    = "delegate_to_ns_group"
        }
      ]
    }
    check = {
      "nios.delegate_to.0.address" = "2.3.3.4"
      "nios.delegate_to.0.name"    = "delegate_to_ns_group"
    }
  }

  step {
    nios {
      name = "{{random}}"
      delegate_to = [
        {
          address = "2.3.4.6"
          name    = "delegate_to_ns_group_update"
        }
      ]
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
      name = "{{random}}"
      delegate_to = [
        {
          address = "2.3.3.4"
          name    = "delegate_to_ns_group"
        }
      ]
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name = "{{random}}"
      delegate_to = [
        {
          address = "2.3.3.4"
          name    = "delegate_to_ns_group"
        }
      ]
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
      delegate_to = [
        {
          address = "2.3.3.4"
          name    = "delegate_to_ns_group"
        }
      ]
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
      delegate_to = [
        {
          address = "2.3.3.4"
          name    = "delegate_to_ns_group"
        }
      ]
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}
