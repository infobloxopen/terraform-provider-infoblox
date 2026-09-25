# Filteroption — uddi list cases
# No legacy list test was found for this object.
# Add list cases here manually.

case "basic" {
  backend           = "uddi"
  min_tf_version    = "1.14.0"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "prereq_code"
      option_space = infoblox_dhcp_optionspace.test.id
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
          option_code  = infoblox_dhcp_optiondefinition.test.id
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
  backend           = "uddi"
  min_tf_version    = "1.14.0"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "prereq_code"
      option_space = infoblox_dhcp_optionspace.test.id
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
          option_code  = infoblox_dhcp_optiondefinition.test.id
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
  backend           = "uddi"
  min_tf_version    = "1.14.0"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_dhcp_optiondefinition" "test" {
    uddi = {
      code         = 234
      name         = "prereq_code"
      option_space = infoblox_dhcp_optionspace.test.id
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
          option_code  = infoblox_dhcp_optiondefinition.test.id
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
