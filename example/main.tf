terraform {
    required_providers {
        hashicups = {
            versions = ["0.3"]
            source = "hashicorp.com/edu/hashicups"
        }
    }
}

provider "pass" {
    host = "dos"
    token = "test123"
}

module "pass" {
    source = "./pass"

    coffee_name = "Packer Spiced Latte"
}

output "pass" {
    value = module.pass.catalog
}