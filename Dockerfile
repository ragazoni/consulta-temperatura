FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]