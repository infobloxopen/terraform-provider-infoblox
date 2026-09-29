case "filters" {
  backend           = "uddi"
  skip_if_env_empty = ["UDDI_INFRA_HOST_LEGACY_ID_1"]
  skip_reason       = "UDDI_INFRA_HOST_LEGACY_ID_1 environment variable must be set for this test to run"

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dns_server" "test" {
      uddi = {
          name = "{{random}}"
      }
  }
    PREREQ

  filter {
    type = "filters"
    values = {
      server = "uddi.server"
    }
  }

  pair_checks = []


  step {
    id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      server = infoblox_dns_server.test.id
    }
  }

}

case "tag_filters" {
  backend           = "uddi"
  skip_if_env_empty = ["UDDI_INFRA_HOST_TAG_KEY_1", "UDDI_INFRA_HOST_TAG_VALUE_1"]
  skip_reason       = "UDDI_INFRA_HOST_TAG_KEY_1 and UDDI_INFRA_HOST_TAG_VALUE_1 environment variables must be set for this test to run"

  prerequisites_hcl = <<-PREREQ
  resource "infoblox_dns_server" "test" {
      uddi = {
          name = "{{random}}"
      }
  }
    PREREQ

  filter {
    type = "tag_filters"
    values = {
      "{{uddi_infra_host_tag_key_1}}" = "{{uddi_infra_host_tag_value_1}}"
    }
  }

  pair_checks = []


  step {
    id = "{{uddi_infra_host_legacy_id_1}}"
    uddi {
      server = infoblox_dns_server.test.id
    }
  }

}
