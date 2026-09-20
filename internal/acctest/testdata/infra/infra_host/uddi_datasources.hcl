case "filters" {
  backend  = "uddi"
  parallel = true

  filter {
    type = "filters"
    values = {
      display_name = "uddi.display_name"
    }
  }

  pair_checks = ["uddi.description", "uddi.display_name", "uddi.ip_space", "uddi.location_id", "uddi.maintenance_mode", "uddi.pool_id", "uddi.serial_number"]

  step {
    uddi {
      display_name = "{{random}}"
    }
  }
}

case "tag_filters" {
  backend  = "uddi"
  parallel = true

  filter {
    type = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.description", "uddi.display_name", "uddi.ip_space", "uddi.location_id", "uddi.maintenance_mode", "uddi.pool_id", "uddi.serial_number"]

  step {
    uddi {
      display_name = "{{random}}"
      tags = {
        tag1 = "{{random2}}"
      }
    }
  }
}
