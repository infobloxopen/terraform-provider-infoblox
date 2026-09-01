# Auto-generated list acceptance-test cases for ZoneDelegated.
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}", address = "10.0.0.1" }]
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
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}", address = "10.0.0.1" }]
    }
  }

  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        fqdn = "nios.fqdn"
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
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      fqdn        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      delegate_to = [{ name = "{{random3}}", address = "10.0.0.1" }]
      ext_attrs   = { Site = "{{random4}}" }
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
