module "network" {
  source       = "./modules/vpc"
  project_name = "victor-portfolio"
  cidr_block   = "10.0.0.0/16"
}
