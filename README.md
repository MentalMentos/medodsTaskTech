Medods web service with go/gin/postgresql/jwt

Запуск:
docker-compose up -d

go run cmd/main.go

Создание нового пользователя:
curl -X POST "http://localhost:8080/auth_v1/login" -H "Content-Type: application/json" -d '{}'

Вход в аккаунт по user_id(пример):
curl -X POST http://localhost:8080/auth_v1/login \
-H "Content-Type: application/json" \
-d '{"user_id": "3fa85f64-5717-4562-b3fc-2c963f66afa6"}'


curl -X POST "http://localhost:8080/auth_v1/refresh" -H "Content-Type: application/json" -d '{}'

Пример:
curl -X POST http://localhost:8080/auth_v1/refresh \
-H "Content-Type: application/json" \
-d '{
    "access_token": "ACCESS_TOKEN_HERE",
    "refresh_token": "REFRESH_TOKEN_HERE"
}'
