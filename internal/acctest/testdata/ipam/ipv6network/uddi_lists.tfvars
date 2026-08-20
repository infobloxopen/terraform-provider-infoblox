# Ipv6network — uddi list cases
case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    }
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
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "filters"
      values = {
        address = "uddi.address"
        space   = "uddi.space"
      }
    }
  }

}

case "tag_filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address = "{{random_ipv6}}"
      cidr    = 128
      # space   = infoblox_ip_space.test.id
      space = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      tags  = { tag1 = "{{random}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type   = "tag_filters"
      values = {
        tag1 = "uddi.tags.tag1"
      }
    }
  }

}
