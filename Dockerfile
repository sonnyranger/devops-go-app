# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /app

# Копируем go.mod и go.sum и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Сборка бинарника
RUN go build -o todo-app main.go

# Stage 2: Run
FROM alpine:latest
WORKDIR /app

# Копируем бинарник из builder
COPY --from=builder /app/todo-app .
COPY index.html .

# Открываем порт
EXPOSE 8080

# Команда запуска
CMD ["./todo-app"]