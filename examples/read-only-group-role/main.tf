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

resource "kubernetes_namespace" "test" {
  metadata {
    name = "test"
  }
}

# Create a read-only role in a "test" namespace
module "read_only_role" {
  source    = "../../"
  name      = "read-only-role"
  namespace = "test"
  rules = [{
    api_groups = [""]
    verbs      = ["get", "list", "watch"]
    resources  = ["*"]
    }, {
    api_groups = ["extensions"]
    verbs      = ["get", "list", "watch"]
    resources  = ["*"]
    }, {
    api_groups = ["apps"]
    verbs      = ["get", "list", "watch"]
    resources  = ["*"]
    }
  ]
  subjects = [{
    kind      = "ServiceAccount"
    name      = "default"
    namespace = "kube-system"
  }]
  depends_on = [
    kubernetes_namespace.test
  ]
}
