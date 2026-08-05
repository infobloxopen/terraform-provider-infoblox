# Auto-generated list acceptance-test cases for RecordPtr.
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
      ipv4addr = "192.168.104.40"
      ptrdname = "host.{{random}}.com"
    }
  }

//need to verify this 
  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        ptrdname = "nios.ptrdname"
      }
    }
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
      ipv4addr = "192.168.104.41"
      ptrdname = "host.{{random}}.com"
    }
  }

  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "filters"
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
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "42.104.168.192.in-addr.arpa"
      ptrdname  = "host.{{random}}.com"
      view      = "default"
      ext_attrs = { Site = "{{random2}}" }
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

case "creator_filter" {
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
      ipv4addr = "192.168.104.43"
      ptrdname = "host.{{random}}.com"
    }
  }

  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        ptrdname = "nios.ptrdname"
        creator  = "nios.creator"
      }
    }
  }

}
