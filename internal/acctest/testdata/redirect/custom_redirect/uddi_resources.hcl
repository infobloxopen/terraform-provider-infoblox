case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "custom-redirect-{{random}}"
      data = "156.2.3.10"
    }
    check = {
      "uddi.name" = "custom-redirect-{{random}}"
      "uddi.data" = "156.2.3.10"
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
      name = "custom-redirect-{{random}}"
      data = "156.2.3.10"
    }
  }
}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "custom-redirect-{{random}}"
      data = "156.2.3.10"
    }
    check = {
      "uddi.name" = "custom-redirect-{{random}}"
    }
  }

  step {
    uddi {
      name = "custom-redirect-{{random2}}"
      data = "156.2.3.10"
    }
    check = {
      "uddi.name" = "custom-redirect-{{random2}}"
    }
  }
}

case "data" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "custom-redirect-{{random}}"
      data = "156.2.3.10"
    }
    check = {
      "uddi.data" = "156.2.3.10"
    }
  }

  step {
    uddi {
      name = "custom-redirect-{{random}}"
      data = "192.168.1.1,192.168.1.2"
    }
    check = {
      "uddi.data" = "192.168.1.1,192.168.1.2"
    }
  }
}
