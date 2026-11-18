# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.

# __generated__ by Terraform
resource "infoblox_forward_looking_delegation" "list_forward_looking_delegation_with_resource_0" {
  provider = infoblox
  uddi = {
    address           = "10.240.0.2"
    cidr              = 31
    federated_pool_id = "federation/federated_pool/5b0ebe4f-90c4-4d50-9c37-9e3d3833d1d3"
    federated_realms  = ["federation/federated_realm/70e8baef-a428-4bef-bb50-8c6813d01781"]
    name              = null
    tags = {
      tf312 = "run1"
    }
  }
  update_trigger = null
}

import {
  to       = infoblox_forward_looking_delegation.list_forward_looking_delegation_with_resource_0
  provider = infoblox
  identity = {
    id = "federation/forward_looking_delegation/09966623-26ea-4bdb-9a13-9c74c4baa453"
  }
}

resource "infoblox_forward_looking_delegation" "list_forward_looking_delegation_with_resource_1" {
  provider = infoblox
  uddi = {
    address           = "10.240.0.0"
    cidr              = 31
    federated_pool_id = "federation/federated_pool/5b0ebe4f-90c4-4d50-9c37-9e3d3833d1d3"
    federated_realms  = ["federation/federated_realm/70e8baef-a428-4bef-bb50-8c6813d01781"]
    name              = null
    tags = {
      tf312 = "run1"
    }
  }
  update_trigger = null
}

import {
  to       = infoblox_forward_looking_delegation.list_forward_looking_delegation_with_resource_1
  provider = infoblox
  identity = {
    id = "federation/forward_looking_delegation/c07ce5d2-b6df-4d02-a5e8-64d2288300f4"
  }
}

resource "infoblox_forward_looking_delegation" "list_forward_looking_delegation_with_resource_2" {
  provider = infoblox
  uddi = {
    address          = "10.10.0.0"
    cidr             = 26
    federated_realms = ["federation/federated_realm/4ec76abc-f328-482d-a221-2ed7138d28dd"]
    name             = null
  }
  update_trigger = null
}

import {
  to       = infoblox_forward_looking_delegation.list_forward_looking_delegation_with_resource_2
  provider = infoblox
  identity = {
    id = "federation/forward_looking_delegation/34748fcf-8aec-419a-bb7b-011a7aaf21c6"
  }
}

resource "infoblox_forward_looking_delegation" "list_forward_looking_delegation_with_resource_3" {
  provider = infoblox
  uddi = {
    address          = "10.10.1.0"
    cidr             = 26
    federated_realms = ["federation/federated_realm/4ec76abc-f328-482d-a221-2ed7138d28dd"]
    name             = null
    tags = {
      site = "Site A"
    }
  }
  update_trigger = null
}

import {
  to       = infoblox_forward_looking_delegation.list_forward_looking_delegation_with_resource_3
  provider = infoblox
  identity = {
    id = "federation/forward_looking_delegation/3a55f2cf-f851-4314-92ca-88a3fd60f4a4"
  }
}

resource "infoblox_forward_looking_delegation" "list_forward_looking_delegation_with_resource_4" {
  provider = infoblox
  uddi = {
    address          = "10.10.1.0"
    cidr             = 26
    federated_realms = ["federation/federated_realm/cd8eca45-b7f2-41fb-9d83-aa4884076f52"]
    name             = null
    tags = {
      site = "Site A"
    }
  }
  update_trigger = null
}

import {
  to       = infoblox_forward_looking_delegation.list_forward_looking_delegation_with_resource_4
  provider = infoblox
  identity = {
    id = "federation/forward_looking_delegation/28357bca-6d13-40fa-a664-4eecdc1bf0cd"
  }
}

resource "infoblox_forward_looking_delegation" "list_forward_looking_delegation_with_resource_5" {
  provider = infoblox
  uddi = {
    address          = "10.10.1.128"
    cidr             = 26
    federated_realms = ["federation/federated_realm/4ec76abc-f328-482d-a221-2ed7138d28dd"]
    name             = null
  }
  update_trigger = null
}

import {
  to       = infoblox_forward_looking_delegation.list_forward_looking_delegation_with_resource_5
  provider = infoblox
  identity = {
    id = "federation/forward_looking_delegation/1cbdd3e1-49f5-4bf9-bfad-36cf5331176d"
  }
}

resource "infoblox_forward_looking_delegation" "list_forward_looking_delegation_with_resource_6" {
  provider = infoblox
  uddi = {
    address          = "10.10.1.192"
    cidr             = 26
    federated_realms = ["federation/federated_realm/4ec76abc-f328-482d-a221-2ed7138d28dd"]
    name             = null
  }
  update_trigger = null
}

import {
  to       = infoblox_forward_looking_delegation.list_forward_looking_delegation_with_resource_6
  provider = infoblox
  identity = {
    id = "federation/forward_looking_delegation/7247e9b8-4787-41c8-9743-6d64c86ad2aa"
  }
}

resource "infoblox_forward_looking_delegation" "list_forward_looking_delegation_with_resource_7" {
  provider = infoblox
  uddi = {
    address          = "2001:db8::"
    cidr             = 48
    federated_realms = ["federation/federated_realm/4ec76abc-f328-482d-a221-2ed7138d28dd"]
    name             = null
    tags = {
      site = "Site A"
    }
  }
  update_trigger = null
}

import {
  to       = infoblox_forward_looking_delegation.list_forward_looking_delegation_with_resource_7
  provider = infoblox
  identity = {
    id = "federation/forward_looking_delegation/da233642-88d3-4da9-ab4e-fa0c74b74ac5"
  }
}

resource "infoblox_forward_looking_delegation" "list_forward_looking_delegation_with_resource_8" {
  provider = infoblox
  uddi = {
    address          = "2001:db8::"
    cidr             = 48
    federated_realms = ["federation/federated_realm/cd8eca45-b7f2-41fb-9d83-aa4884076f52"]
    name             = null
    tags = {
      site = "Site A"
    }
  }
  update_trigger = null
}

import {
  to       = infoblox_forward_looking_delegation.list_forward_looking_delegation_with_resource_8
  provider = infoblox
  identity = {
    id = "federation/forward_looking_delegation/88e40412-f1f1-4322-802c-824c9e9b58b8"
  }
}

resource "infoblox_forward_looking_delegation" "list_forward_looking_delegation_with_resource_9" {
  provider = infoblox
  uddi = {
    address          = "10.10.0.64"
    cidr             = 26
    federated_realms = ["federation/federated_realm/4ec76abc-f328-482d-a221-2ed7138d28dd"]
    name             = null
  }
  update_trigger = null
}

import {
  to       = infoblox_forward_looking_delegation.list_forward_looking_delegation_with_resource_9
  provider = infoblox
  identity = {
    id = "federation/forward_looking_delegation/a243fa66-9f89-4b2a-af5f-c4d5929eaca1"
  }
}

resource "infoblox_forward_looking_delegation" "list_forward_looking_delegation_with_resource_10" {
  provider = infoblox
  uddi = {
    address          = "10.10.1.64"
    cidr             = 26
    federated_realms = ["federation/federated_realm/4ec76abc-f328-482d-a221-2ed7138d28dd"]
    name             = null
    tags = {
      site = "Site A"
    }
  }
  update_trigger = null
}

import {
  to       = infoblox_forward_looking_delegation.list_forward_looking_delegation_with_resource_10
  provider = infoblox
  identity = {
    id = "federation/forward_looking_delegation/2ef5d60f-94b7-4571-a9f9-ccddbcc79b44"
  }
}

resource "infoblox_forward_looking_delegation" "list_forward_looking_delegation_with_resource_11" {
  provider = infoblox
  uddi = {
    address          = "10.10.1.64"
    cidr             = 26
    federated_realms = ["federation/federated_realm/cd8eca45-b7f2-41fb-9d83-aa4884076f52"]
    name             = null
    tags = {
      site = "Site A"
    }
  }
  update_trigger = null
}

import {
  to       = infoblox_forward_looking_delegation.list_forward_looking_delegation_with_resource_11
  provider = infoblox
  identity = {
    id = "federation/forward_looking_delegation/667e3bab-eec1-4201-8d01-604d0ce4f170"
  }
}

resource "infoblox_forward_looking_delegation" "list_forward_looking_delegation_with_resource_12" {
  provider = infoblox
  uddi = {
    address          = "2001:db8:1::"
    cidr             = 48
    federated_realms = ["federation/federated_realm/4ec76abc-f328-482d-a221-2ed7138d28dd"]
    name             = null
    tags = {
      site = "Site A"
    }
  }
  update_trigger = null
}

import {
  to       = infoblox_forward_looking_delegation.list_forward_looking_delegation_with_resource_12
  provider = infoblox
  identity = {
    id = "federation/forward_looking_delegation/abf28977-9049-455c-91f9-b309b42d0866"
  }
}

resource "infoblox_forward_looking_delegation" "list_forward_looking_delegation_with_resource_13" {
  provider = infoblox
  uddi = {
    address          = "2001:db8:1::"
    cidr             = 48
    federated_realms = ["federation/federated_realm/cd8eca45-b7f2-41fb-9d83-aa4884076f52"]
    name             = null
    tags = {
      site = "Site A"
    }
  }
  update_trigger = null
}

import {
  to       = infoblox_forward_looking_delegation.list_forward_looking_delegation_with_resource_13
  provider = infoblox
  identity = {
    id = "federation/forward_looking_delegation/e33dbedc-39ac-4b87-9365-c2cd3ad08c2c"
  }
}


