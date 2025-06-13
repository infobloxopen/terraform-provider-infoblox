provider "google" {
  project = var.project_id
}

locals {
  network_interfaces = [ for i, n in var.networks : {
    network     = n,
    subnetwork  = length(var.sub_networks) > i ? element(var.sub_networks, i) : null
    external_ip = length(var.external_ips) > i ? element(var.external_ips, i) : "NONE"
    }
  ]
  
}

resource "google_compute_instance" "instance" {
  name                = var.goog_cm_deployment_name
  machine_type        = var.machine_type
  labels              = jsondecode(var.labels)
  #deletion_protection = var.deletion_protection
  zone                = var.zone

  tags = ["${var.goog_cm_deployment_name}-deployment"]

  boot_disk {
    initialize_params {
      image  = var.source_image
      type   = var.boot_disk_type
      size   = var.boot_disk_size
    }
  }
  
  dynamic "network_interface" {
    for_each = local.network_interfaces
    content {
      network = network_interface.value.network
      subnetwork = network_interface.value.subnetwork

      dynamic "access_config" {
        for_each = network_interface.value.external_ip == "NONE" ? [] : [1]
        content {
          nat_ip = network_interface.value.external_ip == "EPHEMERAL" ? null : network_interface.value.external_ip
        }
      }
    }
  }

  service_account {
    email = "default"
    scopes = compact([
      "https://www.googleapis.com/auth/cloud.useraccounts.readonly",
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring.write"
      ,var.enable_cloud_api == true ? "https://www.googleapis.com/auth/cloud-platform" : null
    ])
  }

##uncomment for metadata to apply 
  #metadata = var.metadata

}

##this will open ports for the firewall using network tag 
resource "google_compute_firewall" "allow_ports" {
  for_each = {
    "tcp-22"    = { protocol = "tcp", port = "22" }
    "udp-53"    = { protocol = "udp", port = "53" }
    "tcp-443"   = { protocol = "tcp", port = "443" }
    "udp-1194"  = { protocol = "udp", port = "1194" }
    "udp-2114"  = { protocol = "udp", port = "2114" }
    "tcp-8787"  = { protocol = "tcp", port = "8787" }
  }

  name    = "${var.goog_cm_deployment_name}-allow-${each.key}"
  network = element(var.networks, 0)  # Uses the first network in the list

  allow {
    protocol = each.value.protocol
    ports    = [each.value.port]
  }

  source_ranges = ["0.0.0.0/0"]  # Open to all (Modify for security!)

  target_tags = ["${var.goog_cm_deployment_name}-deployment"]
}
