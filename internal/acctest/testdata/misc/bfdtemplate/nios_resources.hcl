# Auto-generated resource acceptance-test cases for Bfdtemplate.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name"                 = "{{random}}"
      "nios.authentication_type"  = "NONE"
      "nios.authentication_key_id" = "1"
      "nios.detection_multiplier" = "3"
      "nios.min_rx_interval"      = "100"
      "nios.min_tx_interval"      = "100"
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

case "authentication_key" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name"                         = "{{random}}"
      "nios.authentication_type"          = "NONE"
      "nios.detection_multiplier"         = "3"
      "nios.min_rx_interval"              = "100"
      "nios.min_tx_interval"              = "100"
    }
  }

  step {
    nios {
      name               = "{{random}}"
      authentication_key = "auth_key_1234"
    }
    check = {
      "nios.name"                         = "{{random}}"
      "nios.authentication_type"          = "NONE"
      "nios.detection_multiplier"         = "3"
      "nios.min_rx_interval"              = "100"
      "nios.min_tx_interval"              = "100"
    }
  }

  step {
    nios {
      name               = "{{random}}"
      authentication_key = "updated_auth_key_1234"
    }
    check = {
      "nios.name"                         = "{{random}}"
      "nios.authentication_type"          = "NONE"
      "nios.detection_multiplier"         = "3"
      "nios.min_rx_interval"              = "100"
      "nios.min_tx_interval"              = "100"
    }
  }

}

case "authentication_key_id" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                  = "{{random}}"
      authentication_key_id = 4
      authentication_type   = "MD5"
      authentication_key    = "1234"
      min_rx_interval       = 1000
      min_tx_interval       = 1000
    }
    check = {
      "nios.authentication_key_id" = "4"
    }
  }

  step {
    nios {
      name                  = "{{random}}"
      authentication_key_id = 5
      authentication_type   = "MD5"
      authentication_key    = "1234"
      min_rx_interval       = 1000
      min_tx_interval       = 1000
    }
    check = {
      "nios.authentication_key_id" = "5"
    }
  }

}

case "authentication_type" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                = "{{random}}"
      authentication_type = "METICULOUS-MD5"
      authentication_key  = "1234"
    }
    check = {
      "nios.authentication_type" = "METICULOUS-MD5"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      authentication_type = "METICULOUS-SHA1"
      authentication_key  = "1234"
    }
    check = {
      "nios.authentication_type" = "METICULOUS-SHA1"
    }
  }

  step {
    nios {
      name                = "{{random}}"
      authentication_type = "SHA1"
      authentication_key  = "1234"
    }
    check = {
      "nios.authentication_type" = "SHA1"
    }
  }

}

case "detection_multiplier" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      detection_multiplier = 4
    }
    check = {
      "nios.detection_multiplier" = "4"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      detection_multiplier = 5
    }
    check = {
      "nios.detection_multiplier" = "5"
    }
  }

}

case "min_rx_interval" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name            = "{{random}}"
      min_rx_interval = 200
    }
    check = {
      "nios.min_rx_interval" = "200"
    }
  }

  step {
    nios {
      name            = "{{random}}"
      min_rx_interval = 300
    }
    check = {
      "nios.min_rx_interval" = "300"
    }
  }

}

case "min_tx_interval" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name            = "{{random}}"
      min_tx_interval = 200
    }
    check = {
      "nios.min_tx_interval" = "200"
    }
  }

  step {
    nios {
      name            = "{{random}}"
      min_tx_interval = 300
    }
    check = {
      "nios.min_tx_interval" = "300"
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
