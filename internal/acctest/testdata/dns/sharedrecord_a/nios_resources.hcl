# Auto-generated resource acceptance-test cases for SharedrecordA.
#
# TODO: These cases use the shared record group "shared_group", which must already
#       exist on the grid. The generated prerequisite is commented out because
#       infoblox_shared_record_group is not implemented in the provider yet.
#       Once it is, restore the prerequisite block and remove this note.
case "basic" {
  backend  = "nios"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  #   nios = {
  #     name = "{{random2}}"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = "shared_group"
    }
    check = {
      "nios.name"                = "{{random}}"
      "nios.ipv4addr"            = "10.0.0.0"
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
  #   nios = {
  #     name = "{{random2}}"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = "shared_group"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  #   nios = {
  #     name = "{{random2}}"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = "shared_group"
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
      shared_record_group = "shared_group"
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
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  #   nios = {
  #     name = "{{random2}}"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = "shared_group"
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
      shared_record_group = "shared_group"
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
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  #   nios = {
  #     name = "{{random4}}"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name                = "{{random}}.example.com"
      ipv4addr            = "10.0.0.1"
      shared_record_group = "shared_group"
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
      shared_record_group = "shared_group"
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
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  #   nios = {
  #     name = "{{random2}}"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = "shared_group"
    }
    check = {
      "nios.ipv4addr" = "10.0.0.0"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.1"
      shared_record_group = "shared_group"
    }
    check = {
      "nios.ipv4addr" = "10.0.0.1"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  #   nios = {
  #     name = "{{random3}}"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = "shared_group"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name                = "{{random2}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = "shared_group"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "shared_record_group" {
  backend  = "nios"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  #   nios = {
  #     name = "{{random2}}"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = "shared_group"
    }
    check = {
      "nios.shared_record_group" = "shared_group"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_shared_record_group_unknown" "parent_sharedrecord_group" {
  #   nios = {
  #     name = "{{random2}}"
  #   }
  # }
  # PREREQ

  step {
    nios {
      name                = "{{random}}"
      ipv4addr            = "10.0.0.0"
      shared_record_group = "shared_group"
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
      shared_record_group = "shared_group"
      ttl                 = 7200
    }
    check = {
      "nios.ttl" = "7200"
    }
  }

}
