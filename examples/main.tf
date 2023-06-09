provider "roztta" {
  // Configurações do provedor, se houver
}

resource "roztta_hello_world" "example" {
  message = "Hello, Terraform!"
}

resource "roztta_gitlab_var" "example" {
  name  = "MY_VARIABLE"
  value = "my_value"
}
