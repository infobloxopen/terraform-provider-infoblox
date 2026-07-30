# Auto-generated resource acceptance-test cases for RecordNs.
case "rdata" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      name_in_zone = "ns"
      rdata        = { dname = "ns4.zone.com." }
      zone         = "dns/auth_zone/0060207e-7664-4742-8eed-2d8e34db0035"
    }
    check = {
      "uddi.rdata.dname" = "ns4.zone.com."
    }
  }

  step {
    uddi {
      name_in_zone = "ns"
      rdata        = { dname = "ns3.zone.com." }
      zone         = "dns/auth_zone/0060207e-7664-4742-8eed-2d8e34db0035"
    }
    check = {
      "uddi.rdata.dname" = "ns3.zone.com."
    }
  }

}
