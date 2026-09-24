case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "10.0.0.240"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
    }
    depends_on = [infoblox_network.test]
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "10.0.0.241"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
      name        = "{{random2}}"
    }
    depends_on = [infoblox_network.test]
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        address = "uddi.address"
      }
    }
  }

}

case "tag_filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr    = 24
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      ip_space    = infoblox_network_view.test.id
      address     = "10.0.0.242"
      match_type  = "mac"
      match_value = "aa:aa:aa:aa:aa:aa"
      tags        = { tag1 = "{{random}}" }
    }
    depends_on = [infoblox_network.test]
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "tag_filters"
      values = {
        tag1 = "uddi.tags.tag1"
      }
    }
  }

}