case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ name = "Microsoft 365" }]
    }
    check = {
      "uddi.name"        = "{{random}}"
      "uddi.description" = ""
      "uddi.readonly"    = "false"
    }
  }

}

case "disappears" {
  backend               = "uddi"
  skip                  = true
  skip_reason           = "t.Skip: Test Skipped due to inconsistent error codes returned by the API [TDDFW-397]"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ name = "Microsoft 365" }]
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ name = "Microsoft 365" }]
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name     = "{{random2}}"
      criteria = [{ name = "Microsoft 365" }]
    }
    check = {
      "uddi.name" = "{{random2}}"
    }
  }

}

case "description" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name        = "{{random}}"
      criteria    = [{ name = "Microsoft 365" }]
      description = "TEST_DESCRIPTION"
    }
    check = {
      "uddi.description" = "TEST_DESCRIPTION"
    }
  }

  step {
    uddi {
      name        = "{{random}}"
      criteria    = [{ name = "Microsoft 365" }]
      description = "TEST_DESCRIPTION_UPDATE"
    }
    check = {
      "uddi.description" = "TEST_DESCRIPTION_UPDATE"
    }
  }

}

case "criteria_name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ name = "Microsoft 365" }]
    }
    check = {
      "uddi.criteria.0.name" = "Microsoft 365"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ name = "163 Cloud" }]
    }
    check = {
      "uddi.criteria.0.name" = "163 Cloud"
    }
  }

}

case "criteria_category" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ category = "Email" }]
    }
    check = {
      "uddi.criteria.0.category" = "Email"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ category = "Communication" }]
    }
    check = {
      "uddi.criteria.0.category" = "Communication"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ name = "Microsoft 365" }]
      tags     = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ name = "Microsoft 365" }]
      tags     = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}
