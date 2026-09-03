# Auto-generated resource acceptance-test cases for Namedacl.
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
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
      name = "{{random}}"
    }
  }

}

case "compartment_id" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name           = "{{random}}"
      compartment_id = "01706."
    }
    check = {
      "uddi.name" = "{{random}}"
      "uddi.compartment_id" = "01706."
    }
  }

  step {
    uddi {
      name           = "{{random}}"
      compartment_id = ""
    }
    check = {
      "uddi.compartment_id" = ""
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "{{random}}"
      comment = "test comment"
    }
    check = {
      "uddi.comment" = "test comment"
    }
  }

  step {
    uddi {
      name    = "{{random}}"
      comment = "updated test comment"
    }
    check = {
      "uddi.comment" = "updated test comment"
    }
  }

}

case "list" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_namedacl" "prereq" {
    uddi = {
      name = "prereq-{{random}}"
      list = [{ element = "ip", access = "allow", address = "10.0.0.0/24" }]
    }
  }
  PREREQ

  step {
    uddi {
      name = "{{random}}"
      list = [{ access = "allow", element = "ip", address = "192.168.11.11" }]
    }
    check = {
      "uddi.list.0.access"  = "allow"
      "uddi.list.0.element" = "ip"
      "uddi.list.0.address" = "192.168.11.11"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      list = [{ access = "deny", element = "any" }]
    }
    check = {
      "uddi.list.0.access"  = "deny"
      "uddi.list.0.element" = "any"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      list = [{ element = "acl", acl = infoblox_namedacl.prereq.id }]
    }
    check = {
      "uddi.list.0.element" = "acl"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "{{random}}"
    }
    check = {
      "uddi.name" = "{{random}}"
    }
  }

  step {
    uddi {
      name = "{{random2}}"
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
      name = "{{random}}"
      tags = { tag1 = "value1", tag2 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
      "uddi.tags.tag2" = "value2"
    }
  }

  step {
    uddi {
      name = "{{random}}"
      tags = { tag2 = "value2changed", tag3 = "value3" }
    }
    check = {
      "uddi.tags.tag2" = "value2changed"
      "uddi.tags.tag3" = "value3"
    }
  }

}
