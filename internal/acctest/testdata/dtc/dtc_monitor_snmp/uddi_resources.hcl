case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      version = "v2c"
    }
    check = {
      "uddi.name"         = "{{random}}"
      "uddi.version"      = "v2c"
      "uddi.comment"      = ""
      "uddi.community"    = "public"
      "uddi.disabled"     = "false"
      "uddi.interval"     = "15"
      "uddi.port"         = "161"
      "uddi.retry_down"   = "1"
      "uddi.retry_up"     = "1"
      "uddi.timeout"      = "10"
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
      name    = "{{random}}"
      version = "v2c"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      version = "v2c"
      comment = "This is a comment"
    }
    check = {
      "uddi.comment" = "This is a comment"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      version = "v2c"
      comment = "This is an updated comment"
    }
    check = {
      "uddi.comment" = "This is an updated comment"
    }
  }

}

case "community" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name      = "{{random}}"
      version   = "v2c"
      community = "private"
    }
    check = {
      "uddi.community" = "private"
    }
  }

  step {
    uddi {
      name      = "{{random}}"
      version   = "v2c"
      community = "trapuser"
    }
    check = {
      "uddi.community" = "trapuser"
    }
  }

}

case "disabled" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      version  = "v2c"
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      version  = "v2c"
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
      version  = "v2c"
      interval = 30
    }
    check = {
      "uddi.interval" = "30"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      version  = "v2c"
      interval = 60
    }
    check = {
      "uddi.interval" = "60"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      version = "v2c"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name    = "{{random2}}"
      version = "v2c"
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
      name    = "{{random}}"
      version = "v2c"
      port    = 10161
    }
    check = {
      "uddi.port" = "10161"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      version = "v2c"
      port    = 10162
    }
    check = {
      "uddi.port" = "10162"
    }
  }

}

case "retry_down" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      version    = "v2c"
      retry_down = 5
    }
    check = {
      "uddi.retry_down" = "5"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      version    = "v2c"
      retry_down = 10
    }
    check = {
      "uddi.retry_down" = "10"
    }
  }

}

case "retry_up" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      version  = "v2c"
      retry_up = 3
    }
    check = {
      "uddi.retry_up" = "3"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      version  = "v2c"
      retry_up = 7
    }
    check = {
      "uddi.retry_up" = "7"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      version = "v2c"
      tags    = { Site = "{{random2}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random2}}"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      version = "v2c"
      tags    = { Site = "{{random3}}" }
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
      version  = "v2c"
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
      version  = "v2c"
      interval = 120
      timeout  = 60
    }
    check = {
      "uddi.timeout" = "60"
    }
  }

}

case "version" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      version = "v1"
    }
    check = {
      "uddi.version" = "v1"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      version = "v2c"
    }
    check = {
      "uddi.version" = "v2c"
    }
  }

}
