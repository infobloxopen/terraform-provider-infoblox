# Auto-generated resource acceptance-test cases for RecordA (uddi).
# Source: terraform-provider-uddi. One labelled `case` block per legacy
# TestAccRecordAResource_<Case>.
# Replayed by the unified provider via acctest.RunResourceCases.

case "basic" {
  # basic — generated from terraform-provider-uddi
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "unified_dns_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { address = "10.0.0.15" }
      zone = unified_dns_zone_auth.test.id
    }
    check = {
      "uddi.rdata.address" = "10.0.0.15"
    }
  }

}

case "disappears" {
  # disappears — generated from terraform-provider-uddi
  backend = "uddi"
  disappears = true
  expect_non_empty_plan = true
  prerequisites_hcl = <<-PREREQ
  resource "unified_dns_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { address = "10.0.0.15" }
      zone = unified_dns_zone_auth.test.id
    }
  }

}

case "comment" {
  # comment — generated from terraform-provider-uddi
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "unified_dns_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { address = "10.0.0.1" }
      zone = unified_dns_zone_auth.test.id
      comment = "some comment"
    }
    check = {
      "uddi.comment" = "some comment"
    }
  }

  step {
    uddi {
      rdata = { address = "10.0.0.1" }
      zone = unified_dns_zone_auth.test.id
      comment = "updated comment"
    }
    check = {
      "uddi.comment" = "updated comment"
    }
  }

}

case "disabled" {
  # disabled — generated from terraform-provider-uddi
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "unified_dns_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { address = "10.0.0.1" }
      zone = unified_dns_zone_auth.test.id
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

  step {
    uddi {
      rdata = { address = "10.0.0.1" }
      zone = unified_dns_zone_auth.test.id
      disabled = false
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

}

case "inheritance_sources" {
  # inheritance_sources — generated from terraform-provider-uddi
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "unified_dns_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { address = "10.0.0.1" }
      zone = unified_dns_zone_auth.test.id
      inheritance_sources = { ttl = { action = "inherit" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "inherit"
    }
  }

  step {
    uddi {
      rdata = { address = "10.0.0.1" }
      zone = unified_dns_zone_auth.test.id
      inheritance_sources = { ttl = { action = "override" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "override"
    }
  }

}

case "name_in_zone" {
  # name_in_zone — generated from terraform-provider-uddi
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "unified_dns_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { address = "10.0.0.1" }
      zone = unified_dns_zone_auth.test.id
      name_in_zone = "xyz"
    }
    check = {
      "uddi.name_in_zone" = "xyz"
    }
  }

  step {
    uddi {
      rdata = { address = "10.0.0.1" }
      zone = unified_dns_zone_auth.test.id
      name_in_zone = "abc"
    }
    check = {
      "uddi.name_in_zone" = "abc"
    }
  }

}

case "rdata" {
  # rdata — generated from terraform-provider-uddi
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "unified_dns_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { address = "10.0.0.1" }
      zone = unified_dns_zone_auth.test.id
    }
    check = {
      "uddi.rdata.address" = "10.0.0.1"
    }
  }

  step {
    uddi {
      rdata = { address = "10.0.0.2" }
      zone = unified_dns_zone_auth.test.id
    }
    check = {
      "uddi.rdata.address" = "10.0.0.2"
    }
  }

}

case "tags" {
  # tags — generated from terraform-provider-uddi
  backend = "uddi"
  skip        = true
  skip_reason = "unsupported config (prerequisites or nested values)"
}

case "ttl" {
  # ttl — generated from terraform-provider-uddi
  backend = "uddi"
  prerequisites_hcl = <<-PREREQ
  resource "unified_dns_zone_auth" "test" {
    uddi = {
      fqdn = "{{random}}.com."
      primary_type = "cloud"
    }
  }
  PREREQ

  step {
    uddi {
      rdata = { address = "10.0.0.1" }
      zone = unified_dns_zone_auth.test.id
      ttl = 60
    }
    check = {
      "uddi.ttl" = "60"
    }
  }

  step {
    uddi {
      rdata = { address = "10.0.0.1" }
      zone = unified_dns_zone_auth.test.id
      ttl = 90
    }
    check = {
      "uddi.ttl" = "90"
    }
  }

}

case "view" {
  # view — generated from terraform-provider-uddi
  backend = "uddi"
  skip        = true
  skip_reason = "unsupported config (prerequisites or nested values)"
}

case "zone" {
  # zone — generated from terraform-provider-uddi
  backend = "uddi"
  skip        = true
  skip_reason = "unsupported config (prerequisites or nested values)"
}

case "options" {
  # options — generated from terraform-provider-uddi
  backend = "uddi"
  skip        = true
  skip_reason = "unsupported config (prerequisites or nested values)"
}
