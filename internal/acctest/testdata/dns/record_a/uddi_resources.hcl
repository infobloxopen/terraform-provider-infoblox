# Auto-generated resource acceptance-test cases for RecordA.
case "basic" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.address" = "{{random_ip}}"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  skip                  = true
  skip_reason           = "Test Skipped due to inconsistent error codes returned by the API [NORTHSTAR-12575]"
  prerequisites_hcl     = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = infoblox_zone_auth.test.id
    }
  }

}

case "comment" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata   = { address = "{{random_ip}}" }
      zone    = infoblox_zone_auth.test.id
      comment = "some comment"
    }
    check = {
      "uddi.comment" = "some comment"
    }
  }

  step {
    uddi {
      rdata   = { address = "{{random_ip}}" }
      zone    = infoblox_zone_auth.test.id
      comment = "updated comment"
    }
    check = {
      "uddi.comment" = "updated comment"
    }
  }

}

case "disabled" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata    = { address = "{{random_ip}}" }
      zone     = infoblox_zone_auth.test.id
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

  step {
    uddi {
      rdata    = { address = "{{random_ip}}" }
      zone     = infoblox_zone_auth.test.id
      disabled = false
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

}

case "inheritance_sources" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata               = { address = "{{random_ip}}" }
      zone                = infoblox_zone_auth.test.id
      inheritance_sources = { ttl = { action = "inherit" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "inherit"
    }
  }

  step {
    uddi {
      rdata               = { address = "{{random_ip}}" }
      zone                = infoblox_zone_auth.test.id
      inheritance_sources = { ttl = { action = "override" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "override"
    }
  }

}

case "name_in_zone" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata        = { address = "{{random_ip}}" }
      zone         = infoblox_zone_auth.test.id
      name_in_zone = "xyz"
    }
    check = {
      "uddi.name_in_zone" = "xyz"
    }
  }

  step {
    uddi {
      rdata        = { address = "{{random_ip}}" }
      zone         = infoblox_zone_auth.test.id
      name_in_zone = "abc"
    }
    check = {
      "uddi.name_in_zone" = "abc"
    }
  }

}

case "rdata" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.address" = "{{random_ip}}"
    }
  }

  step {
    uddi {
      rdata = { address = "{{random_ip2}}" }
      zone  = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.address" = "{{random_ip2}}"
    }
  }

}

case "tags" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = infoblox_zone_auth.test.id
      tags  = { tag1 = "value1" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
    }
  }

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = infoblox_zone_auth.test.id
      tags  = { tag1 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value2"
    }
  }

}

case "ttl" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = infoblox_zone_auth.test.id
      ttl   = 60
    }
    check = {
      "uddi.ttl" = "60"
    }
  }

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = infoblox_zone_auth.test.id
      ttl   = 90
    }
    check = {
      "uddi.ttl" = "90"
    }
  }

}

case "view" {
  backend  = "uddi"
  parallel = true
  step {
    prerequisites_hcl = <<-PREREQ
    resource "infoblox_view" "one" {
      uddi = {
        name = "{{random}}"
      }
    }
    resource "infoblox_view" "two" {
      uddi = {
        name = "{{random2}}"
      }
    }
    resource "infoblox_zone_auth" "test" {
      uddi = {
        fqdn = "{{random3}}.com."
        view = infoblox_view.one.id
        primary_type = "cloud"
      }
    }
    PREREQ

    uddi {
      rdata              = { address = "{{random_ip}}" }
      absolute_name_spec = "a.{{random3}}.com."
      view               = infoblox_view.one.id
    }
    depends_on = [infoblox_zone_auth.test]
  }

  step {
    prerequisites_hcl = <<-PREREQ
    resource "infoblox_view" "one" {
      uddi = {
        name = "{{random}}"
      }
    }
    resource "infoblox_view" "two" {
      uddi = {
        name = "{{random2}}"
      }
    }
    resource "infoblox_zone_auth" "test" {
      uddi = {
        fqdn = "{{random3}}.com."
        view = infoblox_view.two.id
        primary_type = "cloud"
      }
    }
    PREREQ

    uddi {
      rdata              = { address = "{{random_ip}}" }
      absolute_name_spec = "a.{{random3}}.com."
      view               = infoblox_view.two.id
    }
    depends_on = [infoblox_zone_auth.test]
  }

}

case "zone" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "one" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  resource "infoblox_zone_auth" "two" {
    uddi = {
      fqdn = "{{random2}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = infoblox_zone_auth.one.id
    }
  }

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = infoblox_zone_auth.two.id
    }
  }

}

case "options" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random2}}.com."
      primary_type = "cloud"
    }
  }
  resource "infoblox_zone_auth" "rmz" {
    uddi = {
      fqdn = "12.in-addr.arpa."
      primary_type = "cloud"
   }
  }
  PREREQ

  step {
    uddi {
      rdata   = { address = "{{random_ip}}" }
      options = { create_ptr = true, check_rmz = true }
      zone    = infoblox_zone_auth.test.id
    }
    depends_on = [infoblox_zone_auth.rmz, infoblox_zone_auth.test]
    check = {
      "uddi.options.create_ptr" = "true"
      "uddi.options.check_rmz"  = "true"
    }
  }

  step {
    uddi {
      rdata   = { address = "{{random_ip}}" }
      options = { create_ptr = true, check_rmz = false }
      zone    = infoblox_zone_auth.test.id
    }
    depends_on = [infoblox_zone_auth.rmz, infoblox_zone_auth.test]
    check = {
      "uddi.options.create_ptr" = "true"
      "uddi.options.check_rmz"  = "false"
    }
  }

  step {
    uddi {
      rdata   = { address = "{{random_ip}}" }
      options = { create_ptr = false, check_rmz = false }
      zone    = infoblox_zone_auth.test.id
    }
    depends_on = [infoblox_zone_auth.rmz, infoblox_zone_auth.test]
    check = {
      "uddi.options.create_ptr" = "false"
      "uddi.options.check_rmz"  = "false"
    }
  }

}
