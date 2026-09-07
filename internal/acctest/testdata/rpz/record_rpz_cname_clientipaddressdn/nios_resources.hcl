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
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.name"      = "{{random_cidr_network}}.{{random}}.com"
      "nios.canonical" = "{{random2}}.{{random}}.com"
      "nios.rp_zone"   = "{{random}}.com"
      "nios.view"      = "default"
      "nios.disable"   = "false"
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
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
  }

}

case "canonical" {
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
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.canonical" = "{{random2}}.{{random}}.com"
    }
  }

  step {
    nios {
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random3}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.canonical" = "{{random3}}.{{random}}.com"
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
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      comment   = "This is a new rpz cname client ipaddress dn record"
    }
    check = {
      "nios.comment" = "This is a new rpz cname client ipaddress dn record"
    }
  }

  step {
    nios {
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      comment   = "This is an updated rpz cname client ipaddress dn record"
    }
    check = {
      "nios.comment" = "This is an updated rpz cname client ipaddress dn record"
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
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      disable   = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      disable   = true
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
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
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
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.name" = "{{random_cidr_network}}.{{random}}.com"
    }
  }

  step {
    nios {
      name      = "2001:db8::/32.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.name" = "2001:db8::/32.{{random}}.com"
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
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ttl       = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
      ttl       = 0
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
  resource "infoblox_zone_rp" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "{{random_cidr_network}}.${infoblox_zone_rp.test.nios.fqdn}"
      canonical = "{{random2}}.${infoblox_zone_rp.test.nios.fqdn}"
      rp_zone   = infoblox_zone_rp.test.nios.fqdn
    }
    check = {
      "nios.view" = "default"
    }
  }

}
