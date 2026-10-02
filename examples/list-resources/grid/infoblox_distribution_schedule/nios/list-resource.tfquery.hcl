// List the singleton Distribution Schedule
list "infoblox_distribution_schedule" "list_distribution_schedule" {
  provider = infoblox
}

// List the Distribution Schedule with resource details included
list "infoblox_distribution_schedule" "list_distribution_schedule_with_resource" {
  provider         = infoblox
  include_resource = true
}
