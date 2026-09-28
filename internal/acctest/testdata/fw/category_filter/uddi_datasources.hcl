# Hand-written datasource acceptance-test cases for CategoryFilter.
# Legacy mining skipped this object: terraform-provider-bloxone names its test
# functions TestAccCategoryFiltersDataSource_* (plural "Filters"), which does not
# match the miner's TestAcc<Object>DataSource_* regex for singular "CategoryFilter".

case "filters" {
  backend  = "uddi"
  parallel = true

  filter {
    type = "filters"
    values = {
      name = "uddi.name"
    }
  }

  pair_checks = ["uddi.categories", "uddi.description", "uddi.name", "uddi.policies", "uddi.tags"]

  step {
    uddi {
      name       = "{{random}}"
      categories = ["College"]
    }
  }
}

case "tag_filters" {
  backend  = "uddi"
  parallel = true

  filter {
    type = "tag_filters"
    values = {
      tag1 = "uddi.tags.tag1"
    }
  }

  pair_checks = ["uddi.categories", "uddi.description", "uddi.name", "uddi.policies"]

  step {
    uddi {
      name       = "{{random}}"
      categories = ["College"]
      tags       = { tag1 = "{{random2}}" }
    }
  }
}
