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
      "uddi.name"                             = "{{random}}"
      "uddi.items_described.0.item"           = "{{random2}}.com"
      "uddi.items_described.0.description"    = "Example Domain"
      "uddi.items_described.0.status"         = "ACTIVE"
      "uddi.items_described.0.status_details" = ""
      "uddi.description"                      = ""
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

case "items_described_multiple" {
  backend = "uddi"

  step {
    uddi {
      name = "{{random}}"
      items_described = [
        { item = "{{random2}}.com", description = "Example Item 1" },
        { item = "{{random3}}.com", description = "Example Item 2" },
        { item = "{{random4}}.com", description = "Example Item 3" },
      ]
      type = "custom_list"
    }
    check = {
      "uddi.items_described.#"             = "3"
      "uddi.items_described.0.item"        = "{{random2}}.com"
      "uddi.items_described.0.description" = "Example Item 1"
      "uddi.items_described.1.item"        = "{{random3}}.com"
      "uddi.items_described.1.description" = "Example Item 2"
      "uddi.items_described.2.item"        = "{{random4}}.com"
      "uddi.items_described.2.description" = "Example Item 3"
    }
  }

  // Swap element order

  step {
    uddi {
      name = "{{random}}"
      items_described = [
        { item = "{{random4}}.com", description = "Example Item 3" },
        { item = "{{random2}}.com", description = "Example Item 1" },
        { item = "{{random3}}.com", description = "Example Item 2" },
      ]
      type = "custom_list"
    }
    check = {
      "uddi.items_described.#"             = "3"
      "uddi.items_described.0.item"        = "{{random4}}.com"
      "uddi.items_described.0.description" = "Example Item 3"
      "uddi.items_described.1.item"        = "{{random2}}.com"
      "uddi.items_described.1.description" = "Example Item 1"
      "uddi.items_described.2.item"        = "{{random3}}.com"
      "uddi.items_described.2.description" = "Example Item 2"
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

case "items_described_expiry_time" {
  backend = "uddi"

  step {
    uddi {
      name = "{{random}}"
      items_described = [
        {
          item        = "{{random2}}.com"
          description = "Example Item 1"
          expiry_time = "2030-01-01T00:00:00Z"
        },
      ]
      type = "custom_list"
    }
    check = {
      "uddi.items_described.0.expiry_time" = "2030-01-01T00:00:00Z"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      items_described = [
        {
          item        = "{{random2}}.com"
          description = "Example Item 1"
          expiry_time = "2031-06-15T12:30:00Z"
        },
      ]
      type = "custom_list"
    }
    check = {
      "uddi.items_described.0.expiry_time" = "2031-06-15T12:30:00Z"
    }
  }

}
