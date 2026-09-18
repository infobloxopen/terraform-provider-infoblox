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

// Create a Named List and assign it to a Security Policy rule
resource "infoblox_named_list" "example" {
  uddi = {
    name            = "example-named-list"
    type            = "custom_list"
    items_described = [{ item = "tf-domain.com", description = "Example Domain" }]
  }
}

resource "infoblox_security_policy" "example_with_named_list" {
  uddi = {
    name = "example-sp-with-named-list"
    rules = [
      {
        action = "action_block"
        data   = infoblox_named_list.example.uddi.name
        type   = infoblox_named_list.example.uddi.type
      }
    ]
  }
}

// Create an Application Filter and assign it to a Security Policy rule
resource "infoblox_application_filter" "example" {
  uddi = {
    name     = "example-app-filter"
    criteria = [{ name = "Microsoft 365" }]
  }
}

resource "infoblox_security_policy" "example_with_application_filter" {
  uddi = {
    name = "example-sp-with-application-filter"
    rules = [
      {
        action = "action_allow"
        data   = infoblox_application_filter.example.uddi.name
        type   = "application_filter"
      }
    ]
  }
}
