# Auto-generated resource acceptance-test cases for SharedrecordSrv.
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
      name                = "{{random}}.example.com"
      port                = 10
      priority            = 80
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
    }
    check = {
      "nios.name"                = "{{random}}.example.com"
      "nios.port"                = "10"
      "nios.priority"            = "80"
      "nios.shared_record_group" = "{{random3}}"
      "nios.target"              = "{{random2}}.target.com"
      "nios.weight"              = "10"
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
      name                = "{{random}}.example.com"
      port                = 10
      priority            = 80
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
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
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
      comment             = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
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
      name = "{{random3}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
      disable             = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
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
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
      ext_attrs           = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

  step {
    nios {
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
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
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random3}}.target.com"
      weight              = 10
    }
    check = {
      "nios.name" = "{{random}}.example.com"
    }
  }

  step {
    nios {
      name                = "{{random2}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random3}}.target.com"
      weight              = 10
    }
    check = {
      "nios.name" = "{{random2}}.example.com"
    }
  }

}

case "port" {
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
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
    }
    check = {
      "nios.port" = "80"
    }
  }

  step {
    nios {
      name                = "{{random}}.example.com"
      port                = 443
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
    }
    check = {
      "nios.port" = "443"
    }
  }

}

case "priority" {
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
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
    }
    check = {
      "nios.priority" = "10"
    }
  }

  step {
    nios {
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 20
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
    }
    check = {
      "nios.priority" = "20"
    }
  }

}

case "shared_record_group" {
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
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
    }
    check = {
      "nios.shared_record_group" = "{{random3}}"
    }
  }

}

case "target" {
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
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target1.com"
      weight              = 10
    }
    check = {
      "nios.target" = "{{random2}}.target1.com"
    }
  }

  step {
    nios {
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random3}}.target2.com"
      weight              = 10
    }
    check = {
      "nios.target" = "{{random3}}.target2.com"
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
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
      ttl                 = 3600
    }
    check = {
      "nios.ttl" = "3600"
    }
  }

  step {
    nios {
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
      ttl                 = 7200
    }
    check = {
      "nios.ttl" = "7200"
    }
  }

}

case "weight" {
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
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 10
    }
    check = {
      "nios.weight" = "10"
    }
  }

  step {
    nios {
      name                = "{{random}}.example.com"
      port                = 80
      priority            = 10
      shared_record_group = infoblox_sharedrecordgroup.parent_sharedrecord_group.nios.name
      target              = "{{random2}}.target.com"
      weight              = 20
    }
    check = {
      "nios.weight" = "20"
    }
  }

}
