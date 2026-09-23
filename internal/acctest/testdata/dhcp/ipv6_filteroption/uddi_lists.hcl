# Ipv6filteroption — uddi list cases
case "basic" {
  backend        = "uddi"
  parallel = true
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "test_opt"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }
  PREREQ

  step {
    uddi {
      name = "{{random}}"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
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
  parallel = true
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "test_opt"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }
  PREREQ

  step {
    uddi {
      name = "{{random}}"
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
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
  backend        = "uddi"
  parallel = true
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_ipv6_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "test_opt"
      option_space = infoblox_ipv6_dhcp_optionspace.test.id
      type         = "boolean"
    }
  }
  PREREQ

  step {
    uddi {
      name = "{{random}}"
      tags = { tag1 = "{{random3}}" }
      rules = {
        match = "any"
        rules = [{
          compare      = "equals"
          option_code  = infoblox_ipv6_dhcp_optiondefinition.test.id
          option_value = "true"
        }]
      }
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
