# Passo a passo para iniciar o projeto e realizar os testes.

### 1. Criar o container do docker

```
docker compose up -d
```

### 2. Rodar as migrations

```
make migrate
```

### 3. Rodar a aplicação

```
make run
```

### 4. Realizar os testes da aplicação.
Localize o arquivo create_and_list_order.http dentro da pasta `api` localizada na raiz da aplicação, ao localizar esse arquivo será possível fazer os teste criando e listando ao clicar em send conforme o print abaixo.

![image](https://github.com/user-attachments/assets/9725f981-6b1e-4093-a785-c5d2f1497c72)


## Informações de URL's

REST API: http://localhost:8000/orders
GraphQL: http://localhost:8080
gRPC: http://localhost:50051
