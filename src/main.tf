provider "aws" {
  region = "us-east-1"
}

resource "aws_key" "key" {
    key_name = "ec2-key"
    public_key = file("~/.ssh/id_rsa.pub")
}

resource "aws_instance" "t1-v1-ticket-api" {
  ami           = "ami-007855ac798b5175e" # Ubuntu 22.04 LTS
  instance_type = "t2.micro"
  key_name =  aws_key.key.ke_name
y
  tags = {
    Name = "t1-v1-ticket-api"
  }

  vpc_security_group_ids = [aws_security_group.docker_sg.id]
  
  user_data = <<-EOF
              #!/bin/bash
              # primeiro atualizar pacotes e instalar o docker
                sudo apt-get update -y
                amazon-linux-extras install docker -y
                service docker start
                usermod -a -G docker ec2-user
                chkconfig docker on
                # baixar e rodar o container
                docker run -d -p 80:80 --name ticket-api nginx
                EOF

}

resource "aws_security_group" "docker_sg" {
  name        = "allow_ssh_http"
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