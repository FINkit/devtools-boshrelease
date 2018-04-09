resource "google_compute_firewall" "jumpbox-to-nexus" {
  name    = "${var.env_id}-jumpbox-to-nexus"
  network = "${google_compute_network.bbl-network.name}"
  description = "Jumpbox to Nexus for test jobs"

  source_tags = ["${var.env_id}-jumpbox"]

  allow {
    ports    = ["8081"]
    protocol = "tcp"
  }

  target_tags = ["${var.env_id}-internal", "${var.env_id}-bosh-director"]
}
