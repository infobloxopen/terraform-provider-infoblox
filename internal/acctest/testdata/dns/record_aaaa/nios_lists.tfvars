# Auto-generated list acceptance-test cases for RecordAaaa.
case "basic" {
  backend           = "nios"
  min_tf_version    = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      view     = infoblox_zone_auth.test.nios.view
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
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr = "{{random_ipv6}}"
      view     = infoblox_zone_auth.test.nios.view
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        name = "nios.name"
      }
    }
  }

}

case "ext_attr_filters" {
  backend           = "nios"
  min_tf_version    = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "aaaa-record.${infoblox_zone_auth.test.nios.fqdn}"
      ipv6addr  = "{{random_ipv6}}"
      view      = infoblox_zone_auth.test.nios.view
      ext_attrs = { Site = "{{random2}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "ext_attr_filters"
      values = {
        Site = "nios.ext_attrs.Site"
      }
    }
  }

}
