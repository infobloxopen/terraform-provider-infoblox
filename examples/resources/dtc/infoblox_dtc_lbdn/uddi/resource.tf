// Create a DNS View to associate with the LBDN
resource "infoblox_view" "example_view" {
  uddi = {
    name = "example_dns_view"
  }
}

// Create DTC LBDN with Basic Fields
resource "infoblox_dtc_lbdn" "lbdn_basic" {
  uddi = {
    name = "example-lbdn."
    view = infoblox_view.example_view.id
  }
}

// Create DTC LBDN with Optional Fields
resource "infoblox_dtc_lbdn" "lbdn_with_options" {
  uddi = {
    name       = "example-lbdn-advanced."
    view       = infoblox_view.example_view.id
    comment    = "Created by Terraform"
    disabled   = false
    ttl        = 300
    precedence = 5
    tags = {
      Site = "location-1"
    }
  }
}

// Create DTC LBDN linked to an existing DTC Policy
resource "infoblox_dtc_lbdn" "lbdn_with_policy" {
  uddi = {
    name = "example-lbdn-policy."
    view = infoblox_view.example_view.id
    dtc_policy = {
      policy_id = "dtc/policy/<policy-id>"
    }
    tags = {
      Site = "location-1"
    }
  }
}
