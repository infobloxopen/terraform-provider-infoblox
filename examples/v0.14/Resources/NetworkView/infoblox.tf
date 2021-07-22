terraform {
  # Required providers block for Terraform v0.14.7
  required_providers {
    infoblox = {
      source  = "terraform-providers/infoblox"
      version = ">= 1.0"
    }
  }
}

resource "infoblox_network_view" "TestNetworkView" {
  name = "TestNetworkView"
  comment = "New Network View"
  ext_attrs = jsonencode({
    "Tenant ID" = "tf-plugin"
    "Location" = "Test loc."
    "Site" = "Test site"
  })
}

