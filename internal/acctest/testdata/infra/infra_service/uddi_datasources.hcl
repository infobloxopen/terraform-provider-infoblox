case "filters" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  filter {
    type   = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.description", "uddi.desired_state", "uddi.name", "uddi.pool_id", "uddi.service_type"]

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "dns"
    }
  }

}

case "tag_filters" {
  backend           = "uddi"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  filter {
    type   = "tag_filters"
    values = {
      env = "uddi.tags.env"
    }
  }

  pair_checks = ["uddi.description", "uddi.desired_state", "uddi.name", "uddi.pool_id", "uddi.service_type"]

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "dns"
      tags         = { env = "{{random2}}" }
    }
  }

}
