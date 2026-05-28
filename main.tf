terraform {
  required_version = ">= 1.0.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}


module "networking" {
  source = "git::https://github.com/lesepulvedav/tarea_terra_networking.git?ref=v0.1.1"
  
}

module "compute" {
  source = "git::https://github.com/lesepulvedav/lesepulvedav-tarea_terra_compute.git?ref=v0.2.0"
  vpc_id           = module.networking.vpc_id
  public_subnet_id = module.networking.public_subnet_id
  private_subnet_id = module.networking.private_subnet_id
}

module "storage" {
  source = "git::https://github.com/lesepulvedav/tarea_terra_storage.git?ref=v0.2.0"
  
}

output "ip_servidor_web" {
  value = module.compute.instance_public_ip
}

output "nombre_bucket_s3" {
  value = module.storage.bucket_name
}