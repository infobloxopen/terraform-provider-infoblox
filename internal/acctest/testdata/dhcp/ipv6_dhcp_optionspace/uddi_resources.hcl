# Auto-generated resource acceptance-test cases for Ipv6DhcpOptionspace.
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
    }
    check = {
      "uddi.name"     = "{{random}}"
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
      name     = "{{random}}"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
      comment  = "test comment"
    }
    check = {
      "uddi.comment" = "test comment"
    }
  }

  step {
    uddi {
      name     = "{{random}}"
      comment  = "test comment update"
    }
    check = {
      "uddi.comment" = "test comment update"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "{{random}}"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name     = "{{random2}}"
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
      name     = "{{random}}"
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
      tags     = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}
