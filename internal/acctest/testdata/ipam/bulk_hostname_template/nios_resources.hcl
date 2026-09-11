# Auto-generated resource acceptance-test cases for Bulkhostnametemplate.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      template_name   = "{{random}}"
      template_format = "host-$4"
    }
    check = {
      "nios.template_name"   = "{{random}}"
      "nios.template_format" = "host-$4"
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
      template_name   = "{{random}}"
      template_format = "host-$4"
    }
  }

}

case "template_format" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      template_name   = "{{random}}"
      template_format = "server-$4"
    }
    check = {
      "nios.template_format" = "server-$4"
    }
  }

  step {
    nios {
      template_name   = "{{random}}"
      template_format = "server-$3-$4"
    }
    check = {
      "nios.template_format" = "server-$3-$4"
    }
  }

}

case "template_name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      template_name   = "{{random}}"
      template_format = "server-$4"
    }
    check = {
      "nios.template_name" = "{{random}}"
    }
  }

  step {
    nios {
      template_name   = "{{random2}}"
      template_format = "server-$4"
    }
    check = {
      "nios.template_name" = "{{random2}}"
    }
  }

}
