# Networkcontainer — uddi list cases
# No legacy list test was found for this object.
# Add list cases here manually.
case "basic" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      name    = "{{random}}"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name = "uddi.name"
      }
    }
  }

}

case "tag_filters" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_ip_space" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      address = "{{random_ipv4_network}}"
      cidr    = 16
      space   = "ipam/ip_space/1fd490b2-8847-11f1-a8d8-2a72d414108a"
      tags    = { tag1 = "{{random2}}" }
    }
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
