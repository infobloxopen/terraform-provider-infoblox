# Auto-generated resource acceptance-test cases for DtcMonitorSnmp.
case "basic" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name"       = "{{random}}"
      "nios.comment"    = ""
      "nios.community"  = "public"
      "nios.interval"   = "5"
      "nios.port"       = "161"
      "nios.retry_down" = "1"
      "nios.retry_up"   = "1"
      "nios.timeout"    = "15"
      "nios.version"    = "V2C"
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
    }
  }

}

case "comment" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      comment = "This is a comment"
    }
    check = {
      "nios.comment" = "This is a comment"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      comment = "This is an updated comment"
    }
    check = {
      "nios.comment" = "This is an updated comment"
    }
  }

}

case "community" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      community = "private"
    }
    check = {
      "nios.community" = "private"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      community = "trapuser"
    }
    check = {
      "nios.community" = "trapuser"
    }
  }

}

case "context" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      context = "text"
    }
    check = {
      "nios.context" = "text"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      context = "update context"
    }
    check = {
      "nios.context" = "update context"
    }
  }

}

case "engine_id" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      engine_id = "66356e6574776f726B73"
    }
    check = {
      "nios.engine_id" = "66356e6574776f726B73"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      engine_id = "800007DB03000C754120"
    }
    check = {
      "nios.engine_id" = "800007DB03000C754120"
    }
  }

}

case "ext_attrs" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random2}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random2}}"
    }
  }

  step {
    nios {
      name      = "{{random}}"
      ext_attrs = { Site = "{{random3}}" }
    }
    check = {
      "nios.ext_attrs.Site" = "{{random3}}"
    }
  }

}

case "interval" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}"
      interval = 10
    }
    check = {
      "nios.interval" = "10"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      interval = 20
    }
    check = {
      "nios.interval" = "20"
    }
  }

}

case "name" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
    }
    check = {
      "nios.name" = "{{random}}"
    }
  }

  step {
    nios {
      name = "{{random2}}"
    }
    check = {
      "nios.name" = "{{random2}}"
    }
  }

}

case "oids" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      oids = [{ oid = ".2", condition = "EXACT", first = "10" }, { oid = ".02", condition = "RANGE", first = "2", last = "4", type = "INTEGER" }, { oid = ".1", condition = "EXACT", first = "20" }]
    }
    check = {
      "nios.oids.0.oid"       = ".2"
      "nios.oids.0.condition" = "EXACT"
      "nios.oids.0.first"     = "10"
      "nios.oids.1.oid"       = ".02"
      "nios.oids.1.condition" = "RANGE"
      "nios.oids.1.first"     = "2"
      "nios.oids.1.last"      = "4"
      "nios.oids.1.type"      = "INTEGER"
      "nios.oids.2.oid"       = ".1"
      "nios.oids.2.condition" = "EXACT"
      "nios.oids.2.first"     = "20"
    }
  }

  step {
    nios {
      name = "{{random}}"
      oids = [{ oid = ".2", condition = "LEQ", first = "10" }, { oid = ".01", condition = "GEQ", first = "25" }, { oid = ".1", condition = "LEQ", first = "20" }]
    }
    check = {
      "nios.oids.0.oid"       = ".2"
      "nios.oids.0.condition" = "LEQ"
      "nios.oids.0.first"     = "10"
      "nios.oids.1.oid"       = ".01"
      "nios.oids.1.condition" = "GEQ"
      "nios.oids.1.first"     = "25"
      "nios.oids.2.oid"       = ".1"
      "nios.oids.2.condition" = "LEQ"
      "nios.oids.2.first"     = "20"
    }
  }

}

case "port" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name = "{{random}}"
      port = 10161
    }
    check = {
      "nios.port" = "10161"
    }
  }

  step {
    nios {
      name = "{{random}}"
      port = 10162
    }
    check = {
      "nios.port" = "10162"
    }
  }

}

case "retry_down" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name       = "{{random}}"
      retry_down = 5
    }
    check = {
      "nios.retry_down" = "5"
    }
  }

  step {
    nios {
      name       = "{{random}}"
      retry_down = 10
    }
    check = {
      "nios.retry_down" = "10"
    }
  }

}

case "retry_up" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name     = "{{random}}"
      retry_up = 3
    }
    check = {
      "nios.retry_up" = "3"
    }
  }

  step {
    nios {
      name     = "{{random}}"
      retry_up = 7
    }
    check = {
      "nios.retry_up" = "7"
    }
  }

}

case "timeout" {
  backend  = "nios"
  parallel = true

  step {
    nios {
      name    = "{{random}}"
      timeout = 30
    }
    check = {
      "nios.timeout" = "30"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      timeout = 45
    }
    check = {
      "nios.timeout" = "45"
    }
  }

}

# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: requires_resource: infoblox_snmp_user not yet implemented
case "user" {
  backend     = "nios"
  skip        = true
  skip_reason = "requires_resource: infoblox_snmp_user not yet implemented"
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_snmp_user_unknown" "snmpuser_parent" {
    nios = {
      name = "nios_security_snmp_user.snmpuser_parent"
      authentication_protocol = "NONE"
      privacy_protocol = "NONE"
    }
  }
  resource "infoblox_snmp_user_unknown" "snmpuser_parent1" {
    nios = {
      name = "nios_security_snmp_user.snmpuser_parent1"
      authentication_protocol = "NONE"
      privacy_protocol = "NONE"
    }
  }
  PREREQ

  step {
    nios {
      name    = "{{random}}"
      version = "V3"
      user    = infoblox_snmp_user_unknown.snmpuser_parent.nios.name
    }
  }

  step {
    nios {
      name    = "{{random}}"
      version = "V3"
      user    = infoblox_snmp_user_unknown.snmpuser_parent1.nios.name
    }
  }

}

# TODO: auto-extraction incomplete — please verify and fill in manually.
# Reason: requires_resource: infoblox_snmp_user not yet implemented
case "version" {
  backend     = "nios"
  skip        = true
  skip_reason = "requires_resource: infoblox_snmp_user not yet implemented"
  parallel    = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_snmp_user_unknown" "snmpuser_parent" {
    nios = {
      authentication_protocol = "NONE"
      privacy_protocol = "NONE"
    }
  }
  resource "infoblox_snmp_user_unknown" "snmpuser_parent1" {
    nios = {
      authentication_protocol = "NONE"
      privacy_protocol = "NONE"
    }
  }
  PREREQ

  step {
    nios {
      name    = "{{random}}"
      version = "V1"
    }
    check = {
      "nios.version" = "V1"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      version = "V2C"
    }
    check = {
      "nios.version" = "V2C"
    }
  }

  step {
    nios {
      name    = "{{random}}"
      version = "V3"
      user    = infoblox_snmp_user_unknown.snmpuser_parent1.nios.name
    }
    check = {
      "nios.version" = "V3"
    }
  }

}
