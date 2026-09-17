# DtcMonitorPdp — uddi list cases

case "basic" {
  backend        = "uddi"
  parallel       = true
  min_tf_version = "1.14.0"

  step {
    uddi {
      name = "dtc-monitor-pdp-{{random}}"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend        = "uddi"
  parallel       = true
  min_tf_version = "1.14.0"

  step {
    uddi {
      name = "dtc-monitor-pdp-{{random}}"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name = "uddi.name"
      }
    }
  }

}

case "tag_filters" {
  backend        = "uddi"
  parallel       = true
  min_tf_version = "1.14.0"

  step {
    uddi {
      name = "dtc-monitor-pdp-{{random}}"
      tags = { Site = "{{random2}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "tag_filters"
      values = {
        Site = "uddi.tags.Site"
      }
    }
  }

}
