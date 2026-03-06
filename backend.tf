terraform {
  backend "s3" {
    bucket       = "victor-devops-tfstate-bucket"
    key          = "foundation/terraform.tfstate"
    region       = "ca-central-1"
    use_lockfile = true
    encrypt      = true
  }
}
