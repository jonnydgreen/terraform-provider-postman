resource "postman_environment" "example" {
  workspace = "5ce6c911-0563-4c89-b37b-7555836a6238"
  name      = "Example"
}

resource "postman_environment_value" "example" {
  environment = postman_environment.example.id
  key         = "key"
  value       = "value"
}
