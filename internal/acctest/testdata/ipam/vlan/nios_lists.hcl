# Auto-generated list acceptance-test cases for Vlan.
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_vlanview" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      id     = 71
      name   = "{{random}}"
      parent = infoblox_vlanview.test.id
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
  resource "infoblox_vlanview" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      id     = 72
      name   = "{{random}}"
      parent = infoblox_vlanview.test.id
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
  resource "infoblox_vlanview" "test" {
    nios = {
      name          = "{{random2}}"
      start_vlan_id = 50
      end_vlan_id   = 100
    }
  }
  PREREQ

  step {
    nios {
      id        = 73
      name      = "{{random}}"
      parent    = infoblox_vlanview.test.id
      ext_attrs = { Site = "{{random3}}" }
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
