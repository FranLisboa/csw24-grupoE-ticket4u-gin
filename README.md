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

# Como subir o terraform

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

# Aplicar Migrations

Depois de rodar:
```
    serverless deploy
```

E instalar serverless-offline:
```
    npm install serverless-offline --save-dev
```

Rodar:
```
    serverless info --verbose
```
Nos outputs terá a url da database gerada

Depois, intalar o migrations para go, tendo go instalado:
```
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest 
```

Rodar migration na url gerada da database (substitua X pela url):
```
    migrate -database X -path ./src/infrastructure/database/migrations/ up
```
 
