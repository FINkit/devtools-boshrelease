resource "google_compute_firewall" "jumpbox-to-gerrit" {
  name        = "${var.env_id}-jumpbox-to-gerrit"
  network     = "${google_compute_network.bbl-network.name}"
  description = "Jumpbox to Gerrit for test jobs"

  source_tags = ["${var.env_id}-jumpbox"]

  allow {
    ports    = ["8080"]
    protocol = "tcp"
  }

  target_tags = ["${var.env_id}-internal", "${var.env_id}-bosh-director"]
}
