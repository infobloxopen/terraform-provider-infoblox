# Auto-generated resource acceptance-test cases for DtcLbdn (UDDI backend).
# view: default DNS view on env-5 (dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376)
# dtc_policy refs: example-policy-1, example-policy-topology (stable on env-5)
case "basic" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
    }
    check = {
      "uddi.name" = "dtc-lbdn-{{random}}."
      "uddi.view" = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
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
      name = "dtc-lbdn-{{random}}."
      view = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
    }
  }

}

case "comment" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name    = "dtc-lbdn-{{random}}."
      view    = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      comment = "resource-comment"
    }
    check = {
      "uddi.comment" = "resource-comment"
    }
  }

  step {
    uddi {
      name    = "dtc-lbdn-{{random}}."
      view    = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      comment = "resource-comment-update"
    }
    check = {
      "uddi.comment" = "resource-comment-update"
    }
  }

}

case "disabled" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name     = "dtc-lbdn-{{random}}."
      view     = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

  step {
    uddi {
      name     = "dtc-lbdn-{{random}}."
      view     = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      disabled = false
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

}

case "ttl" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      ttl  = 300
    }
    check = {
      "uddi.ttl" = "300"
    }
  }

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      ttl  = 600
    }
    check = {
      "uddi.ttl" = "600"
    }
  }

}

case "tags" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      tags = { Site = "{{random2}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random2}}"
    }
  }

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      tags = { Site = "{{random3}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random3}}"
    }
  }

}

case "dtc_policy" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "dtc-lbdn-{{random}}."
      view       = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      dtc_policy = { policy_id = "dtc/policy/dafaf7a5-307b-4e5c-895e-fd0d922c46fd" }
    }
    check = {
      "uddi.dtc_policy.policy_id" = "dtc/policy/dafaf7a5-307b-4e5c-895e-fd0d922c46fd"
    }
  }

  step {
    uddi {
      name       = "dtc-lbdn-{{random}}."
      view       = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      dtc_policy = { policy_id = "dtc/policy/f088b848-67cb-4b3f-a8fd-86283d8e228d" }
    }
    check = {
      "uddi.dtc_policy.policy_id" = "dtc/policy/f088b848-67cb-4b3f-a8fd-86283d8e228d"
    }
  }

}

case "name" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
    }
    check = {
      "uddi.name" = "dtc-lbdn-{{random}}."
    }
  }

  step {
    uddi {
      name = "dtc-lbdn-{{random2}}."
      view = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
    }
    check = {
      "uddi.name" = "dtc-lbdn-{{random2}}."
    }
  }

}

case "precedence" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name       = "dtc-lbdn-{{random}}."
      view       = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      precedence = 7
    }
    check = {
      "uddi.precedence" = "7"
    }
  }

  step {
    uddi {
      name       = "dtc-lbdn-{{random}}."
      view       = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      precedence = 12
    }
    check = {
      "uddi.precedence" = "12"
    }
  }

}

case "inheritance_sources" {
  backend  = "uddi"
  parallel = true

  step {
    uddi {
      name                = "dtc-lbdn-{{random}}."
      view                = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      ttl                 = 300
      inheritance_sources = { ttl = { action = "override" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "override"
      "uddi.ttl"                            = "300"
    }
  }

  step {
    uddi {
      name                = "dtc-lbdn-{{random}}."
      view                = "dns/view/206a2b2e-44d7-4e36-a376-28b79c5dc376"
      inheritance_sources = { ttl = { action = "inherit" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "inherit"
    }
  }

}
