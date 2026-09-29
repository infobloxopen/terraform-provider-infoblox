# List acceptance-test cases for Distributionschedule.
# Distributionschedule is a grid singleton: the create step adopts the existing schedule.
# Every upgrade group on the grid must be listed, with Default after start_time (see nios_resources.hcl).
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  data "infoblox_upgradegroup" "all" {}

  locals {
    other_groups = [for g in data.infoblox_upgradegroup.all.results : g.nios.name if !contains(["Default", "Grid Master"], g.nios.name)]
  }
  PREREQ

  step {
    nios {
      start_time     = "{{future_time_12h}}"
      upgrade_groups = concat([{ name = "Default", distribution_time = "{{future_time_14h}}" }], [for n in local.other_groups : { name = n, distribution_time = "{{future_time_14h}}" }])
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "include_resource" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  data "infoblox_upgradegroup" "all" {}

  locals {
    other_groups = [for g in data.infoblox_upgradegroup.all.results : g.nios.name if !contains(["Default", "Grid Master"], g.nios.name)]
  }
  PREREQ

  step {
    nios {
      start_time     = "{{future_time_12h}}"
      upgrade_groups = concat([{ name = "Default", distribution_time = "{{future_time_14h}}" }], [for n in local.other_groups : { name = n, distribution_time = "{{future_time_14h}}" }])
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
  }

}
