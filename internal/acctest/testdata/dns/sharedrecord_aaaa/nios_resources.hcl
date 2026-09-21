# Auto-generated resource acceptance-test cases for SharedrecordAaaa.
case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
    }
    check = {
      "nios.name"                = "{{random}}"
      "nios.ipv6addr"            = "2001:db8::1"
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
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      comment             = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
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
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      disable             = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
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
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random4}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}.example.com"
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      ext_attrs           = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name                = "{{random}}.example.com"
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      ext_attrs           = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "ipv6addr" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
    }
    check = {
      "nios.ipv6addr" = "2001:db8::1"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ipv6addr            = "2001:db8::2"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
    }
    check = {
      "nios.ipv6addr" = "2001:db8::2"
    }
  }

}

case "name" {
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
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name                = "{{random2}}"
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
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
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
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
  resource "infoblox_sharedrecordgroup" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      ttl                 = 3600
    }
    check = {
      "nios.ttl" = "3600"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ipv6addr            = "2001:db8::1"
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      ttl                 = 7200
    }
    check = {
      "nios.ttl" = "7200"
    }
  }

}
