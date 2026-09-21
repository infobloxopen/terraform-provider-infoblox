# Auto-generated resource acceptance-test cases for Vlan.
case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlanview" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      id     = 51
      name   = "{{random}}"
      parent = infoblox_vlanview.test.id
    }
    check = {
      "nios.id"       = "51"
      "nios.name"     = "{{random}}"
      "nios.reserved" = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlanview" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      id     = 52
      name   = "{{random}}"
      parent = infoblox_vlanview.test.id
    }
  }

}

case "import" {
  backend  = "nios"
  parallel = true
  import   = true
  import_ignore = ["nios.ext_attrs_all"]
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlanview" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      id     = 53
      name   = "{{random}}"
      parent = infoblox_vlanview.test.id
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlanview" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      id      = 54
      name    = "{{random}}"
      parent  = infoblox_vlanview.test.id
      comment = "Comment for the object"
    }
    check = {
      "nios.comment" = "Comment for the object"
    }
  }

  step {
    nios {
      id      = 54
      name    = "{{random}}"
      parent  = infoblox_vlanview.test.id
      comment = "Updated comment for the object"
    }
    check = {
      "nios.comment" = "Updated comment for the object"
    }
  }

}

case "contact" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlanview" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      id      = 55
      name    = "{{random}}"
      parent  = infoblox_vlanview.test.id
      contact = "contact_FIRST"
    }
    check = {
      "nios.contact" = "contact_FIRST"
    }
  }

  step {
    nios {
      id      = 55
      name    = "{{random}}"
      parent  = infoblox_vlanview.test.id
      contact = "CONTACT_2"
    }
    check = {
      "nios.contact" = "CONTACT_2"
    }
  }

}

case "department" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlanview" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      id         = 56
      name       = "{{random}}"
      parent     = infoblox_vlanview.test.id
      department = "DEPARTMENT"
    }
    check = {
      "nios.department" = "DEPARTMENT"
    }
  }

  step {
    nios {
      id         = 56
      name       = "{{random}}"
      parent     = infoblox_vlanview.test.id
      department = "department_UPDATE"
    }
    check = {
      "nios.department" = "department_UPDATE"
    }
  }

}

case "description" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlanview" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      id          = 57
      name        = "{{random}}"
      parent      = infoblox_vlanview.test.id
      description = "description_INITIAL"
    }
    check = {
      "nios.description" = "description_INITIAL"
    }
  }

  step {
    nios {
      id          = 57
      name        = "{{random}}"
      parent      = infoblox_vlanview.test.id
      description = "DESCRIPTION_UPDATE"
    }
    check = {
      "nios.description" = "DESCRIPTION_UPDATE"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlanview" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      id        = 58
      name      = "{{random}}"
      parent    = infoblox_vlanview.test.id
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      id        = 58
      name      = "{{random}}"
      parent    = infoblox_vlanview.test.id
      ext_attrs = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

}

case "id" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlanview" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      id     = 51
      name   = "{{random}}"
      parent = infoblox_vlanview.test.id
    }
    check = {
      "nios.id" = "51"
    }
  }

  step {
    nios {
      id     = 59
      name   = "{{random3}}"
      parent = infoblox_vlanview.test.id
    }
    check = {
      "nios.id" = "59"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlanview" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      id     = 60
      name   = "{{random}}"
      parent = infoblox_vlanview.test.id
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      id     = 60
      name   = "{{random3}}"
      parent = infoblox_vlanview.test.id
    }
    check = {
      "nios.name" = "{{random3}}"
    }
  }

}

case "parent" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlanview" "one" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  resource "infoblox_vlanview" "two" {
    nios = {
      name          = "{{random3}}"
      start_vlan_id = 51
      end_vlan_id   = 101
    }
  }
  PREREQ

  step {
    nios {
      id     = 61
      name   = "{{random}}"
      parent = infoblox_vlanview.one.id
    }
    depends_on = [infoblox_vlanview.one, infoblox_vlanview.two]
  }

  step {
    nios {
      id     = 61
      name   = "{{random}}"
      parent = infoblox_vlanview.two.id
    }
    depends_on = [infoblox_vlanview.one, infoblox_vlanview.two]
  }

}

case "reserved" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlanview" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      id       = 62
      name     = "{{random}}"
      parent   = infoblox_vlanview.test.id
      reserved = true
    }
    check = {
      "nios.reserved" = "true"
    }
  }

  step {
    nios {
      id       = 62
      name     = "{{random}}"
      parent   = infoblox_vlanview.test.id
      reserved = false
    }
    check = {
      "nios.reserved" = "false"
    }
  }

}
