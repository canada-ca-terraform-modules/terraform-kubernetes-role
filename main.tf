resource "kubernetes_role" "role" {
  count = var.create ? 1 : 0

  metadata {
    name      = var.name
    namespace = var.namespace
  }

  dynamic "rule" {
    for_each = var.rules
    content {
      api_groups = rule.value["api_groups"]
      verbs      = rule.value["verbs"]
      resources  = rule.value["resources"]
    }
  }
}

resource "kubernetes_role_binding" "role-binding" {
  count = var.create ? 1 : 0

  metadata {
    name      = "${var.name}-binding"
    namespace = var.namespace
  }

  role_ref {
    kind      = "Role"
    name      = var.name
    api_group = "rbac.authorization.k8s.io"
  }

  dynamic "subject" {
    for_each = var.subjects
    content {
      kind      = subject.value["kind"]
      name      = subject.value["name"]
      api_group = "rbac.authorization.k8s.io"
    }
  }
}
