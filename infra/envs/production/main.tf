terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.57.0"
    }
  }
  backend "gcs" {
    bucket = "paper-summarizer-tfstate"
    prefix = "terraform/production"
  }
}

provider "google" {
  credentials = file(var.credentials_file_path)
  project     = "paper-summarizer-381022"
  region      = "asia-northeast1"
}

locals {
  services_to_enable = [
    # "cloudresourcemanager.googleapis.com",
    "translate.googleapis.com",
    "cloudbuild.googleapis.com",
    "run.googleapis.com"
  ]
}
resource "google_project_service" "translation_api" {
  project  = var.project_id
  for_each = toset(local.services_to_enable)

  service                    = each.value
  disable_dependent_services = true
  disable_on_destroy         = false
}