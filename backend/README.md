# 🧠 Backend - Gestão de Vendas

Este backend foi desenvolvido em **Go** com o objetivo de aprendizado pessoal. Ele fornece uma API REST para gerenciamento de um mercado genérico, utilizando **PostgreSQL** como banco de dados.

---

## 🔌 Endpoints da API

Base URL: `http://localhost:8081`

---

### 👥 Rotas de Clientes (`/clients`)

| Método | Rota         | Descrição                    | Autenticação | Role requerida |
|--------|--------------|------------------------------|--------------|----------------|
| POST   | `/clients/`  | Criar novo cliente           | Não          | Nenhuma        |
| GET    | `/clients/`  | Listar todos os clientes     | Sim          | admin          |
| GET    | `/clients/:id` | Buscar cliente por ID       | Sim          | admin          |
| PUT    | `/clients/:id` | Atualizar cliente           | Sim          | admin          |
| DELETE | `/clients/:id` | Deletar cliente             | Sim          | admin          |

---

### 🔐 Rotas de Administradores (`/adms`)

> **Acesso restrito:** Apenas usuários com role `admin` e autenticação obrigatória.

| Método | Rota        | Descrição                       |
| ------ | ----------- | ------------------------------- |
| POST   | `/adms/`    | Criar um novo administrador     |
| GET    | `/adms/`    | Listar todos os administradores |
| GET    | `/adms/:id` | Buscar administrador por ID     |
| PUT    | `/adms/:id` | Atualizar administrador         |
| DELETE | `/adms/:id` | Deletar administrador           |

---

### 📦 Rotas de Produtos (`/products`)

| Método | Rota            | Descrição                | Autenticação | Role requerida |
| ------ | --------------- | ------------------------ | ------------ | -------------- |
| POST   | `/products/`    | Criar novo produto       | Sim          | admin          |
| GET    | `/products/`    | Listar todos os produtos | Não          | Nenhuma        |
| GET    | `/products/:id` | Buscar produto por ID    | Não          | Nenhuma        |
| PUT    | `/products/:id` | Atualizar produto        | Sim          | admin          |
| DELETE | `/products/:id` | Deletar produto          | Sim          | admin          |

---

### 📦 Rotas de Pedidos (`/orders`)

| Método | Rota          | Descrição               | Autenticação | Role requerida |
| ------ | ------------- | ----------------------- | ------------ | -------------- |
| POST   | `/orders/`    | Criar novo pedido       | Sim          | Qualquer       |
| GET    | `/orders/`    | Listar todos os pedidos | Sim          | admin          |
| GET    | `/orders/:id` | Buscar pedido por ID    | Sim          | admin          |
| PUT    | `/orders/:id` | Atualizar pedido        | Sim          | admin          |

---

### 🛠 Rotas de Serviços Auxiliares (`/services`)

| Método | Rota                          | Descrição                              | Autenticação | Role requerida |
| ------ | ----------------------------- | -------------------------------------- | ------------ | -------------- |
| GET    | `/services/generate-excel`    | Gera relatório Excel dos dados         | Sim          | admin          |
| POST   | `/services/find-vogal`        | Serviço para encontrar vogais em texto | Não          | Nenhuma        |
| GET    | `/services/rank-clients`      | Ranking dos clientes mais ativos       | Sim          | admin          |
| GET    | `/services/ordes-in-progress` | Lista pedidos em andamento             | Sim          | admin          |
| GET    | `/services/summary`           | Resumo geral dos dados                 | Sim          | admin          |

---

### 🔑 Autenticação

| Método | Rota     | Descrição                           |
| ------ | -------- | ----------------------------------- |
| POST   | `/login` | Autenticação e geração de token JWT |

**Exemplo de corpo para login:**

```json
{
  "email": "usuario@exemplo.com",
  "password": "senha123"
}
```

**Resposta:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

## ⚙️ Observações

* A senha é armazenada criptografada usando `bcrypt`.
* Todas as datas seguem o padrão ISO 8601.
* O campo `countOrders` representa a quantidade de pedidos associados ao cliente.
* As rotas protegidas exigem um token JWT válido no header `Authorization: Bearer <token>`.
* A autorização é feita com base na role do usuário (`admin` para rotas restritas).

---

## 📚 Tecnologias utilizadas

* [Go](https://golang.org/)
* [PostgreSQL](https://www.postgresql.org/)
* [GORM](https://gorm.io/)
* [Gin Framework](https://gin-gonic.com/)

---

## 📁 Estrutura de pastas (resumo)

```
backend/
├── controllers/
├── models/
├── routes/
├── config/
├── main.go
└── .env
```

---

## 🚧 Projeto em desenvolvimento

Este projeto é de caráter **experimental e educacional**, sendo aprimorado continuamente.

---
