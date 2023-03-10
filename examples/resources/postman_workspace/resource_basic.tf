# Manage postman workspace.
resource "postman_workspace" "example" {
  name = "My Workspace"
  type = "personal"
}
