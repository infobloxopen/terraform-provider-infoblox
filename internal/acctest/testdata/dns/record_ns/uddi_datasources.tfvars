# Auto-generated datasource acceptance-test cases for RecordNs.
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
      zone  = "dns/auth_zone/0060207e-7664-4742-8eed-2d8e34db0035"
      rdata = { dname = "example.com." }
    }
  }

}
