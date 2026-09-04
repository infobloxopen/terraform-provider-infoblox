# Auto-generated resource acceptance-test cases for DtcLbdn (UDDI backend).
# dtc_policy refs: example-policy-1, example-policy-topology 
case "basic" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test_view" {
    uddi = {
      name = "view-{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "$${infoblox_view.test_view.id}"
    }
    check = {
      "uddi.name" = "dtc-lbdn-{{random}}."
    }
  }

}

case "disappears" {
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl     = <<-PREREQ
  resource "infoblox_view" "test_view" {
    uddi = {
      name = "view-{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "$${infoblox_view.test_view.id}"
    }
  }

}

case "comment" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test_view" {
    uddi = {
      name = "view-{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      name    = "dtc-lbdn-{{random}}."
      view    = "$${infoblox_view.test_view.id}"
      comment = "resource-comment"
    }
    check = {
      "uddi.comment" = "resource-comment"
    }
  }

  step {
    uddi {
      name    = "dtc-lbdn-{{random}}."
      view    = "$${infoblox_view.test_view.id}"
      comment = "resource-comment-update"
    }
    check = {
      "uddi.comment" = "resource-comment-update"
    }
  }

}

case "disabled" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test_view" {
    uddi = {
      name = "view-{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      name     = "dtc-lbdn-{{random}}."
      view     = "$${infoblox_view.test_view.id}"
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

  step {
    uddi {
      name     = "dtc-lbdn-{{random}}."
      view     = "$${infoblox_view.test_view.id}"
      disabled = false
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

}

case "ttl" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test_view" {
    uddi = {
      name = "view-{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "$${infoblox_view.test_view.id}"
      ttl  = 300
    }
    check = {
      "uddi.ttl" = "300"
    }
  }

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "$${infoblox_view.test_view.id}"
      ttl  = 600
    }
    check = {
      "uddi.ttl" = "600"
    }
  }

}

case "tags" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test_view" {
    uddi = {
      name = "view-{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "$${infoblox_view.test_view.id}"
      tags = { Site = "{{random2}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random2}}"
    }
  }

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "$${infoblox_view.test_view.id}"
      tags = { Site = "{{random3}}" }
    }
    check = {
      "uddi.tags.Site" = "{{random3}}"
    }
  }

}

case "dtc_policy" {
  backend           = "uddi"
  parallel          = true
  skip              = false
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test_view" {
    uddi = {
      name = "view-{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      name       = "dtc-lbdn-{{random}}."
      view       = "$${infoblox_view.test_view.id}"
      dtc_policy = { policy_id = "dtc/policy/dafaf7a5-307b-4e5c-895e-fd0d922c46fd" }
    }
    check = {
      "uddi.dtc_policy.policy_id" = "dtc/policy/dafaf7a5-307b-4e5c-895e-fd0d922c46fd"
    }
  }

  step {
    uddi {
      name       = "dtc-lbdn-{{random}}."
      view       = "$${infoblox_view.test_view.id}"
      dtc_policy = { policy_id = "dtc/policy/f088b848-67cb-4b3f-a8fd-86283d8e228d" }
    }
    check = {
      "uddi.dtc_policy.policy_id" = "dtc/policy/f088b848-67cb-4b3f-a8fd-86283d8e228d"
    }
  }

}

case "name" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test_view" {
    uddi = {
      name = "view-{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      name = "dtc-lbdn-{{random}}."
      view = "$${infoblox_view.test_view.id}"
    }
    check = {
      "uddi.name" = "dtc-lbdn-{{random}}."
    }
  }

  step {
    uddi {
      name = "dtc-lbdn-{{random2}}."
      view = "$${infoblox_view.test_view.id}"
    }
    check = {
      "uddi.name" = "dtc-lbdn-{{random2}}."
    }
  }

}

case "precedence" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test_view" {
    uddi = {
      name = "view-{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      name       = "dtc-lbdn-{{random}}."
      view       = "$${infoblox_view.test_view.id}"
      precedence = 7
    }
    check = {
      "uddi.precedence" = "7"
    }
  }

  step {
    uddi {
      name       = "dtc-lbdn-{{random}}."
      view       = "$${infoblox_view.test_view.id}"
      precedence = 12
    }
    check = {
      "uddi.precedence" = "12"
    }
  }

}

case "inheritance_sources" {
  backend           = "uddi"
  parallel          = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_view" "test_view" {
    uddi = {
      name = "view-{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      name                = "dtc-lbdn-{{random}}."
      view                = "$${infoblox_view.test_view.id}"
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
      view                = "$${infoblox_view.test_view.id}"
      inheritance_sources = { ttl = { action = "inherit" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "inherit"
    }
  }

}
