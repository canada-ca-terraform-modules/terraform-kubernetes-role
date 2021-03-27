[![Build status](https://github.com/canada-ca-terraform-modules/terraform-kubernetes-role/actions/workflows/terraform.yml/badge.svg)](https://github.com/canada-ca-terraform-modules/terraform-kubernetes-role/actions/workflows/terraform.yml)

# Terraform Kubernetes Role
Creates Kubernetes Role and RoleBinding objects in a namespace.  Supports `User` and `Group` subjects for the RoleBinding.

# Requirements

No requirements.

# Providers

| Name | Version |
|------|---------|
| kubernetes | n/a |

# Modules

No Modules.

# Resources

| Name |
|------|
| [kubernetes_role](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/role) |
| [kubernetes_role_binding](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/role_binding) |

# Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| create | Should the Role and RoleBinding objects be created? | `bool` | `true` | no |
| name | Name of the Kubernetes Role to create | `string` | n/a | yes |
| namespace | Kubernetes namespace to create the Role in.  The namespace must exist. | `string` | n/a | yes |
| rules | The role's rule, which should include lists of `api_groups`, `verbs` and `resources` | <pre>list(object({<br>    api_groups = list(string)<br>    verbs      = list(string)<br>    resources  = list(string)<br>  }))</pre> | n/a | yes |
| subjects | The role binding's subject.  To use the `default` namespace for a kind of `User` or `Group`, pass a `null` namespace. | <pre>list(object({<br>    kind      = string<br>    name      = string<br>    namespace = string<br>  }))</pre> | n/a | yes |

# Outputs

| Name | Description |
|------|-------------|
| role\_binding\_name | Name of the RoleBinding object |
| role\_name | Name of the Role object |

# Local testing
You can use k3d in the devcontainer to run a small test cluster with the examples:
```sh
make cluster_create_dev
make test
```
