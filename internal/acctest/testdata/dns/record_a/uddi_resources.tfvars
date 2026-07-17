# Auto-generated resource acceptance-test cases for RecordA (uddi).
case "basic" {
  # basic — generated from terraform-provider-uddi
  backend = "uddi"
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
      rdata = { address = "10.0.0.15" }
      zone  = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.address" = "10.0.0.15"
    }
  }

}

case "disappears" {
  # disappears — generated from terraform-provider-uddi
  backend = "uddi"
  disappears = true
  expect_non_empty_plan = true
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
      rdata = { address = "10.0.0.15" }
      zone  = infoblox_zone_auth.test.id
    }
  }

}

case "comment" {
  # comment — generated from terraform-provider-uddi
  backend = "uddi"
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
      rdata   = { address = "10.0.0.1" }
      zone    = infoblox_zone_auth.test.id
      comment = "some comment"
    }
    check = {
      "uddi.comment" = "some comment"
    }
  }

  step {
    uddi {
      rdata   = { address = "10.0.0.1" }
      zone    = infoblox_zone_auth.test.id
      comment = "updated comment"
    }
    check = {
      "uddi.comment" = "updated comment"
    }
  }

}

case "disabled" {
  # disabled — generated from terraform-provider-uddi
  backend = "uddi"
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
      rdata    = { address = "10.0.0.1" }
      zone     = infoblox_zone_auth.test.id
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

  step {
    uddi {
      rdata    = { address = "10.0.0.1" }
      zone     = infoblox_zone_auth.test.id
      disabled = false
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

}

case "inheritance_sources" {
  # inheritance_sources — generated from terraform-provider-uddi
  backend = "uddi"
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
      rdata               = { address = "10.0.0.1" }
      zone                = infoblox_zone_auth.test.id
      inheritance_sources = { ttl = { action = "inherit" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "inherit"
    }
  }

  step {
    uddi {
      rdata               = { address = "10.0.0.1" }
      zone                = infoblox_zone_auth.test.id
      inheritance_sources = { ttl = { action = "override" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "override"
    }
  }

}

case "name_in_zone" {
  # name_in_zone — generated from terraform-provider-uddi
  backend = "uddi"
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
      rdata        = { address = "10.0.0.1" }
      zone         = infoblox_zone_auth.test.id
      name_in_zone = "xyz"
    }
    check = {
      "uddi.name_in_zone" = "xyz"
    }
  }

  step {
    uddi {
      rdata        = { address = "10.0.0.1" }
      zone         = infoblox_zone_auth.test.id
      name_in_zone = "abc"
    }
    check = {
      "uddi.name_in_zone" = "abc"
    }
  }

}

case "rdata" {
  # rdata — generated from terraform-provider-uddi
  backend = "uddi"
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
      rdata = { address = "10.0.0.1" }
      zone  = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.address" = "10.0.0.1"
    }
  }

  step {
    uddi {
      rdata = { address = "10.0.0.2" }
      zone  = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.address" = "10.0.0.2"
    }
  }

}

case "tags" {
  # tags — generated from terraform-provider-uddi
  backend = "uddi"
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
      rdata = { address = "10.0.0.1" }
      zone  = infoblox_zone_auth.test.id
      tags  = { tag1 = "value1" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
    }
  }

  step {
    uddi {
      rdata = { address = "10.0.0.1" }
      zone  = infoblox_zone_auth.test.id
      tags  = { tag1 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value2"
    }
  }

}

case "ttl" {
  # ttl — generated from terraform-provider-uddi
  backend = "uddi"
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
      rdata = { address = "10.0.0.1" }
      zone  = infoblox_zone_auth.test.id
      ttl   = 60
    }
    check = {
      "uddi.ttl" = "60"
    }
  }

  step {
    uddi {
      rdata = { address = "10.0.0.1" }
      zone  = infoblox_zone_auth.test.id
      ttl   = 90
    }
    check = {
      "uddi.ttl" = "90"
    }
  }

}

case "view" {
  # view — generated from terraform-provider-uddi
  backend = "uddi"
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
      fqdn = "test.com."
      view = infoblox_view.one.id
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata              = { address = "10.0.0.1" }
      absolute_name_spec = "a.test.com."
      view               = infoblox_view.one.id
    }
  }

  step {
    uddi {
      rdata              = { address = "10.0.0.1" }
      absolute_name_spec = "a.test.com."
      view               = infoblox_view.two.id
    }
  }

}

case "zone" {
  # zone — generated from terraform-provider-uddi
  backend = "uddi"
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
      rdata = { address = "10.0.0.1" }
      zone  = infoblox_zone_auth.one.id
    }
  }

  step {
    uddi {
      rdata = { address = "10.0.0.1" }
      zone  = infoblox_zone_auth.two.id
    }
  }

}

case "options" {
  # options — generated from terraform-provider-uddi
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "{{random2}}.com."
      primary_type = "cloud"
      view = infoblox_view.test.id
    }
  }
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  resource "infoblox_zone_auth" "rmz" {
    uddi = {
      fqdn = "10.in-addr.arpa."
      primary_type = "cloud"
      view = infoblox_view.test.id
    }
  }
  PREREQ

  step {
    uddi {
      rdata   = { address = "10.0.0.1" }
      options = { create_ptr = true, check_rmz = true }
      zone    = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.options.create_ptr" = "true"
      "uddi.options.check_rmz"  = "true"
    }
  }

  step {
    uddi {
      rdata   = { address = "10.0.0.1" }
      options = { create_ptr = true, check_rmz = false }
      zone    = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.options.create_ptr" = "true"
      "uddi.options.check_rmz"  = "false"
    }
  }

  step {
    uddi {
      rdata   = { address = "10.0.0.1" }
      options = { create_ptr = false, check_rmz = false }
      zone    = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.options.create_ptr" = "false"
      "uddi.options.check_rmz"  = "false"
    }
  }

}
