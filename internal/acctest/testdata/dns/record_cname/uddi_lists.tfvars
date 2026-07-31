# RecordCname — uddi list cases
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name_in_zone = "{{random}}"
      rdata        = { cname = "target.example.com." }
      zone         = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
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
      rdata        = { cname = "target.example.com." }
      zone         = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
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
      rdata        = { cname = "target.example.com." }
      zone         = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
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
