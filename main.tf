terraform {
  required_version = ">= 1.0"
  required_providers {
    local = {
      source  = "hashicorp/local"
      version = "~> 2.0"
    }
  }
}

resource "local_file" "example" {
  content  = var.content
  filename = "${path.module}/output/${var.filename}"
}

resource "local_file" "info" {
  content  = "Environment: ${var.environment}\nGenerated at: ${timestamp()}"
  filename = "${path.module}/output/info.txt"
}
