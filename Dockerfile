# Etapa 1: Build
FROM golang:1.22-alpine AS builder

# Configurar diretório de trabalho
WORKDIR /app

# Copiar os arquivos de dependência
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar todo o código-fonte para o container
COPY . .

# Compilar o binário (desabilitar CGO para criar um binário estático)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Etapa 2: Imagem enxuta para rodar o app
FROM alpine:latest

# Instalar certificados necessários para requisições HTTPS e tzdata para fuso horário
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copiar o binário compilado da etapa 1
COPY --from=builder /app/main .

# O Dokploy irá expor essa porta e gerenciar variáveis de ambiente .env pelo seu painel administrativo
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./main"]
