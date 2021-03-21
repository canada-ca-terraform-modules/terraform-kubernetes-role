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
