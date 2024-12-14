# GO Simple API

Este é um projeto simples em Go que cria uma API simples. A sua publicação é dada através da criação de um pod Kubernetes utilizando Helm

## Funcionalidades

- **Rota Raiz (`/`)**: Exibe uma página HTML com links para as rotas disponíveis.
- **Rota `/people`**: Retorna uma lista de pessoas em formato JSON.

# Instalação

### Clone este repositório:

    git clone https://github.com/gdrocha/simple-api-go.git
    cd simple-api-go

### Instale as dependências (se necessário):

    go mod tidy

### Executando o projeto

#### Execute a aplicação:

    go run main.go

    Acesse a API em http://localhost:8081/