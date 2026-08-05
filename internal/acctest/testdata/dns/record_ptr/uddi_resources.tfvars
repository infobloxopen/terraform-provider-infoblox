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
      "uddi.rdata.dname" = "domain.com."
    }
  }

  step {
    uddi {
      name_in_zone = "1.0.0"
      rdata        = { dname = "apple.com." }
      zone         = infoblox_zone_auth.test.id
    }
    check = {
      "uddi.rdata.dname" = "apple.com."
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
      "uddi.name_in_zone" = "1.0.0"
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
      "uddi.name_in_zone" = "2.0.0"
    }
  }

}
