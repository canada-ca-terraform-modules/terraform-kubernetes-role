
# Part of a hack for module-to-module dependencies.
# https://github.com/hashicorp/terraform/issues/1178#issuecomment-449158607
# https://github.com/hashicorp/terraform/issues/1178#issuecomment-473091030
# Make sure to add this null_resource.dependency_getter to the `depends_on`
# attribute to all resource(s) that will be constructed first within this
# module:
resource "null_resource" "dependency_getter" {
  triggers = {
    my_dependencies = join(",", var.dependencies)
  }

  lifecycle {
    ignore_changes = [
      triggers["my_dependencies"],
    ]
  }
}

resource "kubernetes_role" "role" {

  metadata {
    name      = var.name
    namespace = var.namespace
  }

  # Users
  dynamic "rule" {
    for_each = var.rules
    content {
      api_groups = rule.api_groups
      verbs      = rule.verbs
      resources  = rule.resources
    }
  }

  depends_on = [
    null_resource.dependency_getter,
  ]
}

resource "kubernetes_role_binding" "role-binding" {

  metadata {
    name      = "${var.name}-binding"
    namespace = var.namespace
  }

  role_ref {
    kind      = "Role"
    name      = var.name
    api_group = "rbac.authorization.k8s.io"
  }

  # Role binding subject
  dynamic "subject" {
    for_each = var.subjects
    content {
      kind      = subject.kind
      name      = subject.name
      api_group = "rbac.authorization.k8s.io"
    }
  }

  depends_on = [
    null_resource.dependency_getter,
  ]
}

# Part of a hack for module-to-module dependencies.
# https://github.com/hashicorp/terraform/issues/1178#issuecomment-449158607
resource "null_resource" "dependency_setter" {
  depends_on = [

  ]
}
