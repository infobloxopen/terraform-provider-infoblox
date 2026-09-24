# Hand-authored resource acceptance-test cases for RecordRpzPtr.
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
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.ipv4addr" = "{{random_ip}}"
      "nios.ptrdname" = "{{random2}}.{{random}}.com"
      "nios.rp_zone"  = "{{random}}.com"
      "nios.view"     = "default"
      "nios.disable"  = "false"
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
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
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
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
      comment  = "This is a new rpz ptr record"
    }
    check = {
      "nios.comment" = "This is a new rpz ptr record"
    }
  }

  step {
    nios {
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
      comment  = "This is a updated rpz ptr record"
    }
    check = {
      "nios.comment" = "This is a updated rpz ptr record"
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
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
      disable  = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
      disable  = true
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
      ptrdname  = "{{random2}}.{{random}}.com"
      ipv4addr  = "{{random_ip}}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      ptrdname  = "{{random2}}.{{random}}.com"
      ipv4addr  = "{{random_ip}}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ext_attrs = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

}

case "ipv4addr" {
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
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.ipv4addr" = "{{random_ip}}"
    }
  }

  step {
    nios {
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip_2}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.ipv4addr" = "{{random_ip_2}}"
    }
  }

}

case "ipv6addr" {
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
      ptrdname = "{{random2}}.{{random}}.com"
      ipv6addr = "{{random_ipv6}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.ipv6addr" = "{{random_ipv6}}"
    }
  }

  step {
    nios {
      ptrdname = "{{random2}}.{{random}}.com"
      ipv6addr = "{{random_ipv6_2}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.ipv6addr" = "{{random_ipv6_2}}"
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
      ptrdname = "{{random2}}.{{random}}.com"
      name     = "1.0.10.10.in-addr.arpa.{{random}}.com"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.name" = "1.0.10.10.in-addr.arpa.{{random}}.com"
    }
  }

  step {
    nios {
      ptrdname = "{{random2}}.{{random}}.com"
      name     = "2.0.10.10.in-addr.arpa.{{random}}.com"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.name" = "2.0.10.10.in-addr.arpa.{{random}}.com"
    }
  }

}

case "ptrdname" {
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
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.ptrdname" = "{{random2}}.{{random}}.com"
    }
  }

  step {
    nios {
      ptrdname = "{{random3}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.ptrdname" = "{{random3}}.{{random}}.com"
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
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.rp_zone" = "{{random}}.com"
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
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
      ttl      = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
      ttl      = 0
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
      ptrdname = "{{random2}}.{{random}}.com"
      ipv4addr = "{{random_ip}}"
      rp_zone  = infoblox_zone_rp.test.nios.fqdn
      view     = infoblox_view.custom_view.nios.name
    }
    check = {
      "nios.view" = "{{random3}}"
    }
  }

}
