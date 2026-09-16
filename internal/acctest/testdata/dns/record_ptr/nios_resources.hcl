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
  parallel = false
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "reverse" {
    nios = {
      fqdn        = "192.168.11.0/24"
      zone_format = "IPV4"
      view        = "default"
    }
  }
  PREREQ

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      ipv4addr = "192.168.11.10"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.ipv4addr" = "192.168.11.10"
      "nios.ptrdname" = "{{random3}}.com"
      "nios.view"     = "default"
    }
  }

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      ipv4addr = "192.168.11.20"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.ipv4addr" = "192.168.11.20"
      "nios.ptrdname" = "{{random3}}.com"
      "nios.view"     = "default"
    }
  }

}

case "ipv6addr" {
  backend  = "nios"
  parallel = false
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "reverse" {
    nios = {
      fqdn        = "2002:5598::/64"
      zone_format = "IPV6"
      view        = "default"
    }
  }
  PREREQ

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      ipv6addr = "2002:5598::10"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.ipv6addr" = "2002:5598::10"
      "nios.ptrdname" = "{{random3}}.com"
      "nios.view"     = "default"
    }
  }

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      ipv6addr = "2002:5598::20"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.ipv6addr" = "2002:5598::20"
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

case "func_call" {
  backend  = "nios"
  parallel = false
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_network" "test" {
    nios = {
      network      = "192.168.12.0/24"
      network_view = "default"
    }
  }
  resource "infoblox_zone_auth" "reverse" {
    nios = {
      fqdn        = "192.168.12.0/24"
      zone_format = "IPV4"
      view        = "default"
    }
  }
  PREREQ

  step {
    depends_on = [infoblox_network.test, infoblox_zone_auth.reverse]
    nios {
      dynamic_allocation = { network = infoblox_network.test.nios.network, network_view = "default" }
      ptrdname           = "{{random3}}.com"
      view               = "default"
      comment            = "Original Function Call"
    }
    check = {
      "nios.ptrdname" = "{{random3}}.com"
      "nios.view"     = "default"
      "nios.comment"  = "Original Function Call"
      "nios.creator"  = "STATIC"
      "nios.disable"  = "false"
    }
  }

  step {
    depends_on = [infoblox_network.test, infoblox_zone_auth.reverse]
    nios {
      dynamic_allocation = { network = infoblox_network.test.nios.network, network_view = "default" }
      ptrdname           = "updated-{{random3}}.com"
      view               = "default"
      comment            = "Function Call with Update"
    }
    check = {
      "nios.ptrdname" = "updated-{{random3}}.com"
      "nios.comment"  = "Function Call with Update"
    }
  }

}

case "func_call_ipv6" {
  backend  = "nios"
  parallel = false
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_ipv6_network" "test" {
    nios = {
      network      = "2001:db8:abcd:15::/64"
      network_view = "default"
    }
  }
  resource "infoblox_zone_auth" "reverse" {
    nios = {
      fqdn        = "2001:db8:abcd:15::/64"
      zone_format = "IPV6"
      view        = "default"
    }
  }
  PREREQ

  step {
    depends_on = [infoblox_ipv6_network.test, infoblox_zone_auth.reverse]
    nios {
      dynamic_allocation = { network = "2001:db8:abcd:15::/64", network_view = "default" }
      ptrdname           = "{{random3}}.com"
      view               = "default"
      comment            = "IPv6 Function Call"
    }
    check = {
      "nios.ptrdname" = "{{random3}}.com"
      "nios.view"     = "default"
      "nios.comment"  = "IPv6 Function Call"
      "nios.creator"  = "STATIC"
      "nios.disable"  = "false"
    }
  }

  step {
    depends_on = [infoblox_ipv6_network.test, infoblox_zone_auth.reverse]
    nios {
      dynamic_allocation = { network = "2001:db8:abcd:15::/64", network_view = "default" }
      ptrdname           = "updated-{{random3}}.com"
      view               = "default"
      comment            = "IPv6 Function Call with Update"
    }
    check = {
      "nios.ptrdname" = "updated-{{random3}}.com"
      "nios.comment"  = "IPv6 Function Call with Update"
    }
  }

}

case "name_arpa_ipv4" {
  backend  = "nios"
  parallel = false
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "reverse" {
    nios = {
      fqdn        = "192.168.13.0/24"
      zone_format = "IPV4"
      view        = "default"
    }
  }
  PREREQ

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      name     = "50.13.168.192.in-addr.arpa"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.name"     = "50.13.168.192.in-addr.arpa"
      "nios.ptrdname" = "{{random3}}.com"
      "nios.view"     = "default"
      "nios.creator"  = "STATIC"
      "nios.disable"  = "false"
    }
  }

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      name     = "50.13.168.192.in-addr.arpa"
      ptrdname = "{{random4}}.com"
      view     = "default"
    }
    check = {
      "nios.name"     = "50.13.168.192.in-addr.arpa"
      "nios.ptrdname" = "{{random4}}.com"
    }
  }

}

case "name_arpa_ipv6" {
  backend  = "nios"
  parallel = false
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "reverse" {
    nios = {
      fqdn        = "2002:5597::/64"
      zone_format = "IPV6"
      view        = "default"
    }
  }
  PREREQ

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      name     = "0.5.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.7.9.5.5.2.0.0.2.ip6.arpa"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
    check = {
      "nios.name"     = "0.5.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.7.9.5.5.2.0.0.2.ip6.arpa"
      "nios.ptrdname" = "{{random3}}.com"
      "nios.view"     = "default"
      "nios.creator"  = "STATIC"
      "nios.disable"  = "false"
    }
  }

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      name     = "0.5.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.7.9.5.5.2.0.0.2.ip6.arpa"
      ptrdname = "{{random4}}.com"
      view     = "default"
    }
    check = {
      "nios.name"     = "0.5.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.7.9.5.5.2.0.0.2.ip6.arpa"
      "nios.ptrdname" = "{{random4}}.com"
    }
  }

}
