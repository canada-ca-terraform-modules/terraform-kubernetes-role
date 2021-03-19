variable "create" {
  description = "Should the Role and RoleBinding objects be created"
  type        = bool
  default     = true
}

variable "name" {
  description = "Name of the Kubernetes Role to create"
  type        = string
}

variable "namespace" {
  description = "Kubernetes namespace to create the Role in.  The namespace must exist."
  type        = string
}

variable "rules" {
  description = "The role's rule, which should include lists of `api_groups`, `verbs` and `resources`"
  type = object({
    api_groups = list(string)
    verbs      = list(string)
    resources  = list(string)

  })
  default = {
    api_groups = []
    verbs      = []
    resources  = []
  }
}

variable "subject" {
  description = "The role binding's subject, which should include lists of `kind` and `name`"
  type = object({
    kind = list(string)
    name = list(string)
  })
  default = {
    kind = []
    name = []
  }
}
