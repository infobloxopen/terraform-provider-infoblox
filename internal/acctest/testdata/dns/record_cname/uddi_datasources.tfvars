# Auto-generated datasource acceptance-test cases for RecordCname.
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
      zone  = infoblox_zone_auth.test.id
      rdata = { cname = "c" }
    }
  }

}
