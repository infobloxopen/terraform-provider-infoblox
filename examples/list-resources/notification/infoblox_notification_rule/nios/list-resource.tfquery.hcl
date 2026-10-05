// List specific Notification Rules using filters
list "infoblox_notification_rule" "list_notification_rules_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_notification_rule"
    }
  }
  limit = 10
}

// List Notification Rules with resource details included
list "infoblox_notification_rule" "list_notification_rules_with_resource" {
  provider         = infoblox
  include_resource = true
}
