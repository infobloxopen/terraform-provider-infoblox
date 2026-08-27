# Auto-generated resource acceptance-test cases for RecordRpzNaptr.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
    }
    check = {
      "nios.name"        = "{{random2}}.{{random}}.com"
      "nios.rp_zone"     = "{{random}}.com"
      "nios.order"       = "10"
      "nios.preference"  = "10"
      "nios.replacement" = "."
      "nios.view"        = "default"
      "nios.disable"     = "false"
      "nios.services"    = ""
      "nios.flags"       = ""
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
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      comment     = "This is a new rpz naptr record"
    }
    check = {
      "nios.comment" = "This is a new rpz naptr record"
    }
  }

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      comment     = "This is an updated rpz naptr record"
    }
    check = {
      "nios.comment" = "This is an updated rpz naptr record"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      disable     = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      disable     = true
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
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      ext_attrs   = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      ext_attrs   = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

}

case "flags" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      flags       = "U"
    }
    check = {
      "nios.flags" = "U"
    }
  }

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      flags       = "S"
    }
    check = {
      "nios.flags" = "S"
    }
  }

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      flags       = "P"
    }
    check = {
      "nios.flags" = "P"
    }
  }

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      flags       = "A"
    }
    check = {
      "nios.flags" = "A"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
    }
    check = {
      "nios.name" = "{{random2}}.{{random}}.com"
    }
  }

  step {
    nios {
      name        = "{{random3}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
    }
    check = {
      "nios.name" = "{{random3}}.{{random}}.com"
    }
  }

}

case "order" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
    }
    check = {
      "nios.order" = "10"
    }
  }

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 20
      preference  = 10
      replacement = "."
    }
    check = {
      "nios.order" = "20"
    }
  }

}

case "preference" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
    }
    check = {
      "nios.preference" = "10"
    }
  }

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 20
      replacement = "."
    }
    check = {
      "nios.preference" = "20"
    }
  }

}

case "regexp" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      regexp      = ""
    }
    check = {
      "nios.regexp" = ""
    }
  }

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      regexp      = "!^.*$!sip:jdoe@corpxyz.com!"
    }
    check = {
      "nios.regexp" = "!^.*$!sip:jdoe@corpxyz.com!"
    }
  }

}

case "replacement" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
    }
    check = {
      "nios.replacement" = "."
    }
  }

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "test.com"
    }
    check = {
      "nios.replacement" = "test.com"
    }
  }

}

case "rp_zone" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
    }
    check = {
      "nios.rp_zone" = "{{random}}.com"
    }
  }

}

case "services" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      services    = "http+E2U"
    }
    check = {
      "nios.services" = "http+E2U"
    }
  }

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      services    = "SIPS+D2T"
    }
    check = {
      "nios.services" = "SIPS+D2T"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      ttl         = 10
    }
    check = {
      "nios.ttl" = "10"
    }
  }

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      ttl         = 0
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
  PREREQ

  step {
    nios {
      name        = "{{random2}}.{{random}}.com"
      rp_zone     = infoblox_zone_rp.test.nios.fqdn
      order       = 10
      preference  = 10
      replacement = "."
      view        = infoblox_view.custom_view.nios.name
    }
    check = {
      "nios.view" = "{{random3}}"
    }
  }

}
