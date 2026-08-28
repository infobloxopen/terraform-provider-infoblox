# Auto-generated resource acceptance-test cases for DtcServer.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      host = "{{random_ip}}"
    }
    check = {
      "nios.name"                    = "{{random}}"
      "nios.host"                    = "{{random_ip}}"
      "nios.auto_create_host_record" = "true"
      "nios.disable"                 = "false"
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
      name = "{{random}}"
      host = "{{random_ip}}"
    }
  }

}

case "auto_create_host_record" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                    = "{{random}}"
      host                    = "{{random_ip}}"
      auto_create_host_record = false
    }
    check = {
      "nios.auto_create_host_record" = "false"
    }
  }

  step {
    nios {
      name                    = "{{random}}"
      host                    = "{{random_ip}}"
      auto_create_host_record = true
    }
    check = {
      "nios.auto_create_host_record" = "true"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      host    = "{{random_ip}}"
      comment = "initial comment"
    }
    check = {
      "nios.comment" = "initial comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      host    = "{{random_ip}}"
      comment = "updated comment"
    }
    check = {
      "nios.comment" = "updated comment"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      host    = "{{random_ip}}"
      disable = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      host    = "{{random_ip}}"
      disable = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      host      = "{{random_ip}}"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      host      = "{{random_ip}}"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "host" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      host = "{{random_ip}}"
    }
    check = {
      "nios.host" = "{{random_ip}}"
    }
  }

  step {
    nios {
      name = "{{random}}"
      host = "{{random_ip2}}"
    }
    check = {
      "nios.host" = "{{random_ip2}}"
    }
  }

}

case "monitors" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}"
      host     = "{{random_ip}}"
      monitors = [
        { host = "3.2.2.2",   monitor = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHBz:https" },
        { host = "3.231.2.2", monitor = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http" },
      ]
    }
    check = {
      "nios.monitors.#"         = "2"
      "nios.monitors.0.host"    = "3.2.2.2"
      "nios.monitors.0.monitor" = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHBz:https"
      "nios.monitors.1.host"    = "3.231.2.2"
      "nios.monitors.1.monitor" = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHA:http"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      host     = "{{random_ip}}"
      monitors = [
        { host = "3.2.2.2", monitor = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHBz:https" },
      ]
    }
    check = {
      "nios.monitors.#"         = "1"
      "nios.monitors.0.host"    = "3.2.2.2"
      "nios.monitors.0.monitor" = "dtc:monitor:http/ZG5zLmlkbnNfbW9uaXRvcl9odHRwJGh0dHBz:https"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      host = "{{random_ip}}"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
      host = "{{random_ip}}"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "sni_hostname" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name         = "{{random}}"
      host         = "{{random_ip}}"
      sni_hostname = "{{random2}}"
    }
    check = {
      "nios.sni_hostname" = "{{random2}}"
    }
  }

  step {
    nios {
      name         = "{{random}}"
      host         = "{{random_ip}}"
      sni_hostname = "{{random3}}-update"
    }
    check = {
      "nios.sni_hostname" = "{{random3}}-update"
    }
  }

}
