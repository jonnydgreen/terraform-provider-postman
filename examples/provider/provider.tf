terraform {
  required_providers {
    postman = {
      source = "jonnydgreen/postman"
    }
  }
}

provider "postman" {}

resource "postman_workspace" "example" {
  name = "test3"
  type = "personal"
  description = "test3 desc"
}
