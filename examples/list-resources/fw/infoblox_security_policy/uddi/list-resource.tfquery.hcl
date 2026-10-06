// List specific Security Policies using filters
list "infoblox_security_policy" "list_security_policy_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example-security-policy"
    }
  }
  limit = 10
}

// List specific Security Policies using Tags
list "infoblox_security_policy" "list_security_policy_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Security Policies with resource details included
list "infoblox_security_policy" "list_security_policy_with_resource" {
  provider         = infoblox
  include_resource = true
}
