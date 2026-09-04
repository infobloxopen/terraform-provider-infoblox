# Auto-generated datasource acceptance-test cases for RecordPtr.
case "filters" {
  backend = "nios"

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      name     = "nios.name"
      ptrdname = "nios.ptrdname"
      view     = "nios.view"
    }
  }

  pair_checks = ["nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.ipv4addr", "nios.ipv6addr", "nios.name", "nios.ptrdname", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name     = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
  }

}

case "ext_attr_filters" {
  backend = "nios"

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    nios = {
      fqdn = "{{random}}.com"
    }
  }
  PREREQ

  filter {
    type   = "ext_attr_filters"
    values = {
      Site = "nios.ext_attrs.Site"
    }
  }

  pair_checks = ["nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.ipv4addr", "nios.ipv6addr", "nios.name", "nios.ptrdname", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    nios {
      name      = "{{random2}}.${infoblox_zone_auth.test.nios.fqdn}"
      ptrdname  = "{{random3}}.com"
      view      = "default"
      ext_attrs = { Site = "{{random4}}" }
    }
  }

}

# Reverse-mapped PTR lookup: ipv4addr is only populated when the record lives in an
# authoritative reverse zone, so this case provisions its own IPV4 zone_auth.
case "ipv4addr_filters" {
  backend = "nios"

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "reverse" {
    nios = {
      fqdn        = "192.168.125.0/24"
      zone_format = "IPV4"
      view        = "default"
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      ipv4addr = "nios.ipv4addr"
      ptrdname = "nios.ptrdname"
      view     = "nios.view"
    }
  }

  pair_checks = ["nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.ipv4addr", "nios.ipv6addr", "nios.name", "nios.ptrdname", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      ipv4addr = "192.168.125.10"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
  }

}

# The case Ipv6addr_filters also provisions IPv6 reverse zone as a pre-req
case "ipv6addr_filters" {
  backend = "nios"

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "reverse" {
    nios = {
      fqdn        = "2002:5525::/64"
      zone_format = "IPV6"
      view        = "default"
    }
  }
  PREREQ

  filter {
    type   = "filters"
    values = {
      ipv6addr = "nios.ipv6addr"
      ptrdname = "nios.ptrdname"
      view     = "nios.view"
    }
  }

  pair_checks = ["nios.comment", "nios.creator", "nios.ddns_principal", "nios.ddns_protected", "nios.disable", "nios.forbid_reclamation", "nios.ipv4addr", "nios.ipv6addr", "nios.name", "nios.ptrdname", "nios.ttl", "nios.use_ttl", "nios.view"]

  step {
    depends_on = [infoblox_zone_auth.reverse]
    nios {
      ipv6addr = "2002:5525::10"
      ptrdname = "{{random3}}.com"
      view     = "default"
    }
  }

}
