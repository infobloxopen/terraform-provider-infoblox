# Resource acceptance-test cases for Distributionschedule.
// TODO : Objects to be present in the grid for testing
// Upgrade Groups - example_upgrade_dependent_group1, example_upgrade_dependent_group2
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

# Mirrors the legacy test: schedules the example_upgrade_dependent_group1/2 fixtures and a created
# group (plus Default), then moves their distribution_time.
case "upgrade_groups" {
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
      start_time     = "{{future_time_12h}}"
      upgrade_groups = [for n in ["example_upgrade_dependent_group1", "example_upgrade_dependent_group2", infoblox_upgrade_group.test.nios.name, "Default"] : { name = n, distribution_time = "{{future_time_14h}}" }]
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
      upgrade_groups = [for n in ["example_upgrade_dependent_group1", "example_upgrade_dependent_group2", infoblox_upgrade_group.test.nios.name, "Default"] : { name = n, distribution_time = "{{future_time_16h}}" }]
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
