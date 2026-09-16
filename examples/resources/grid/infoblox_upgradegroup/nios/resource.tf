// Create Upgrade Group with Basic Fields
resource "infoblox_upgradegroup" "upgradegroup_basic_fields" {
  nios = {
    name = "example-upgradegroup"
  }
}

// Create Upgrade Group with Additional Fields
resource "infoblox_upgradegroup" "upgradegroup_with_additional_config" {
  nios = {
    name                         = "upgradegroup-additional-fields"
    comment                      = "Example Upgrade Group for Grid members"
    distribution_policy          = "SIMULTANEOUSLY"
    distribution_dependent_group = "example_distribution_dependent_group"
    distribution_time            = "2026-09-01T02:00:00"
    upgrade_policy               = "SEQUENTIALLY"
    upgrade_dependent_group      = "example_upgrade_dependent_group"
    upgrade_time                 = "2026-09-02T02:00:00"
    members = [
      {
        member = "infoblox.localdomain"
      }
    ]
  }
}
