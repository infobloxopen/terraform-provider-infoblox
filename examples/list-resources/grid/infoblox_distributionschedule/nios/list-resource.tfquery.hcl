// List the singleton Distribution Schedule
list "infoblox_distributionschedule" "list_distributionschedule" {
  provider = infoblox
}

// List the Distribution Schedule with resource details included
list "infoblox_distributionschedule" "list_distributionschedule_with_resource" {
  provider         = infoblox
  include_resource = true
}
