# Auto-generated datasource acceptance-test cases for Vlanview.
case "filters" {
  backend = "nios"

  filter {
    type   = "filters"
    values = {
      end_vlan_id = "nios.end_vlan_id"
      name        = "nios.name"
    }
  }

  pair_checks = ["nios.allow_range_overlapping", "nios.comment", "nios.end_vlan_id", "nios.name", "nios.pre_create_vlan", "nios.start_vlan_id", "nios.vlan_name_prefix"]

  step {
    nios {
      end_vlan_id   = 15
      name          = "{{random}}"
      start_vlan_id = 10
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.allow_range_overlapping", "nios.comment", "nios.end_vlan_id", "nios.name", "nios.pre_create_vlan", "nios.start_vlan_id", "nios.vlan_name_prefix"]

  step {
    nios {
      end_vlan_id   = 15
      name          = "{{random}}"
      start_vlan_id = 10
      ext_attrs     = { Site = "{{random}}" }
    }
  }

}
