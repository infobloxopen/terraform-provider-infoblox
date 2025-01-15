# Zone Auth Resource

The `infoblox_zone_auth` resource associates authoritative zone with a DNS View.The resource represents the ‘zone_auth’ WAPI object in NIOS.

The following list describes the parameters you can define in the resource block of the zone auth object:

* `fqdn`: required, specifies the name of this DNS zone. For a reverse zone, this is in “address/cidr” format.
For other zones, this is in FQDN format. This value can be in unicode format.
Example: `10.1.0.0/24` for reverse zone and `zone1.com` for forward zone.
* `view`: optional, specifies The name of the DNS view in which the zone resides. If value is not specified, `default` will be considered as default DNS view Example: `external`.
* `zone_format`: optional, determines the format of corresponding zone. Valid values are `FORWARD`, `IPV4` and `IPV6`. Default value: `FORWARD`.
* `ns_group`: optional, specifies the name server group that serves DNS for this zone. Example: `demoGrp`.
* `restart_if_needed`: optional, restarts the member service. It is boolean value, based on requirement value changes.
* `soa_default_ttl`: The Time to Live (TTL) value of the SOA record of this zone. This value is the number of seconds that data is cached. Default value: `28800`.
* `soa_expire`: This setting defines the amount of time, in seconds, after which the secondary server stops giving out answers about the zone because the zone data is too old to be useful. Default value: `2419200`.
* `soa_negative_ttl`: The negative Time to Live (TTL) value of the SOA of the zone indicates how long a secondary server can cache data for “Does Not Respond” responses. Default value: `900`.
* `soa_refresh`: This indicates the interval at which a secondary server sends a message to the primary server for a zone to check that its data is current, and retrieve fresh data if it is not. Default value: `10800`.
* `soa_retry`: This indicates how long a secondary server must wait before attempting to recontact the primary server after a connection failure between the two servers occurs. Default value: `3600`.
* `comment`: optional, description of the zone. Example: `custom reverse zone`.
* `ext_attrs`: optional, set of the Extensible attributes of the zone, as a map in JSON format. Example: `jsonencode({})`.

!> For a reverse zone, the corresponding 'zone_format' value should be set. And 'fqdn' once set cannot be updated.

### Examples of a Zone Auth Block

```hcl
# Forward mapping zone, with minimal set of parameters
resource "infoblox_zone_auth" "zone1" {
  fqdn = "test3.com"
  view = "default"
  zone_format = "FORWARD"
  comment = "Zone Auth created newly"
  ext_attrs = jsonencode({
    Location = "AcceptanceTerraform"
  })
}

# IPV4 reverse mapping zone, with full set of parameters
resource "infoblox_zone_auth" "zone2" {
  fqdn = "10.0.0.0/24"
  view = "default"
  zone_format = "IPV4"
  ns_group = "nsgroup1"
  restart_if_needed = true
  soa_default_ttl = 37000
  soa_expire = 92000
  soa_negative_ttl = 900
  soa_refresh = 2100
  soa_retry = 800
  comment = "IPV4 reverse zone auth created"
  ext_attrs = jsonencode({
    Location = "TestTerraform"
  })
}

# IPV6 reverse mapping zone, with minimal set of parameters
resource "infoblox_zone_auth" "zone3" {
  fqdn = "2002:1100::/64"
  view = "non_defaultview"
  zone_format = "IPV6"
  ns_group = "nsgroup2"
  comment = "IPV6 reverse zone auth created"
  ext_attrs = jsonencode({
    Location = "Random TF location"
  })
}
```
