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
    id = "dns/host/{{uddi_infra_host_legacy_id_1}}"
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
    id = "dns/host/{{uddi_infra_host_legacy_id_1}}"
    uddi {
      server        = infoblox_dns_server.test.id
      absolute_name = "{{random}}."
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        absolute_name = "uddi.absolute_name"
      }
    }
  }

}

case "tag_filters" {
  backend           = "uddi"
  min_tf_version    = "1.14.0"
  skip              = true
  skip_if_env_empty = ["UDDI_INFRA_HOST_DISPLAY_NAME_1", "UDDI_DNS_SERVICE_TAG_KEY_1", "UDDI_DNS_SERVICE_TAG_VALUE_1"]
  skip_reason       = "UDDI_INFRA_HOST_DISPLAY_NAME_1, UDDI_DNS_SERVICE_TAG_KEY_1 and UDDI_DNS_SERVICE_TAG_VALUE_1 environment variables must be set for this test to run"

  prerequisites_hcl = <<-PREREQ
    resource "infoblox_dns_server" "test" {
        uddi = {
            name = "{{random}}"
        }
    }
      PREREQ

  step {
    id = "dns/host/{{uddi_infra_host_legacy_id_1}}"
    uddi {
      server        = infoblox_dns_server.test.id
      absolute_name = "{{random}}."
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "tag_filters"
      values = {
        "location" = "uddi.tags.location"
      }
    }
  }

}
