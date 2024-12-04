# csw24-grupoE-ticket4u-gin

## Alunos 
- Bernardo Muller
- David Campos
- Francisco Lisboa
- Gabriel Reis
- Guilherme Poglia
  
# Dependências e Tecnologias
[Docker](https://www.docker.com/)

[Go](https://go.dev/)

[Terraform](https://www.terraform.io/)

[AWS](https://aws.amazon.com/)

# Como executar localmente o programa

Para executar o programa localmente é necessário o uso de docker. 

Comando para executar: 

```
docker compose up
```

Caso decida parar algum container, basta executar:

```
docker compose stop
```

Rode as migrations
```
goose -dir ./infrastructure/database/migrations postgres 'host=db port=5432 user=admin password=admin dbname=postgres sslmode=disable' up
```

# Como subir o terraform (Utilizado para EC2)

Adicionar suas credenciais em X

Iniciar o terraform usando:
```
terraform init
```

Caso queira que o plano de criação seja demonstrado, execute:

```
terraform plan
```

Para de fato subir a infraestrutura:

```
terraform apply
```

# Acessando o Swagger

Para acessar o swagger, baster ir em URL/swagger/index.html 

ex: http://localhost:8080/swagger/index.html

# Fazer deploy na AWS Lambda com banco de dados RDS

# Opção 1
Rodar localmente o workflow do Github Actions com uma ferramenta como [Act](https://github.com/nektos/act)

Configurando os secrets utilizados:

```
AWS_ACCESS_KEY_ID=sua_aws_access_key_id
AWS_SECRET_ACCESS_KEY=sua_aws_secret_access_key
AWS_SESSION_TOKEN=sua_aws_session_token
AWS_ROLE=sua_aws_iam_role
```

# Opção 2

Instalar [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)

Instalar [Node Runtime](https://nodejs.org)

Instalar o framework serverless

```
npm install -g serverless@3
```

Instalar o framework serverless-offline

```
serverless plugin install -n serverless-offline --config serverless-lambda.yml
```
Instalar Goose, uma biblioteca para fazer migrations

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Configurar credenciais da AWS

```
export AWS_ACCESS_KEY_ID=sua_aws_access_key_id
export AWS_SECRET_ACCESS_KEY=sua_aws_secret_access_key
export AWS_SESSION_TOKEN=sua_aws_session_token
export AWS_REGION=us-east-1
export AWS_ROLE=sua_aws_iam_role
```

Buildar o Projeto Go. Se não estiver em ambiente Linux, rode o arquvio em um terminal Git Bash

```
cd src
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../bootstrap
cd ..
```

Gere um .zip do arquivo gerado. Se estiver em ambiente Linux:

```
zip -j function.zip bootstrap
```
Após isso, rode o script para fazer deploy. Se não estiver em ambiente Linux, rode o arquvio em um terminal Git Bash

```
chmod +x deploy-aws.sh
./deploy-aws.sh
```

Após finalizar o deploy, você pode consultar a URL do seu Api Gateway pelo portal web da AWS ou pela CLI
Se deseja fazer pela CLI, utilize o comando:

```
aws configure
```
Colocando suas credenciais da AWS e a região us-east-1

Se estiver utilizando contas que possuem session tokens, configure essa variável também com o seguinte comando:

```
aws configure set aws_session_token <seu_token>
```

Para pegar a URL, acesse as variáveis de Stack da sua Lambda

```
aws cloudformation describe-stacks --stack-name serverless-5-dev --query "Stacks[0].Outputs"  
```

E procure pelo conjunto chave-valor com chave 'HttpApiUrl', que será sua URL para acessar a aplicação

# Fazer deploy usando ECS

Instalar o framework serverless
```
npm install -g serverless@3
```

Instalar Goose, uma biblioteca para fazer migrations
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Configure as chaves da AWS (se ainda não realizado)
```
aws configure
aws configure set aws_session_token <seu_token>
export AWS_ROLE=sua_aws_iam_role
```

Após isso, rode o script para fazer deploy da ECS e do RDS. Se não estiver em ambiente Linux, rode o arquvio em um terminal Git Bash
```
chmod +x deploy-rds-ecs.sh
./deploy-rds-ecs.sh
```

Para saber o IP público da task do ECS olhar no console da AWS, ou pegar por linha de comando
Listar tasks do ECS:
```
aws ecs list-tasks --cluster csw24-grupo-e-ticket4u-gin-ecs-2
```

Pegar o ARN da task para ver seus detalhes
```
aws ecs describe-tasks --cluster csw24-grupo-e-ticket4u-gin-ecs-2 --tasks <task-arn>
```

Dentro de 'attachments' e depois dentro de 'details', pegar o valor do atributo 'networkInterfaceId'
```
aws ec2 describe-network-interfaces --network-interface-ids <networkInterfaceId>
```

Após isso procurar pelo ip público e usar como url base para a aplicação
```
http://<ip_publico>:8080/
```
