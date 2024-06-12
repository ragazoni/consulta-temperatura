# Etapa de construção
FROM golang:1.22 AS builder

# Instala ca-certificates
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates


WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

# Etapa final
FROM scratch
WORKDIR /app
COPY --from=builder /app/main .

# Expor a porta 8080
EXPOSE 8080

# Comando de entrada
ENTRYPOINT ["./main"]
