provider "azurerm" {
    features {}
}

resource "azurerm_resource_group" "rg1" {
    name = "${local.res_prefix}_rg1"
    location = "ukwest"
}

resource "azurerm_virtual_network" "vnet1" {
    location = "ukwest"
    resource_group_name = azurerm_resource_group.rg1.name

    name = "${local.res_prefix}_vnet1"
    address_space =  [infoblox_ipv4_network_container.v4nc_1.cidr]
}

resource "azurerm_subnet" "net1" {
    resource_group_name = azurerm_resource_group.rg1.name
    virtual_network_name = azurerm_virtual_network.vnet1.name

    name = "${local.res_prefix}_net1"
    address_prefixes =[infoblox_ipv4_network.subnet1.cidr]
}

resource "azurerm_public_ip" "pub_addr1" {
    resource_group_name = azurerm_resource_group.rg1.name
    location = "ukwest"

    name = "${local.res_prefix}_pub_addr1"
    allocation_method = "Dynamic"
    domain_name_label = "a132"
}

resource "azurerm_network_interface" "ni1" {
    resource_group_name = azurerm_resource_group.rg1.name
    location = "ukwest"

    name = "${local.res_prefix}_ni1"
    ip_configuration {
        name = "${local.res_prefix}_ipconfiguration1"
        subnet_id = azurerm_subnet.net1.id
        private_ip_address_allocation = "Static"
        private_ip_address =infoblox_ipv4_allocation.alloc1.ip_addr

        public_ip_address_id = azurerm_public_ip.pub_addr1.id
    }
}

resource "azurerm_virtual_machine" "vm1" {
    resource_group_name = azurerm_resource_group.rg1.name
    location = "ukwest"

    name = "${local.res_prefix}_vm1"
    network_interface_ids = [azurerm_network_interface.ni1.id]
    primary_network_interface_id = azurerm_network_interface.ni1.id
    vm_size = "Standard_A6"
    delete_os_disk_on_termination = true
    delete_data_disks_on_termination = true

    storage_image_reference {
        publisher = "Canonical"
        offer = "UbuntuServer"
        sku = "16.04-LTS"
        version = "latest"
    }

    storage_os_disk {
        name = "main_storage"
        caching = "ReadWrite"
        create_option = "FromImage"
    }

    os_profile {
        computer_name  = "ubuntu"
        admin_username = "ubuntu"
        admin_password="JKLhdsa&^52128"
    }

    os_profile_linux_config {
        disable_password_authentication = false
    }
}
