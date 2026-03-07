resource "aws_ecr_repository" "cloud-api" {
  name                 = var.api_ecr_repository_name
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_lifecycle_policy" "cloud_api_policy" {
  repository = aws_ecr_repository.cloud-api.name

  policy = jsonencode({
    rules = [
      {
        rulePriority = 1
        description  = "Keep only the last 5 images"
        selection    = {
          tagStatus    = "any"
          countType    = "imageCountMoreThan"
          countNumber  = 5
        }
        action       = {
          type = "expire"
        }
      }
    ]
  })
}