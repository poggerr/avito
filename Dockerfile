# Используем официальный образ Golang
FROM golang:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /go/src/app

# Копируем файлы go.mod и go.sum внутрь контейнера
COPY go.mod go.sum ./

# Запускаем команду go mod download для загрузки зависимостей
RUN go mod download

# Копируем все файлы проекта внутрь контейнера
COPY . .

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

# Собираем приложение
RUN go build -o main ./cmd/avito/

# Запускаем приложение при старте контейнера
CMD ["./main"]
