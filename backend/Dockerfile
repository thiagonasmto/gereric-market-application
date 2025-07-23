# Etapa de build
FROM golang:1.24-alpine AS builder

WORKDIR /gestao-vendas/backend

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Etapa final da imagem
FROM alpine:3.22

WORKDIR /gestao-vendas/backend

RUN apk add --no-cache ca-certificates netcat-openbsd

COPY --from=builder /gestao-vendas/backend/main .
COPY .env .

EXPOSE 8080

CMD ["./main"]
