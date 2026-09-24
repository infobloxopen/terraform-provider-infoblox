# Auto-generated resource acceptance-test cases for Vlanview.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      end_vlan_id   = 15
      name          = "{{random}}"
      start_vlan_id = 10
    }
    check = {
      "nios.end_vlan_id"             = "15"
      "nios.name"                    = "{{random}}"
      "nios.start_vlan_id"           = "10"
      "nios.allow_range_overlapping" = "false"
      "nios.pre_create_vlan"         = "false"
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
      end_vlan_id   = 15
      name          = "{{random}}"
      start_vlan_id = 10
    }
  }

}

case "import" {
  backend  = "nios"
  parallel = true
  import   = true
  import_ignore = ["nios.ext_attrs_all"]

  step {
    nios {
      end_vlan_id   = 15
      name          = "{{random}}"
      start_vlan_id = 10
    }
  }

}

case "allow_range_overlapping" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      end_vlan_id             = 15
      name                    = "{{random}}"
      start_vlan_id           = 10
      allow_range_overlapping = false
    }
    check = {
      "nios.allow_range_overlapping" = "false"
    }
  }

  step {
    nios {
      end_vlan_id             = 15
      name                    = "{{random}}"
      start_vlan_id           = 10
      allow_range_overlapping = true
    }
    check = {
      "nios.allow_range_overlapping" = "true"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      end_vlan_id   = 15
      name          = "{{random}}"
      start_vlan_id = 10
      comment       = "Comment for the Vlan view object"
    }
    check = {
      "nios.comment" = "Comment for the Vlan view object"
    }
  }

  step {
    nios {
      end_vlan_id   = 15
      name          = "{{random}}"
      start_vlan_id = 10
      comment       = "Updated comment for the Vlan view object"
    }
    check = {
      "nios.comment" = "Updated comment for the Vlan view object"
    }
  }

}

case "end_vlan_id" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      end_vlan_id   = 4094
      name          = "{{random}}"
      start_vlan_id = 1
    }
    check = {
      "nios.end_vlan_id" = "4094"
    }
  }

  step {
    nios {
      end_vlan_id   = 1
      name          = "{{random}}"
      start_vlan_id = 1
    }
    check = {
      "nios.end_vlan_id" = "1"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      end_vlan_id   = 15
      name          = "{{random}}"
      start_vlan_id = 10
      ext_attrs     = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      end_vlan_id   = 15
      name          = "{{random}}"
      start_vlan_id = 10
      ext_attrs     = { Site = "{{random3}}" }
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
      end_vlan_id   = 15
      name          = "{{random}}"
      start_vlan_id = 10
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      end_vlan_id   = 15
      name          = "{{random2}}"
      start_vlan_id = 10
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "pre_create_vlan" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      end_vlan_id     = 15
      name            = "{{random}}"
      start_vlan_id   = 10
      pre_create_vlan = true
    }
    check = {
      "nios.pre_create_vlan" = "true"
    }
  }

}

case "start_vlan_id" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      end_vlan_id   = 4094
      name          = "{{random}}"
      start_vlan_id = 1
    }
    check = {
      "nios.start_vlan_id" = "1"
    }
  }

  step {
    nios {
      end_vlan_id   = 4094
      name          = "{{random}}"
      start_vlan_id = 4094
    }
    check = {
      "nios.start_vlan_id" = "4094"
    }
  }

}

case "vlan_name_prefix" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      end_vlan_id      = 15
      name             = "{{random}}"
      start_vlan_id    = 10
      vlan_name_prefix = "prefixCaseInsensitive"
      pre_create_vlan  = true
    }
    check = {
      "nios.vlan_name_prefix" = "prefixCaseInsensitive"
    }
  }

}
