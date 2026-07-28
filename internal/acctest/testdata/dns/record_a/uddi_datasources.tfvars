# Auto-generated datasource acceptance-test cases for RecordA (uddi).
case "filters" {
  # filters — generated from terraform-provider-uddi
  backend = "uddi"

  filter {
    type   = "filters"
    values = {
      name_in_zone = "uddi.name_in_zone"
      zone         = "uddi.zone"
    }
  }

  step {
    uddi {
      zone  = infoblox_zone_auth.test.id
      rdata = { address = "10.0.0.15" }
    }
  }

}

case "tag_filters" {
  # tag_filters — generated from terraform-provider-uddi
  backend = "uddi"

  filter {
    type   = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  step {
    uddi {
      zone  = infoblox_zone_auth.test.id
      rdata = { address = "10.0.0.15" }
    }
  }

}
