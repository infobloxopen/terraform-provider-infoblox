# TODO: The following prerequisites MUST exist in the CSP environment before running these tests:
#   - A valid AWS IAM role ARN that the CSP discovery service can assume

# CloudDiscoveryProvider — uddi datasource cases
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name               = "{{random}}"
      provider_type      = "Amazon Web Services"
      account_preference = "single"
      credential_preference = {
        access_identifier_type = "role_arn"
        credential_type        = "dynamic"
      }
      source_configs = [{
        credential_config = {
          access_identifier = "arn:aws:iam::123456789012:role/infoblox_discovery"
        }
      }]
      tags = { "site" = "datasource-test" }
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }
}
