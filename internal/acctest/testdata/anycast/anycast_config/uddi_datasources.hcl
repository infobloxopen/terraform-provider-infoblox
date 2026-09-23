case "filters" {
  backend = "uddi"

  filter {
    type = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.comment", "uddi.compartment_id", "uddi.name"]

  step {
    uddi {
      name               = "{{random}}"
            service            = "NTP"
            anycast_ip_address = "{{random_ip}}"
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type   = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.account_id", "uddi.anycast_ip_address", "uddi.anycast_ipv6_address", "uddi.created_at", "uddi.description", "uddi.is_configured", "uddi.name", "uddi.runtime_status", "uddi.service", "uddi.updated_at"]

  step {
    uddi {
      anycast_ip_address = "{{random_ip}}"
      name               = "{{random}}"
      service            = "DNS"
      tags               = { tag1 = "value1" }
    }
  }

}
