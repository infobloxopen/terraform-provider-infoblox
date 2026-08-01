# Auto-generated list acceptance-test cases for RecordPtr.
case "basic" {
  backend = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      ipv4addr = "192.168.104.22"
      ptrdname = "{{random}}.example.com"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      ipv4addr = "192.168.104.22"
      ptrdname = "{{random}}.example.com"
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
  backend = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      name      = "22.104.168.192.in-addr.arpa"
      ptrdname  = "{{random}}.example.com"
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
  backend = "nios"
  min_tf_version = "1.14.0"

  step {
    nios {
      ipv4addr = "192.168.104.22"
      ptrdname = "{{random}}.example.com"
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
