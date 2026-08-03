# Auto-generated resource acceptance-test cases for RecordTxt.
case "rdata" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      name_in_zone = "txt"
      rdata        = { text = "abc" }
      zone         = infoblox_zone_auth.test.id
      #zone  = "dns/auth_zone/491ca52a-b154-4411-a684-0faf1d118719"
    }
    check = {
      "uddi.rdata.text" = "abc"
    }
  }

  step {
    uddi {
      name_in_zone = "txt"
      rdata        = { text = "xyz" }
      zone         = infoblox_zone_auth.test.id
      #zone  = "dns/auth_zone/491ca52a-b154-4411-a684-0faf1d118719"
    }
    check = {
      "uddi.rdata.text" = "xyz"
    }
  }

}
