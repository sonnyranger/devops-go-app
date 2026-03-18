# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /app

# Если появятся зависимости, создадим go.mod/go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Собираем бинарник
RUN go build -o todo-app main.go

# Stage 2: Run
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/todo-app .
COPY index.html .

EXPOSE 8080

CMD ["./todo-app"]