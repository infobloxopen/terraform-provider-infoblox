# RecordNs — uddi list cases
# No legacy list test was found for this object.
# Add list cases here manually.

case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    uddi {
      name_in_zone = "{{random2}}"
      zone         = "dns/auth_zone/6ab79060-d813-4de7-be60-84f228583684"
      rdata        = { dname = "ns1.example.com" }
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
      name_in_zone = "{{random2}}"
      zone         = "dns/auth_zone/6ab79060-d813-4de7-be60-84f228583684"
      rdata        = { dname = "ns1.example.com" }
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
      }
    }
  }

}

case "tag_filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    uddi {
      name_in_zone = "{{random2}}"
      zone         = "dns/auth_zone/6ab79060-d813-4de7-be60-84f228583684"
      rdata        = { dname = "ns1.example.com" }
      tags         = { tag1 = "{{random4}}" }
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
