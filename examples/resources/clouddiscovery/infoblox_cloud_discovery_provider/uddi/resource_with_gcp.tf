// Manage a DNS View ( Required as Parent )
resource "infoblox_view" "example_gcp" {
  uddi = {
    name = "example_dns_view_gcp"
  }
}

// Manage a GCP Cloud Discovery Provider
resource "infoblox_cloud_discovery_provider" "example_gcp" {
  uddi = {
    name               = "example_provider_gcp"
    provider_type      = "Google Cloud Platform"
    account_preference = "single"
    credential_preference = {
      access_identifier_type = "project_id"
      credential_type        = "dynamic"
    }
    source_configs = [
      {
        credential_config = {
          access_identifier = "my-bloxone-example-2024"
        }
      }
    ]
    destination_types_enabled = [
      "IPAM/DHCP",
      "DNS"
    ]
    destinations = [
      {
        config           = {}
        destination_type = "IPAM/DHCP"
      },
      {
        config = {
          dns = {
            view_id = infoblox_view.example_gcp.id
            # Optional: filter which DNS zones are synced
            zone_filters = [
              {
                action    = "include"
                wildcards = ["*.gcp.example.com"]
              }
            ]
          }
        }
        destination_type = "DNS"
      }
    ]

    # Other Optional fields

    tags = {
      site = "Site B"
    }
  }
}
