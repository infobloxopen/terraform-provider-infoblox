// Create an Admin User with required fields only
resource "infoblox_admin_user" "basic" {
  nios = {
    name         = "tf-admin-example"
    password     = "SecurePassword123!"
    admin_groups = ["admin-group"]
  }
}

// Create an Admin User with additional configuration
resource "infoblox_admin_user" "full" {
  nios = {
    name         = "tf-admin-full"
    password     = "SecurePassword123!"
    admin_groups = ["admin-group"]
    comment      = "Example admin user managed by Terraform"
    email        = "admin@example.com"
    disable      = false
    auth_type    = "LOCAL"
    auth_method  = "KEYPAIR"
    time_zone    = "US/Pacific"
    ext_attrs = {
      Site = "us-west"
    }
  }
}

// Create an Admin User with SSH key authentication
resource "infoblox_admin_user" "with_ssh_keys" {
  nios = {
    name         = "tf-admin-ssh"
    password     = "SecurePassword123!"
    admin_groups = ["admin-group"]
    auth_method  = "KEYPAIR"
    ssh_keys = [
      {
        key_name  = "my-ssh-key"
        key_type  = "RSA"
        key_value = "ssh-rsa AAAA...your-public-key-here..."
      }
    ]
  }
}

// Create an Admin User with certificate-based authentication
// Prerequisites: the CA certificate must exist in NIOS and its WAPI reference must be known.
resource "infoblox_admin_user" "with_cert_auth" {
  nios = {
    name                              = "tf-admin-cert"
    password                          = "SecurePassword123!"
    admin_groups                      = ["admin-group"]
    enable_certificate_authentication = true
    ca_certificate_issuer             = "cacertificate/ZXhhbXBsZQ:CN%3D%22example.com%22"
    client_certificate_serial_number  = "abc123def456"
  }
}
