# Auto-generated resource acceptance-test cases for Ruleset.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      type = "NXDOMAIN"
    }
    check = {
      "nios.name"     = "{{random}}"
      "nios.type"     = "NXDOMAIN"
       "nios.disabled" = "false"
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
      type = "NXDOMAIN"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      type    = "BLACKLIST"
      comment = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      type    = "BLACKLIST"
      comment = "Updated comment"
    }
    check = {
      "nios.comment" = "Updated comment"
    }
  }

}

case "disabled" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}"
      type     = "NXDOMAIN"
      disabled = true
    }
    check = {
      "nios.disabled" = "true"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      type     = "NXDOMAIN"
      disabled = false
    }
    check = {
      "nios.disabled" = "false"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      type = "NXDOMAIN"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
      type = "NXDOMAIN"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "nxdomain_rules" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name           = "{{random}}"
      type           = "NXDOMAIN"
      nxdomain_rules = [{ action = "PASS", pattern = "example.com" }]
    }
    check = {
      "nios.nxdomain_rules.0.action"  = "PASS"
      "nios.nxdomain_rules.0.pattern" = "example.com"
    }
  }

  step {
    nios {
      name           = "{{random}}"
      type           = "NXDOMAIN"
      nxdomain_rules = [{ action = "MODIFY", pattern = "test.com" }]
    }
    check = {
      "nios.nxdomain_rules.0.action"  = "MODIFY"
      "nios.nxdomain_rules.0.pattern" = "test.com"
    }
  }

  step {
    nios {
      name           = "{{random}}"
      type           = "NXDOMAIN"
      nxdomain_rules = [{ action = "REDIRECT", pattern = "redirect-test.com" }]
    }
    check = {
      "nios.nxdomain_rules.0.action"  = "REDIRECT"
      "nios.nxdomain_rules.0.pattern" = "redirect-test.com"
    }
  }

}
