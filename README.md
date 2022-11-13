# Microservices go gRPC
[![Run in Insomnia}](https://insomnia.rest/images/run.svg)](https://insomnia.rest/run/?label=API%20alunos&uri=https%3A%2F%2Fraw.githubusercontent.com%2Fdanielsdev%2Fapi-go-gin%2Fmain%2Finsomnia%2FInsomnia_2022-10-28.json)

Esse projeto é uma API REST para gerenciar alunos de uma escola com dois serviços


## API
Esse serviço é responsável por receber e responder as requisições do cliente, e se comunicar com o manager usando gRPC

## Manager
Esse serviço é responsável por gerenciar os dados das requisições vindas da API no banco de dados, somente essa camada se comunica diretamente com o banco.

## Fluxo da aplicação

![Alt text](docs/images/fluxo.png?raw=true "Fluxo da aplicação")

## Executar o projeto

É bem simples executar esse projeto, basta você seguir esses passos:

### Primeiro passo
  - Copiar .env.example para .env com as variáveis de ambiente de ambos os serviços, você pode usar o comando: `cp .env.example .env`

### Segundo passo
  - Fazer o build dos containers com o comando: `docker-compose up -d --build`

## Acessar a API
 - http://localhost:{APP_PORT}/

 ## Acessar o manager
 - http://localhost:{APP_MANAGER_PORT}/

## Acessar o banco de dados
 - http://localhost:{DB_PORT}/

## Acessar o pgadmin (gerenciador do BD)
 - http://localhost:{PG_ADMIN_PORT}/

## Tecnologias
 - **[Golang](https://go.dev/)**
 - **[Gin](https://github.com/gin-gonic/gin)**
 - **[GORM](https://gorm.io/)**
 - **[Docker](https://www.docker.com/)**
 - **[gRPC](https://grpc.io/)**
 - **[PostgreSQL](https://www.postgresql.org/)**

## Arquitetura
 - REST
 - gRPC
 - Soft delete

## Containers docker
 - golang_app
 - golang_manager
 - postgres
 - pgadmin-compose

## Endpoints
 - `GET /alunos` Retorna a lista de alunos
 - `GET /alunos/{ID}` Retorna as informações de um aluno
 - `POST /alunos` Cria um aluno, exemlo de json a ser enviado no body:
   ```json
    {
		"nome": "Daniel",
		"cpf": "23034153066",
		"rg": "140535925",
		"matricula": "120931"
    }

   ```
 - `PATCH /alunos/{ID}` Edita os dados de um aluno, exemlo de json a ser enviado no body:
    ```json
    {
		"nome": "Daniel test",
    }
   ```
   Observação: note que por ser usado o método `PATCH`, não é necessário enviar todos os dados do recurso aluno, apenas os campos que deseja editar

 - `DELETE /alunos/{ID}` Deleta um aluno 
 - `GET /alunos/busca/{CPF}` Busca as informações de um aluno pelo CPF
