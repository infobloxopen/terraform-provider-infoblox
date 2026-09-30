# Data source acceptance-test cases for Distributionschedule.
case "read" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_upgrade_group" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  pair_checks = [
    "id",
    "nios.upgrade_groups.0.name",
    "nios.upgrade_groups.0.distribution_time",
  ]

  filter {
    type   = "filters"
    values = {}
  }

  step {
    nios {
      active         = false
      start_time     = "{{future_time_12h}}"
      upgrade_groups = [{ name = "Default", distribution_time = "{{future_time_14h}}" }, { name = infoblox_upgrade_group.test.nios.name, distribution_time = "{{future_time_14h}}" }]
    }
  }
}
