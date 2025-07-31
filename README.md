Данное API по скоращению ссылок предоставляет следующие возможности:

+ Создавать сокращенную ссылку
+ Изменять сокращенную ссылку и/или оригинальный ресурс
+ Переходить по сокращенной ссылке к оригинальному ресурсу
+ Удалять сокращенную сссылку
+ Получать все ссылки
+ Получать статистику по дням или месяцам

# Запросы и ответы

- Создание ссылки `POST /link`

Request example:
```json
{
    "url": "https://testlink.com"
}
```

Response example:
```json
{
    "ID": 5,
    "CreatedAt": "2025-07-31T22:44:16.827182+03:00",
    "UpdatedAt": "2025-07-31T22:44:16.827182+03:00",
    "DeletedAt": null,
    "url": "https://testlink.com",
    "hash": "grjrS",
    "Stats": null
}
```

- Изменение ссылки `POST /link/{link_id}`

Request example:
```json
{
    "url": "https://testlink.com",
    "hash": "grjrO"
}
```

Response example:
```json
{
    "ID": 5,
    "CreatedAt": "2025-07-31T22:44:16.827182+03:00",
    "UpdatedAt": "2025-07-31T22:49:17.599312+03:00",
    "DeletedAt": null,
    "url": "https://testlink.com",
    "hash": "grjrO",
    "Stats": null
}
```

- Переход по ссылке `GET /link/{hash}`

Если переданный хэш есть в базе данных, то произойдет redirect на оригинальный ресурс

- Удаление ссылки `GET /link/{link_id}`

Если link_id есть в базе данных, то ссылка будет удалена (выполняется soft delete)

- Получание всех ссылок `GET /link?limit=5&offset=0`

Response example:
```json
{
    "count": 4,
    "links": [
        {
            "ID": 2,
            "CreatedAt": "2025-07-08T15:34:37.320935+03:00",
            "UpdatedAt": "2025-07-08T15:34:37.320935+03:00",
            "DeletedAt": null,
            "url": "https://google.com",
            "hash": "IJKYb",
            "Stats": null
        },
        {
            "ID": 3,
            "CreatedAt": "2025-07-08T15:39:09.53441+03:00",
            "UpdatedAt": "2025-07-16T16:08:21.96572+03:00",
            "DeletedAt": null,
            "url": "https://pkg.go.dev",
            "hash": "pLeru",
            "Stats": null
        },
        {
            "ID": 4,
            "CreatedAt": "2025-07-08T18:49:52.301595+03:00",
            "UpdatedAt": "2025-07-08T18:49:52.301595+03:00",
            "DeletedAt": null,
            "url": "https://stepik.org",
            "hash": "njNpF",
            "Stats": null
        },
        {
            "ID": 5,
            "CreatedAt": "2025-07-31T22:44:16.827182+03:00",
            "UpdatedAt": "2025-07-31T22:49:17.599312+03:00",
            "DeletedAt": null,
            "url": "https://testlink.com",
            "hash": "grjrO",
            "Stats": null
        }
    ]
}
```

- Регистрация пользователя `POST /auth/register`

Request example:
```json
{
    "name": "test",
    "email": "test@test.test",
    "password": "testpassword"
}
```

Response example:
```json
{
    "token": "*JWT*",
}
```

- Логин пользователя `POST /auth/login`

Request example:
```json
{
    "email": "test@test.test",
    "password": "testpassword"
}
```

Response example:
```json
{
    "token": "*JWT*",
}
```

- Получение статистики `GET /stat?from=date&to=date&by=day`

date в формате YYYY-MM-DD, by=day или month

Response example (by=day):
```json
[
    {
        "period": "2025-07-18",
        "sum": 2
    },
    {
        "period": "2025-07-19",
        "sum": 1
    },
    {
        "period": "2025-07-20",
        "sum": 3
    }
]
```

Another response example (by=month):
```json
[
    {
        "period": "2025-07",
        "sum": 6
    }
]
```