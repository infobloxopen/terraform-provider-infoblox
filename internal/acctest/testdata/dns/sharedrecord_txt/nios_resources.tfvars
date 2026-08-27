# Auto-generated resource acceptance-test cases for SharedrecordTxt.
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
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      text                = "This is a shared record TXT record"
    }
    check = {
      "nios.name"    = "{{random}}"
      "nios.comment" = ""
      "nios.disable" = "false"
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
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      text                = "example txt for sharedrecord:txt"
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
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      text                = "This is a shared record TXT record"
      comment             = "Shared TXT Record Comment"
    }
    check = {
      "nios.comment" = "Shared TXT Record Comment"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      text                = "This is a shared record TXT record"
      comment             = "Shared TXT Record Comment Updated"
    }
    check = {
      "nios.comment" = "Shared TXT Record Comment Updated"
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
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      text                = "This is a shared record TXT record"
      disable             = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      text                = "This is a shared record TXT record"
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
  resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
    nios = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      text                = "This is a shared record TXT record"
      ext_attrs           = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      text                = "This is a shared record TXT record"
      ext_attrs           = { Site = "{{random4}}" }
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
  resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
    nios = {
      name = "{{random3}}"
    }
  }
  PREREQ

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      text                = "This is a shared record TXT record"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name                = "{{random2}}"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      text                = "This is a shared record TXT record"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "text" {
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
      name                = "{{random}}"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      text                = "{{random2}}"
    }
    check = {
      "nios.text" = "{{random2}}"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      text                = "{{random3}}"
    }
    check = {
      "nios.text" = "{{random3}}"
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
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      text                = "This is a shared record TXT record"
      ttl                 = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      shared_record_group = infoblox_shared_record_group_unknown.parent_sharedrecord_group.nios.name
      text                = "This is a shared record TXT record"
      ttl                 = 20
    }
    check = {
      "nios.ttl" = "20"
    }
  }

}
