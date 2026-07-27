# Auto-generated resource acceptance-test cases for RecordA.
case "basic" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
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
  skip                  = true
  skip_reason           = "Test Skipped due to inconsistent error codes returned by the API [NORTHSTAR-12575]"
  # prerequisites_hcl     = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata   = { address = "{{random_ip}}" }
      zone    = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      comment = "some comment"
    }
    check = {
      "uddi.comment" = "some comment"
    }
  }

  step {
    uddi {
      rdata   = { address = "{{random_ip}}" }
      zone    = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      comment = "updated comment"
    }
    check = {
      "uddi.comment" = "updated comment"
    }
  }

}

case "disabled" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata    = { address = "{{random_ip}}" }
      zone     = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

  step {
    uddi {
      rdata    = { address = "{{random_ip}}" }
      zone     = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      disabled = false
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

}

case "inheritance_sources" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata               = { address = "{{random_ip}}" }
      zone                = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      inheritance_sources = { ttl = { action = "inherit" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "inherit"
    }
  }

  step {
    uddi {
      rdata               = { address = "{{random_ip}}" }
      zone                = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      inheritance_sources = { ttl = { action = "override" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "override"
    }
  }

}

case "name_in_zone" {
  backend = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata        = { address = "{{random_ip}}" }
      zone         = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      name_in_zone = "xyz"
    }
    check = {
      "uddi.name_in_zone" = "xyz"
    }
  }

  step {
    uddi {
      rdata        = { address = "{{random_ip}}" }
      zone         = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      name_in_zone = "abc"
    }
    check = {
      "uddi.name_in_zone" = "abc"
    }
  }

}

case "rdata" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
    }
    check = {
      "uddi.rdata.address" = "{{random_ip}}"
    }
  }

  step {
    uddi {
      rdata = { address = "{{random_ip2}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
    }
    check = {
      "uddi.rdata.address" = "{{random_ip2}}"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      tags  = { tag1 = "value1" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
    }
  }

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      tags  = { tag1 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value2"
    }
  }

}

case "ttl" {
  backend = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      ttl   = 60
    }
    check = {
      "uddi.ttl" = "60"
    }
  }

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      ttl   = 90
    }
    check = {
      "uddi.ttl" = "90"
    }
  }

}

case "view" {
  backend = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_view" "one" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # resource "infoblox_view" "two" {
  #   uddi = {
  #     name = "{{random2}}"
  #   }
  # }
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "test.com."
  #     view = infoblox_view.one.id
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata              = { address = "{{random_ip}}" }
      absolute_name_spec = "10.in-addr.arpa."
      view               = "dns/view/28b9c115-8d5f-416e-979f-e7e71d80a3a3"
    }
    check = {
      "uddi.view" = "dns/view/28b9c115-8d5f-416e-979f-e7e71d80a3a3"
    }
  }

  step {
    uddi {
      rdata              = { address = "{{random_ip}}" }
      absolute_name_spec = "zone-xoeivh.com."
      # view               = infoblox_view.two.id
      view = "dns/view/ce528cc5-7482-4278-835f-801fb4f884fe"
    }
    # depends_on = [infoblox_zone_auth.test]
    check = {
      "uddi.view" = "dns/view/ce528cc5-7482-4278-835f-801fb4f884fe"
    }
  }

}

case "zone" {
  backend = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "one" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # resource "infoblox_zone_auth" "two" {
  #   uddi = {
  #     fqdn = "{{random2}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      # zone  = infoblox_zone_auth.one.id
      zone = "dns/auth_zone/c75d3700-05b5-4ff8-a413-dfa0bcb5b020"
    }
    check = {
      "uddi.zone" = "dns/auth_zone/c75d3700-05b5-4ff8-a413-dfa0bcb5b020"
    }
  }

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      #  zone  = infoblox_zone_auth.two.id
      zone = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      ttl  = 90
    }
    check = {
      "uddi.zone" = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
    }
  }

}

case "options" {
  backend = "uddi"
  parallel = true
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random2}}.com."
  #     primary_type = "cloud"
  #     view = infoblox_view.test.id
  #   }
  # }
  # resource "infoblox_view" "test" {
  #   uddi = {
  #     name = "{{random}}"
  #   }
  # }
  # resource "infoblox_zone_auth" "rmz" {
  #   uddi = {
  #     fqdn = "10.in-addr.arpa."
  #     primary_type = "cloud"
  #     view = infoblox_view.test.id
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      # zone    = infoblox_zone_auth.test.id
      zone    = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      options = { create_ptr = true, check_rmz = true }
    }
    # depends_on = [infoblox_zone_auth.rmz, infoblox_zone_auth.test]
    check = {
      "uddi.options.create_ptr" = "true"
      "uddi.options.check_rmz"  = "true"
    }
  }

  step {
    uddi {
      rdata   = { address = "{{random_ip}}" }
      zone    = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      options = { create_ptr = true, check_rmz = false }
      # zone    = infoblox_zone_auth.test.id
    }
    # depends_on = [infoblox_zone_auth.rmz, infoblox_zone_auth.test]
    check = {
      "uddi.options.create_ptr" = "true"
      "uddi.options.check_rmz"  = "false"
    }
  }

  step {
    uddi {
      rdata   = { address = "{{random_ip}}" }
      zone    = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      options = { create_ptr = false, check_rmz = false }
      # zone    = infoblox_zone_auth.test.id
    }
    # depends_on = [infoblox_zone_auth.rmz, infoblox_zone_auth.test]
    check = {
      "uddi.options.create_ptr" = "false"
      "uddi.options.check_rmz"  = "false"
    }
  }

}
