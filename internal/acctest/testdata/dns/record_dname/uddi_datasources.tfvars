# Auto-generated datasource acceptance-test cases for RecordDname.
case "filters" {
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      name_in_zone = "uddi.name_in_zone"
      zone         = "uddi.zone"
    }
  }

  step {
    uddi {
      name_in_zone = "{{random2}}"
      zone         = infoblox_zone_auth.test.id
      rdata        = { target = "example.com." }
    }
  }

}
