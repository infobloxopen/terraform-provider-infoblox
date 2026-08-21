// Create DTC LBDN with Basic Fields
// The view field is the resource identifier for the DNS view (e.g. dns/view/<uuid>)
resource "infoblox_dtc_lbdn" "lbdn_basic" {
  uddi = {
    name = "example-lbdn."
    view = "dns/view/<view-id>"
  }
}

// Create DTC LBDN with Optional Fields
resource "infoblox_dtc_lbdn" "lbdn_with_options" {
  uddi = {
    name       = "example-lbdn-advanced."
    view       = "dns/view/<view-id>"
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
    view = "dns/view/<view-id>"
    dtc_policy = {
      policy_id = "dtc/policy/<policy-id>"
    }
    tags = {
      Site = "location-1"
    }
  }
}
