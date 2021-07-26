# Infoblox IPAM Driver for Terraform

## Resources

There are resources for the following objects, supported by the plugin:

-   Network view
-   Network container
-   Network
-   A-record
-   AAAA-record
-   PTR-record
-   CNAME-record
-   Host record

Network container and network resources have two versions: IPv4 and IPv6. In
addition, there are two operations which are implemented as resources:
IP address allocation and IP address association with a network host
(ex. VM in a cloud environment); they have two versions as well: IPv4
and IPv6.

To work with DNS records a user must ensure that appropriate DNS zones
exist on the NIOS side, because currently the plugin does not support
creating a DNS zone.

Every resource has common attributes: 'comment' and 'ext_attrs'.
'comment' is text which describes the resource. 'ext_attrs' is a set of
NIOS Extensible Attributes attached to the resource, read more on this
attribute in a separate clause.

For DNS-related resources there is 'ttl' attribute as well, it specifies
TTL value (in seconds) for appropriate record. There is no default
value, zone's TTL is used if omitted. TTL value of 0 (zero) means
caching should be disabled for this record.

All the resources have 'comment' and 'ext_attrs' attributes,
additionally DNS-related records have 'ttl' attribute. They are all
optional. In this document, a resource's description implies that there
may be no explicit note in the appropriate clauses.

## Data sources

There are data sources for the following objects:

- A-record
- CNAME-record
- IPv4 Network
