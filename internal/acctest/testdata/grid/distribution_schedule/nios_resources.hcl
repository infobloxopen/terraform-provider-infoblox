# Resource acceptance-test cases for Distributionschedule.
case "basic" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_upgrade_group" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      active         = false
      start_time     = "{{future_time_12h}}"
      upgrade_groups = [{ name = "Default", distribution_time = "{{future_time_14h}}" }, { name = infoblox_upgrade_group.test.nios.name, distribution_time = "{{future_time_14h}}" }]
    }
    check = {
      "nios.start_time" = "{{future_time_12h}}"
      "nios.active"     = "false"
    }
  }

}

case "active" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_upgrade_group" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      active         = true
      start_time     = "{{future_time_12h}}"
      upgrade_groups = [{ name = "Default", distribution_time = "{{future_time_14h}}" }, { name = infoblox_upgrade_group.test.nios.name, distribution_time = "{{future_time_14h}}" }]
    }
    check = {
      "nios.active" = "true"
    }
  }

  step {
    nios {
      active         = false
      start_time     = "{{future_time_12h}}"
      upgrade_groups = [{ name = "Default", distribution_time = "{{future_time_14h}}" }, { name = infoblox_upgrade_group.test.nios.name, distribution_time = "{{future_time_14h}}" }]
    }
    check = {
      "nios.active" = "false"
    }
  }

}

case "start_time" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_upgrade_group" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      start_time     = "{{future_time_6h}}"
      upgrade_groups = [{ name = "Default", distribution_time = "{{future_time_8h}}" }, { name = infoblox_upgrade_group.test.nios.name, distribution_time = "{{future_time_8h}}" }]
    }
    check = {
      "nios.start_time" = "{{future_time_6h}}"
    }
  }

  step {
    nios {
      start_time     = "{{future_time_10h}}"
      upgrade_groups = [{ name = "Default", distribution_time = "{{future_time_12h}}" }, { name = infoblox_upgrade_group.test.nios.name, distribution_time = "{{future_time_12h}}" }]
    }
    check = {
      "nios.start_time" = "{{future_time_10h}}"
    }
  }

}

case "upgrade_groups" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_upgrade_group" "test1" {
    nios = {
      name = "{{random}}-1"
    }
  }

  resource "infoblox_upgrade_group" "test2" {
    nios = {
      name = "{{random}}-2"
    }
  }
  PREREQ

  step {
    nios {
      start_time     = "{{future_time_12h}}"
      upgrade_groups = [{ name = infoblox_upgrade_group.test1.nios.name, distribution_time = "{{future_time_14h}}" }, { name = infoblox_upgrade_group.test2.nios.name, distribution_time = "{{future_time_14h}}" }, { name = "Default", distribution_time = "{{future_time_14h}}" }]
    }
    check = {
      "nios.upgrade_groups.0.name"              = "{{random}}-1"
      "nios.upgrade_groups.0.distribution_time" = "{{future_time_14h}}"
      "nios.upgrade_groups.1.name"              = "{{random}}-2"
      "nios.upgrade_groups.1.distribution_time" = "{{future_time_14h}}"
    }
  }

  step {
    nios {
      start_time     = "{{future_time_12h}}"
      upgrade_groups = [{ name = infoblox_upgrade_group.test1.nios.name, distribution_time = "{{future_time_16h}}" }, { name = infoblox_upgrade_group.test2.nios.name, distribution_time = "{{future_time_16h}}" }, { name = "Default", distribution_time = "{{future_time_16h}}" }]
    }
    check = {
      "nios.upgrade_groups.0.name"              = "{{random}}-1"
      "nios.upgrade_groups.0.distribution_time" = "{{future_time_16h}}"
      "nios.upgrade_groups.1.name"              = "{{random}}-2"
      "nios.upgrade_groups.1.distribution_time" = "{{future_time_16h}}"
    }
  }

}
