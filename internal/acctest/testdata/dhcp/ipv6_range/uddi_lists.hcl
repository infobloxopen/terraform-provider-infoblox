# Ipv6range — uddi list cases
case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr    = 64
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space = infoblox_network_view.test.id
      start = "2001:db8:{{random_hextet}}:{{random_int}}::8"
      end   = "2001:db8:{{random_hextet}}:{{random_int}}::20"
    }
    depends_on = [infoblox_ipv6_network.test]
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
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr    = 64
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space = infoblox_network_view.test.id
      start = "2001:db8:{{random_hextet}}:{{random_int}}::8"
      end   = "2001:db8:{{random_hextet}}:{{random_int}}::20"
      comment = "{{random2}}"
    }
    depends_on = [infoblox_ipv6_network.test]
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        start   = "uddi.start"
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
  resource "infoblox_ipv6_network" "test" {
    uddi = {
      address = "2001:db8:{{random_hextet}}:{{random_int}}::"
      cidr    = 64
      space   = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space = infoblox_network_view.test.id
      start = "2001:db8:{{random_hextet}}:{{random_int}}::8"
      end   = "2001:db8:{{random_hextet}}:{{random_int}}::20"
      tags  = { tag1 = "{{random2}}" }
    }
    depends_on = [infoblox_ipv6_network.test]
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
