## Запуск

``docker compose build`` -> 
``docker compose up``
Или ``docker compose up --build``



## Маршруты

1. http://localhost:8080/auth/get-token

body:
```json
{
   "gu_id": string
}
```
response:
```json
{
    "access_token": string,
    "refresh_token": string
}
```

2. http://localhost:8080/auth/refresh

body:
```json
{
   "refresh_token": string
}
```
response:
```json
{
    "access_token": string,
    "refresh_token": string
}
```