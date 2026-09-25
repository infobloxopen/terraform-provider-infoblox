# Auto-generated list acceptance-test cases for Vlanrange.
#
# Each case creates its own parent VLAN View; the range's start/end must fall
# within the view's start_vlan_id/end_vlan_id window.

case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        name = "nios.name"
      }
    }
  }

}

case "ext_attr_filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlan_view" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      name          = "{{random}}"
      start_vlan_id = 61
      end_vlan_id   = 71
      vlan_view     = infoblox_vlan_view.test.id
      ext_attrs     = { Site = "{{random3}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "ext_attr_filters"
      values = {
        Site = "nios.ext_attrs.Site"
      }
    }
  }

}
