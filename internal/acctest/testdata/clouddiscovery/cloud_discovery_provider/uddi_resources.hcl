# TODO: The following prerequisites MUST exist in the CSP environment before running these tests:
#   - A valid AWS IAM role ARN that the CSP discovery service can assume

# CloudDiscoveryProvider — uddi resource cases
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
    }
    check = {
      "uddi.name"               = "{{random}}"
      "uddi.provider_type"      = "Amazon Web Services"
      "uddi.account_preference" = "single"
      "uddi.desired_state"      = "enabled"
      "uddi.sync_interval"      = "Auto"
    }
  }
}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

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
    }
  }
}

case "description" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name               = "{{random}}"
      provider_type      = "Amazon Web Services"
      account_preference = "single"
      description        = "Initial description"
      credential_preference = {
        access_identifier_type = "role_arn"
        credential_type        = "dynamic"
      }
      source_configs = [{
        credential_config = {
          access_identifier = "arn:aws:iam::123456789012:role/infoblox_discovery"
        }
      }]
    }
    check = {
      "uddi.description" = "Initial description"
    }
  }

  step {
    uddi {
      name               = "{{random}}"
      provider_type      = "Amazon Web Services"
      account_preference = "single"
      description        = "Updated description"
      credential_preference = {
        access_identifier_type = "role_arn"
        credential_type        = "dynamic"
      }
      source_configs = [{
        credential_config = {
          access_identifier = "arn:aws:iam::123456789012:role/infoblox_discovery"
        }
      }]
    }
    check = {
      "uddi.description" = "Updated description"
    }
  }
}

case "desired_state" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name               = "{{random}}"
      provider_type      = "Amazon Web Services"
      account_preference = "single"
      desired_state      = "disabled"
      credential_preference = {
        access_identifier_type = "role_arn"
        credential_type        = "dynamic"
      }
      source_configs = [{
        credential_config = {
          access_identifier = "arn:aws:iam::123456789012:role/infoblox_discovery"
        }
      }]
    }
    check = {
      "uddi.desired_state" = "disabled"
    }
  }
}

case "tags" {
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
      tags = { "site" = "Site A" }
    }
    check = {
      "uddi.tags.site" = "Site A"
    }
  }

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
      tags = { "site" = "Site B" }
    }
    check = {
      "uddi.tags.site" = "Site B"
    }
  }
}
