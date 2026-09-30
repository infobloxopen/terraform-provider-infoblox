// Create a Security Policy with basic fields
resource "infoblox_security_policy" "example_basic" {
  uddi = {
    name = "example-security-policy"
  }
}

// Create a Named List and Application Filter to reference in the policy rules
resource "infoblox_named_list" "example" {
  uddi = {
    name            = "example-named-list"
    type            = "custom_list"
    items_described = [{ item = "tf-domain.com", description = "Example Domain" }]
  }
}

resource "infoblox_application_filter" "example" {
  uddi = {
    name     = "example-app-filter"
    criteria = [{ name = "Microsoft 365" }]
  }
}

// Create a Security Policy with all configurable fields
resource "infoblox_security_policy" "example_full" {
  uddi = {
    name                  = "example-security-policy-full"
    description           = "Security policy created by Terraform"
    default_action        = "action_allow"
    default_redirect_name = ""
    ecs                   = true
    onprem_resolve        = false
    safe_search           = false
    dfps                  = [530499]
    tags = {
      Site = "location-1"
    }
    rules = [
      {
        action = "action_block"
        data   = infoblox_named_list.example.uddi.name
        type   = infoblox_named_list.example.uddi.type
      },
      {
        action = "action_allow"
        data   = infoblox_application_filter.example.uddi.name
        type   = "application_filter"
      }
    ]
  }
}
