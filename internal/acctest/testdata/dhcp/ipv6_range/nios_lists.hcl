# Ipv6range — nios list cases
case "basic" {
  backend           = "nios"
  min_tf_version    = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_ipv6_network" "test_network" {
    nios = {
      network      = "14::/64"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test_network.nios.network
      start_addr   = "14::140"
      end_addr     = "14::141"
      network_view = infoblox_ipv6_network.test_network.nios.network_view
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend           = "nios"
  min_tf_version    = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_ipv6_network" "test_network" {
    nios = {
      network      = "14::/64"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test_network.nios.network
      start_addr   = "14::142"
      end_addr     = "14::143"
      network_view = infoblox_ipv6_network.test_network.nios.network_view
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        start_addr   = "nios.start_addr"
        network_view = "nios.network_view"
      }
    }
  }

}

case "ext_attr_filters" {
  backend           = "nios"
  min_tf_version    = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test_view" {
    nios = {
      name = "{{random_view}}"
    }
  }
  resource "infoblox_ipv6_network" "test_network" {
    nios = {
      network      = "14::/64"
      network_view = infoblox_network_view.test_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      network      = infoblox_ipv6_network.test_network.nios.network
      start_addr   = "14::144"
      end_addr     = "14::145"
      network_view = infoblox_ipv6_network.test_network.nios.network_view
      ext_attrs    = { Site = "{{random}}" }
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
