
# 🛒 Gestão de Vendas

Aplicação web com backend em Go e frontend em ReactJS + TypeScript, utilizando PostgreSQL como banco de dados.

---

# 🌐 Deploy na Render

Esta aplicação está disponível publicamente via Render!  
Acesse através do link: [https://gereric-market-application-1.onrender.com](https://gereric-market-application-1.onrender.com)

---

## 🚀 Como rodar o projeto localmente

### 1. Clone o repositório

```bash
git clone https://github.com/seu_usuario/gereric-market-application.git
cd gereric-market-application
````

---

### 2. Configure o PostgreSQL

Certifique-se de ter o **PostgreSQL** instalado e em execução na sua máquina.

Crie um banco de dados com as seguintes credenciais (já configuradas no `.env` do backend):

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=12345
DB_NAME=gestao_vendas
```

No terminal ou em um cliente SQL (como pgAdmin ou DBeaver), execute:

```sql
CREATE DATABASE gestao_vendas;
```

> 🔐 As credenciais e o nome do banco já estão definidos no arquivo `backend/.env`. Se necessário, ajuste de acordo com o seu ambiente.

---

### 3. Rodando o Backend

> Requisitos: [Go instalado](https://go.dev/dl/)

```bash
cd backend
go mod tidy       # Instala as dependências
go run main.go    # Inicia o servidor backend (geralmente em http://localhost:8080)
```

---

### 4. Rodando o Frontend

> Requisitos: [Node.js](https://nodejs.org/) e [npm](https://www.npmjs.com/) instalados

```bash
cd frontend
npm install       # Instala as dependências
npm run dev       # Inicia o frontend (geralmente em http://localhost:5173)
```

---

### ✅ Acesso à Aplicação

* Frontend: [http://localhost:5173](http://localhost:5173)
* API Backend: [http://localhost:8081](http://localhost:8081)

---

### 📚 Documentação da API

A documentação completa das rotas e estrutura da API pode ser encontrada no [README do backend](./backend/README.md).

```

Se quiser que eu atualize outro trecho ou crie uma versão traduzida para inglês também, posso ajudar!
```
