# Hand-authored list acceptance-test cases for RecordRpzPtr.
case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
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
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = { ptrdname = "nios.ptrdname" }
    }
  }

}

case "ext_attr_filters" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      ptrdname  = "{{random2}}.{{random}}.com"
      ipv4addr  = "{{random_ip}}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ext_attrs = { Site = "{{random3}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "ext_attr_filters"
      values = { Site = "nios.ext_attrs.Site" }
    }
  }

}
