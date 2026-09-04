# DtcLbdn — uddi list cases
case "basic" {
  backend        = "uddi"
  parallel       = true
  min_tf_version = "1.14.0"

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
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
      name = "dtc-lbdn-{{random}}."
      view = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = { name = "uddi.name" }
    }
  }

}

case "tag_filters" {
  backend        = "uddi"
  parallel       = true
  min_tf_version = "1.14.0"

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      tags = { Site = "{{random2}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "tag_filters"
      values = { Site = "uddi.tags.Site" }
    }
  }

}
