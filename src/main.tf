provider "aws" {
  region = "us-east-1"
}

resource "aws_db_instance" "my_postgres_db" {
  identifier              = "my-postgres-db"
  instance_class          = "db.t3.micro"
  engine                  = "postgres"
  engine_version          = "16.3"
  username                = "postgres"
  password                = "postgres"
  db_name                 = "postgres"
  allocated_storage       = 20
  publicly_accessible     = true
  multi_az                = false
  backup_retention_period = 7

  # Apply the security group
  vpc_security_group_ids = [aws_security_group.my_rds_security_group.id]
}

resource "aws_security_group" "my_rds_security_group" {
  name        = "my-rds-security-group"
  description = "Allow access to Postgres RDS instance"

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # This allows access from anywhere; restrict in production
  }
}

output "database_url" {
  value = "postgresql://postgres:${aws_db_instance.my_postgres_db.endpoint}:5432/postgres"
}