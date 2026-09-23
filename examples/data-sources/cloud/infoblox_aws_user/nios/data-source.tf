// Retrieve a specific AWS User by filters
data "infoblox_aws_user" "aws_user_by_filters" {
  filters = {
    name = "aws-user"
  }
}

// Retrieve all AWS Users
data "infoblox_aws_user" "get_all_aws_users" {}
