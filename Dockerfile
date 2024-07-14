# Dockerfile

# Usando a imagem oficial do Golang
FROM golang:1.17-alpine AS builder

# Configurando o diretório de trabalho dentro do container
WORKDIR /app

# Copiando o go.mod e go.sum para baixar as dependências
COPY go.mod go.sum ./

# Baixando as dependências do projeto
RUN go mod download

# Copiando o restante do código-fonte para o diretório de trabalho
COPY . .

# Compilando o aplicativo
RUN go build -o dance-moves-api ./cmd

# Usando uma imagem menor para reduzir o tamanho final da imagem Docker
FROM alpine:latest

# Copiando o executável compilado do estágio anterior
COPY --from=builder /app/dance-moves-api /dance-moves-api

# Expondo a porta 8080 para o exterior
EXPOSE 8080

# Comando para iniciar o aplicativo
CMD ["/dance-moves-api"]
