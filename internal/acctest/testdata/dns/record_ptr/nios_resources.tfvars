# Auto-generated resource acceptance-test cases for RecordPtr.
case "basic" {
  backend = "nios"
  parallel = true

  step {
    nios {
      ipv4addr = "192.168.104.22"
      ptrdname = "{{random}}.example.com"
    }
    check = {
      "nios.ipv4addr"           = "192.168.104.22"
      "nios.ptrdname"           = "{{random}}.example.com"
      "nios.view"               = "default"
      "nios.name"               = "22.104.168.192.in-addr.arpa"
      "nios.creator"            = "STATIC"
      "nios.ddns_protected"     = "false"
      "nios.disable"            = "false"
      "nios.forbid_reclamation" = "false"
    }
  }

}

case "disappears" {
  backend = "nios"
  disappears = true
  expect_non_empty_plan = true
  parallel = true

  step {
    nios {
      ipv4addr = "192.168.104.23"
      ptrdname = "{{random}}.example.com"
    }
  }

}

case "comment" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name     = "23.104.168.192.in-addr.arpa"
      ptrdname = "{{random}}.example.com"
      view     = "default"
      comment  = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name     = "23.104.168.192.in-addr.arpa"
      ptrdname = "{{random}}.example.com"
      view     = "default"
      comment  = "This is an updated comment"
    }
    check = {
      "nios.comment" = "This is an updated comment"
    }
  }

}

case "creator" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name     = "24.104.168.192.in-addr.arpa"
      ptrdname = "{{random}}.example.com"
      view     = "default"
      creator  = "STATIC"
    }
    check = {
      "nios.creator" = "STATIC"
    }
  }

  step {
    nios {
      name     = "24.104.168.192.in-addr.arpa"
      ptrdname = "{{random}}.example.com"
      view     = "default"
      creator  = "DYNAMIC"
    }
    check = {
      "nios.creator" = "DYNAMIC"
    }
  }

}

case "ddns_principal" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name           = "25.104.168.192.in-addr.arpa"
      ptrdname       = "{{random}}.example.com"
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
      name           = "25.104.168.192.in-addr.arpa"
      ptrdname       = "{{random}}.example.com"
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
  backend = "nios"
  parallel = true

  step {
    nios {
      name           = "26.104.168.192.in-addr.arpa"
      ptrdname       = "{{random}}.example.com"
      view           = "default"
      ddns_protected = false
    }
    check = {
      "nios.ddns_protected" = "false"
    }
  }

  step {
    nios {
      name           = "26.104.168.192.in-addr.arpa"
      ptrdname       = "{{random}}.example.com"
      view           = "default"
      ddns_protected = true
    }
    check = {
      "nios.ddns_protected" = "true"
    }
  }

}

case "disable" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name     = "27.104.168.192.in-addr.arpa"
      ptrdname = "{{random}}.example.com"
      view     = "default"
      disable  = false
    }
    check = {
      "nios.disable" = "false"
    }
  }

  step {
    nios {
      name     = "27.104.168.192.in-addr.arpa"
      ptrdname = "{{random}}.example.com"
      view     = "default"
      disable  = true
    }
    check = {
      "nios.disable" = "true"
    }
  }

}

case "ext_attrs" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name      = "28.104.168.192.in-addr.arpa"
      ptrdname  = "{{random}}.example.com"
      view      = "default"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "28.104.168.192.in-addr.arpa"
      ptrdname  = "{{random}}.example.com"
      view      = "default"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "forbid_reclamation" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name               = "29.104.168.192.in-addr.arpa"
      ptrdname           = "{{random}}.example.com"
      view               = "default"
      forbid_reclamation = true
    }
    check = {
      "nios.forbid_reclamation" = "true"
    }
  }

  step {
    nios {
      name               = "29.104.168.192.in-addr.arpa"
      ptrdname           = "{{random}}.example.com"
      view               = "default"
      forbid_reclamation = false
    }
    check = {
      "nios.forbid_reclamation" = "false"
    }
  }

}

case "ipv4addr" {
  backend = "nios"
  parallel = true

  step {
    nios {
      ipv4addr = "192.168.104.30"
      ptrdname = "{{random}}.example.com"
      view     = "default"
    }
    check = {
      "nios.ipv4addr" = "192.168.104.30"
    }
  }

  step {
    nios {
      ipv4addr = "192.168.104.31"
      ptrdname = "{{random}}.example.com"
      view     = "default"
    }
    check = {
      "nios.ipv4addr" = "192.168.104.31"
    }
  }

}

case "ipv6addr" {
  backend = "nios"
  parallel = true

  step {
    nios {
      ipv6addr = "2001:db8::24"
      ptrdname = "test.example.com"
      view     = "default"
    }
    check = {
      "nios.ipv6addr" = "2001:db8::24"
      "nios.ptrdname" = "test.example.com"
      "nios.view"     = "default"
    }
  }

  step {
    nios {
      ipv6addr = "2001:db8::25"
      ptrdname = "test.example.com"
      view     = "default"
    }
    check = {
      "nios.ipv6addr" = "2001:db8::25"
      "nios.ptrdname" = "test.example.com"
      "nios.view"     = "default"
    }
  }

}

case "name" {
  backend = "nios"
  parallel = true

  step {
    nios {
      name     = "32.104.168.192.in-addr.arpa"
      ptrdname = "{{random}}.example.com"
      view     = "default"
    }
    check = {
      "nios.name" = "32.104.168.192.in-addr.arpa"
    }
  }

  step {
    nios {
      name     = "33.104.168.192.in-addr.arpa"
      ptrdname = "{{random}}.example.com"
      view     = "default"
    }
    check = {
      "nios.name" = "33.104.168.192.in-addr.arpa"
    }
  }

}

case "ptrdname" {
  backend = "nios"
  parallel = true

  step {
    nios {
      ptrdname = "{{random}}.example.com"
      ipv4addr = "192.168.104.34"
      view     = "default"
    }
    check = {
      "nios.ptrdname" = "{{random}}.example.com"
    }
  }

  step {
    nios {
      ptrdname = "{{random2}}.example.com"
      ipv4addr = "192.168.104.34"
      view     = "default"
    }
    check = {
      "nios.ptrdname" = "{{random2}}.example.com"
    }
  }

}

case "ttl" {
  backend = "nios"
  parallel = true

  step {
    nios {
      ipv6addr = "2001:db8::26"
      ptrdname = "{{random}}.example.com"
      view     = "default"
      ttl      = 300
    }
    check = {
      "nios.ttl" = "300"
    }
  }

  step {
    nios {
      ipv6addr = "2001:db8::26"
      ptrdname = "{{random}}.example.com"
      view     = "default"
      ttl      = 600
    }
    check = {
      "nios.ttl" = "600"
    }
  }

}
