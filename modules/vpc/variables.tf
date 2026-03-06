variable "project_name" {
  description = "The name of the project, used for tagging resources"
  type        = string
}

variable "cidr_block" {
    description = "The CIDR block for the VPC"
    type        = string
}