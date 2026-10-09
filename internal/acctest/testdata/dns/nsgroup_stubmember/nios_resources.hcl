# Auto-generated resource acceptance-test cases for NsgroupStubmember.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      stub_members = [{ name = "{{grid_member_hostname}}" }]
    }
    check = {
      "nios.name"                = "{{random}}"
      "nios.stub_members.#"      = "1"
      "nios.stub_members.0.name" = "{{grid_member_hostname}}"
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
      name         = "{{random}}"
      stub_members = [{ name = "{{grid_member_hostname}}" }]
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      stub_members = [{ name = "{{grid_member_hostname}}" }]
      comment      = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      stub_members = [{ name = "{{grid_member_hostname}}" }]
      comment      = "This comment is updated"
    }
    check = {
      "nios.comment" = "This comment is updated"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      ext_attrs    = { Site = "{{random2}}" }
      stub_members = [{ name = "{{grid_member_hostname}}" }]
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      ext_attrs    = { Site = "{{random3}}" }
      stub_members = [{ name = "{{grid_member_hostname}}" }]
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
      name         = "{{random}}"
      stub_members = [{ name = "{{grid_member_hostname}}" }]
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name         = "{{random2}}"
      stub_members = [{ name = "{{grid_member_hostname}}" }]
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "stub_members" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      stub_members = [{ name = "{{grid_member_hostname}}" }]
    }
    check = {
      "nios.stub_members.0.name" = "{{grid_member_hostname}}"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      stub_members = [{ name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.stub_members.0.name" = "{{grid_master_hostname}}"
    }
  }

}
