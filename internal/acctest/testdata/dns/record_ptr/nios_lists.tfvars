# Auto-generated list acceptance-test cases for RecordPtr.
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "reverse" {
    nios = {
      fqdn        = "192.168.121.0/24"
      zone_format = "IPV4"
      view        = "default"
    }
  }
  PREREQ

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      ipv4addr = "192.168.121.10"
      ptrdname = "host.{{random}}.com"
      view     = "default"
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
  resource "infoblox_zone_auth" "reverse" {
    nios = {
      fqdn        = "192.168.122.0/24"
      zone_format = "IPV4"
      view        = "default"
    }
  }
  PREREQ

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      ipv4addr = "192.168.122.10"
      ptrdname = "host.{{random}}.com"
      view     = "default"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        ptrdname = "nios.ptrdname"
      }
    }
  }

}

case "ext_attr_filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "reverse" {
    nios = {
      fqdn        = "192.168.123.0/24"
      zone_format = "IPV4"
      view        = "default"
    }
  }
  PREREQ

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      ipv4addr  = "192.168.123.10"
      ptrdname  = "host.{{random}}.com"
      view      = "default"
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

case "creator_filter" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "reverse" {
    nios = {
      fqdn        = "192.168.124.0/24"
      zone_format = "IPV4"
      view        = "default"
    }
  }
  PREREQ

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      ipv4addr = "192.168.124.10"
      ptrdname = "host.{{random}}.com"
      view     = "default"
      creator  = "STATIC"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        ptrdname = "nios.ptrdname"
        creator  = "nios.creator"
      }
    }
  }

}
