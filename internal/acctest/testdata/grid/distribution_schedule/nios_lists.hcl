# List acceptance-test cases for Distributionschedule.
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_upgrade_group" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      start_time     = "{{future_time_12h}}"
      upgrade_groups = [{ name = "Default", distribution_time = "{{future_time_14h}}" }, { name = infoblox_upgrade_group.test.nios.name, distribution_time = "{{future_time_14h}}" }]
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
  resource "infoblox_upgrade_group" "test" {
    nios = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    nios {
      start_time     = "{{future_time_12h}}"
      upgrade_groups = [{ name = "Default", distribution_time = "{{future_time_14h}}" }, { name = infoblox_upgrade_group.test.nios.name, distribution_time = "{{future_time_14h}}" }]
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
  }

}
