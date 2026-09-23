# TODO : Add Support for Filters for AnyCast Config - Requires Schema Update

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
