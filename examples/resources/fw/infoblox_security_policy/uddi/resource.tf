// Create a Security Policy with basic fields
resource "infoblox_security_policy" "example_basic" {
  uddi = {
    name = "example-security-policy"
  }
}

// Create a Security Policy with additional fields
resource "infoblox_security_policy" "example_full" {
  uddi = {
    name           = "example-security-policy-full"
    description    = "Security policy created by Terraform"
    default_action = "action_allow"
    ecs            = true
    onprem_resolve = false
    safe_search    = false
    tags = {
      Site = "location-1"
    }
  }
}

// Create a Security Policy with DNS Forwarding Proxies assigned
resource "infoblox_security_policy" "example_with_dfps" {
  uddi = {
    name = "example-security-policy-with-dfps"
    dfps = [530499]
  }
}
