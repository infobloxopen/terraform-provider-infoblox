// Create a Cloud Discovery Provider for Amazon Web Services (basic)
resource "infoblox_cloud_discovery_provider" "aws_basic" {
  uddi = {
    name               = "example-aws-provider"
    provider_type      = "Amazon Web Services"
    account_preference = "single"
    credential_preference = {
      access_identifier_type = "role_arn"
      credential_type        = "dynamic"
    }
    source_configs = [
      {
        credential_config = {
          access_identifier = "arn:aws:iam::123456789012:role/infoblox_discovery"
        }
      }
    ]
  }
}

// Create a Cloud Discovery Provider with additional fields
resource "infoblox_cloud_discovery_provider" "aws_advanced" {
  uddi = {
    name               = "example-aws-provider-advanced"
    provider_type      = "Amazon Web Services"
    account_preference = "single"
    description        = "AWS discovery provider for production accounts"
    desired_state      = "enabled"
    sync_interval      = "Auto"
    credential_preference = {
      access_identifier_type = "role_arn"
      credential_type        = "dynamic"
    }
    source_configs = [
      {
        credential_config = {
          access_identifier = "arn:aws:iam::123456789012:role/infoblox_discovery"
          region            = "us-east-1"
        }
      }
    ]
    destination_types_enabled = ["IPAM/DHCP", "DNS"]
    tags = {
      environment = "production"
      site        = "Site A"
    }
  }
}
