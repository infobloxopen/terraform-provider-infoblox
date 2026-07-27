# Auto-generated datasource acceptance-test cases for RecordA.
case "filters" {
  backend = "uddi"

  filter {
    type = "filters"
    values = {
      name_in_zone = "uddi.name_in_zone"
      zone         = "uddi.zone"
    }
  }

  step {
    uddi {

      # zone  = infoblox_zone_auth.test.id
      rdata        = { address = "{{random_ip}}" }
      zone         = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      name_in_zone = "{{random}}"
    }
  }

}

case "tag_filters" {
  backend = "uddi"

  filter {
    type = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  step {
    uddi {
      # zone  = infoblox_zone_auth.test.id
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      tags  = { tag1 = "{{random2}}" }
    }
  }

}
