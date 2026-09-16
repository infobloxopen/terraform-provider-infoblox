# RecordCaa — uddi list cases
# No legacy list test was found for this object.
# Add list cases here manually.
case "basic" {
  backend           = "uddi"
  min_tf_version    = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { tag = "issue", value = "ca.example.com" }
      zone  = infoblox_zone_auth.test.id
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend           = "uddi"
  min_tf_version    = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      name_in_zone = "{{random2}}"
      zone         = infoblox_zone_auth.test.id
      rdata        = { tag = "issue", value = "ca.example.com" }
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
  backend           = "uddi"
  min_tf_version    = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { tag = "issue", value = "ca.example.com" }
      zone  = infoblox_zone_auth.test.id
      tags  = { tag1 = "{{random3}}" }
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
