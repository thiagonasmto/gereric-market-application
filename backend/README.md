# 🧠 Backend - Gestão de Vendas

Este backend foi desenvolvido em **Go** com o objetivo de aprendizado pessoal. Ele fornece uma API REST para gerenciamento de um mercado genérico, utilizando **PostgreSQL** como banco de dados.

---

## 🔌 Endpoints da API

Base URL: `http://localhost:8081`

### 📄 Listar todos os clientes

**GET** `/clients/`

**Resposta:**

```json
[
  {
    "id": "141a44d1-f04d-409f-936a-fa39abb2b978",
    "name": "name",
    "email": "email@gmail.com",
    "password": "$2a$10$NIyo0rd6bZqFnHxLupNzb...",
    "createdAt": "2025-07-21T21:11:32.008906-03:00",
    "updatedAt": "2025-07-21T21:11:32.008906-03:00",
    "countOrders": 0
  }
]
````

---

### 🔍 Buscar cliente por ID

**GET** `/clients/:id`

**Resposta:**

```json
{
  "id": "141a44d1-f04d-409f-936a-fa39abb2b978",
  "name": "name",
  "email": "email@gmail.com",
  "password": "$2a$10$NIyo0rd6bZqFnHxLupNzb...",
  "createdAt": "2025-07-21T21:11:32.008906-03:00",
  "updatedAt": "2025-07-21T21:11:32.008906-03:00",
  "countOrders": 0
}
```

---

### ➕ Criar novo cliente

**POST** `/clients/`

**Corpo da requisição:**

```json
{
  "name": "name",
  "email": "email@gmail.com",
  "password": "senha123"
}
```

**Resposta:**

```json
{
  "id": "141a44d1-f04d-409f-936a-fa39abb2b978",
  "name": "name",
  "email": "email@gmail.com",
  "password": "$2a$10$NIyo0rd6bZqFnHxLupNzb...",
  "createdAt": "2025-07-21T21:11:32.008906-03:00",
  "updatedAt": "2025-07-21T21:11:32.008906-03:00",
  "countOrders": 0
}
```

---

### ✏️ Atualizar cliente existente

**PUT** `/clients/:id`

**Corpo da requisição:**

```json
{
  "name": "user3",
  "email": "user4@gmail.com",
  "password": "novaSenha"
}
```

**Resposta:**

```json
{
  "id": "52d8d8ed-f3f1-4a1e-9952-d8857beb15c1",
  "name": "user3",
  "email": "user4@gmail.com",
  "password": "$2a$10$lsa0BCKawE2oPQAMfmok...",
  "createdAt": "2025-07-18T23:18:22.438798-03:00",
  "updatedAt": "2025-07-18T23:32:53.0246658-03:00",
  "countOrders": 0
}
```

---

### ❌ Deletar cliente

**DELETE** `/clients/:id`

**Resposta:**

```json
{
  "message": "Usuário deletado com sucesso"
}
```

---

## ⚙️ Observações

* A senha é armazenada criptografada usando `bcrypt`.
* Todas as datas seguem o padrão ISO 8601.
* O campo `countOrders` representa a quantidade de pedidos associados ao cliente.

---

## 📚 Tecnologias utilizadas

* [Go](https://golang.org/)
* [PostgreSQL](https://www.postgresql.org/)
* [GORM](https://gorm.io/)
* [Echo Framework](https://echo.labstack.com/) (ou o framework que estiver usando)

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
