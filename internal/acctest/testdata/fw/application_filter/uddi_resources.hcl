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

case "criteria_subcategory" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ category = "Email", subcategory = "Collaboration" }]
    }
    check = {
      "uddi.criteria.0.category"    = "Email"
      "uddi.criteria.0.subcategory" = "Collaboration"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ category = "Email", subcategory = "Productivity" }]
    }
    check = {
      "uddi.criteria.0.category"    = "Email"
      "uddi.criteria.0.subcategory" = "Productivity"
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
      tags     = { tag1 = "{{random2}}", tag2 = "{{random3}}" }
    }
    check = {
      "uddi.tags.tag1" = "{{random2}}"
      "uddi.tags.tag2" = "{{random3}}"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      criteria = [{ name = "Microsoft 365" }]
      tags     = { tag2 = "{{random4}}", tag3 = "{{random5}}" }
    }
    check = {
      "uddi.tags.tag2" = "{{random4}}"
      "uddi.tags.tag3" = "{{random5}}"
    }
  }

}
