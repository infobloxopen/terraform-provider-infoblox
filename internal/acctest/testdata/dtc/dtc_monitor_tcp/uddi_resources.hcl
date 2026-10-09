# Auto-generated resource acceptance-test cases for DtcMonitorTcp (uddi).
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      port = 49152
    }
    check = {
      "uddi.name" = "{{random}}"
      "uddi.port" = "49152"
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
      port = 49152
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      port    = 49152
      comment = "DTC TCP monitor test"
    }
    check = {
      "uddi.comment" = "DTC TCP monitor test"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      port    = 49152
      comment = "updated DTC TCP monitor comment"
    }
    check = {
      "uddi.comment" = "updated DTC TCP monitor comment"
    }
  }

}

case "disabled" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      port     = 49152
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      port     = 49152
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
      name     = "{{random}}"
      port     = 49152
      interval = 10
    }
    check = {
      "uddi.interval" = "10"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      port     = 49152
      interval = 30
    }
    check = {
      "uddi.interval" = "30"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      port = 49152
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name = "{{random2}}"
      port = 49152
    }
    check = {
      "uddi.name" = "{{random2}}"
    }
  }

}

case "port" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      port = 49152
    }
    check = {
      "uddi.port" = "49152"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      port = 49153
    }
    check = {
      "uddi.port" = "49153"
    }
  }

}

case "retry_down" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      port       = 49152
      retry_down = 3
    }
    check = {
      "uddi.retry_down" = "3"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      port       = 49152
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
      name     = "{{random}}"
      port     = 49152
      retry_up = 4
    }
    check = {
      "uddi.retry_up" = "4"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      port     = 49152
      retry_up = 6
    }
    check = {
      "uddi.retry_up" = "6"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
      port = 49152
      tags = { Site = "{{random2}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random2}}"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      port = 49152
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
      name     = "{{random}}"
      port     = 49152
      interval = 60
      timeout  = 30
    }
    check = {
      "uddi.timeout" = "30"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      port     = 49152
      interval = 120
      timeout  = 60
    }
    check = {
      "uddi.timeout" = "60"
    }
  }

}
