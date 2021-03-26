# Name of the Role object
output "role_name" {
  value = var.name
}

# Name of the RoleBinding object
output "role_binding_name" {
  value = "${var.name}-binding"
}
