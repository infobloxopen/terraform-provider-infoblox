case "basic" {
  backend           = "uddi"
  min_tf_version    = "1.14.0"
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
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }

}

case "filters" {
  backend           = "uddi"
  min_tf_version    = "1.14.0"
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
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        server = "uddi.server"
      }
    }
  }

}

case "tag_filters" {
  backend           = "uddi"
  min_tf_version    = "1.14.0"
  skip_if_env_empty = ["UDDI_INFRA_HOST_TAG_KEY_1", "UDDI_INFRA_HOST_TAG_VALUE_1"]
  skip_reason       = "UDDI_INFRA_HOST_TAG_KEY_1 and UDDI_INFRA_HOST_TAG_VALUE_1 environment variables must be set for this test to run"
  prerequisites_hcl = <<-PREREQ
    data "infoblox_infra_hosts" "test" {
        filters = {
            display_name = "{{uddi_infra_host_display_name_1}}"
            }
        }
    resource "infoblox_dns_server" "test" {
        name = {{random}}
    }
      PREREQ

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "tag_filters"
      values = {
        "{{uddi_infra_host_tag_key_1}}" = "{{uddi_infra_host_tag_value_1}}"
      }
    }
  }

}
