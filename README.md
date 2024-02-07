# Impossible Cloud SRE Code Challenge

> This repository contains the recipes for the Impossible Cloud SRE Code Challenge

## How To

### Create EKS

```sh
export AWS_DEFAULT_REGION='us-east-2'
export AWS_ACCESS_KEY_ID=VVVVVVVVVVVVVVVVVVVVV
export AWS_SECRET_ACCESS_KEY=XXXXXXXXXXXXXXXXXXXXX
export TF_VAR_region=$AWS_DEFAULT_REGION

terraform -chdir=eks init &&\
terraform -chdir=eks plan &&\
terraform -chdir=eks apply -auto-approve

aws eks \
--region $(terraform -chdir=eks output -raw region) \
update-kubeconfig \
--name $(terraform -chdir=eks output -raw cluster_name)
```

### Build and Push the Sample App to ECR

```sh
export AWS_ACCOUNT_ID=$(aws sts get-caller-identity --query 'Account' --output text)
export ECR_REGISTRY="${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com"
aws ecr get-login-password --region "$AWS_DEFAULT_REGION" | docker login --username AWS --password-stdin "$ECR_REGISTRY"

export IMAGE_NAME=$(terraform -chdir=eks output -raw aws_ecr_repository_app_url):latest

docker build -t $IMAGE_NAME app/ && docker push $IMAGE_NAME

export IMAGE_NAME=$(terraform -chdir=eks output -raw aws_ecr_repository_load_url):latest

docker build -t $IMAGE_NAME app/ && docker push $IMAGE_NAME

```

### Deploy Sample App to EKS cluster

```sh
helm install \
--set 'nodeSelector.eks\.amazonaws\.com/nodegroup=node-group-1-20240207114437297600000019' \ # CHANGE ME
-f helm/app/values.yaml app helm/app
```

### Run LoadTest Kubernetes Job

```sh
kubectl apply -f load/manifest/job.yaml
```

### Clean up

```sh
terraform -chdir=eks destroy
```
