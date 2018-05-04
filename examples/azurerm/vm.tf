# create a resource group if it doesn't exist
resource "azurerm_resource_group" "terraform" {
    name = "a132terraform"
    location = "ukwest"
}

# create virtual network
resource "azurerm_virtual_network" "vnet" {
    name = "tfvnet"
    address_space =  ["${infoblox_network.demo_network.cidr}"]
    location = "ukwest"
    resource_group_name = "${azurerm_resource_group.terraform.name}"
}

# create subnet
resource "azurerm_subnet" "subnet" {
    name = "tfsub"
    resource_group_name = "${azurerm_resource_group.terraform.name}"
    virtual_network_name = "${azurerm_virtual_network.vnet.name}"
    address_prefix ="10.0.0.0/24"
    #network_security_group_id = "${azurerm_network_security_group.nsg.id}"
}

# create public IPs
resource "azurerm_public_ip" "ip" {
    name = "tfip"
    location = "ukwest"
    resource_group_name = "${azurerm_resource_group.terraform.name}"
    public_ip_address_allocation = "dynamic"
    domain_name_label = "a132"

    tags {
        environment = "staging"
    }
}

# create network interface
resource "azurerm_network_interface" "ni" {
    name = "tfni"
    location = "ukwest"
    resource_group_name = "${azurerm_resource_group.terraform.name}"

    ip_configuration {
        name = "ipconfiguration"
        subnet_id = "${azurerm_subnet.subnet.id}"
        private_ip_address_allocation = "static"
        private_ip_address ="${infoblox_ip_allocation.demo_allocation.ip_addr}"
        public_ip_address_id = "${azurerm_public_ip.ip.id}"
    }
}

# create virtual machine
resource "azurerm_virtual_machine" "vm" {
    name = "${infoblox_ip_allocation.demo_allocation.host_name}"
    location = "ukwest"
    resource_group_name = "${azurerm_resource_group.terraform.name}"
    network_interface_ids = ["${azurerm_network_interface.ni.id}"]
    vm_size = "Standard_A6"

    storage_image_reference {
        publisher = "Canonical"
        offer = "UbuntuServer"
        sku = "16.04-LTS"
        version = "latest"
    }

    storage_os_disk {
        name = "myosdisk"
        caching = "ReadWrite"
        create_option = "FromImage"
    }

    os_profile {
        computer_name  = "ubuntu"
        admin_username = "ssastry"
        admin_password="pass12345!"
    }

    os_profile_linux_config {
        disable_password_authentication = false
           }

    tags {
        environment = "Terraform Demo"
    }
}

