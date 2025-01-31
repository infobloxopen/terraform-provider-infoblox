resource "infoblox_ip_association" "association1" {
  internal_id = infoblox_ip_allocation.allocation1.id # which of the existing host records to deal with

  # enable_dhcp = false // this is the default
  mac_addr = "12:00:43:fe:9a:8c" # the address will be assigned but DHCP configuration will not contain it
}

resource "infoblox_ip_association" "association2" {
  internal_id = infoblox_ip_allocation.allocation4.id # let's choose allocation resource with enable_dns = false

  # enable_dhcp = false # this is the default
  duid = "00:43:d2:0a:11:e6" # the address will be assigned but DHCP configuration will not contain it
}

resource "infoblox_ip_association" "association3" {
  internal_id = infoblox_ip_allocation.allocation5.id

  enable_dhcp = true # all systems go

  mac_addr = "12:43:fd:ba:9c:c9"
  duid = "00:43:d2:0a:11:e6"
}

resource "infoblox_ip_association" "association4" {
  internal_id = infoblox_ip_allocation.allocation3.id

  enable_dhcp = true # all systems go

  # DHCP will be enabled for IPv4 ...
  mac_addr = "09:01:03:d3:db:2a"

  # ... and disabled for IPv6
  # duid = "10:2a:9f:dd:3e:0a"
  # yes, DUID will be de-associated, but you can uncomment this later
  # if you will decide to enable DHCP for IPv6
}
