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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
    }
  }
}

case "account_preference" {
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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
    }
    check = {
      "uddi.account_preference" = "single"
    }
  }

  step {
    uddi {
      name               = "{{random}}"
      provider_type      = "Amazon Web Services"
      account_preference = "auto_discover_multiple"
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
    check = {
      "uddi.account_preference" = "auto_discover_multiple"
    }
  }

}

case "additional_config" {
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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
      additional_config = {
        object_type = {
          objects = [{
            category = {
              excluded = true
              id       = "storage"
            }
            }
          ]
        }
        forward_zone_enabled = false
      }
    }
    check = {
      "uddi.additional_config.object_type.objects.0.category.excluded" = "single"
      "uddi.additional_config.object_type.objects.0.category.id"       = "storage"
      "uddi.additional_config.forward_zone_enabled"                    = "false"
    }
  }

  step {
    uddi {
      name               = "{{random}}"
      provider_type      = "Amazon Web Services"
      account_preference = "auto_discover_multiple"
      credential_preference = {
        access_identifier_type = "role_arn"
        credential_type        = "dynamic"
      }
      source_configs = [{
        credential_config = {
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
      additional_config = {
        object_type = {
          objects = [{
            category = {
              excluded = true
              id       = "security"
            }
            }
          ]
        }
        forward_zone_enabled = true
      }
    }
    check = {
      "uddi.additional_config.object_type.objects.0.category.excluded" = "single"
      "uddi.additional_config.object_type.objects.0.category.id"       = "security"
      "uddi.additional_config.forward_zone_enabled"                    = "true"
    }
  }

}

case "credential_preference" {
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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
    }
    check = {
      "uddi.credential_preference.credential_type" = "dynamic"
    }
  }

}

case "destinations" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_network_view" "test" {
      uddi = {
        name = "{{random}}"
      }
    }
    PREREQ

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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
      destination_types_enabled = ["IPAM/DHCP"]
      destinations = [
        {
          config           = {}
          destination_type = "IPAM/DHCP"
        }
      ]
    }
    check = {
      "uddi.destinations.0.destination_type" = "IPAM/DHCP"
    }
  }

  step {
    uddi {
      name               = "{{random}}"
      provider_type      = "Amazon Web Services"
      account_preference = "auto_discover_multiple"
      credential_preference = {
        access_identifier_type = "role_arn"
        credential_type        = "dynamic"
      }
      source_configs = [{
        credential_config = {
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
      destination_types_enabled = ["DNS"]
      destinations = [
        {
          config           = {}
          destination_type = "IPAM/DHCP"
        },
        {
          config = {
            dns = {
              view_id = infoblox_network_view.test.id
            }
          }
          destination_type = "DNS"
        }
      ]
    }
    check = {
      "uddi.destinations.0.destination_type" = "IPAM/DHCP"
      "uddi.destinations.1.destination_type" = "DNS"
    }
  }

}

case "destinations_with_zone_filters" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_network_view" "test" {
      uddi = {
        name = "{{random}}"
      }
    }
    PREREQ

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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
      destination_types_enabled = ["IPAM/DHCP", "DNS"]
      destinations = [
        {
          config           = {}
          destination_type = "IPAM/DHCP"
        },
        {
          config = {
            dns = {
              view_id = infoblox_network_view.test.id
              zone_filters = [
                {
                  action    = "include"
                  wildcards = ["*.example.com"]
                }
              ]
            }
          }
          destination_type = "DNS"
        }
      ]
    }
    check = {
      "uddi.destinations.0.destination_type" = "IPAM/DHCP"
    }
  }

  step {
    uddi {
      name               = "{{random}}"
      provider_type      = "Amazon Web Services"
      account_preference = "auto_discover_multiple"
      credential_preference = {
        access_identifier_type = "role_arn"
        credential_type        = "dynamic"
      }
      source_configs = [{
        credential_config = {
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
      destination_types_enabled = ["DNS"]
      destinations = [
        {
          config           = {}
          destination_type = "IPAM/DHCP"
        },
        {
          config = {
            dns = {
              view_id = infoblox_network_view.test.id
              zone_filters = [
                {
                  action    = "exclude"
                  wildcards = ["private.*", "internal.*"]
                }
              ]
            }
          }
        }
        destination_type = "DNS"
      }
    ]
  }
  check = {
    "uddi.destinations.0.destination_type" = "IPAM/DHCP"
    "uddi.destinations.1.destination_type" = "DNS"
  }
}

}

case "name" {
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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name               = "{{random2}}"
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
    check = {
      "uddi.name" = "{{random2}}"
    }
  }

}

case "provider_type" {
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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
        restricted_to_accounts = ["{{random_arn}}"]
      }]
    }
    check = {
      "uddi.provider_type" = "Amazon Web Services"
    }
  }

  step {
    uddi {
      name               = "{{random}}"
      provider_type      = "Google Cloud Platform"
      account_preference = "single"
      credential_preference = {
        access_identifier_type = "role_arn"
        credential_type        = "dynamic"
      }
      source_configs = [{
        credential_config = {
          access_identifier = "{{random_arn2}}"
        }
        restricted_to_accounts = ["{{random_arn2}}"]
      }]
    }
    check = {
      "uddi.provider_type" = "Google Cloud Platform"
    }
  }

  step {
    uddi {
      name               = "{{random}}"
      provider_type      = "Microsoft Azure"
      account_preference = "single"
      credential_preference = {
        access_identifier_type = "role_arn"
        credential_type        = "dynamic"
      }
      source_configs = [{
        credential_config = {
          access_identifier = "{{random_arn3}}"
        }
        restricted_to_accounts = ["{{random_arn3}}"]
      }]
    }
    check = {
      "uddi.provider_type" = "Microsoft Azure"
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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
    }
    check = {
      "uddi.desired_state" = "disabled"
    }
  }

  step {
    uddi {
      name               = "{{random}}"
      provider_type      = "Amazon Web Services"
      account_preference = "single"
      desired_state      = "enabled"
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
    check = {
      "uddi.desired_state" = "enabled"
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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
      tags = { "site" = "Site B" }
    }
    check = {
      "uddi.tags.site" = "Site B"
    }
  }
}

case "source_configs" {
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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
    }
    check = {
      "uddi.source_configs.0.credential_config.access_identifier" = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
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
          access_identifier = "arn:aws:iam::{{random_arn2}}:role/infoblox_discovery"
        }
      }]
    }
    check = {
      "uddi.source_configs.0.credential_config.access_identifier" = "arn:aws:iam::{{random_arn2}}:role/infoblox_discovery"
    }
  }

}

case "sync_interval" {
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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
      sync_interval = "15"
    }
    check = {
      "uddi.sync_interval" = "15"
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
          access_identifier = "arn:aws:iam::{{random_arn}}:role/infoblox_discovery"
        }
      }]
      sync_interval = "Auto"
    }
    check = {
      "uddi.sync_interval" = "Auto"
    }
  }

}
