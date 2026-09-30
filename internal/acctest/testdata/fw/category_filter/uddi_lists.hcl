case "filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    uddi {
      name       = "{{random}}"
      categories = ["College"]
    }
  }

  step {
    query    = true
    provider = infoblox
    filter {
      type = "filters"
      values = {
        name = "uddi.name"
      }
    }
    include_resource = true
  }
}

case "tag_filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    uddi {
      name       = "{{random}}"
      categories = ["College"]
      tags       = { tag1 = "{{random2}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "tag_filters"
      values = {
        tag1 = "uddi.tags.tag1"
      }
    }
  }
}
