# Auto-generated datasource acceptance-test cases for RecordTxt.
case "filters" {
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
      # zone  = infoblox_zone_auth.test.id
      zone  = "dns/auth_zone/491ca52a-b154-4411-a684-0faf1d118719"
      rdata = { text = "xyz" }
    }
  }

}
