# Auto-generated list acceptance-test cases for RecordDname.
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random2}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name   = infoblox_zone_auth.test.nios.fqdn
      target = "{{random}}.com"
      view   = infoblox_zone_auth.test.nios.view
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random2}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name   = infoblox_zone_auth.test.nios.fqdn
      target = "{{random}}.com"
      view   = infoblox_zone_auth.test.nios.view
    }
  }

  step {
    query    = true
    provider = infoblox
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
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random3}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = infoblox_zone_auth.test.nios.fqdn
      target    = "{{random}}.com"
      ext_attrs = { Site = "{{random4}}" }
      view      = infoblox_zone_auth.test.nios.view
    }
  }

  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "ext_attr_filters"
      values = {
        Site = "nios.ext_attrs.Site"
      }
    }
  }

}
