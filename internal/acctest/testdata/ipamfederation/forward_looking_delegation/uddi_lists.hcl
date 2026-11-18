# ForwardLookingDelegation — uddi list cases
case "basic" {
  backend        = "uddi"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm" "test" {
    uddi = {
      name = "{{random}}-realm"
    }
  }
  PREREQ

  step {
    uddi {
      address          = "{{random_ipv4_network}}"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
    }
  }

  step {
    query    = true
    provider = infoblox
    limit    = 5
  }
}

case "filters" {
  backend        = "uddi"
  min_tf_version = "1.14.0"
  prerequisites_hcl = <<-PREREQ
  resource "infoblox_federated_realm" "test" {
    uddi = {
      name = "{{random}}-realm"
    }
  }
  PREREQ

  step {
    uddi {
      address          = "{{random_ipv4_network}}"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
      name             = "{{random}}"
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
      name = "{{random}}-realm"
    }
  }
  PREREQ

  step {
    uddi {
      address          = "{{random_ipv4_network}}"
      cidr             = 16
      federated_realms = [infoblox_federated_realm.test.id]
      tags             = { tag1 = "{{random2}}" }
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
