# Auto-generated resource acceptance-test cases for DtcPool (uddi).
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name   = "{{random}}"
      method = "round_robin"
    }
    check = {
      "uddi.name"   = "{{random}}"
      "uddi.method" = "round_robin"
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
      name   = "{{random}}"
      method = "round_robin"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      method  = "round_robin"
      comment = "pool testing"
    }
    check = {
      "uddi.comment" = "pool testing"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      method  = "round_robin"
      comment = "updated pool comment"
    }
    check = {
      "uddi.comment" = "updated pool comment"
    }
  }

}

case "disabled" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      method   = "round_robin"
      disabled = false
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      method   = "round_robin"
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

}

case "method" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name   = "{{random}}"
      method = "round_robin"
    }
    check = {
      "uddi.method" = "round_robin"
    }
  }

  step {
    uddi {
      name   = "{{random}}"
      method = "ratio"
    }
    check = {
      "uddi.method" = "ratio"
    }
  }

}

case "pool_availability" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name              = "{{random}}"
      method            = "round_robin"
      pool_availability = "any"
    }
    check = {
      "uddi.pool_availability" = "any"
    }
  }

  step {
    uddi {
      name              = "{{random}}"
      method            = "round_robin"
      pool_availability = "all"
    }
    check = {
      "uddi.pool_availability" = "all"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name   = "{{random}}"
      method = "round_robin"
      tags   = { Site = "{{random2}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random2}}"
    }
  }

  step {
    uddi {
      name   = "{{random}}"
      method = "round_robin"
      tags   = { Site = "{{random3}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random3}}"
    }
  }

}

case "ttl" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name   = "{{random}}"
      method = "round_robin"
      ttl    = 30
    }
    check = {
      "uddi.ttl" = "30"
    }
  }

  step {
    uddi {
      name   = "{{random}}"
      method = "round_robin"
      ttl    = 60
    }
    check = {
      "uddi.ttl" = "60"
    }
  }

}
