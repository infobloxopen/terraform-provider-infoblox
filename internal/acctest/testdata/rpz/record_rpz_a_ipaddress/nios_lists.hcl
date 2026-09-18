case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "10.20.0.0/16.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.20.0.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "10.20.1.0/24.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr = "10.20.1.1"
      rp_zone  = infoblox_zone_rp.test_zone.nios.fqdn
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name = "nios.name"
      }
    }
  }

}

case "ext_attr_filters" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test_zone" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "10.20.2.0/24.${infoblox_zone_rp.test_zone.nios.fqdn}"
      ipv4addr  = "10.20.2.1"
      rp_zone   = infoblox_zone_rp.test_zone.nios.fqdn
      ext_attrs = { Site = "{{random2}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "ext_attr_filters"
      values = {
        Site = "nios.ext_attrs.Site"
      }
    }
  }

}
