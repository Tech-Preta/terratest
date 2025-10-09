output "file_path" {
  description = "Path to the created file"
  value       = local_file.example.filename
}

output "file_content" {
  description = "Content of the file"
  value       = local_file.example.content
}

output "environment" {
  description = "Environment name"
  value       = var.environment
}
