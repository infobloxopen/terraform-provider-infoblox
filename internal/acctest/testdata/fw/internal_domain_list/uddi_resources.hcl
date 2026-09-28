# Auto-generated resource acceptance-test cases for InternalDomainList.
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name             = "{{random}}"
      internal_domains = ["example.somedomain.com"]
    }
    check = {
      "uddi.name" = "{{random}}"
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
      name             = "{{random}}"
      internal_domains = ["example.somedomain.com"]
    }
  }

}

case "description" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name             = "{{random}}"
      internal_domains = ["example.somedomain.com"]
      description      = "TEST_DESCRIPTION"
    }
    check = {
      "uddi.description" = "TEST_DESCRIPTION"
    }
  }

  step {
    uddi {
      name             = "{{random}}"
      internal_domains = ["example.somedomain.com"]
      description      = "TEST_DESCRIPTION_UPDATE"
    }
    check = {
      "uddi.description" = "TEST_DESCRIPTION_UPDATE"
    }
  }

}

case "internal_domains" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name             = "{{random}}"
      internal_domains = ["example.somedomain.com"]
    }
    check = {
      "uddi.internal_domains.0" = "example.somedomain.com"
    }
  }

  step {
    uddi {
      name             = "{{random}}"
      internal_domains = ["example.newdomain.com"]
    }
    check = {
      "uddi.internal_domains.0" = "example.newdomain.com"
    }
  }

}

case "internal_domains_multiple" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name             = "{{random}}"
      internal_domains = ["example.somedomain.com", "dev.somedomain.com", "internal.somedomain.com", "test.somedomain.com"]
    }
    check = {
      "uddi.internal_domains.0" = "example.somedomain.com"
      "uddi.internal_domains.1" = "dev.somedomain.com"
      "uddi.internal_domains.2" = "internal.somedomain.com"
      "uddi.internal_domains.3" = "test.somedomain.com"
    }
  }

  step {
    uddi {
      name             = "{{random}}"
      internal_domains = ["test.newdomain.com", "internal.newdomain.com", "newdomain.com"]
    }
    check = {
      "uddi.internal_domains.0" = "test.newdomain.com"
      "uddi.internal_domains.1" = "internal.newdomain.com"
      "uddi.internal_domains.2" = "newdomain.com"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name             = "{{random}}"
      internal_domains = ["example.somedomain.com"]
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name             = "{{random2}}"
      internal_domains = ["example.somedomain.com"]
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
      name             = "{{random}}"
      internal_domains = ["example.somedomain.com"]
      tags             = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      name             = "{{random}}"
      internal_domains = ["example.somedomain.com"]
      tags             = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}
