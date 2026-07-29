# Auto-generated resource acceptance-test cases for RecordCname.
case "rdata" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      name_in_zone = "cname"
      rdata        = { cname = "c1" }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.cname" = "c1"
    }
  }

  step {
    uddi {
      name_in_zone = "cname"
      rdata        = { cname = "c2" }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.cname" = "c2"
    }
  }

}
