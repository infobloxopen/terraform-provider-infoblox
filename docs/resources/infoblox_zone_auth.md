# Authoritative Zone Resource

The `infoblox_zone_auth` resource creates a DNS zone with SOA and NS records for the given FQDN.

## Parameter Reference

The following list describes the parameters you can define in the resource block for the zone:

* `fqdn`: required, specifies the fully qualified domain name for the DNS zone. Example: `prod123.example.com`
* `ns_group`: required, name of Nameserver Group in Infoblox; will create NS records
* `restart_if_needed`: optional, restart the member service if necessary for changes to take effect
* `comment`: required, description of this Authoritative Zone; max 256 characters
* `soa_default_ttl`: optional, Time To Live (TTL) of the SOA record, in seconds; default `3600`
* `soa_expire`: optional, time in seconds for secondary servers to stop answering about the zone because the data is stale; default `2419200` (1 week)
* `soa_negative_ttl`: optional, time in seconds for secondary servers to cache data for 'Does Not Respond' responses; default `900`
* `soa_refresh`: optional, interval in seconds for secondary servers to check the primary server for fresh data about the zone; default `10800`
* `soa_retry`: optional, interval in seconds for secondary servers to wait before recontacting primary server about the zone after failure; default `3600`
* `ext_attrs`: optional, a set of NIOS extensible attributes that are attached to the record, using jsonencode().

## Example Usage

```hcl
resource "infoblox_zone_auth" "prod123_zone" {
  fqdn              = "prod123.example.com"
  ns_group          = "Prod_NS_Group"
  restart_if_needed = true
  comment           = "Managed by Terraform"
  soa_default_ttl   = 3600
  soa_expire        = 2419200
  soa_negative_ttl  = 900
  soa_refresh       = 10800
  soa_retry         = 3600
}
```

## Terraform Import

Authoritative Zone resources can be imported using the Zone Reference from the Infoblox NIOS API.

Example query to see the Zone Refs for an Infoblox instance. Find the Ref for your zone:

```
> curl -k -u $USER:$PASSWORD -X GET "https://infoblox-dns.example.com/wapi/v2.11/zone_auth?_return_as_object=1"
```

Then import into your Terraform state:

```
> terraform import infoblox_zone_auth.prod123_zone zone_auth/kZWZhdWx0LmlvLnZG5zLnpvbmUkLl9RrdG0udXxLnNhbmRib3gxMQ:prod123.example.com/default
```