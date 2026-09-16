# Auto-generated list acceptance-test cases for RecordRpzCnameIpaddress.
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  parallel       = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }
}

case "filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  parallel       = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
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
  backend        = "nios"
  min_tf_version = "1.14.0"
  parallel       = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "10.0.0.1.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "10.0.0.1"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
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
