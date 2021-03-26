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

# A role that does not get created because `create=false`
module "no_create_role" {
  source    = "../../"
  name      = "no-create-role"
  namespace = "default"
  create    = false
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
