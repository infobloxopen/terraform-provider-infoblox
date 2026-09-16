// Create an AWS User with Basic Fields
resource "infoblox_awsuser" "aws_user_basic" {
  nios = {
    access_key_id     = "AKIAexample1"
    account_id        = "337773173961"
    name              = "aws-user"
    secret_access_key = "S1JGWfwcZWESkfpyhxigL9A/u96mY"
  }
}

// Create an AWS User with Additional Fields
resource "infoblox_awsuser" "aws_user_additional" {
  nios = {
    access_key_id     = "AKIAexample2"
    account_id        = "337773173962"
    name              = "aws-user-2"
    secret_access_key = "S1JGWfwcZWkfpyhxigL9A/ua6mZ"
    govcloud_enabled  = true
  }
}
