case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
    }
    check = {
      "nios.name"      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      "nios.canonical" = "test-cname.example.com"
      "nios.rp_zone"   = "rpz-test.infoblox.com"
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

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
    }
  }

}

case "canonical" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname1.example.com"
      rp_zone   = "rpz-test.infoblox.com"
    }
    check = {
      "nios.canonical" = "test-cname1.example.com"
    }
  }

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname2.example.com"
      rp_zone   = "rpz-test.infoblox.com"
    }
    check = {
      "nios.canonical" = "test-cname2.example.com"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
      comment   = "This is a new rpz cname client ipaddress dn record"
    }
    check = {
      "nios.comment" = "This is a new rpz cname client ipaddress dn record"
    }
  }

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
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

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
      disable   = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
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

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
    }
    check = {
      "nios.name" = "{{random_cidr_network}}.rpz-test.infoblox.com"
    }
  }

  step {
    nios {
      name      = "2001:db8::/32.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
    }
    check = {
      "nios.name" = "2001:db8::/32.rpz-test.infoblox.com"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
      ttl       = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
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

  step {
    nios {
      name      = "{{random_cidr_network}}.rpz-test.infoblox.com"
      canonical = "test-cname.example.com"
      rp_zone   = "rpz-test.infoblox.com"
    }
    check = {
      "nios.view" = "default"
    }
  }

}
