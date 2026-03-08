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

## Cleanup

To avoid incurring costs, make sure to destroy the infrastructure when you are done:

```bash
terraform destroy
```

## GoLang Sample API

In addition to the Terraform code for the AWS Foundation Infrastructure, this repository also includes a sample GoLang API that demonstrates how to build a simple RESTful API using Go. The API is designed to be lightweight and easy to deploy on the AWS infrastructure provisioned by the Terraform code.

Inside `go-api` directory, you will find the GoLang code for the sample API, there's also a Dockerfile for containerizing the application, which has been used by the `.github/workflows/deploy.yaml` GitHub Actions workflow, to run tests, build and push the Docker image to Amazon ECR.

### Running the GoLang API Locally

To run the GoLang API locally, follow these steps:

1. Navigate to the `go-api` directory:
   ```bash
   cd go-api
   ```
2. Make sure the dependencies are installed:
   ```bash
   go mod tidy
   ```
3. Run the API:
   ```bash
   go run cmd/server/main.go
   ```

Optionally if you want to run the Tests, you can use the following command:

```bash
go test ./...
```

### Building and Running the Docker Container

To build and run the Docker container for the GoLang API, follow these steps:

1. Navigate to the `go-api` directory:
   ```bash
   cd go-api
   ```
2. Build the Docker image:
   ```bash
   docker build -t cloud-api:latest .
   ```
3. Run the Docker container:
   ```bash
   docker run -p 8080:8080 cloud-api:latest
   ```

### Testing the Running API

Once the API is running, you can test it by sending a request to the endpoint. For example, you can use `curl` to send a GET request:

```bash
curl http://localhost:8080/health
```

You can change the server port that the API listen on by setting the `SERVER_PORT` environment variable, for example:

```bash
export SERVER_PORT=9090
go run cmd/server/main.go
```

Now the server will run on port 9090, instead of the default 8080.

Sample response:

```json
{
  "status": "UP",
  "uptime": "5.917135483s",
  "go_version": "go1.26.1",
  "timestamp": "2026-03-07T21:38:29.581957704-05:00"
}
```
