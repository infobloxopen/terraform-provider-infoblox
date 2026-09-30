# Resource acceptance-test cases for Distributionschedule.
// TODO : Objects to be present in the grid for testing
// Upgrade Groups - example_upgrade_dependent_group1, example_upgrade_dependent_group2
# Distributionschedule is a grid singleton (noCreate/noDelete): Create adopts the grid's
# existing schedule (found via the lookup hook) instead of creating one.
# NIOS validates the whole schedule on every update: every upgrade group on the grid must be
# listed ("Missing upgrade groups"), and 'Default' needs a distribution_time later than
# start_time. Steps list Default first, then every other group read via the upgradegroup
# data source ('Grid Master' is upgraded as part of Default and cannot be scheduled).
case "basic" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  data "infoblox_upgradegroup" "all" {}

  locals {
    other_groups = [for g in data.infoblox_upgradegroup.all.results : g.nios.name if !contains(["Default", "Grid Master"], g.nios.name)]
  }
  PREREQ

  step {
    nios {
      active         = false
      start_time     = "{{future_time_12h}}"
      upgrade_groups = concat([{ name = "Default", distribution_time = "{{future_time_14h}}" }], [for n in local.other_groups : { name = n, distribution_time = "{{future_time_14h}}" }])
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
  data "infoblox_upgradegroup" "all" {}

  locals {
    other_groups = [for g in data.infoblox_upgradegroup.all.results : g.nios.name if !contains(["Default", "Grid Master"], g.nios.name)]
  }
  PREREQ

  step {
    nios {
      active         = true
      start_time     = "{{future_time_12h}}"
      upgrade_groups = concat([{ name = "Default", distribution_time = "{{future_time_14h}}" }], [for n in local.other_groups : { name = n, distribution_time = "{{future_time_14h}}" }])
    }
    check = {
      "nios.active" = "true"
    }
  }

  step {
    nios {
      active         = false
      start_time     = "{{future_time_12h}}"
      upgrade_groups = concat([{ name = "Default", distribution_time = "{{future_time_14h}}" }], [for n in local.other_groups : { name = n, distribution_time = "{{future_time_14h}}" }])
    }
    check = {
      "nios.active" = "false"
    }
  }

}

case "start_time" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  data "infoblox_upgradegroup" "all" {}

  locals {
    other_groups = [for g in data.infoblox_upgradegroup.all.results : g.nios.name if !contains(["Default", "Grid Master"], g.nios.name)]
  }
  PREREQ

  step {
    nios {
      start_time     = "{{future_time_6h}}"
      upgrade_groups = concat([{ name = "Default", distribution_time = "{{future_time_8h}}" }], [for n in local.other_groups : { name = n, distribution_time = "{{future_time_8h}}" }])
    }
    check = {
      "nios.start_time" = "{{future_time_6h}}"
    }
  }

  step {
    nios {
      start_time     = "{{future_time_10h}}"
      upgrade_groups = concat([{ name = "Default", distribution_time = "{{future_time_12h}}" }], [for n in local.other_groups : { name = n, distribution_time = "{{future_time_12h}}" }])
    }
    check = {
      "nios.start_time" = "{{future_time_10h}}"
    }
  }

}

# Mirrors the legacy test: schedules example_upgrade_dependent_group1/2 (grid fixtures) and a
# newly created group, then moves their distribution_time. The remaining groups (Default included)
# are given the same time so the case does not depend on times left on the grid by earlier runs.
case "upgrade_groups" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  # The schedule writes distribution_time onto the group itself; ignore it here so the
  # group does not drift (which would also defer the data source read below).
  resource "infoblox_upgradegroup" "test" {
    nios = {
      name = "{{random}}"
    }
    lifecycle {
      ignore_changes = [nios.distribution_time]
    }
  }

  data "infoblox_upgradegroup" "all" {
    depends_on = [infoblox_upgradegroup.test]
  }

  locals {
    named_groups = ["example_upgrade_dependent_group1", "example_upgrade_dependent_group2", infoblox_upgradegroup.test.nios.name]
    other_groups = [for g in data.infoblox_upgradegroup.all.results : g.nios.name if !contains(concat(["Grid Master"], local.named_groups), g.nios.name)]
  }
  PREREQ

  step {
    nios {
      start_time     = "{{future_time_12h}}"
      upgrade_groups = [for n in concat(local.named_groups, local.other_groups) : { name = n, distribution_time = "{{future_time_14h}}" }]
    }
    check = {
      "nios.upgrade_groups.0.name"              = "example_upgrade_dependent_group1"
      "nios.upgrade_groups.0.distribution_time" = "{{future_time_14h}}"
      "nios.upgrade_groups.1.name"              = "example_upgrade_dependent_group2"
      "nios.upgrade_groups.1.distribution_time" = "{{future_time_14h}}"
      "nios.upgrade_groups.2.name"              = "{{random}}"
      "nios.upgrade_groups.2.distribution_time" = "{{future_time_14h}}"
    }
  }

  step {
    nios {
      start_time     = "{{future_time_12h}}"
      upgrade_groups = [for n in concat(local.named_groups, local.other_groups) : { name = n, distribution_time = "{{future_time_16h}}" }]
    }
    check = {
      "nios.upgrade_groups.0.name"              = "example_upgrade_dependent_group1"
      "nios.upgrade_groups.0.distribution_time" = "{{future_time_16h}}"
      "nios.upgrade_groups.1.name"              = "example_upgrade_dependent_group2"
      "nios.upgrade_groups.1.distribution_time" = "{{future_time_16h}}"
      "nios.upgrade_groups.2.name"              = "{{random}}"
      "nios.upgrade_groups.2.distribution_time" = "{{future_time_16h}}"
    }
  }

}
