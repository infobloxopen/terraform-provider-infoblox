# Auto-generated datasource acceptance-test cases for Vlan.
case "filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlanview" "test" {
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

  pair_checks = ["nios.comment", "nios.contact", "nios.department", "nios.description", "nios.reserved"]

  step {
    nios {
      id     = 99
      name   = "{{random}}"
      parent = infoblox_vlanview.test.id
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlanview" "test" {
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

  pair_checks = ["nios.comment", "nios.contact", "nios.department", "nios.description", "nios.reserved"]

  step {
    nios {
      id        = 100
      name      = "{{random}}"
      parent    = infoblox_vlanview.test.id
      ext_attrs = { Site = "{{random3}}" }
    }
  }

}
