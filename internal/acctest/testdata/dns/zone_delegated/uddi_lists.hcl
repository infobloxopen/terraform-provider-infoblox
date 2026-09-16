# ZoneDelegated — uddi list cases

case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn         = "{{random3}}.com."
      primary_type = "cloud"
      view         = infoblox_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "{{random2}}.{{random3}}.com."
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn         = "{{random3}}.com."
      primary_type = "cloud"
      view         = infoblox_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "{{random2}}.{{random3}}.com."
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
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
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn         = "{{random3}}.com."
      primary_type = "cloud"
      view         = infoblox_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      fqdn               = "{{random2}}.{{random3}}.com."
      delegation_servers = [{ address = "12.0.0.0", fqdn = "ns1.com." }]
      tags               = { tag1 = "{{random2}}" }
      view               = infoblox_view.test.id
    }
    depends_on = [infoblox_view.test, infoblox_zone_auth.test]
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
