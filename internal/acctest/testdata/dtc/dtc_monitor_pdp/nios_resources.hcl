# Auto-generated resource acceptance-test cases for DtcMonitorPdp.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name"       = "{{random}}"
      "nios.port"       = "2123"
      "nios.retry_down" = "1"
      "nios.retry_up"   = "1"
      "nios.timeout"    = "15"
      "nios.interval"   = "5"
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
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
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
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}"
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
      interval = 4
    }
    check = {
      "nios.interval" = "4"
    }
  }

  step {
    nios {
      name     = "{{random}}"
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
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
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
      port = 2314
    }
    check = {
      "nios.port" = "2314"
    }
  }

  step {
    nios {
      name = "{{random}}"
      port = 4321
    }
    check = {
      "nios.port" = "4321"
    }
  }

}

case "retry_down" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name       = "{{random}}"
      retry_down = 5
    }
    check = {
      "nios.retry_down" = "5"
    }
  }

  step {
    nios {
      name       = "{{random}}"
      retry_down = 3
    }
    check = {
      "nios.retry_down" = "3"
    }
  }

}

case "retry_up" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}"
      retry_up = 2
    }
    check = {
      "nios.retry_up" = "2"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      retry_up = 4
    }
    check = {
      "nios.retry_up" = "4"
    }
  }

}

case "timeout" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      timeout = 20
    }
    check = {
      "nios.timeout" = "20"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      timeout = 25
    }
    check = {
      "nios.timeout" = "25"
    }
  }

}
