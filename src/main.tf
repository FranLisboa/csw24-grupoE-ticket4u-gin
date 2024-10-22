terraform {
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

resource "tls_private_key" "rsa-4096" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

variable key_name{}

resource "aws_key_pair" "key_pair" {
  key_name   = var.key_name
  public_key = tls_private_key.rsa-4096.public_key_openssh
}

resource "local_file" "private_key" {
  content  = tls_private_key.rsa-4096.private_key_pem
  filename = var.key_name
}

resource "aws_security_group" "sec_group" {
  name        = "new_sec_group"
  description = "Allow SSH and HTTP inbound traffic"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "t1_ticket_api_v1" {
  ami           = "ami-0866a3c8686eaeeba"
  instance_type = "t2.micro"
  key_name      = aws_key_pair.key_pair.key_name

  tags = {
    Name = "t1_ticket_api_v1"
  }

  vpc_security_group_ids = [aws_security_group.sec_group.id]
  
  user_data = <<-EOF
              #!/bin/bash
              # Update package index
              sudo apt-get update -y
              sudo ufw allow 8080

              # Install Docker

              sudo apt-get update
              sudo apt-get install ca-certificates curl
              sudo install -m 0755 -d /etc/apt/keyrings
              sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
              sudo chmod a+r /etc/apt/keyrings/docker.asc

              # Add the repository to Apt sources:
              echo \
                "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
                $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
                sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
              sudo apt-get update

              # Install Docker Engine
              sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y

              # Start Docker service
              sudo systemctl start docker
              sudo systemctl enable docker
              
              # Clone github
              mkdir repo
              cd repo
              git clone https://github.com/GuilhermePoglia/csw24-grupoE-ticket4u-gin.git

              # Enter repo
              cd csw24-grupoE-ticket4u-gin
              cd src

              # Build Docker image
              sudo docker compose up db -d
              sudo docker compose up app -d
              EOF
}
