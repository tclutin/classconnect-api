# сlassconnect-api

## Установка

1. Склонируйте репозиторий
```bash
git clone github.com/tclutin/classconnect-api
cd classconnect-api
```
2. После этого, переименуйте файл .env_example в .env и настройте его, если это необходимо.
3. Запустите приложение
```bash
docker-compose up
```
4. Затем подключитесь к контейнеру и выполните миграцию
```bash
docker exec <name of container with app> goose -dir ./migrations postgres "postgresql://postgres:postgres@db:5432/classconnect-api" up
```