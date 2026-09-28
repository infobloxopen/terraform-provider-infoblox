case "basic" {
  backend  = "uddi"
  parallel = true
  skip_if_env_empty = ["UDDI_INFRA_HOST_LEGACY_ID_1"]
  skip_reason       = "UDDI_INFRA_HOST_LEGACY_ID_1 environment variable must be set for this test to run"

  prerequisites_hcl = <<-PREREQ
    resource "infoblox_dns_server" "test" {
        uddi = {
              name = "{{random}}"
            }
    }
  PREREQ

  step {
      id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      server = infoblox_dns_server.test.id
    }
    check = {
      "uddi.server"                        = infoblox_dns_server.test.id
    }
  }

}

case "absolute_name" {
  backend  = "uddi"
  parallel = true
  skip_if_env_empty = ["UDDI_INFRA_HOST_LEGACY_ID_1"]
  skip_reason       = "UDDI_INFRA_HOST_LEGACY_ID_1 environment variable must be set for this test to run"

  prerequisites_hcl = <<-PREREQ
    resource "infoblox_dns_server" "test" {
        name = {{random}}
    }
  PREREQ

  step {
      id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      server = infoblox_dns_server.test.id
      absolute_name = "{{random2}}."
    }
    check = {
      "uddi.absolute_name"                        = "{{random}}."
    }
  }

step {
    id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      server = infoblox_dns_server.test.id
      absolute_name = "{{random3}}."
    }
    check = {
      "uddi.absolute_name"                        = "{{random3}}."
    }
  }

}


case "server" {
  backend  = "uddi"
  parallel = true
  skip_if_env_empty = ["UDDI_INFRA_HOST_LEGACY_ID_1"]
  skip_reason       = "UDDI_INFRA_HOST_LEGACY_ID_1 environment variable must be set for this test to run"

  prerequisites_hcl = <<-PREREQ
    resource "infoblox_dns_server" "test" {
        name = {{random}}
    }
    resource "infoblox_dns_server" "test2" {
            name = {{random2}}
        }
  PREREQ

  step {
      id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      server = infoblox_dns_server.test.id
    }
    check = {
      "uddi.server"                        = infoblox_dns_server.test.id
    }
  }

step {
    id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      server = infoblox_dns_server.test2.id
    }
    check = {
      "uddi.server"                        = infoblox_dns_server.test2.id
    }
  }

}