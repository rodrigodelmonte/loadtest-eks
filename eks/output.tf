output "cluster_endpoint" {
  description = "Endpoint for EKS control plane"
  value       = module.eks.cluster_endpoint
}

output "cluster_security_group_id" {
  description = "Security group ids attached to the cluster control plane"
  value       = module.eks.cluster_security_group_id
}

output "region" {
  description = "AWS region"
  value       = var.region
}

output "cluster_name" {
  description = "Kubernetes Cluster Name"
  value       = module.eks.cluster_name
}


output "aws_ecr_repository_app_url" {
  description = "ECR Repository App URL"
  value       = aws_ecr_repository.app.repository_url
}

output "aws_ecr_repository_load_url" {
  description = "ECR Repository App URL"
  value       = aws_ecr_repository.loadtest.repository_url
}
