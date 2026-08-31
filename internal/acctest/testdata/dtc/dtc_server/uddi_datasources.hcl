# DtcServer — uddi datasource cases

case "filters" {
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.address", "uddi.auto_create_response_records", "uddi.comment", "uddi.disabled", "uddi.endpoint_type", "uddi.fqdn", "uddi.name"]

  step {
    uddi {
      name          = "{{random}}"
      address       = "{{random_ip}}"
      endpoint_type = "address"
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type   = "tag_filters"
    values = {
      env = "uddi.tags_all.env"
    }
  }

  pair_checks = ["uddi.address", "uddi.auto_create_response_records", "uddi.comment", "uddi.disabled", "uddi.endpoint_type", "uddi.fqdn", "uddi.name"]

  step {
    uddi {
      name          = "{{random}}"
      address       = "{{random_ip}}"
      endpoint_type = "address"
      tags          = { env = "{{random2}}" }
    }
  }

}
