# Auto-generated datasource acceptance-test cases for Upgradegroup.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.distribution_dependent_group", "nios.distribution_policy", "nios.distribution_time", "nios.name", "nios.upgrade_dependent_group", "nios.upgrade_policy", "nios.upgrade_time"]

  step {
    nios {
      name = "{{random}}"
    }
  }

}
