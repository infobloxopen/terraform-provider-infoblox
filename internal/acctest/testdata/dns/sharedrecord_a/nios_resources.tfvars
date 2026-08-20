# Auto-generated resource acceptance-test cases for SharedrecordA.
case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
    }
    check = {
      "nios.name"                = "{{random}}"
      "nios.ipv4addr"            = "10.0.0.0"
      "nios.shared_record_group" = "{{random2}}"
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
  resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      comment             = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      comment             = "This is an updated comment"
    }
    check = {
      "nios.comment" = "This is an updated comment"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      disable             = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      disable             = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
    nios = {
      name = "{{random4}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}.example.com"
      ipv4addr            = "10.0.0.1"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      ext_attrs           = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name                = "{{random}}.example.com"
      ipv4addr            = "10.0.0.1"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      ext_attrs           = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "ipv4addr" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
    }
    check = {
      "nios.ipv4addr" = "10.0.0.0"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.1"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
    }
    check = {
      "nios.ipv4addr" = "10.0.0.1"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
    nios = {
      name = "{{random3}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name                = "{{random2}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "shared_record_group" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
    }
    check = {
      "nios.shared_record_group" = "{{random2}}"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      ttl                 = 3600
    }
    check = {
      "nios.ttl" = "3600"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      ttl                 = 7200
    }
    check = {
      "nios.ttl" = "7200"
    }
  }

}
