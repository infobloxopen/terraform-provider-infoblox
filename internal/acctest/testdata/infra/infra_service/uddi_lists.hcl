case "basic" {
  backend           = "uddi"
  parallel          = true
  min_tf_version    = "1.14.0"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "dns"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend           = "uddi"
  parallel          = true
  min_tf_version    = "1.14.0"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_infra_host" "test" {
      uddi = {
        display_name  = "{{random}}"
        serial_number = "{{random_int}}"
        tags          = { "host/serial_number" = "{{random_int}}" }
      }
    }
    PREREQ

  step {
    uddi {
      name         = "{{random}}"
      pool_id      = infoblox_infra_host.test.uddi.pool_id
      service_type = "dns"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter           = "name=='{{random}}'"
  }

}
