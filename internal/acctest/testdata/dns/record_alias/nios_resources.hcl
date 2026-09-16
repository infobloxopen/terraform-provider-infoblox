# Auto-generated resource acceptance-test cases for RecordAlias.
case "basic" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.name"        = "{{random2}}.{{random}}.com"
      "nios.target_name" = "server.example.com"
      "nios.target_type" = "A"
      "nios.view"        = "default"
      "nios.creator"     = "STATIC"
      "nios.disable"     = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl     = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
    }
  }

}

case "comment" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
      comment     = "This is a sample comment."
    }
    check = {
      "nios.comment" = "This is a sample comment."
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
      comment     = "This is an updated comment."
    }
    check = {
      "nios.comment" = "This is an updated comment."
    }
  }

}

case "creator" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
      creator     = "STATIC"
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

}

case "disable" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
      disable     = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
      disable     = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "ext_attrs" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
      ext_attrs   = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
      ext_attrs   = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

}

case "name" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.name" = "{{random2}}.{{random}}.com"
    }
  }

  step {
    nios {
      name        = "{{random3}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.name" = "{{random3}}.{{random}}.com"
    }
  }

}

case "target_name" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.target_name" = "server.example.com"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "updated-server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.target_name" = "updated-server.example.com"
    }
  }

}

case "target_type" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.target_type" = "A"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "AAAA"
      view        = infoblox_zone_auth.test.nios.view
    }
    check = {
      "nios.target_type" = "AAAA"
    }
  }

}

case "ttl" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
      ttl         = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_zone_auth.test.nios.view
      ttl         = 0
    }
    check = {
      "nios.ttl" = "0"
    }
  }

}

case "view" {
  backend           = "nios"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "view1" {
    nios = {
      name = "{{random3}}"
    }
  }
  resource "infoblox_view" "view2" {
    nios = {
      name = "{{random4}}"
    }
  }
  resource "infoblox_zone_auth" "zone1" {
    nios = {
      fqdn = "{{random}}.com"
      view = infoblox_view.view1.nios.name
    }
  }
  resource "infoblox_zone_auth" "zone2" {
    nios = {
      fqdn = "{{random}}.com"
      view = infoblox_view.view2.nios.name
    }
  }
  PREREQ

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.zone1.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_view.view1.nios.name
    }
    check = {
      "nios.view" = "{{random3}}"
    }
  }

  step {
    nios {
      name        = "{{random2}}.${infoblox_zone_auth.zone2.nios.fqdn}"
      target_name = "server.example.com"
      target_type = "A"
      view        = infoblox_view.view2.nios.name
    }
    check = {
      "nios.view" = "{{random4}}"
    }
  }

}
