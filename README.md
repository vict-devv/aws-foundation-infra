# AWS Foundation Infrastructure

This repository contains the Terraform code for the AWS Foundation Infrastructure, which provides a baseline set of resources and configurations to support the deployment of applications and services on AWS. The infrastructure includes components such as VPCs, subnets, security groups, IAM roles, and other foundational resources necessary for building and operating applications in the cloud.

## Important Note

I don't want to incur costs as I'm using my personal AWS account for this project. Therefore, I have intentionally limited the resources provisioned by this Terraform configuration to ensure that it remains within the Free-Tier limits, or don't have cost at all. 

This means that the infrastructure may not include all the components typically found in a production environment, but it will provide a solid foundation for learning and experimentation without incurring additional costs.

## Getting Started

To get started with the AWS Foundation Infrastructure, follow these steps:

1. Clone the repository:
   ```bash
   git clone
   ```
2. Navigate to the project directory:
   ```bash
   cd aws-foundation-infrastructure
   ```
3. Review the Terraform configuration files to understand the resources being provisioned.
4. Initialize the Terraform working directory:
   ```bash
   terraform init
   ```
5. Plan the infrastructure deployment to see the changes that will be made:
   ```bash
   terraform plan
   ```
6. Apply the Terraform configuration to provision the infrastructure:
   ```bash
   terraform apply
   ```
