terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.0.0"
    }
  }
  required_version = "~> 0.14.0"
}

provider "kubernetes" {
  config_path = "~/.kube/config"
}

# Create an admin Role in the default namespace
module "admin-role" {
  source    = "../../"
  name      = "admin-role"
  namespace = "default"
  rules = [{
    api_groups = ["*"]
    verbs      = ["*"]
    resources  = ["*"]
  }]
  subjects = [{
    kind = "User"
    name = "Admin"
  }]
}
