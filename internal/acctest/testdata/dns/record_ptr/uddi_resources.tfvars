# Auto-generated resource acceptance-test cases for RecordPtr.
case "rdata" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "10.in-addr.arpa."
      primary_type = "cloud"
      view = infoblox_view.test.id
    }
  }
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      name_in_zone = "1.0.0"
      rdata        = { dname = "domain.com." }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.dname"     = "domain.com."
      "uddi.name_in_zone"    = "1.0.0"
      "uddi.options.address" = "10.0.0.1"
    }
  }

  step {
    uddi {
      name_in_zone = "1.0.0"
      rdata        = { dname = "apple.com." }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.dname"     = "apple.com."
      "uddi.name_in_zone"    = "1.0.0"
      "uddi.options.address" = "10.0.0.1"
    }
  }

}

case "options" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "10.in-addr.arpa."
      primary_type = "cloud"
      view = infoblox_view.test.id
    }
  }
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      rdata   = { dname = "domain.com." }
      options = { address = "10.0.0.1" }
      view    = infoblox_zone_auth.test.uddi.view
    }
    depends_on = [infoblox_zone_auth.test]
    check = {
      "uddi.name_in_zone"    = "1.0.0"
      "uddi.options.address" = "10.0.0.1"
      "uddi.rdata.dname"     = "domain.com."
    }
  }

  step {
    uddi {
      rdata   = { dname = "domain.com." }
      options = { address = "10.0.0.2" }
      view    = infoblox_zone_auth.test.uddi.view
    }
    depends_on = [infoblox_zone_auth.test]
    check = {
      "uddi.name_in_zone"    = "2.0.0"
      "uddi.options.address" = "10.0.0.2"
      "uddi.rdata.dname"     = "domain.com."
    }
  }

}

case "absolute_name_spec" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "10.in-addr.arpa."
      primary_type = "cloud"
      view = infoblox_view.test.id
    }
  }
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      absolute_name_spec = "3.0.0.10.in-addr.arpa."
      view               = infoblox_zone_auth.test.uddi.view
      rdata              = { dname = "domain.com." }
    }
    depends_on = [infoblox_zone_auth.test]
    check = {
      "uddi.absolute_name_spec" = "3.0.0.10.in-addr.arpa."
      "uddi.name_in_zone"       = "3.0.0"
      "uddi.options.address"    = "10.0.0.3"
      "uddi.rdata.dname"        = "domain.com."
    }
  }

  step {
    uddi {
      absolute_name_spec = "3.0.0.10.in-addr.arpa."
      view               = infoblox_zone_auth.test.uddi.view
      rdata              = { dname = "apple.com." }
    }
    depends_on = [infoblox_zone_auth.test]
    check = {
      "uddi.absolute_name_spec" = "3.0.0.10.in-addr.arpa."
      "uddi.name_in_zone"       = "3.0.0"
      "uddi.options.address"    = "10.0.0.3"
      "uddi.rdata.dname"        = "apple.com."
    }
  }

}

# IPv6 reverse mapping: the zone covers 2001:db8:0:1::/64, so the record is addressed by
# options.address and UDDI derives the 16 nibble labels of name_in_zone from it.
case "reverse_mapping_ipv6" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "1.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa."
      primary_type = "cloud"
      view = infoblox_view.test.id
    }
  }
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      rdata   = { dname = "domain.com." }
      options = { address = "2001:db8:0:1::50" }
      view    = infoblox_zone_auth.test.uddi.view
    }
    depends_on = [infoblox_zone_auth.test]
    check = {
      "uddi.name_in_zone"    = "0.5.0.0.0.0.0.0.0.0.0.0.0.0.0.0"
      "uddi.options.address" = "2001:db8:0:1::50"
      "uddi.rdata.dname"     = "domain.com."
    }
  }

  step {
    uddi {
      rdata   = { dname = "apple.com." }
      options = { address = "2001:db8:0:1::50" }
      view    = infoblox_zone_auth.test.uddi.view
    }
    depends_on = [infoblox_zone_auth.test]
    check = {
      "uddi.name_in_zone"    = "0.5.0.0.0.0.0.0.0.0.0.0.0.0.0.0"
      "uddi.options.address" = "2001:db8:0:1::50"
      "uddi.rdata.dname"     = "apple.com."
    }
  }

}

case "ttl" {
  backend  = "uddi"
  parallel = true
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_zone_auth" "test" {
    uddi = {
      fqdn = "10.in-addr.arpa."
      primary_type = "cloud"
      view = infoblox_view.test.id
    }
  }
  resource "infoblox_view" "test" {
    uddi = {
      name = "{{random}}"
    }
  }
  PREREQ

  step {
    uddi {
      name_in_zone = "5.0.0"
      rdata        = { dname = "domain.com." }
      zone         = infoblox_zone_auth.test.id
      ttl          = 300
    }
    check = {
      "uddi.ttl"             = "300"
      "uddi.name_in_zone"    = "5.0.0"
      "uddi.options.address" = "10.0.0.5"
      "uddi.rdata.dname"     = "domain.com."
    }
  }

  step {
    uddi {
      name_in_zone = "5.0.0"
      rdata        = { dname = "domain.com." }
      zone         = infoblox_zone_auth.test.id
      ttl          = 600
    }
    check = {
      "uddi.ttl"             = "600"
      "uddi.name_in_zone"    = "5.0.0"
      "uddi.options.address" = "10.0.0.5"
      "uddi.rdata.dname"     = "domain.com."
    }
  }

}
