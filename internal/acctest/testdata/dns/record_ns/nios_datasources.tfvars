# Auto-generated datasource acceptance-test cases for RecordNs.
case "filters" {
  backend           = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
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

  pair_checks = ["nios.ms_delegation_name", "nios.name", "nios.nameserver", "nios.view"]

  step {
    nios {
      name       = infoblox_zone_auth.test_zone.nios.fqdn
      nameserver = "{{random2}}.${infoblox_zone_auth.test_zone.nios.fqdn}"
      addresses  = [{ address = "{{random_ip}}", auto_create_ptr = false }]
      view       = infoblox_zone_auth.test_zone.nios.view
    }
  }

}
