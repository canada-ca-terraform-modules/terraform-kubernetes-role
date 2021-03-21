# Name of the Role object
output "role_name" {
  value = var.name
}

# Name of the RoleBinding object
output "role_binding_name" {
  value = "${var.name}-binding"
}

# Part of a hack for module-to-module dependencies.
# https://github.com/hashicorp/terraform/issues/1178#issuecomment-449158607
# https://github.com/hashicorp/terraform/issues/1178#issuecomment-473091030
output "depended_on" {
  value = "${null_resource.dependency_setter.id}-${timestamp()}"
}
