[![Build status](https://github.com/patheard/terraform-kubernetes-role/actions/workflows/terraform.yml/badge.svg)](https://github.com/patheard/terraform-kubernetes-role/actions/workflows/terraform.yml)

# Terraform Kubernetes Role
Creates Kubernetes Role and RoleBinding objects in a namespace.  Supports `User` and `Group` subjects for the RoleBinding.

# Dependencies
None.

# Resources

| Name |
|------|
| [kubernetes_role](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/role) |
| [kubernetes_role_binding](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/role_binding) |
| [null_resource](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) |

# Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| create | Should the Role and RoleBinding objects be created? | `bool` | `true` | no |
| dependencies | Allows module to module dependencies | `list` | `[]` | no |
| name | Name of the Kubernetes Role to create | `string` | n/a | yes |
| namespace | Kubernetes namespace to create the Role in.  The namespace must exist. | `string` | n/a | yes |
| rules | The role's rule, which should include lists of `api_groups`, `verbs` and `resources` | <pre>list(object({<br>    api_groups = list(string)<br>    verbs      = list(string)<br>    resources  = list(string)<br>  }))</pre> | n/a | yes |
| subjects | The role binding's subject, which should include lists of `kind` and `name` | <pre>list(object({<br>    kind = string<br>    name = string<br>  }))</pre> | n/a | yes |

# Outputs

| Name | Description |
|------|-------------|
| depended\_on | Part of a hack for module-to-module dependencies. https://github.com/hashicorp/terraform/issues/1178#issuecomment-449158607 https://github.com/hashicorp/terraform/issues/1178#issuecomment-473091030 |
| role\_binding\_name | Name of the RoleBinding object |
| role\_name | Name of the Role object |

# Local testing
You can use k3d in the devcontainer to run a small test cluster with the examples:
```sh
make cluster_create_dev
make test
```
