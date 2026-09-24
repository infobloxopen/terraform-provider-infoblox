# Auto-generated resource acceptance-test cases for SharedrecordCname.
case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random3}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      canonical           = "{{random2}}.com"
    }
    check = {
      "nios.name"                = "{{random}}"
      "nios.canonical"           = "{{random2}}.com"
      "nios.shared_record_group" = "{{random3}}"
      "nios.disable"             = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random3}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      canonical           = "{{random2}}.com"
    }
  }

}

case "canonical" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random4}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      canonical           = "{{random2}}.com"
    }
    check = {
      "nios.canonical" = "{{random2}}.com"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      canonical           = "{{random3}}.com"
    }
    check = {
      "nios.canonical" = "{{random3}}.com"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random3}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      canonical           = "{{random2}}.com"
      comment             = "Example Shared CNAME Record Comment"
    }
    check = {
      "nios.comment" = "Example Shared CNAME Record Comment"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      canonical           = "{{random2}}.com"
      comment             = "Example Shared CNAME Record Comment Updated"
    }
    check = {
      "nios.comment" = "Example Shared CNAME Record Comment Updated"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random3}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      canonical           = "{{random2}}.com"
      disable             = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      canonical           = "{{random2}}.com"
      disable             = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random3}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      canonical           = "{{random2}}.com"
      ext_attrs           = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      canonical           = "{{random2}}.com"
      ext_attrs           = { Site = "{{random5}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random5}}"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random4}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      canonical           = "{{random3}}.com"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name                = "{{random2}}"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      canonical           = "{{random3}}.com"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random3}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      canonical           = "{{random2}}.com"
      ttl                 = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      canonical           = "{{random2}}.com"
      ttl                 = 20
    }
    check = {
      "nios.ttl" = "20"
    }
  }

}
