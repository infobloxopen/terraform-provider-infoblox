case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ name = "Microsoft 365" }]
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ name = "Microsoft 365" }]
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = { name = "uddi.name" }
    }
  }

}

case "tag_filters" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ name = "Microsoft 365" }]
      tags     = { Site = "{{random2}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "tag_filters"
      values = { Site = "uddi.tags.Site" }
    }
  }

}
