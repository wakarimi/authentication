# Сервис аутентификации

# Запуск
```yml
version: "3"

services:
  database:
    image: postgres:latest
    container_name: wakarimi-authentication-db
    ports:
      - "5432:5432"
    volumes:
      - /data/postgresql:/var/lib/postgresql
      - /data/postgresql/data:/var/lib/postgresql/data
    environment:
      - POSTGRES_DB=wakarimi-authentication-db
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=password
    restart: unless-stopped
    networks:
      - wakarimi-network

  wakarimi-authentication: 
    image: zalimannard/wakarimi-authentication
    container_name: wakarimi-authentication
    ports:
      - "8020:8020"
    environment:
      - APP_LOGGING_LEVEL=DEBUG
      - APP_REFRESH_KEY=key1
      - APP_ACCESS_KEY=key2
      # - APP_LOGGING_LEVEL=INFO # TRACE, DEBUG, INFO, WARN, ERROR, FATAL
      # - HTTP_PORT=8020
      - DB_HOST=database
      - DB_PORT=5432
      - DB_NAME=wakarimi-authentication-db
      - DB_USER=user
      - DB_PASSWORD=password
      # - DB_READ_TIMEOUT=1s
      # - DB_WRITE_TIMEOUT=1s
      # - DB_CHARSET=UTF-8
    restart: unless-stopped
    networks:
      - wakarimi-network
    depends_on:
      - database

networks:
  wakarimi-network:
    external: true

```
# Эндпоинты

## Регистрация

```
POST /api/accounts/sign-up
```

Входные данные:
```json
{
    "username": "string",
    "password": "string"
}
```

## Вход

```
POST /api/auth/sign-in
```

Входные данные:
```json
{
    "username": "string",
    "password": "string",
    "fingerprint": "string"
}
```

Выходные данные:
```json
{
    "refreshToken": "string",
    "accessToken": "string"
}
```

## Выход

```
POST /api/auth/sign-out
```

Требуется хедер X-Device-ID со значением устройства, с которого нужно выйти

## Выход со всех устройств

```
POST /api/auth/sign-out-all
```

Требуется хедер X-Account-ID со значением аккаунта, с которого нужно выйти

## Смена пароля

```
POST /api/accounts/change-password
```

Входные данные:
```json
{
    "oldPassword": "string",
    "newPassword": "string"
}
```

Требуется хедер X-Account-ID со значением аккаунта, которому нужно поменять пароль

## Обновление токенов

```
POST /api/tokens/refresh
```

Входные данные:
```json
{
    "refreshToken": "string"
}
```

Выходные данные:
```json
{
    "refreshToken": "string",
    "accessToken": "string"
}
```

## Проверка токена

```
POST /api/tokens/verify
```

Входные данные:
```json
{
    "accessToken": "string"
}
```

Выходные данные:
```json
{
    "valid": false
}
```
или
```json
{
    "valid": true,
    "accountId": "int",
    "deviceId": "int",
    "roles": [
        "string"
    ],
    "issuedAt": "int",
    "expiryAt": "int"
}
```

## Получение деталей своего аккаунта

```
GET /api/accounts/me
```

Требуется хедер X-Account-ID со значением аккаунта, который запрашивает свои данные

Выходные данные:
```json
{
    "id": "int",
    "username": "string",
    "roles": [
        "string"
    ]
}
```

## Получение списка указанных аккаунтов

```
GET /api/accounts?ids=1,2
```

Выходные данные:
```json
{
    "accounts": [
        {
            "id": "int",
            "username": "string"
        }
    ]
}
```

## Назначение роли

```
POST /api/accounts/3/roles
```

Требуется хедер X-Account-ID со значением аккаунта, который запрашивает назначение роли для проверки прав

Входные данные:
```json
{
    "roleName": "roleName"
}
```

## Снятие роли

```
DELETE /api/accounts/3/roles
```

Требуется хедер X-Account-ID со значением аккаунта, который запрашивает снятие роли для проверки прав

Входные данные:
```json
{
    "roleName": "roleName"
}
```
