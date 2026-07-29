# ZoneAuth — uddi list cases
# No legacy list test was found for this object.
# Add list cases here manually.

case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
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
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        fqdn = "uddi.fqdn"
      }
    }
  }

}

case "tag_filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"

  step {
    uddi {
      fqdn         = "{{random}}.com."
      primary_type = "cloud"
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
