# Auto-generated datasource acceptance-test cases for Vlanrange.
#
# Each case creates its own parent VLAN View; the range's start/end must fall
# within the view's start_vlan_id/end_vlan_id window.

case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      name = "nios.name"
    }
  }

  pair_checks = ["nios.comment", "nios.end_vlan_id", "nios.name", "nios.pre_create_vlan", "nios.start_vlan_id", "nios.vlan_name_prefix", "nios.vlan_view"]

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.end_vlan_id", "nios.name", "nios.pre_create_vlan", "nios.start_vlan_id", "nios.vlan_name_prefix", "nios.vlan_view"]

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
      ext_attrs     = { Site = "{{random3}}" }
    }
  }

}
