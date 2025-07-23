# Etapa 1: Build com Node.js
FROM node:22.15.0 AS builder

WORKDIR /app

COPY package*.json ./
COPY tsconfig*.json ./
COPY .env .
RUN npm install

COPY . .

# Compila TypeScript e builda com Vite
RUN npm run build

# Etapa 2: Nginx para servir a aplicação
FROM nginx:stable-alpine

RUN rm -rf /usr/share/nginx/html/*

# Corrigido: pasta dist é o resultado do Vite
COPY --from=builder /app/dist /usr/share/nginx/html

# Copia o arquivo nginx.conf customizado para evitar 404 nas rotas do React Router
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
