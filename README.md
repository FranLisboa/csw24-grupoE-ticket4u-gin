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

Para executar o programa é necessário o uso de docker. 

Primeiro, comece executando o banco de dados através de 

```
docker compose up db -d
```

Após isso, execute o seguinte programa para iniciar o app:

```
docker compose up app -d
```

Caso decida parar algum container, basta executar:
```
docker compose stop
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
chmod +x deploy-aws
./deploy-aws
```

Após finalizar o deploy, você pode consultar a URL do seu Api Gateway pelo portal web da AWS ou pela CLI
Se deseja fazer pela CLI, utilize o comando:

```
aws configure
```
Colocando suas credenciais da AWS

Se estiver utilizando contas que possuem session tokens, configure essa variável também com o seguinte comando:

```
aws configure set aws_session_token <seu_token>
```

Para pegar a URL, acesse as variáveis de Stack da sua Lambda
```
aws cloudformation describe-stacks --stack-name serverless-5-dev --query "Stacks[0].Outputs"  
```

E procure pelo conjunto chave-valor com chave 'HttpApiUrl', que será sua URL para acessar a aplicação