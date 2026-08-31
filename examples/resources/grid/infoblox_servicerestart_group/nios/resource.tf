# // Create a Service Restart Group with Basic Fields
# resource "infoblox_servicerestart_group" "grid_servicerestart_group_with_basic_fields" {
#   nios = {
#     name    = "example_grid_service_restart_group"
#     service = "DNS"
#   }
# }

# // Create a Service Restart Group with Additional Fields
# resource "infoblox_servicerestart_group" "grid_servicerestart_group_with_additional_fields" {
#   nios = {
#     name    = "example_grid_service_restart_group_additional"
#     service = "DHCP"
#     comment = "Comment for GRID Service Restart Group"


#     ext_attrs = {
#       Site = "location-1"
#     }

#     members = ["infoblox.172_28_83_113"]
#     mode    = "SEQUENTIAL"

#     recurring_schedule = {
#       services = ["ALL"]
#       mode     = "SIMULTANEOUS"
#       force    = false
#       disabled = false

#       schedule = {
#         minutes_past_hour = 6
#         repeat            = "ONCE"
#         day_of_month      = 30
#         month             = 1
#         year              = 2027
#         hour_of_day       = 20
#       }
#     }
#   }
# }

resource "infoblox_servicerestart_group" "test" {
  nios = {
    name    = "grid_service_restart_group_additional"
    service = "DHCP"

    recurring_schedule = {
      services = ["ALL"]
      mode     = "SIMULTANEOUS"
      force    = false

      # schedule = {
      #   minutes_past_hour = 6
      #   repeat            = "ONCE"
      #   day_of_month      = 30
      #   month             = 1
      #   year              = 2027
      #   hour_of_day       = 20
      # }
      schedule = {
        weekdays          = ["MONDAY"]
        frequency         = "WEEKLY"
        every             = 1
        minutes_past_hour = 6
        hour_of_day       = 20
        repeat            = "RECUR"
        # day_of_month, month, year intentionally dropped
      }
    }
  }
}
