# Auto-generated resource acceptance-test cases for Extensibleattributedef.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      type = "STRING"
    }
    check = {
      "nios.name" = "{{random}}"
      "nios.type" = "STRING"
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
      type = "STRING"
    }
  }

}

case "allowed_object_types" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name                 = "{{random}}"
      allowed_object_types = ["NetworkContainer", "IPv6NetworkContainer", "Network"]
      type                 = "STRING"
    }
    check = {
      "nios.allowed_object_types.#" = "3"
      "nios.allowed_object_types.0" = "NetworkContainer"
      "nios.allowed_object_types.1" = "IPv6NetworkContainer"
      "nios.allowed_object_types.2" = "Network"
    }
  }

  step {
    nios {
      name                 = "{{random}}"
      allowed_object_types = ["NetworkContainer", "IPv6NetworkContainer", "Network", "IPv6Network", "FixedAddress", "IPv6FixedAddress"]
      type                 = "STRING"
    }
    check = {
      "nios.allowed_object_types.#" = "6"
      "nios.allowed_object_types.0" = "NetworkContainer"
      "nios.allowed_object_types.1" = "IPv6NetworkContainer"
      "nios.allowed_object_types.2" = "Network"
      "nios.allowed_object_types.3" = "IPv6Network"
      "nios.allowed_object_types.4" = "FixedAddress"
      "nios.allowed_object_types.5" = "IPv6FixedAddress"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a sample comment"
      type    = "URL"
    }
    check = {
      "nios.comment" = "This is a sample comment"
      "nios.type"    = "URL"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "This is an updated comment"
      type    = "URL"
    }
    check = {
      "nios.comment" = "This is an updated comment"
      "nios.type"    = "URL"
    }
  }

}

case "default_value" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name          = "{{random}}"
      default_value = "STRING"
      type          = "STRING"
    }
    check = {
      "nios.default_value" = "STRING"
    }
  }

  step {
    nios {
      name          = "{{random}}"
      default_value = "9945"
      type          = "STRING"
    }
    check = {
      "nios.default_value" = "9945"
    }
  }

}

case "flags" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name  = "{{random}}"
      flags = "C"
      type  = "STRING"
    }
    check = {
      "nios.flags" = "C"
    }
  }

  step {
    nios {
      name  = "{{random}}"
      flags = "CL"
      type  = "STRING"
    }
    check = {
      "nios.flags" = "CL"
    }
  }

}

case "list_values" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name        = "{{random}}"
      list_values = [{ value = "value1" }, { value = "value2" }, { value = "value3" }]
      type        = "STRING"
    }
    check = {
      "nios.list_values.#"       = "3"
      "nios.list_values.0.value" = "value1"
      "nios.list_values.1.value" = "value2"
      "nios.list_values.2.value" = "value3"
    }
  }

  step {
    nios {
      name        = "{{random}}"
      list_values = [{ value = "value1" }, { value = "value2" }, { value = "value3" }, { value = "value4" }, { value = "value5" }]
      type        = "STRING"
    }
    check = {
      "nios.list_values.#"       = "5"
      "nios.list_values.0.value" = "value1"
      "nios.list_values.1.value" = "value2"
      "nios.list_values.2.value" = "value3"
      "nios.list_values.3.value" = "value4"
      "nios.list_values.4.value" = "value5"
    }
  }

}

case "max" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      max  = 100
      type = "INTEGER"
    }
    check = {
      "nios.max"  = "100"
      "nios.type" = "INTEGER"
    }
  }

  step {
    nios {
      name = "{{random}}"
      max  = 200
      type = "INTEGER"
    }
    check = {
      "nios.max"  = "200"
      "nios.type" = "INTEGER"
    }
  }

}

case "min" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      min  = 10
      type = "INTEGER"
    }
    check = {
      "nios.min" = "10"
    }
  }

  step {
    nios {
      name = "{{random}}"
      min  = 5
      type = "INTEGER"
    }
    check = {
      "nios.min" = "5"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      type = "INTEGER"
    }
    check = {
      "nios.name" = "{{random}}"
      "nios.type" = "INTEGER"
    }
  }

  step {
    nios {
      name = "{{random2}}"
      type = "INTEGER"
    }
    check = {
      "nios.name" = "{{random2}}"
      "nios.type" = "INTEGER"
    }
  }

}

case "type" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      type = "STRING"
    }
    check = {
      "nios.type" = "STRING"
    }
  }

}
