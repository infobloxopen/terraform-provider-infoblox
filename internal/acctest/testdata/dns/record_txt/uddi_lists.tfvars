# RecordTxt — uddi list cases
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name_in_zone = "{{random}}"
      rdata        = { text = "sample text" }
      zone         = "dns/auth_zone/491ca52a-b154-4411-a684-0faf1d118719"
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
      name_in_zone = "{{random}}"
      rdata        = { text = "sample text" }
      zone         = "dns/auth_zone/491ca52a-b154-4411-a684-0faf1d118719"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name_in_zone = "uddi.name_in_zone"
        zone         = "uddi.zone"
      }
    }
  }

}

case "tag_filters" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name_in_zone = "{{random}}"
      rdata        = { text = "sample text" }
      zone         = "dns/auth_zone/491ca52a-b154-4411-a684-0faf1d118719"
      tags         = { tag1 = "{{random2}}" }
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
