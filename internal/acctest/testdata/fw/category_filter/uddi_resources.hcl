# Hand-written resource acceptance-test cases for CategoryFilter.
# Legacy mining skipped this object: terraform-provider-bloxone names its test
# functions TestAccCategoryFiltersResource_* (plural "Filters"), which does not
# match the miner's TestAcc<Object>Resource_* regex for singular "CategoryFilter".

case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      categories = ["College"]
    }
    check = {
      "uddi.name"         = "{{random}}"
      "uddi.categories.0" = "College"
      "uddi.description"  = ""
    }
  }
}

case "disappears" {
  backend               = "uddi"
  skip                  = true
  skip_reason           = "t.Skip: Test Skipped due to inconsistent error codes returned by the API [TDDFW-397]"
  disappears            = true
  expect_non_empty_plan = true

  step {
    uddi {
      name       = "{{random}}"
      categories = ["College"]
    }
  }
}

case "categories" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      categories = ["College"]
    }
    check = {
      "uddi.categories.0" = "College"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      categories = ["Tutoring"]
    }
    check = {
      "uddi.categories.0" = "Tutoring"
    }
  }
}

case "description" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name        = "{{random}}"
      categories  = ["College"]
      description = "Test Description"
    }
    check = {
      "uddi.description" = "Test Description"
    }
  }

  step {
    uddi {
      name        = "{{random}}"
      categories  = ["College"]
      description = "Updated Test Description"
    }
    check = {
      "uddi.description" = "Updated Test Description"
    }
  }
}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      categories = ["College"]
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name       = "{{random2}}"
      categories = ["College"]
    }
    check = {
      "uddi.name" = "{{random2}}"
    }
  }
}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "{{random}}"
      categories = ["College"]
      tags = {
        tag1 = "value1"
        tag2 = "value2"
      }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      name       = "{{random}}"
      categories = ["College"]
      tags = {
        tag2 = "value2changed"
        tag3 = "value3"
      }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }
}
