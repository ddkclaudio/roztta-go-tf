terraform {
  required_providers {
    roztta = {
      source  = "roztta.com/terraform/roztta"
      version = "1.0"
    }
  }
}

resource "roztta_hello_world" "example" {
  message = "Hello, World!"
}

output "message" {
  value = roztta_hello_world.example.message
}
