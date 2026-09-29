case "filters" {
  backend  = "uddi"
  parallel = true

  filter {
    type = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.account_preference", "uddi.additional_config", "uddi.credential_preference", "uddi.description", "uddi.desired_state", "uddi.destination_types_enabled", "uddi.destinations", "uddi.is_disabled", "uddi.labs_provider", "uddi.name", "uddi.provider_type", "uddi.source_configs", "uddi.sync_interval", "uddi.tags"]

  step {
    uddi {
      name               = "tf_acc_{{random_int}}"
      provider_type      = "Amazon Web Services"
      account_preference = "single"
      credential_preference = {
        access_identifier_type = "role_arn"
        credential_type        = "dynamic"
      }
      source_configs = [{
        credential_config = {
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
    }
  }
}

case "tag_filters" {
  backend  = "uddi"
  parallel = true

  filter {
    type = "tag_filters"
    values = {
      name = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.account_preference", "uddi.additional_config", "uddi.credential_preference", "uddi.description", "uddi.desired_state", "uddi.destination_types_enabled", "uddi.destinations", "uddi.is_disabled", "uddi.labs_provider", "uddi.name", "uddi.provider_type", "uddi.source_configs", "uddi.sync_interval", "uddi.tags"]

  step {
    uddi {
      name               = "tf_acc_{{random_int}}"
      provider_type      = "Amazon Web Services"
      account_preference = "single"
      credential_preference = {
        access_identifier_type = "role_arn"
        credential_type        = "dynamic"
      }
      source_configs = [{
        credential_config = {
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
      tags = { tag1 = "{{random}}" }
    }
  }
}