# Auto-generated resource acceptance-test cases for RecordPtr.
case "rdata" {
  backend = "uddi"
  parallel = true

  step {
    uddi {
      name_in_zone = "ptr-test-1"
      rdata        = { dname = "domain.com." }
      zone         = "dns/auth_zone/0060207e-7664-4742-8eed-2d8e34db0035"
    }
    check = {
      "uddi.rdata.dname" = "domain.com."
    }
  }

  step {
    uddi {
      name_in_zone = "ptr-test-1"
      rdata        = { dname = "apple.com." }
      zone         = "dns/auth_zone/0060207e-7664-4742-8eed-2d8e34db0035"
    }
    check = {
      "uddi.rdata.dname" = "apple.com."
    }
  }

}
