# Address — uddi list cases
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
      address = "10.0.0.1"
      space   = infoblox_network_view.test.id
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
      address = "10.0.0.1"
      space   = infoblox_network_view.test.id
      comment = "{{random}}"
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
        comment = "uddi.comment"
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
      address = "10.0.0.1"
      space   = infoblox_network_view.test.id
      tags    = { tag1 = "{{random}}" }
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
