# ReservedBlock — uddi list cases
case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_federated_realm" "test" {
      uddi = {
        name = "{{random}}"
      }
    }
    PREREQ

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 24
      federated_realm = infoblox_federated_realm.test.id
      name            = "{{random}}"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name = "uddi.name"
      }
    }
  }

}

case "filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_federated_realm" "test" {
      uddi = {
        name = "{{random}}"
      }
    }
    PREREQ

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 24
      federated_realm = infoblox_federated_realm.test.id
      name            = "{{random}}"
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "filters"
      values = {
        name = "uddi.name"
      }
    }
  }

}

case "tag_filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
    resource "infoblox_federated_realm" "test" {
      uddi = {
        name = "{{random}}"
      }
    }
    PREREQ

  step {
    uddi {
      address         = "{{random_ipv4_network}}"
      cidr            = 24
      federated_realm = infoblox_federated_realm.test.id
      tags            = { tag1 = "{{random2}}" }
    }
  }

  step {
    query            = true
    provider         = infoblox
    include_resource = true
    filter {
      type = "tag_filters"
      values = {
        tag1 = "uddi.tags.tag1"
      }
    }
  }

}
