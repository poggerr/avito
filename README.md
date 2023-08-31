# AvitoTest

Это сервис на языке Golang, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)

## Требования

- Golang
- PostgreSQL
- Docker

## Установка и настройка

1. Склонируйте репозиторий:
   ```shell
   git clone https://github.com/avito-tech/backend-trainee-assignment-2023.git

2. Установите зависимости:
   ```shell
   go mod download

## Запуск сервиса

1. Запустите базу данных PostgreSQL и сервис с помощью Docker:
   ```shell
   docker-compose up --build

## Использование

### REST API Запросы

1. Создание сегмента:
   ```json
   POST localhost:8080/api/segment/create
   Content-Type: application/json

   {
      "segment": "AVITO_VOICE_MESSAGES"
   }

2. Добавление пользователя в сегмент:
   ```json
   POST localhost:8080/api/user/segment
   Content-Type: application/json

   {
     "add" : ["AVITO_VOICE_MESSAGES"],
     "delete" : [],
     "user" : "2376e110-e40d-41d0-85ba-22db804c4f51"
   }
   
3. Удаление сегмента
   ```json
   POST localhost:8080/api/segment/delete
   Content-Type: application/json

   {
      "segment": "AVITO_VOICE_MESSAGES"
   }
   
4. Удаление пользователя из сегмента
   ```json
   POST localhost:8080/api/user/segment
   Content-Type: application/json

   {
      "add" : [],
      "delete" : ["AVITO_VOICE_MESSAGES"],
      "user" : "2376e110-e40d-41d0-85ba-22db804c4f51"
   }

5. Вывод сегментов пользователя 
   ```json
   GET localhost:8080/api/segment/2376e110-e40d-41d0-85ba-22db804c4f51
   Content-Type: text/plain
   
6. Создание отчета в формате CSV пользовательских сегментов
   ```json
   POST localhost:8080/api/segment/csv/2376e110-e40d-41d0-85ba-22db804c4f51
   Content-Type: application/json

   {
      "period": "30-2023"
   }

## Тестирование

Вы можете запустить unit-тесты с помощью следующей команды:
   ```shell
   go test
   ```
## Вклад

Если вы нашли ошибку или хотите внести улучшения, пожалуйста, создайте issue или отправьте pull request.

## Лицензия

Этот проект лицензирован под MIT License - подробности см. в файле [LICENSE](LICENSE).

