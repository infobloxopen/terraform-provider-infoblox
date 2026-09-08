// Retrieve a specific AWS User by filters
data "infoblox_awsuser" "aws_user_by_filters" {
  filters = {
    name = "aws-user"
  }
}

// Retrieve all AWS Users
data "infoblox_awsuser" "get_all_aws_users" {}
