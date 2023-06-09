provider "roztta" {
  roztta_token    = "your_access_key_here"
  gitlab_token    = "your_gitlab_token_here"    // Opcional
  bitbucket_token = "your_bitbucket_token_here" // Opcional
}

resource "roztta_hello_world" "example" {
  message = "Hello, Terraform!"
}

resource "roztta_gitlab_var" "example" {
  name  = "MY_VARIABLE"
  value = data.roztta_gitlab_var.my_var.value
}

data "roztta_gitlab_var" "my_var" {
  name = "my_variable_name"
}

output "my_var_value" {
  value = data.roztta_gitlab_var.my_var.value
}
