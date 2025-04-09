//example for fixed address with maximal parameters and using next available ip function
//ipv4addr not specified and network is given so next available ip in the network will be allocated
resource "infoblox_ipv4_fixed_address" "fix4"{
    client_identifier_prepend_zero=true
    comment= "fixed address"
    dhcp_client_identifier="23"
    disable= true
    ext_attrs = jsonencode({
        "Site": "Blr"
    })
    match_client = "CLIENT_ID"
    name = "fixed_address_1"
    network = "18.0.0.0/24"
    network_view = "default"
    options {
        name         = "dhcp-lease-time"
        value        = "43200"
        vendor_class = "DHCP"
        num          = 51
        use_option   = true
    }
    options {
        name = "routers"
        num = "3"
        use_option = true
        value = "18.0.0.2"
        vendor_class = "DHCP"
    }
    use_options = true
    depends_on=[infoblox_ipv4_network.net4]
}
resource "infoblox_ipv4_network" "net4" {
    cidr = "18.0.0.0/24"
}
//creates a fixed address by explicitly providing the `ipv4addr` value instead of using the next available IP function.
resource "infoblox_ipv4_fixed_address" "fix3"{
    ipv4addr        = "17.0.0.9"
    mac = "00:0C:24:2E:8F:2A"
    options {
        name         = "dhcp-lease-time"
        value        = "43200"
        vendor_class = "DHCP"
        num          = 51
        use_option   = true
    }
    depends_on=[infoblox_ipv4_network.net5]
}
resource "infoblox_ipv4_network" "net5" {
    cidr = "17.0.0.0/24"
}