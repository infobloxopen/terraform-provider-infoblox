# Auto-generated resource acceptance-test cases for NamedList.
case "basic" {
  backend = "uddi"

  step {
    uddi {
      name            = "{{random}}"
      items_described = [{ item = "{{random2}}.com", description = "Example Domain" }]
      type            = "custom_list"
    }
    check = {
      "uddi.name"                          = "{{random}}"
      "uddi.items_described.0.item"        = "{{random2}}.com"
      "uddi.items_described.0.description" = "Example Domain"
      "uddi.description"                   = ""
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true

  step {
    uddi {
      name            = "{{random}}"
      items_described = [{ item = "{{random2}}.com", description = "Example Domain" }]
      type            = "custom_list"
    }
  }

}

case "name" {
  backend = "uddi"

  step {
    uddi {
      name            = "{{random}}"
      items_described = [{ item = "{{random3}}.com", description = "Example Domain" }]
      type            = "custom_list"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name            = "{{random2}}"
      items_described = [{ item = "{{random3}}.com", description = "Example Domain" }]
      type            = "custom_list"
    }
    check = {
      "uddi.name" = "{{random2}}"
    }
  }

}

case "items_described" {
  backend = "uddi"

  step {
    uddi {
      name            = "{{random}}"
      items_described = [{ item = "{{random2}}.com", description = "Example Item 1" }]
      type            = "custom_list"
    }
    check = {
      "uddi.items_described.0.item"        = "{{random2}}.com"
      "uddi.items_described.0.description" = "Example Item 1"
    }
  }

  step {
    uddi {
      name            = "{{random}}"
      items_described = [{ item = "{{random3}}.com", description = "Example Item 2" }]
      type            = "custom_list"
    }
    check = {
      "uddi.items_described.0.item"        = "{{random3}}.com"
      "uddi.items_described.0.description" = "Example Item 2"
    }
  }

}

case "description" {
  backend = "uddi"

  step {
    uddi {
      name            = "{{random}}"
      items_described = [{ item = "{{random2}}.com", description = "Example Domain" }]
      description     = "Test Description"
      type            = "custom_list"
    }
    check = {
      "uddi.description" = "Test Description"
    }
  }

  step {
    uddi {
      name            = "{{random}}"
      items_described = [{ item = "{{random2}}.com", description = "Example Domain" }]
      description     = "Updated Test Description"
      type            = "custom_list"
    }
    check = {
      "uddi.description" = "Updated Test Description"
    }
  }

}

case "confidence" {
  backend = "uddi"

  step {
    uddi {
      name             = "{{random}}"
      items_described  = [{ item = "{{random2}}.com", description = "Example Domain" }]
      confidence_level = "HIGH"
      type             = "custom_list"
    }
    check = {
      "uddi.confidence_level" = "HIGH"
    }
  }

  step {
    uddi {
      name             = "{{random}}"
      items_described  = [{ item = "{{random2}}.com", description = "Example Domain" }]
      confidence_level = "MEDIUM"
      type             = "custom_list"
    }
    check = {
      "uddi.confidence_level" = "MEDIUM"
    }
  }

  step {
    uddi {
      name             = "{{random}}"
      items_described  = [{ item = "{{random2}}.com", description = "Example Domain" }]
      confidence_level = "LOW"
      type             = "custom_list"
    }
    check = {
      "uddi.confidence_level" = "LOW"
    }
  }

}

case "type" {
  backend = "uddi"

  step {
    uddi {
      name            = "{{random}}"
      items_described = [{ item = "{{random2}}.com", description = "Example Domain" }]
      type            = "custom_list"
    }
    check = {
      "uddi.type" = "custom_list"
    }
  }

}

case "threat_level" {
  backend = "uddi"

  step {
    uddi {
      name            = "{{random}}"
      items_described = [{ item = "{{random2}}.com", description = "Example Domain" }]
      threat_level    = "HIGH"
      type            = "custom_list"
    }
    check = {
      "uddi.threat_level" = "HIGH"
    }
  }

  step {
    uddi {
      name            = "{{random}}"
      items_described = [{ item = "{{random2}}.com", description = "Example Domain" }]
      threat_level    = "MEDIUM"
      type            = "custom_list"
    }
    check = {
      "uddi.threat_level" = "MEDIUM"
    }
  }

  step {
    uddi {
      name            = "{{random}}"
      items_described = [{ item = "{{random2}}.com", description = "Example Domain" }]
      threat_level    = "LOW"
      type            = "custom_list"
    }
    check = {
      "uddi.threat_level" = "LOW"
    }
  }

}
