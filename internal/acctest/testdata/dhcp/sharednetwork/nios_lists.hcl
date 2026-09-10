# Sharednetwork — nios list cases
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "201.1.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network.id }]
      network_view = infoblox_network_view.test_view.nios.name
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
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "201.1.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network.id }]
      network_view = infoblox_network_view.test_view.nios.name
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name         = "nios.name"
        network_view = "nios.network_view"
      }
    }
  }

}

case "ext_attr_filters" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_network" "test_network" {
    nios = {
      network      = "201.1.0.0/24"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name         = "{{random}}"
      networks     = [{ ref = infoblox_network.test_network.id }]
      network_view = infoblox_network_view.test_view.nios.name
      ext_attrs    = { Site = "{{random2}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "ext_attr_filters"
      values = {
        Site = "nios.ext_attrs.Site"
      }
    }
  }

}
