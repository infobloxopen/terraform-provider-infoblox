// List specific AWS Users using filters
list "infoblox_awsuser" "list_aws_users_by_filters" {
  provider = infoblox
  config {
    filters = {
      name = "aws-user"
    }
  }
  limit = 10
}

// List all AWS Users with resource details included
list "infoblox_awsuser" "list_all_aws_users" {
  provider         = infoblox
  include_resource = true
}
