// Manage a DNS View ( Required as Parent )
resource "bloxone_dns_view" "example_azure" {
  name = "example_dns_view_azure"
}

// Manage an Azure Cloud Discovery Provider
resource "infoblox_cloud_discovery_provider" "example_azure" {
  name               = "example_provider_azure"
  provider_type      = "Microsoft Azure"
  account_preference = "auto_discover_multiple"
  credential_preference = {
    access_identifier_type = "tenant_id"
    credential_type        = "dynamic"
  }
  source_configs = [
    {
      credential_config = {
        access_identifier = "xyz98765-4321-abcd-efgh-ijklmnopqrst"
      }
      restricted_to_accounts = ["12345678-abcd-efgh-ijkl-901234567890"]
    }
  ]
  destination_types_enabled = [
    "IPAM/DHCP",
    "ACCOUNTS",
    "DNS"
  ]
  destinations = [
    {
      config           = {}
      destination_type = "IPAM/DHCP"
    },
    {
      config           = {}
      destination_type = "ACCOUNTS"
    },
    {
      config = {
        dns = {
          view_id = bloxone_dns_view.example_azure.id
          # Optional: filter which DNS zones are synced
          zone_filters = [
            {
              action    = "exclude"
              wildcards = ["*.test.azure.com", "*.staging.azure.com"]
            }
          ]
        }
      }
      destination_type = "DNS"
    }
  ]

  # Other Optional fields

  tags = {
    site = "Site A"
  }
}
