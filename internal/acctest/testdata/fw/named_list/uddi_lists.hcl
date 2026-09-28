# NamedList — uddi list cases
case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    uddi {
      name            = "{{random}}"
      items_described = [{ item = "{{random2}}.com", description = "Example Domain" }]
      type            = "custom_list"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "tag_filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    uddi {
      name            = "{{random}}"
      items_described = [{ item = "{{random2}}.com", description = "Example Domain" }]
      type            = "custom_list"
      tags            = { tag1 = "{{random3}}" }
    }
  }

  step {
    query             = true
    provider          = infoblox
    include_resource  = true
    filter {
      type   = "tag_filters"
      values = {
        tag1 = "uddi.tags.tag1"
      }
    }
  }

}
