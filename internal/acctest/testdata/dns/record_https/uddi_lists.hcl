case "basic" {
  backend           = "uddi"
  parallel          = true
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
      rdata        = { target_name = "{{random}}.com" }
      zone         = infoblox_zone_auth.test.id
      name_in_zone = "{{random2}}"
    }
  }

  # HTTPS records require a zone+name filter here: the UDDI API validates
  # target_name server-side on every list response, so an unfiltered query
  # 400s if any orphaned HTTPS record on the shared grid has an empty
  # target_name. No other record type triggers this behaviour.
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

case "filters" {
  backend           = "uddi"
  parallel          = true
  skip_if_env_empty = ["UDDI_AUTH_ZONE_ID_1"]
  skip_reason       = "UDDI_AUTH_ZONE_ID_1 environment variable must be set for this test to run"
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
      rdata        = { target_name = "{{random}}.com" }
      zone         = infoblox_zone_auth.test.id
      name_in_zone = "{{random2}}"
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
  backend           = "uddi"
  parallel          = true
  skip_if_env_empty = ["UDDI_AUTH_ZONE_ID_1"]
  skip_reason       = "UDDI_AUTH_ZONE_ID_1 environment variable must be set for this test to run"
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
      rdata = { target_name = "{{random}}.com" }
      zone  = "{{uddi_auth_zone_id_1}}"
      tags  = { tag1 = "{{random2}}" }
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
