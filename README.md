# GO Simple API

Este é um projeto simples em Go que cria uma API simples. A sua publicação é dada através da criação de um pod Kubernetes utilizando Helm

## Funcionalidades

- **Rota Raiz (`/`)**: Exibe uma página HTML com links para as rotas disponíveis.
- **Rota `/people`**: Retorna uma lista de pessoas em formato JSON.

# Instalação

## Clone este repositório:

    git clone https://github.com/gdrocha/go-http-server.git
    cd go-http-server

## Executando o projeto localmente
    go mod download
    go run src/main.go

## Executando o projeto no cluster kubernetes com minikube
    
### Certifique-se que você possua um namespace chamado go
    create namespace go
    
### Instale o chart
    cd charts && \ helm upgrade --install -f Values.yaml -n go go-http-server .

Caso necessário, faça um port-forward da aplicação 
    
    kubectl port-forward service/go-http-server -n go 8081:8081