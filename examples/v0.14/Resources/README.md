* Infoblox NIOS Resource Examples are provided under Infoblox directory. Going forward this would be a standard practise. The cloud provider(AWS, VMWare and Azure) directories would contain an use case rather than a specific resource example. It would be a addressed in subsequent release.

* Infoblox strongly recommends that you use the infoblox_ip_association resource enhanced to provide better support for managing Host records created using the infoblox_ip_allocation resource. The Infoblox_ipv4_association and Infoblox_ipv6_association resources will be deprecated in an upcoming release and will not be supported.
* If the IP address allocation operation was done using the infoblox_ip_allocation resource, for the IP address association operation, you are advised to use the infoblox_ip_association resource only and not infoblox_ipv4_association or infoblox_ipv6_association resource in a .tf file.

# While using IPv6 Features make sure to consider the following:
* IPv6 Allocation and Association through Host Record mandates unique DUID in .tf file.
* For all the cloud providers(AWS, Azure and VMWare), DUID will be updated with MAC address of the interface in NIOS.
