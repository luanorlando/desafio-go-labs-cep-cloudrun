# --------------------------------------------------------
# Estágio 1: Base de Desenvolvimento e Testes
# --------------------------------------------------------
FROM golang:1.25.2-alpine AS builder

WORKDIR /app

# Instala ferramentas básicas necessárias
RUN apk add --no-cache git gcc musl-dev

# Copia os arquivos de dependências primeiro (otimiza o cache do Docker)
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código fonte do projeto
COPY . .

# Comando padrão deste estágio será rodar os testes
CMD ["go", "test", "./...", "-v"]


# --------------------------------------------------------
# Estágio 2: Compilação do Executável (Novo Estágio)
# --------------------------------------------------------
FROM builder AS compiler

# Compila o binário de forma estática a partir da pasta cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/main.go


# --------------------------------------------------------
# Estágio 3: Imagem final de Execução (Produção/Local)
# --------------------------------------------------------
FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copia o binário real gerado no estágio anterior para a raiz do WORKDIR
COPY --from=compiler /app/main .

EXPOSE 8080

# Inicia o servidor que manterá o contêiner de pé
CMD ["./main"]
