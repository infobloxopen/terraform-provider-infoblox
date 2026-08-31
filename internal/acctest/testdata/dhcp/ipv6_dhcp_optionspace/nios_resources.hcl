# Auto-generated resource acceptance-test cases for Ipv6DhcpOptionspace.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      enterprise_number = 5896
      name              = "{{random}}"
    }
    check = {
      "nios.enterprise_number" = "5896"
      "nios.name"              = "{{random}}"
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
      enterprise_number = 5896
      name              = "{{random}}"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      enterprise_number = 5896
      name              = "{{random}}"
      comment           = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      enterprise_number = 5896
      name              = "{{random}}"
      comment           = "This is an updated comment"
    }
    check = {
      "nios.comment" = "This is an updated comment"
    }
  }

}

case "enterprise_number" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      enterprise_number = 5896
      name              = "{{random}}"
    }
    check = {
      "nios.enterprise_number" = "5896"
    }
  }

  step {
    nios {
      enterprise_number = 5123
      name              = "{{random}}"
    }
    check = {
      "nios.enterprise_number" = "5123"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      enterprise_number = 5896
      name              = "{{random}}"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      enterprise_number = 5896
      name              = "{{random2}}"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}
