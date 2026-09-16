# Auto-generated resource acceptance-test cases for Upgradegroup.
// TODO : Objects to be present in the grid for testing
// Upgrade/Distribution Dependent Group - example_upgrade_dependent_group1, example_upgrade_dependent_group2
// Grid Members - set NIOS_GRID_MEMBER_HOSTNAME and NIOS_GRID_MEMBER_2_HOSTNAME env vars

case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name"                = "{{random}}"
      "nios.distribution_policy" = "SIMULTANEOUSLY"
      "nios.upgrade_policy"      = "SEQUENTIALLY"
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
      name = "{{random}}"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "This is an updated comment"
    }
    check = {
      "nios.comment" = "This is an updated comment"
    }
  }

}

case "distribution_dependent_group" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                         = "{{random}}"
      distribution_dependent_group = "example_upgrade_dependent_group1"
    }
    check = {
      "nios.distribution_dependent_group" = "example_upgrade_dependent_group1"
    }
  }

  step {
    nios {
      name                         = "{{random}}"
      distribution_dependent_group = "example_upgrade_dependent_group2"
    }
    check = {
      "nios.distribution_dependent_group" = "example_upgrade_dependent_group2"
    }
  }

}

case "distribution_policy" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      distribution_policy = "SEQUENTIALLY"
    }
    check = {
      "nios.distribution_policy" = "SEQUENTIALLY"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      distribution_policy = "SIMULTANEOUSLY"
    }
    check = {
      "nios.distribution_policy" = "SIMULTANEOUSLY"
    }
  }

}

case "distribution_time" {
  backend = "nios"

  step {
    nios {
      name              = "{{random}}"
      distribution_time = "{{future_time_24h}}"
      members           = [{ member = "{{grid_member_hostname}}" }]
    }
    check = {
      "nios.distribution_time" = "{{future_time_24h}}"
    }
  }

  step {
    nios {
      name              = "{{random}}"
      distribution_time = "{{future_time_48h}}"
      members           = [{ member = "{{grid_member_hostname}}" }]
    }
    check = {
      "nios.distribution_time" = "{{future_time_48h}}"
    }
  }

}

case "members" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      members = [{ member = "{{grid_member_hostname}}" }]
    }
    check = {
      "nios.members.0.member" = "{{grid_member_hostname}}"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      members = [{ member = "{{grid_member_2_hostname}}" }]
    }
    check = {
      "nios.members.0.member" = "{{grid_member_2_hostname}}"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "upgrade_dependent_group" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                    = "{{random}}"
      upgrade_dependent_group = "example_upgrade_dependent_group1"
    }
    check = {
      "nios.upgrade_dependent_group" = "example_upgrade_dependent_group1"
    }
  }

  step {
    nios {
      name                    = "{{random}}"
      upgrade_dependent_group = "example_upgrade_dependent_group2"
    }
    check = {
      "nios.upgrade_dependent_group" = "example_upgrade_dependent_group2"
    }
  }

}

case "upgrade_policy" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name           = "{{random}}"
      upgrade_policy = "SIMULTANEOUSLY"
    }
    check = {
      "nios.upgrade_policy" = "SIMULTANEOUSLY"
    }
  }

  step {
    nios {
      name           = "{{random}}"
      upgrade_policy = "SEQUENTIALLY"
    }
    check = {
      "nios.upgrade_policy" = "SEQUENTIALLY"
    }
  }

}

case "upgrade_time" {
  backend = "nios"

  step {
    nios {
      name         = "{{random}}"
      upgrade_time = "{{future_time_30h}}"
      members      = [{ member = "{{grid_member_hostname}}" }]
    }
    check = {
      "nios.upgrade_time" = "{{future_time_30h}}"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      upgrade_time = "{{future_time_54h}}"
      members      = [{ member = "{{grid_member_hostname}}" }]
    }
    check = {
      "nios.upgrade_time" = "{{future_time_54h}}"
    }
  }

}
