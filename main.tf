module "network" {
  source       = "./modules/vpc"
  project_name = var.project_name
  cidr_block   = var.cidr_block
}
