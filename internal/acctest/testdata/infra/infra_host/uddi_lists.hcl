case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    uddi {
      display_name = "{{random}}"
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
  min_tf_version = "1.14.0"

  step {
    uddi {
      display_name = "{{random}}"
    }
  }

  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        display_name = "uddi.display_name"
      }
    }
  }

}

case "tag_filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    uddi {
      display_name = "{{random}}"
      tags = { Site = "{{random2}}" }
    }
  }

  step {
    query    = true
    provider = infoblox
    include_resource = true
    filter {
      type   = "tag_filters"
      values = {
        Site = "uddi.tags.Site"
      }
    }
  }

}
