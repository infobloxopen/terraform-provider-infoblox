# InfraService — uddi list cases
#  TODO: Objects to be present in the grid for testing
#   - infra/pool : infra/pool/ch4sbtxrcat2wkppegumcuizdyll5nur

case "basic" {
  backend        = "uddi"
  parallel       = true
  min_tf_version = "1.14.0"

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = "infra/pool/ch4sbtxrcat2wkppegumcuizdyll5nur"
      service_type = "dns"
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
  parallel       = true
  min_tf_version = "1.14.0"

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = "infra/pool/ch4sbtxrcat2wkppegumcuizdyll5nur"
      service_type = "dns"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter           = "name=='{{random}}'"
  }

}
