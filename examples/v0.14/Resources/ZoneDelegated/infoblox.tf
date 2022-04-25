# Zone Delegated

resource "infoblox_zone_delegated" "subdomain" {

  fqdn = "subdomain.example.com"

  delegate_to {
    address = "205.251.197.208"
    name = "ns-1488.awsdns-58.org"
  }

  delegate_to {
    address = "205.251.199.242"
    name = "ns-2034.awsdns-62.co.uk"
  }

}
