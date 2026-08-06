# Auto-generated list acceptance-test cases for RecordNs.
case "basic" {
  backend           = "nios"
  min_tf_version    = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name       = infoblox_zone_auth.test_zone.nios.fqdn
      nameserver = "{{random2}}.${infoblox_zone_auth.test_zone.nios.fqdn}"
      addresses  = [{ address = "{{random_ip}}", auto_create_ptr = false }]
      view       = infoblox_zone_auth.test_zone.nios.view
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend           = "nios"
  min_tf_version    = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name       = infoblox_zone_auth.test_zone.nios.fqdn
      nameserver = "{{random2}}.${infoblox_zone_auth.test_zone.nios.fqdn}"
      addresses  = [{ address = "{{random_ip}}", auto_create_ptr = false }]
      view       = infoblox_zone_auth.test_zone.nios.view
    }
  }

  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        name       = "nios.name"
        nameserver = "nios.nameserver"
      }
    }
  }

}
