# Auto-generated resource acceptance-test cases for Range.
case "basic" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space = infoblox_network_view.test.id
      start = "10.0.0.8"
      end   = "10.0.0.20"
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.end"          = "10.0.0.20"
      "uddi.start"        = "10.0.0.8"
      "uddi.disable_dhcp" = "false"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space = infoblox_network_view.test.id
      start = "10.0.0.8"
      end   = "10.0.0.20"
    }
    depends_on = [infoblox_network.test]
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space   = infoblox_network_view.test.id
      start   = "10.0.0.8"
      end     = "10.0.0.20"
      comment = "this range is created by terraform"
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.comment" = "this range is created by terraform"
    }
  }

  step {
    uddi {
      space   = infoblox_network_view.test.id
      start   = "10.0.0.8"
      end     = "10.0.0.20"
      comment = "this range was created by terraform"
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.comment" = "this range was created by terraform"
    }
  }

}

case "disable_dhcp" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space        = infoblox_network_view.test.id
      start        = "10.0.0.8"
      end          = "10.0.0.20"
      disable_dhcp = true
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.disable_dhcp" = "true"
    }
  }

  step {
    uddi {
      space        = infoblox_network_view.test.id
      start        = "10.0.0.8"
      end          = "10.0.0.20"
      disable_dhcp = false
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.disable_dhcp" = "false"
    }
  }

}

# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: requires_resource: infoblox_dhcp_option_group not yet implemented
case "dhcp_options" {
  backend     = "uddi"
  skip        = true
  skip_reason = "requires_resource: infoblox_dhcp_option_group not yet implemented"
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_network_view.test.id
    }
  }
  resource "infoblox_dhcp_optiondefinition" "test" {
    uddi = {
      code = 234
      name = "test_dhcp_option_code"
      option_space = infoblox_dhcp_optionspace.test.id
      type = "boolean"
    }
  }
  resource "infoblox_dhcp_option_group_unknown" "test" {
    uddi = {
      name = "og-{{random}}"
      protocol = "ip4"
    }
  }
  resource "infoblox_dhcp_optionspace" "test" {
    uddi = {
      name = "{{random}}"
      protocol = "ip4"
    }
  }
  PREREQ

  step {
    uddi {
      space        = infoblox_network_view.test.id
      start        = "10.0.0.10"
      end          = "10.0.0.20"
      dhcp_options = [{ type = "option", option_code = infoblox_dhcp_optiondefinition.test.id, option_value = true }]
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.dhcp_options.#"              = "1"
      "uddi.dhcp_options.0.option_value" = "true"
    }
  }

  step {
    uddi {
      space        = infoblox_network_view.test.id
      start        = "10.0.0.10"
      end          = "10.0.0.20"
      dhcp_options = [{ type = "group", group = infoblox_dhcp_option_group_unknown.test.id }]
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.dhcp_options.#" = "1"
    }
  }

}

case "end" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space = infoblox_network_view.test.id
      start = "10.0.0.8"
      end   = "10.0.0.20"
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.end" = "10.0.0.20"
    }
  }

  step {
    uddi {
      space = infoblox_network_view.test.id
      start = "10.0.0.8"
      end   = "10.0.0.29"
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.end" = "10.0.0.29"
    }
  }

}

case "exclusion_ranges" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space            = infoblox_network_view.test.id
      start            = "10.0.0.8"
      end              = "10.0.0.20"
      exclusion_ranges = [{ end = "10.0.0.16", start = "10.0.0.12" }]
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.exclusion_ranges.0.start" = "10.0.0.12"
      "uddi.exclusion_ranges.0.end"   = "10.0.0.16"
    }
  }

  step {
    uddi {
      space            = infoblox_network_view.test.id
      start            = "10.0.0.8"
      end              = "10.0.0.20"
      exclusion_ranges = [{ end = "10.0.0.16", start = "10.0.0.14" }]
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.exclusion_ranges.0.start" = "10.0.0.14"
      "uddi.exclusion_ranges.0.end"   = "10.0.0.16"
    }
  }

}

case "inheritance_sources" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      start               = "10.0.0.8"
      end                 = "10.0.0.20"
      space               = infoblox_network_view.test.id
      inheritance_sources = { dhcp_options = { action = "inherit" } }
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.inheritance_sources.dhcp_options.action" = "inherit"
    }
  }

  step {
    uddi {
      start               = "10.0.0.8"
      end                 = "10.0.0.20"
      space               = infoblox_network_view.test.id
      inheritance_sources = { dhcp_options = { action = "block" } }
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.inheritance_sources.dhcp_options.action" = "block"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space = infoblox_network_view.test.id
      start = "10.0.0.8"
      end   = "10.0.0.20"
      name  = "range-test"
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.name" = "range-test"
    }
  }

  step {
    uddi {
      space = infoblox_network_view.test.id
      start = "10.0.0.8"
      end   = "10.0.0.20"
      name  = "range-test-1"
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.name" = "range-test-1"
    }
  }

}

case "space" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "one" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network_view" "two" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 8
      space = infoblox_network_view.one.id
    }
  }
  PREREQ

  step {
    uddi {
      space = infoblox_network_view.one.id
      start = "10.0.0.8"
      end   = "10.0.0.20"
    }
    depends_on = [infoblox_network.test]
  }

  step {
    uddi {
      space = infoblox_network_view.one.id
      start = "10.0.0.8"
      end   = "10.0.0.20"
    }
    depends_on = [infoblox_network.test]
  }

}

case "start" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space = infoblox_network_view.test.id
      start = "10.0.0.8"
      end   = "10.0.0.20"
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.start" = "10.0.0.8"
    }
  }

  step {
    uddi {
      space = infoblox_network_view.test.id
      start = "10.0.0.12"
      end   = "10.0.0.20"
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.start" = "10.0.0.12"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_network" "test" {
    uddi = {
      address = "10.0.0.0"
      cidr = 24
      space = infoblox_network_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      space = infoblox_network_view.test.id
      start = "10.0.0.8"
      end   = "10.0.0.20"
      tags  = { site = "NA" }
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.tags.site" = "NA"
    }
  }

  step {
    uddi {
      space = infoblox_network_view.test.id
      start = "10.0.0.8"
      end   = "10.0.0.20"
      tags  = { site = "CA" }
    }
    depends_on = [infoblox_network.test]
    check = {
      "uddi.tags.site" = "CA"
    }
  }

}
