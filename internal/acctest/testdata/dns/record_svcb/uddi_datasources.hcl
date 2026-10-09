case "filters" {
  backend           = "uddi"
  skip_if_env_empty = ["UDDI_AUTH_ZONE_ID_1"]
  skip_reason       = "UDDI_AUTH_ZONE_ID_1 environment variable must be set for this test to run"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  filter {
    type = "filters"
    values = {
      name_in_zone = "uddi.name_in_zone"
      zone         = "uddi.zone"
    }
  }

  pair_checks = ["uddi.absolute_name_spec", "uddi.comment", "uddi.disabled", "uddi.name_in_zone", "uddi.ttl", "uddi.type", "uddi.view", "uddi.zone"]

  step {
    uddi {
      name_in_zone = "{{random2}}"
      zone         = "{{uddi_auth_zone_id_1}}"
      rdata        = { target_name = "example.com." }
    }
  }

}
