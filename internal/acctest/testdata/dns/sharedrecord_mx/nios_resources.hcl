# Auto-generated resource acceptance-test cases for SharedrecordMx.
#
# TODO: These cases use the shared record group "shared_group", which must already
#       exist on the grid. The generated prerequisite is commented out because
#       infoblox_shared_record_group is not implemented in the provider yet.
#       Once it is, restore the prerequisite block and remove this note.
case "basic" {
  backend     = "nios"
  parallel    = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  # nios = {
  # name = "{{random3}}"
  # }
  # }
  # PREREQ

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
    }
    check = {
      "nios.mail_exchanger"      = "{{random2}}.example.com"
      "nios.name"                = "{{random}}.example.com"
      "nios.preference"          = "10"
      "nios.shared_record_group" = "shared_group"
      "nios.disable"             = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  # nios = {
  # name = "{{random3}}"
  # }
  # }
  # PREREQ

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
    }
  }

}

case "comment" {
  backend     = "nios"
  parallel    = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  # nios = {
  # name = "{{random3}}"
  # }
  # }
  # PREREQ

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
      comment             = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
      comment             = "This is an updated comment"
    }
    check = {
      "nios.comment" = "This is an updated comment"
    }
  }

}

case "disable" {
  backend     = "nios"
  parallel    = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  # nios = {
  # name = "{{random3}}"
  # }
  # }
  # PREREQ

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
      disable             = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
      disable             = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "ext_attrs" {
  backend     = "nios"
  parallel    = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  # nios = {
  # name = "{{random5}}"
  # }
  # }
  # PREREQ

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
      ext_attrs           = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
      ext_attrs           = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

}

case "mail_exchanger" {
  backend     = "nios"
  parallel    = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  # nios = {
  # name = "{{random4}}"
  # }
  # }
  # PREREQ

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
    }
    check = {
      "nios.mail_exchanger" = "{{random2}}.example.com"
    }
  }

  step {
    nios {
      mail_exchanger      = "{{random3}}.example.com"
      name                = "example.com"
      preference          = 10
      shared_record_group = "shared_group"
    }
    check = {
      "nios.mail_exchanger" = "{{random3}}.example.com"
    }
  }

}

case "name" {
  backend     = "nios"
  parallel    = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  # nios = {
  # name = "{{random4}}"
  # }
  # }
  # PREREQ

  step {
    nios {
      mail_exchanger      = "{{random3}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
    }
    check = {
      "nios.name" = "{{random}}.example.com"
    }
  }

  step {
    nios {
      mail_exchanger      = "{{random3}}.example.com"
      name                = "{{random2}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
    }
    check = {
      "nios.name" = "{{random2}}.example.com"
    }
  }

}

case "preference" {
  backend     = "nios"
  parallel    = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  # nios = {
  # name = "{{random3}}"
  # }
  # }
  # PREREQ

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
    }
    check = {
      "nios.preference" = "10"
    }
  }

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 20
      shared_record_group = "shared_group"
    }
    check = {
      "nios.preference" = "20"
    }
  }

}

case "shared_record_group" {
  backend     = "nios"
  parallel    = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  # nios = {
  # name = "{{random3}}"
  # }
  # }
  # PREREQ

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
    }
    check = {
      "nios.shared_record_group" = "shared_group"
    }
  }

}

case "ttl" {
  backend     = "nios"
  parallel    = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  # nios = {
  # name = "{{random3}}"
  # }
  # }
  # PREREQ

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
      ttl                 = 3600
    }
    check = {
      "nios.ttl" = "3600"
    }
  }

  step {
    nios {
      mail_exchanger      = "{{random2}}.example.com"
      name                = "{{random}}.example.com"
      preference          = 10
      shared_record_group = "shared_group"
      ttl                 = 4200
    }
    check = {
      "nios.ttl" = "4200"
    }
  }

}
