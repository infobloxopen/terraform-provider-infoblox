# Auto-generated resource acceptance-test cases for DtcMonitorIcmp.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name"       = "{{random}}"
      "nios.interval"   = "5"
      "nios.retry_down" = "1"
      "nios.retry_up"   = "1"
      "nios.timeout"    = "15"
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
      comment = "Updated comment"
    }
    check = {
      "nios.comment" = "Updated comment"
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
      interval = 20
    }
    check = {
      "nios.interval" = "20"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      interval = 30
    }
    check = {
      "nios.interval" = "30"
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

case "retry_down" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name       = "{{random}}"
      retry_down = 2
    }
    check = {
      "nios.retry_down" = "2"
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
      retry_up = 5
    }
    check = {
      "nios.retry_up" = "5"
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
      timeout = 30
    }
    check = {
      "nios.timeout" = "30"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      timeout = 45
    }
    check = {
      "nios.timeout" = "45"
    }
  }

}
