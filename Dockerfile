# Stage 1: Build
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .
RUN go build -o todo-app main.go

# Stage 2: Run
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/todo-app .
COPY index.html .

EXPOSE 8080
CMD ["./todo-app"]