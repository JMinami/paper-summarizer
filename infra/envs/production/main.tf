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
    "iamcredentials.googleapis.com",
    "run.googleapis.com",
    "iam.googleapis.com",
  ]
}

resource "google_project_iam_binding" "project" {
  project = var.project_id
  role    = "roles/run.admin"
  members = [
    "serviceAccount:github-actions@paper-summarizer-381022.iam.gserviceaccount.com"
  ]
}

resource "google_project_iam_binding" "storage" {
  project = var.project_id
  role    = "roles/storage.admin"
  members = [
    "serviceAccount:github-actions@paper-summarizer-381022.iam.gserviceaccount.com"
  ]
}

resource "google_project_iam_binding" "iam" {
  project = var.project_id
  role    = "roles/iam.serviceAccountUser"
  members = [
    "serviceAccount:github-actions@paper-summarizer-381022.iam.gserviceaccount.com"
  ]
}

resource "google_project_service" "translation_api" {
  project  = var.project_id
  for_each = toset(local.services_to_enable)

  service                    = each.value
  disable_dependent_services = true
  disable_on_destroy         = false
}


resource "google_service_account" "github-actions" {
  project      = var.project_id
  account_id   = "github-actions"
  display_name = "A service account for GitHub Actions"
}



resource "google_service_account_iam_binding" "github-actions" {
  service_account_id = google_service_account.github-actions.name
  role               = "roles/iam.workloadIdentityUser"
  members = [
    "principalSet://iam.googleapis.com/projects/571905713425/locations/global/workloadIdentityPools/github-actions/*"
  ]
}

resource "google_project_service" "project" {
  project = var.project_id
  service = "iamcredentials.googleapis.com"
}

resource "google_iam_workload_identity_pool" "github-actions" {
  workload_identity_pool_id = "github-actions"
}

resource "google_iam_workload_identity_pool_provider" "github-actions-provider" {
  workload_identity_pool_id          = google_iam_workload_identity_pool.github-actions.workload_identity_pool_id
  workload_identity_pool_provider_id = "github-actions-provider"
  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.repository" = "assertion.repository"
    "attribute.actor"      = "assertion.actor"
    "attribute.aud"        = "assertion.aud"
  }
  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
}