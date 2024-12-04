module "ecs_cluster" {
  source  = "terraform-aws-modules/ecs/aws//modules/cluster"

  cluster_name = "csw24-grupo-e-ticket4u-gin-ecs-2"

  fargate_capacity_providers = {
    FARGATE = {
      default_capacity_provider_strategy = {
        weight = 50
        base   = 20
      }
    }
    FARGATE_SPOT = {
      default_capacity_provider_strategy = {
        weight = 50
      }
    }
  } 

  create_cloudwatch_log_group = false 
  create_task_exec_iam_role = false 
}

resource "aws_ecs_service" "ecs_service" {

  name = "csw24-grupo-e-ticket4u-gin-ecs"
  cluster = module.ecs_cluster.arn
  launch_type = "FARGATE"
  task_definition = aws_ecs_task_definition.task_definition.arn
  enable_execute_command = true
    
  network_configuration {
    subnets          = module.vpc.public_subnets
    security_groups  = [aws_security_group.ecs_service_sg.id]
    assign_public_ip = true
  }
  
  desired_count = 1 

}

resource "aws_security_group" "ecs_service_sg" {
  name        = "ecs-service-sg"
  description = "Security group for ECS service"
  vpc_id      = module.vpc.vpc_id

}

resource "aws_security_group_rule" "egress_all" {
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.ecs_service_sg.id
}

resource "aws_security_group_rule" "ingress_http" {
  type              = "ingress"
  from_port         = 8080
  to_port           = 8080
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.ecs_service_sg.id
}

resource "aws_ecs_task_definition" "task_definition" {
  family                   = "ecs-task"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  task_role_arn = var.AWS_ROLE 
  execution_role_arn = var.AWS_ROLE 

  container_definitions = jsonencode([
    {
      name = "ecs"
      image = "franlisboa/csw24-grupo-e-ticket4u-gin:latest"
      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
        }
      ]
      environment = [ 
        {
          name = "DATABASE_URL"
          value = var.DATABASE_URL
        },
        {
          name = "RUN_MODE"
          value = "local"
        }
      ]
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-group         = aws_cloudwatch_log_group.ecs_log_group.name
          awslogs-region        = "us-east-1"  # Adjust the region as needed
          awslogs-stream-prefix = "ecs"
        }
      }
    }
  ])
}

resource "aws_cloudwatch_log_group" "ecs_log_group" {
  name              = "/ecs/csw24-grupo-e-ticket4u-gin"
  retention_in_days = 7  # Adjust the retention period as needed
}

variable "DATABASE_URL" {
  type = string
}

variable "AWS_ROLE" {
  type = string
}