terraform {
  backend "remote" {
    hostname = "localhost:8443"
    organization = "default"

    token = "token"

    workspaces {
      name = "default"
    }
  }
}

resource "terraform_data" "example" {
  input = "Hello, OpenAtlas!"
}
