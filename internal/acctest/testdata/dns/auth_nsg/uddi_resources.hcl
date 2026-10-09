# Auto-generated resource acceptance-test cases for AuthNsg.
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    uddi {
      name = "{{random}}"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      comment = "COMMENT_REPLACE_ME"
    }
    check = {
      "uddi.comment" = "COMMENT_REPLACE_ME"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      comment = "COMMENT_UPDATE_REPLACE_ME"
    }
    check = {
      "uddi.comment" = "COMMENT_UPDATE_REPLACE_ME"
    }
  }

}

case "external_primaries" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_auth_nsg" "one" {
    uddi = {
      name = "{{random2}}"
    }
  }
  PREREQ

  step {
    uddi {
      name               = "{{random}}"
      external_primaries = [{ address = "1.1.1.1", fqdn = "a.com.", type = "primary" }, { nsg = infoblox_auth_nsg.one.id, type = "nsg" }]
    }
    check = {
      "uddi.external_primaries.#"         = "2"
      "uddi.external_primaries.0.address" = "1.1.1.1"
      "uddi.external_primaries.0.fqdn"    = "a.com."
      "uddi.external_primaries.0.type"    = "primary"
      "uddi.external_primaries.1.type"    = "nsg"
    }
  }

  step {
    uddi {
      name               = "{{random}}"
      external_primaries = [{ address = "2.2.2.2", fqdn = "b.com.", type = "primary" }, { nsg = infoblox_auth_nsg.one.id, type = "nsg" }]
    }
    check = {
      "uddi.external_primaries.#"         = "2"
      "uddi.external_primaries.0.address" = "2.2.2.2"
      "uddi.external_primaries.0.fqdn"    = "b.com."
      "uddi.external_primaries.0.type"    = "primary"
      "uddi.external_primaries.1.type"    = "nsg"
    }
  }

}

case "external_secondaries" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                 = "{{random}}"
      external_secondaries = [{ address = "1.1.1.1", fqdn = "a.com." }]
    }
    check = {
      "uddi.external_secondaries.#"         = "1"
      "uddi.external_secondaries.0.address" = "1.1.1.1"
      "uddi.external_secondaries.0.fqdn"    = "a.com."
    }
  }

  step {
    uddi {
      name                 = "{{random}}"
      external_secondaries = [{ address = "2.2.2.2", fqdn = "b.com." }]
    }
    check = {
      "uddi.external_secondaries.#"         = "1"
      "uddi.external_secondaries.0.address" = "2.2.2.2"
      "uddi.external_secondaries.0.fqdn"    = "b.com."
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name = "nsg2"
    }
    check = {
      "uddi.name" = "nsg2"
    }
  }

}

case "nsgs" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_auth_nsg" "nsg1" {
    uddi = {
      name = "{{random2}}"
    }
  }
  resource "infoblox_auth_nsg" "nsg2" {
    uddi = {
      name = "{{random3}}"
    }
  }
  PREREQ

  step {
    uddi {
      name = "{{random}}"
      nsgs = [infoblox_auth_nsg.nsg1.id]
    }
    check = {
      "uddi.nsgs.#" = "1"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      nsgs = [infoblox_auth_nsg.nsg1.id, infoblox_auth_nsg.nsg2.id]
    }
    check = {
      "uddi.nsgs.#" = "2"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      tags = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      tags = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}
