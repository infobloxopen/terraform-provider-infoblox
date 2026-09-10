# DtcMonitorPdp — uddi resource cases

case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "dtc-monitor-pdp-{{random}}"
    }
    check = {
      "uddi.name"       = "dtc-monitor-pdp-{{random}}"
      "uddi.disabled"   = "false"
      "uddi.interval"   = "15"
      "uddi.port"       = "2123"
      "uddi.retry_down" = "1"
      "uddi.retry_up"   = "1"
      "uddi.timeout"    = "10"
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
      name = "dtc-monitor-pdp-{{random}}"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "dtc-monitor-pdp-{{random}}"
      comment = "This is a comment"
    }
    check = {
      "uddi.comment" = "This is a comment"
    }
  }

  step {
    uddi {
      name    = "dtc-monitor-pdp-{{random}}"
      comment = "This is an updated comment"
    }
    check = {
      "uddi.comment" = "This is an updated comment"
    }
  }

}

case "disabled" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "dtc-monitor-pdp-{{random}}"
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

  step {
    uddi {
      name     = "dtc-monitor-pdp-{{random}}"
      disabled = false
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

}

case "interval" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "dtc-monitor-pdp-{{random}}"
      interval = 10
    }
    check = {
      "uddi.interval" = "10"
    }
  }

  step {
    uddi {
      name     = "dtc-monitor-pdp-{{random}}"
      interval = 20
    }
    check = {
      "uddi.interval" = "20"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "dtc-monitor-pdp-{{random}}"
    }
    check = {
      "uddi.name" = "dtc-monitor-pdp-{{random}}"
    }
  }

  step {
    uddi {
      name = "dtc-monitor-pdp-{{random2}}"
    }
    check = {
      "uddi.name" = "dtc-monitor-pdp-{{random2}}"
    }
  }

}

case "port" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "dtc-monitor-pdp-{{random}}"
      port = 2314
    }
    check = {
      "uddi.port" = "2314"
    }
  }

  step {
    uddi {
      name = "dtc-monitor-pdp-{{random}}"
      port = 4321
    }
    check = {
      "uddi.port" = "4321"
    }
  }

}

case "retry_down" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "dtc-monitor-pdp-{{random}}"
      retry_down = 3
    }
    check = {
      "uddi.retry_down" = "3"
    }
  }

  step {
    uddi {
      name       = "dtc-monitor-pdp-{{random}}"
      retry_down = 5
    }
    check = {
      "uddi.retry_down" = "5"
    }
  }

}

case "retry_up" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "dtc-monitor-pdp-{{random}}"
      retry_up = 2
    }
    check = {
      "uddi.retry_up" = "2"
    }
  }

  step {
    uddi {
      name     = "dtc-monitor-pdp-{{random}}"
      retry_up = 4
    }
    check = {
      "uddi.retry_up" = "4"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "dtc-monitor-pdp-{{random}}"
      tags = { Site = "{{random2}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random2}}"
    }
  }

  step {
    uddi {
      name = "dtc-monitor-pdp-{{random}}"
      tags = { Site = "{{random3}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random3}}"
    }
  }

}

case "timeout" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "dtc-monitor-pdp-{{random}}"
      interval = 30
      timeout  = 15
    }
    check = {
      "uddi.timeout" = "15"
    }
  }

  step {
    uddi {
      name     = "dtc-monitor-pdp-{{random}}"
      interval = 30
      timeout  = 20
    }
    check = {
      "uddi.timeout" = "20"
    }
  }

}
