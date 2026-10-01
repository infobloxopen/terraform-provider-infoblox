// Manage a DNS View ( Required as Parent )
resource "infoblox_view" "example" {
  name = "example_dns_view"
}

// Manage an AWS Cloud Discovery Provider
resource "infoblox_cloud_discovery_provider" "example_aws" {
  name               = "example_provider_aws"
  provider_type      = "Amazon Web Services"
  account_preference = "single"
  credential_preference = {
    access_identifier_type = "role_arn"
    credential_type        = "dynamic"
  }
  source_configs = [
    {
      credential_config = {
        access_identifier = "arn:aws:iam::123456789012:role/role-name"
      }
    }
  ]
  destination_types_enabled = [
    "IPAM/DHCP",
    "DNS"
  ]
  # Other Optional fields
  destinations = [
    {
      config           = {}
      destination_type = "IPAM/DHCP"
    },
    {
      config = {
        dns = {
          view_id = infoblox_view.example.id
          # Optional: filter which DNS zones are synced
          zone_filters = [
            {
              action    = "include"
              wildcards = ["*.example.com", "*.internal.example.com"]
            }
          ]
        }
      }
      destination_type = "DNS"
    }
  ]

  tags = {
    Site = "location-1"
  }

}
