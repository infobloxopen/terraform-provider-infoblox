# Auto-generated list acceptance-test cases for RecordCaa.
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
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
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag  = 0
      ca_tag   = "issue"
      ca_value = "digicert.com"
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
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ca_flag   = 0
      ca_tag    = "issue"
      ca_value  = "digicert.com"
      ext_attrs = { Site = "{{random3}}" }
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
