# Data source acceptance-test cases for Distributionschedule.
# Mirrors the legacy test: adopt the grid's singleton schedule, then read it back with no
# filter (empty filters map) and pair-check the data source result against the resource.
# active stays false so the schedule is left deactivated (the legacy test's deactivate step).
# See nios_resources.hcl for why every upgrade group is listed with Default first.
case "read" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  data "infoblox_upgradegroup" "all" {}

  locals {
    other_groups = [for g in data.infoblox_upgradegroup.all.results : g.nios.name if !contains(["Default", "Grid Master"], g.nios.name)]
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
      upgrade_groups = concat([{ name = "Default", distribution_time = "{{future_time_14h}}" }], [for n in local.other_groups : { name = n, distribution_time = "{{future_time_14h}}" }])
    }
  }
}
