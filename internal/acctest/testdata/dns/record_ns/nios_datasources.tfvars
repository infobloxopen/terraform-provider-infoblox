# Auto-generated datasource acceptance-test cases for RecordNs.
case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
      view = "default"
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      name       = "nios.name"
      nameserver = "nios.nameserver"
    }
  }

  step {
    nios {
      name       = "${infoblox_zone_auth.test_zone.nios.fqdn}"
      nameserver = "{{random2}}.${infoblox_zone_auth.test_zone.nios.fqdn}"
      addresses  = [{address = "20.0.0.0", auto_create_ptr = false}]
      view       = "default"
    }
    depends_on = [infoblox_zone_auth.test_zone]
  }

}
