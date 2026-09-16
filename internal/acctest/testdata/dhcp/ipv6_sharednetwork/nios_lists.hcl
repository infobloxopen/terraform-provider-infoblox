# Ipv6sharednetwork — nios list cases
case "basic" {
  backend        = "nios"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
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
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random}}"
      networks = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
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
  resource "infoblox_ipv6_network" "test1" {
    nios = {
      network = "{{random_ipv6_network}}"
    }
  }
  resource "infoblox_ipv6_network" "test2" {
    nios = {
      network = "{{random_ipv6_network2}}"
    }
  }
  PREREQ

  step {
    nios {
      name      = "{{random}}"
      networks  = ["${infoblox_ipv6_network.test1.id}", "${infoblox_ipv6_network.test2.id}"]
      ext_attrs = { Site = "{{random2}}" }
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
