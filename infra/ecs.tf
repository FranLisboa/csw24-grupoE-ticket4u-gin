module "ecs_cluster" {
  source  = "terraform-aws-modules/ecs/aws//modules/cluster"

  cluster_name = "csw24-grupo-e-ticket4u-gin-ecs"

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

module "ecs_service" {
  source = "terraform-aws-modules/ecs/aws//modules/service"

  cluster_arn = module.ecs_cluster.arn

  security_group_rules = {
        egress_all = {
          type        = "egress"
          from_port   = 0
          to_port     = 0
          protocol    = "-1"
          cidr_blocks = ["0.0.0.0/0"]
        }
  }

  create_task_exec_iam_role = false 
  create_tasks_iam_role = false
  
  subnet_ids = module.vpc.public_subnets
  name = "csw24-grupo-e-ticket4u-gin-ecs"
    container_definitions = {
          cpu       = 512
          memory    = 1024
          essential = true
          image     = "franlisboa/csw24-grupo-e-ticket4u-gin:latest"
          port_mappings = [
            {
              name          = "ecs-sample"
              containerPort = 80
              protocol      = "tcp"
            }
          ]
          # Example image used requires access to write to root filesystem
          readonly_root_filesystem = true
        
          memory_reservation = 100
          create_task_exec_iam_role = false 

          public_ip = true

          network_configuration = {
            subnets          = module.vpc.public_subnets
            assign_public_ip = true      
          }

        ingress_http = {
          type        = "ingress"
          from_port   = 80
          to_port     = 80
          protocol    = "tcp"
          cidr_blocks = ["0.0.0.0/0"]
        }

      enable_cloudwatch_logging = false   
      create_cloudwatch_log_group = false 
      create_task_exec_iam_role = false 
      create_tasks_iam_role = false
      task_exec_iam_role_arn = "arn:aws:iam::948675409837:role/LabRole"
      
      network_configuration = {
        subnets          = module.vpc.public_subnets
        assign_public_ip = true      
      }
    }
}
