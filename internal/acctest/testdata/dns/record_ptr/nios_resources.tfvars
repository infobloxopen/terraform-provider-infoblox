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
    depends_on = [infoblox_zone_auth.test]
    nios {
      name     = "ptrrecord.{{random}}.com"
      ptrdname = "host.{{random}}.com"
      view     = "default"
    }
    check = {
      "nios.name"               = "ptrrecord.{{random}}.com"
      "nios.ptrdname"           = "host.{{random}}.com"
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
    depends_on = [infoblox_zone_auth.test]
    nios {
      name     = "ptrrecord.{{random}}.com"
      ptrdname = "host.{{random}}.com"
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
    depends_on = [infoblox_zone_auth.test]
    nios {
      name     = "ptrrecord.{{random}}.com"
      ptrdname = "host.{{random}}.com"
      view     = "default"
      comment  = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name     = "ptrrecord.{{random}}.com"
      ptrdname = "host.{{random}}.com"
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
    depends_on = [infoblox_zone_auth.test]
    nios {
      name     = "ptrrecord.{{random}}.com"
      ptrdname = "host.{{random}}.com"
      view     = "default"
      creator  = "STATIC"
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

  step {
    nios {
      name     = "ptrrecord.{{random}}.com"
      ptrdname = "host.{{random}}.com"
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
    depends_on = [infoblox_zone_auth.test]
    nios {
      name           = "ptrrecord.{{random}}.com"
      ptrdname       = "host.{{random}}.com"
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
      name           = "ptrrecord.{{random}}.com"
      ptrdname       = "host.{{random}}.com"
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
    depends_on = [infoblox_zone_auth.test]
    nios {
      name           = "ptrrecord.{{random}}.com"
      ptrdname       = "host.{{random}}.com"
      view           = "default"
      ddns_protected = false
    }
    check = {
      "nios.ddns_protected" = "false"
    }
  }

  step {
    nios {
      name           = "ptrrecord.{{random}}.com"
      ptrdname       = "host.{{random}}.com"
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
    depends_on = [infoblox_zone_auth.test]
    nios {
      name     = "ptrrecord.{{random}}.com"
      ptrdname = "host.{{random}}.com"
      view     = "default"
      disable  = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name     = "ptrrecord.{{random}}.com"
      ptrdname = "host.{{random}}.com"
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
    depends_on = [infoblox_zone_auth.test]
    nios {
      name      = "ptrrecord.{{random}}.com"
      ptrdname  = "host.{{random}}.com"
      view      = "default"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "ptrrecord.{{random}}.com"
      ptrdname  = "host.{{random}}.com"
      view      = "default"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
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
    depends_on = [infoblox_zone_auth.test]
    nios {
      name               = "ptrrecord.{{random}}.com"
      ptrdname           = "host.{{random}}.com"
      view               = "default"
      forbid_reclamation = true
    }
    check = {
      "nios.forbid_reclamation" = "true"
    }
  }

  step {
    nios {
      name               = "ptrrecord.{{random}}.com"
      ptrdname           = "host.{{random}}.com"
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
    depends_on = [infoblox_zone_auth.test]
    nios {
      name     = "ptrrecord1.{{random}}.com"
      ptrdname = "host.{{random}}.com"
      view     = "default"
    }
    check = {
      "nios.name" = "ptrrecord1.{{random}}.com"
    }
  }

  step {
    nios {
      name     = "ptrrecord2.{{random}}.com"
      ptrdname = "host.{{random}}.com"
      view     = "default"
    }
    check = {
      "nios.name" = "ptrrecord2.{{random}}.com"
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
    depends_on = [infoblox_zone_auth.test]
    nios {
      name     = "ptrrecord1.{{random}}.com"
      ptrdname = "host.{{random}}.com"
      view     = "default"
    }
    check = {
      "nios.name"     = "ptrrecord1.{{random}}.com"
      "nios.ptrdname" = "host.{{random}}.com"
      "nios.view"     = "default"
    }
  }

  step {
    nios {
      name     = "ptrrecord2.{{random}}.com"
      ptrdname = "host.{{random}}.com"
      view     = "default"
    }
    check = {
      "nios.name"     = "ptrrecord2.{{random}}.com"
      "nios.ptrdname" = "host.{{random}}.com"
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
    depends_on = [infoblox_zone_auth.test]
    nios {
      name     = "ptrrecord1.{{random}}.com"
      ptrdname = "host.{{random}}.com"
      view     = "default"
    }
    check = {
      "nios.name" = "ptrrecord1.{{random}}.com"
    }
  }

  step {
    nios {
      name     = "ptrrecord2.{{random}}.com"
      ptrdname = "host.{{random}}.com"
      view     = "default"
    }
    check = {
      "nios.name" = "ptrrecord2.{{random}}.com"
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
    depends_on = [infoblox_zone_auth.test]
    nios {
      name     = "ptrrecord.{{random}}.com"
      ptrdname = "host.{{random}}.com"
      view     = "default"
    }
    check = {
      "nios.ptrdname" = "host.{{random}}.com"
    }
  }

  step {
    nios {
      name     = "ptrrecord.{{random}}.com"
      ptrdname = "host2.{{random}}.com"
      view     = "default"
    }
    check = {
      "nios.ptrdname" = "host2.{{random}}.com"
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
    depends_on = [infoblox_zone_auth.test]
    nios {
      name     = "ptrrecord.{{random}}.com"
      ptrdname = "host.{{random}}.com"
      view     = "default"
      ttl      = 300
    }
    check = {
      "nios.ttl" = "300"
    }
  }

  step {
    nios {
      name     = "ptrrecord.{{random}}.com"
      ptrdname = "host.{{random}}.com"
      view     = "default"
      ttl      = 600
    }
    check = {
      "nios.ttl" = "600"
    }
  }

}
