# RecordNaptr — uddi list cases
case "basic" {
  backend  = "uddi"
  parallel = true
  min_tf_version = "1.14.0"
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
      rdata        = { order = 10, preference = 10, replacement = ".", services = "SIP+D2U" }
      zone         = infoblox_zone_auth.test.id
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
  min_tf_version = "1.14.0"
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
      rdata        = { order = 10, preference = 10, replacement = ".", services = "SIP+D2U" }
      zone         = infoblox_zone_auth.test.id
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
  min_tf_version = "1.14.0"
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
      rdata        = { order = 10, preference = 10, replacement = ".", services = "SIP+D2U" }
      zone         = infoblox_zone_auth.test.id
      tags         = { tag1 = "{{random3}}" }
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
