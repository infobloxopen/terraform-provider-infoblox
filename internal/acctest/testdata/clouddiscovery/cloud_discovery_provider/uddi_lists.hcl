case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

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

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }
}

case "filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

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

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name = "uddi.name"
      }
    }
  }

}

case "tag_filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

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
      tags = { tag1 = "{{random2}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "tag_filters"
      values = {
        tag1 = "uddi.tags.tag1"
      }
    }
  }

}
