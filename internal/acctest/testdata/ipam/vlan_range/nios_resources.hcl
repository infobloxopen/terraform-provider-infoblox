# Auto-generated resource acceptance-test cases for Vlanrange.
#
# Every case creates its own parent VLAN View, because a VLAN Range is required
# to live inside one. The range's start/end must fall within the view's own
# start_vlan_id/end_vlan_id window (50-100 below).

case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
    }
    check = {
      "nios.name"          = "{{random}}"
      "nios.start_vlan_id" = "61"
      "nios.end_vlan_id"   = "71"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
    }
  }

}

case "import" {
  backend       = "nios"
  parallel      = true
  import        = true
  import_ignore = ["nios.ext_attrs_all"]
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
      comment       = "Comment for the object"
    }
    check = {
      "nios.comment" = "Comment for the object"
    }
  }

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
      comment       = "Updated comment for the object"
    }
    check = {
      "nios.comment" = "Updated comment for the object"
    }
  }

}

case "end_vlan_id" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
    }
    check = {
      "nios.end_vlan_id" = "71"
    }
  }

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 81
      vlan_view     = infoblox_vlan_view.test.id
    }
    check = {
      "nios.end_vlan_id" = "81"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
      ext_attrs     = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
      ext_attrs     = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name          = "{{random3}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
    }
    check = {
      "nios.name" = "{{random3}}"
    }
  }

}

case "pre_create_vlan" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  # pre_create_vlan is create-only (ImmutableBool), so there is no update step.
  step {
    nios {
      name            = "{{random}}"
      start_vlan_id   = 61
      end_vlan_id     = 71
      vlan_view       = infoblox_vlan_view.test.id
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
    }
    check = {
      "nios.start_vlan_id" = "61"
    }
  }

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 51
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
    }
    check = {
      "nios.start_vlan_id" = "51"
    }
  }

}

case "vlan_name_prefix" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  # vlan_name_prefix is create-only (ImmutableString) and only takes effect when
  # pre_create_vlan is set, so there is no update step.
  step {
    nios {
      name             = "{{random}}"
      start_vlan_id    = 61
      end_vlan_id      = 71
      vlan_view        = infoblox_vlan_view.test.id
      pre_create_vlan  = true
      vlan_name_prefix = "prefixCaseInsensitive"
    }
    check = {
      "nios.vlan_name_prefix" = "prefixCaseInsensitive"
    }
  }

}

case "vlan_view" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "one" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  resource "infoblox_vlan_view" "two" {
    nios = {
      name          = "{{random3}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.one.id
    }
    check_pair = {
      "nios.vlan_view" = infoblox_vlan_view.one.id
    }
  }

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.two.id
    }
    check_pair = {
      "nios.vlan_view" = infoblox_vlan_view.two.id
    }
  }

}
