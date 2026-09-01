# Auto-generated resource acceptance-test cases for RecordRpzTxt.
case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name    = "txt-record.${infoblox_zone_rp.test.nios.fqdn}"
      text    = "Record Text"
      rp_zone = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.text"    = "Record Text"
      "nios.name"    = "txt-record.{{random}}.com"
      "nios.rp_zone" = "{{random}}.com"
      "nios.view"    = "default"
      "nios.disable" = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name    = "txt-record.${infoblox_zone_rp.test.nios.fqdn}"
      text    = "Record Text"
      rp_zone = infoblox_zone_rp.test.nios.fqdn
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name    = "txt-record.${infoblox_zone_rp.test.nios.fqdn}"
      text    = "Record Text"
      rp_zone = infoblox_zone_rp.test.nios.fqdn
      comment = "This is a new rpz txt record"
    }
    check = {
      "nios.comment" = "This is a new rpz txt record"
    }
  }

  step {
    nios {
      name    = "txt-record.${infoblox_zone_rp.test.nios.fqdn}"
      text    = "Record Text"
      rp_zone = infoblox_zone_rp.test.nios.fqdn
      comment = "This is an updated rpz txt record"
    }
    check = {
      "nios.comment" = "This is an updated rpz txt record"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name    = "txt-record.${infoblox_zone_rp.test.nios.fqdn}"
      text    = "Record Text"
      rp_zone = infoblox_zone_rp.test.nios.fqdn
      disable = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name    = "txt-record.${infoblox_zone_rp.test.nios.fqdn}"
      text    = "Record Text"
      rp_zone = infoblox_zone_rp.test.nios.fqdn
      disable = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "txt-record.${infoblox_zone_rp.test.nios.fqdn}"
      text      = "Record Text"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      name      = "txt-record.${infoblox_zone_rp.test.nios.fqdn}"
      text      = "Record Text"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ext_attrs = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name    = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      text    = "Record Text"
      rp_zone = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.name" = "{{random2}}.{{random}}.com"
    }
  }

  step {
    nios {
      name    = "{{random3}}.${infoblox_zone_rp.test.nios.fqdn}"
      text    = "Record Text"
      rp_zone = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.name" = "{{random3}}.{{random}}.com"
    }
  }

}

case "rp_zone" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name    = "txt-record.${infoblox_zone_rp.test.nios.fqdn}"
      text    = "Record Text"
      rp_zone = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.rp_zone" = "{{random}}.com"
    }
  }

}

case "text" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name    = "txt-record.${infoblox_zone_rp.test.nios.fqdn}"
      text    = "Record Text"
      rp_zone = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.text" = "Record Text"
    }
  }

  step {
    nios {
      name    = "txt-record.${infoblox_zone_rp.test.nios.fqdn}"
      text    = "Updated Record Text"
      rp_zone = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.text" = "Updated Record Text"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name    = "txt-record.${infoblox_zone_rp.test.nios.fqdn}"
      text    = "Record Text"
      rp_zone = infoblox_zone_rp.test.nios.fqdn
      ttl     = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name    = "txt-record.${infoblox_zone_rp.test.nios.fqdn}"
      text    = "Record Text"
      rp_zone = infoblox_zone_rp.test.nios.fqdn
      ttl     = 0
    }
    check = {
      "nios.ttl" = "0"
    }
  }

}

case "view" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "custom_view" {
    nios = {
      name = "{{random3}}"
    }
  }
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
      view = infoblox_view.custom_view.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name    = "txt-record.${infoblox_zone_rp.test.nios.fqdn}"
      text    = "Record Text"
      rp_zone = infoblox_zone_rp.test.nios.fqdn
      view    = infoblox_view.custom_view.nios.name
    }
    check = {
      "nios.view" = "{{random3}}"
    }
  }

}
