# Auto-generated resource acceptance-test cases for RecordPtr.
case "basic" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.name"               = "{{random2}}.{{random}}.com"
      "nios.ptrdname"           = "{{random3}}.com"
      "nios.view"               = "default"
      "nios.creator"            = "STATIC"
      "nios.ddns_protected"     = "false"
      "nios.disable"            = "false"
      "nios.forbid_reclamation" = "false"
    }
  }

}

case "disappears" {
  backend               = "nios"
  disappears            = true
  expect_non_empty_plan = true
  parallel              = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
      comment  = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
      comment  = "This is an updated comment"
    }
    check = {
      "nios.comment" = "This is an updated comment"
    }
  }

}

case "creator" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
      creator  = "STATIC"
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
      creator  = "DYNAMIC"
    }
    check = {
      "nios.creator" = "DYNAMIC"
    }
  }

}

case "ddns_principal" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname       = "{{random3}}.com"
      view           = "default"
      creator        = "DYNAMIC"
      ddns_principal = "host/myhost.example.com@EXAMPLE.COM"
    }
    check = {
      "nios.ddns_principal" = "host/myhost.example.com@EXAMPLE.COM"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname       = "{{random3}}.com"
      view           = "default"
      creator        = "DYNAMIC"
      ddns_principal = "host/otherhost.example.net@EXAMPLE.NET"
    }
    check = {
      "nios.ddns_principal" = "host/otherhost.example.net@EXAMPLE.NET"
    }
  }

}

case "ddns_protected" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname       = "{{random3}}.com"
      view           = "default"
      ddns_protected = false
    }
    check = {
      "nios.ddns_protected" = "false"
    }
  }

  step {
    nios {
      name           = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname       = "{{random3}}.com"
      view           = "default"
      ddns_protected = true
    }
    check = {
      "nios.ddns_protected" = "true"
    }
  }

}

case "disable" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
      disable  = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
      disable  = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname  = "{{random3}}.com"
      view      = "default"
      ext_attrs = { Site = "{{random4}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random4}}"
    }
  }

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname  = "{{random3}}.com"
      view      = "default"
      ext_attrs = { Site = "{{random5}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random5}}"
    }
  }

}

case "forbid_reclamation" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname           = "{{random3}}.com"
      view               = "default"
      forbid_reclamation = true
    }
    check = {
      "nios.forbid_reclamation" = "true"
    }
  }

  step {
    nios {
      name               = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname           = "{{random3}}.com"
      view               = "default"
      forbid_reclamation = false
    }
    check = {
      "nios.forbid_reclamation" = "false"
    }
  }

}

case "ipv4addr" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.name" = "{{random2}}.{{random}}.com"
    }
  }

  step {
    nios {
      name     = "{{random4}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.name" = "{{random4}}.{{random}}.com"
    }
  }

}

case "ipv6addr" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.name"     = "{{random2}}.{{random}}.com"
      "nios.ptrdname" = "{{random3}}.com"
      "nios.view"     = "default"
    }
  }

  step {
    nios {
      name     = "{{random4}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.name"     = "{{random4}}.{{random}}.com"
      "nios.ptrdname" = "{{random3}}.com"
      "nios.view"     = "default"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.name" = "{{random2}}.{{random}}.com"
    }
  }

  step {
    nios {
      name     = "{{random4}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.name" = "{{random4}}.{{random}}.com"
    }
  }

}

case "ptrdname" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.ptrdname" = "{{random3}}.com"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random4}}.com"
      view     = "default"
    }
    check = {
      "nios.ptrdname" = "{{random4}}.com"
    }
  }

}

case "ttl" {
  backend  = "nios"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
      ttl      = 300
    }
    check = {
      "nios.ttl" = "300"
    }
  }

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
      ttl      = 600
    }
    check = {
      "nios.ttl" = "600"
    }
  }

}

case "reverse_mapping" {
  backend  = "nios"
  parallel = false
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "reverse" {
    nios = {
      fqdn        = "192.168.10.0/24"
      zone_format = "IPV4"
      view        = "default"
    }
  }
  PREREQ

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      ipv4addr = "192.168.10.50"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.ipv4addr" = "192.168.10.50"
      "nios.ptrdname" = "{{random3}}.com"
      "nios.view"     = "default"
      "nios.creator"  = "STATIC"
      "nios.disable"  = "false"
    }
  }

}

case "reverse_mapping_ipv6" {
  backend  = "nios"
  parallel = false
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "reverse" {
    nios = {
      fqdn        = "2002:5599::/64"
      zone_format = "IPV6"
      view        = "default"
    }
  }
  PREREQ

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      ipv6addr = "2002:5599::50"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.ipv6addr" = "2002:5599::50"
      "nios.ptrdname" = "{{random3}}.com"
      "nios.view"     = "default"
      "nios.creator"  = "STATIC"
      "nios.disable"  = "false"
    }
  }

}
