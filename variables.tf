variable "project_name" {
  description = "The name of the project, used for tagging resources"
  type        = string
  default     = "victor-portfolio"
}

variable "cidr_block" {
  description = "The CIDR block for the VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "api_ecr_repository_name" {
  description = "The name of the ECR repository for the API"
  type        = string
  default     = "cloud-api"
}
