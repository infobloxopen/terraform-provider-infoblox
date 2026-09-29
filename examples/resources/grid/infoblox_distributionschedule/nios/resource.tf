// The Distribution Schedule is a grid singleton: each resource below manages the same
// object, so use one of them. All times must be in the future, and every upgrade group's
// distribution_time must be after start_time.

// Update Distribution Schedule with Basic Fields
resource "infoblox_distributionschedule" "basic_schedule" {
  nios = {
    active     = true
    start_time = "2027-01-15T20:30:00"
  }
}

// Update Distribution Schedule with Additional Fields
resource "infoblox_distributionschedule" "schedule_with_upgrade_groups" {
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
