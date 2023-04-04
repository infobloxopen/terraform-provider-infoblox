# SRV-record Data Source

Use the data source to retrieve the following information for an SRV-record from the corresponding object in NIOS:

* `priority`: a priority number, as described in RFC2782.
* `weight`: a weight number, as described in RFC2782.
* `zone`: the zone which the record belongs to.
* `ttl`: the "time to live" value of the record, in seconds. Example: `1800`.
* `comment`: the description of the record. This is a regular comment. Example: `spare node for the service`.
* `ext_attrs`: the set of extensible attributes of the record, if any. The content is formatted as a JSON map. Example: `{"Owner”: "State Library”, "Expires”: "never”}`.

The following list describes the parameters you must define in an `infoblox_srv_record` data source block:

* `dns_view`: optional, specifies the DNS view which the record's zone belongs to. If a value is not specified, the name `default` is used as the DNS view.
* `name`: required, specifies the record's name in the format, defined in RFC2782 document. Example: `_http._tcp.acme.com`
* `target`: required, specifies an FQDN of the host which is responsible for providing the service specified by `name`. Example: `www.acme.com`
* `port`: required, specifies a port number (0..65535) on the `target` host which the service expects requests on.

### Example of the SRV-record Data Source Block

```hcl
resource "infoblox_srv_record" "rec2" {
    dns_view = "nondefault_dnsview1"
    name = "_sip._udp.example2.org"
    priority = 12
    weight = 10
    port = 5060
    target = "sip.example2.org"
    ttl = 3600
    comment = "example SRV record"
    ext_attrs = jsonencode({
        "Location" = "65.8665701230204, -37.00791763398113"
    })
}

data "infoblox_srv_record" "ds1" {
    dns_view = "nondefault_dnsview1" // not 'default' thus must be specified
    name = "_sip._udp.example2.org"
    port = 5060
    target = "sip.example2.org"

    depends_on = [infoblox_srv_record.rec2]
}

output "srv_rec2_priority" {
  value = data.infoblox_srv_record.ds1.priority
}

output "srv_rec2_weight" {
  value = data.infoblox_srv_record.ds1.weight
}

output "srv_rec2_zone" {
  value = data.infoblox_srv_record.ds1.zone
}

output "srv_rec2_ttl" {
  value = data.infoblox_srv_record.ds1.ttl
}

output "srv_rec2_comment" {
  value = data.infoblox_srv_record.ds1.comment
}

output "srv_rec2_ext_attrs" {
  value = data.infoblox_srv_record.ds1.ext_attrs
}

```
