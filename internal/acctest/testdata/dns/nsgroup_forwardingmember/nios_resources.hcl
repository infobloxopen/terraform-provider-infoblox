# Auto-generated resource acceptance-test cases for NsgroupForwardingmember.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name               = "{{random}}"
      forwarding_servers = [{ name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.name"                      = "{{random}}"
      "nios.forwarding_servers.0.name" = "{{grid_master_hostname}}"
      "nios.comment"                   = ""
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
      name               = "{{random}}"
      forwarding_servers = [{ name = "{{grid_master_hostname}}" }]
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name               = "{{random}}"
      forwarding_servers = [{ name = "{{grid_master_hostname}}" }]
      comment            = "this is an comment"
    }
    check = {
      "nios.comment" = "this is an comment"
    }
  }

  step {
    nios {
      name               = "{{random}}"
      forwarding_servers = [{ name = "{{grid_master_hostname}}" }]
      comment            = "this is an updated comment"
    }
    check = {
      "nios.comment" = "this is an updated comment"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name               = "{{random}}"
      forwarding_servers = [{ name = "{{grid_master_hostname}}" }]
      ext_attrs          = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name               = "{{random}}"
      forwarding_servers = [{ name = "{{grid_master_hostname}}" }]
      ext_attrs          = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "forwarding_servers" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name               = "{{random}}"
      forwarding_servers = [{ name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.forwarding_servers.0.name" = "{{grid_master_hostname}}"
    }
  }

  step {
    nios {
      name               = "{{random}}"
      forwarding_servers = [{ name = "{{grid_member_hostname}}", use_override_forwarders = true, forward_to = [{ name = "forwarder.com", address = "2.3.4.5" }] }]
    }
    check = {
      "nios.forwarding_servers.0.name" = "{{grid_member_hostname}}"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name               = "{{random}}"
      forwarding_servers = [{ name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name               = "{{random2}}"
      forwarding_servers = [{ name = "{{grid_master_hostname}}" }]
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}
