// Update Distribution Schedule with Basic Fields
resource "infoblox_distribution_schedule" "basic_schedule" {
  nios = {
    active     = true
    start_time = "2027-01-15T20:30:00"
  }
}

// Update Distribution Schedule with Additional Fields
resource "infoblox_distribution_schedule" "schedule_with_upgrade_groups" {
  nios = {
    active     = true
    start_time = "2027-01-15T20:00:00"
    upgrade_groups = [
      {
        name              = "Default"
        distribution_time = "2027-01-15T22:30:00"
      }
    ]
  }
}
