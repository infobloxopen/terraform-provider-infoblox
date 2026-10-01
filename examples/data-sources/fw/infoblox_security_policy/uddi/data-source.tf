// Retrieve Security Policies filtered by an attribute
data "infoblox_security_policy" "get_security_policy_using_filters" {
  filters = {
    name = "example-security-policy"
  }
}

// Retrieve Security Policies filtered by tag
data "infoblox_security_policy" "get_security_policy_using_tag_filters" {
  tag_filters = {
    Site = "location-1"
  }
}

// Retrieve all Security Policies
data "infoblox_security_policy" "get_all_security_policies" {}
