# Auto-generated resource acceptance-test cases for DtcMonitorTcp.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      port = 49152
    }
    check = {
      "nios.name"       = "{{random}}"
      "nios.interval"   = "5"
      "nios.timeout"    = "15"
      "nios.retry_down" = "1"
      "nios.retry_up"   = "1"
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
      port = 49152
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      port    = 49152
      comment = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      port    = 49152
      comment = "This is an updated comment"
    }
    check = {
      "nios.comment" = "This is an updated comment"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      port      = 49152
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      port      = 49152
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "interval" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}"
      port     = 49152
      interval = 4
    }
    check = {
      "nios.interval" = "4"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      port     = 49152
      interval = 10
    }
    check = {
      "nios.interval" = "10"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      port = 49152
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
      port = 49152
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "port" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      port = 49152
    }
  }

  step {
    nios {
      name = "{{random}}"
      port = 49153
    }
  }

}

case "retry_down" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name       = "{{random}}"
      port       = 49152
      retry_down = 3
    }
    check = {
      "nios.retry_down" = "3"
    }
  }

  step {
    nios {
      name       = "{{random}}"
      port       = 49152
      retry_down = 5
    }
    check = {
      "nios.retry_down" = "5"
    }
  }

}

case "retry_up" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}"
      port     = 49152
      retry_up = 4
    }
    check = {
      "nios.retry_up" = "4"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      port     = 49152
      retry_up = 6
    }
    check = {
      "nios.retry_up" = "6"
    }
  }

}

case "timeout" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      port    = 49152
      timeout = 30
    }
    check = {
      "nios.timeout" = "30"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      port    = 49152
      timeout = 40
    }
    check = {
      "nios.timeout" = "40"
    }
  }

}
