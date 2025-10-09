variable "content" {
  description = "Content to write to the file"
  type        = string
  default     = "Hello from Terratest!"
}

variable "filename" {
  description = "Name of the file to create"
  type        = string
  default     = "example.txt"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "test"
}
