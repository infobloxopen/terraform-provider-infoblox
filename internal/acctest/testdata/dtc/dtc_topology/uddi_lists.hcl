# Auto-generated list acceptance-test cases for DtcTopology (UDDI backend).
case "basic" {
  backend        = "uddi"
  parallel       = true
  min_tf_version = "1.14.0"

  step {
    uddi {
      name    = "topology-{{random}}"
      sources = [{ name = "src-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] }]
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
      name    = "topology-{{random}}"
      sources = [{ name = "src-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] }]
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
      name    = "topology-{{random}}"
      sources = [{ name = "src-{{random}}", source = "subnet", subnets = ["10.0.0.0/8"] }]
      tags    = { Site = "{{random2}}" }
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
