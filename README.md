# Avito Segments Service

Go-сервис для управления сегментами пользователей:
- создание и удаление сегментов;
- добавление/удаление сегментов у пользователя;
- получение сегментов пользователя;
- выгрузка истории операций в CSV за указанный период.

## Технологии

- Go
- PostgreSQL
- Chi Router
- Docker / Docker Compose

## Быстрый старт (Docker)

```bash
docker compose up --build
```

Сервис стартует на `http://localhost:8080`, PostgreSQL на `localhost:5432`.

## Локальный запуск (без Docker)

1. Поднять PostgreSQL и выполнить `dumps/dump.sql`.
2. Задать переменные окружения:

```bash
export RUN_ADDRESS=":8080"
export DATABASE_URI="host=localhost user=avito password=password dbname=avito sslmode=disable"
```

3. Запустить приложение:

```bash
make run
```

## Основные API endpoint-ы

- `POST /api/segment/create` - создать сегмент
- `POST /api/segment/delete` - удалить сегмент
- `POST /api/user/segment` - добавить/удалить сегменты пользователя
- `GET /api/segment/{id}` - получить сегменты пользователя
- `POST /api/segment/csv/{id}` - сформировать CSV-отчет

Подробный demo-сценарий запросов: `req.http`.

## Примеры ответов об ошибках

```json
{
  "error": "invalid JSON body"
}
```

```json
{
  "error": "segment already exists"
}
```

## Тесты и команды

```bash
make test
make test-unit
make build
```

## Что важно для ревью кода

- Конфигурация читается из флагов и env (`internal/config`).
- HTTP-слой и маршрутизация разделены (`internal/app`, `internal/routers`).
- Логирование запросов добавлено через middleware (`internal/logger`).

## Лицензия

MIT, подробности в [LICENSE](LICENSE).

