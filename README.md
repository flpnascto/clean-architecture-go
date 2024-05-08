# Go com Clean Architecture

## :notebook_with_decorative_cover: Sobre o Projeto

Aplicação desenvolvida no curso **Pós Graduação Go Expert - Full Cycle** na linguagem Go aplicando conceitos e técnicas de Clean Architecture.

### :sparkles: Funcionalidades
- Criação de ordens de compra com id, preço, taxa e preço final.
- Lista de ordens criadas
- Manipulação de eventos

### :computer: Tecnologias Aplicadas
* Go
* Webserver
* GraphQL
* gRPC
* RabbitMQ
* Banco de Dados MySQL
* Docker

## :arrow_forward: Executando a aplicação

Para executar a aplicação localmente siga as instruções abaixo.

### Pré-requisitos

Primeiramente é necessário que possua instalado as seguintes ferramentas: Go, Git e Docker.
Além disto é bom ter um editor para trabalhar com o código como VSCode.

### Instalação

1. Faça uma cópia do repositório (HTTPS ou SSH)
   ```sh
   git clone https://github.com/flpnascto/clean-architecture-go
   ```
   ```sh
   git clone git@github.com:flpnascto/clean-architecture-go.git
   ```
2. Acesse a pasta do repositório local e instale os pacotes necessários
   ```sh
   go mod tidy
   ```
3. Inicializar o serviço do docker e executar o docker-compose
   ```sh
   docker compose up -d
   ```
4. Executando as aplicações

    1. Acesse o diretório `./cmd/ordersystem`
    2. Execute o comando `go run main.go wire_gen.go`
    3. O web server executa na port 8000
    4. O gRPC executa na porta 50051
    5. O GraphQL executa na porta 8080
    6. O RabbitMQ pode ser acessado na porta 15672

