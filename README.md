Данное API по сокращению ссылок предоставляет следующие возможности:

+ Создавать сокращенную ссылку
+ Изменять сокращенную ссылку и/или оригинальный ресурс
+ Переходить по сокращенной ссылке к оригинальному ресурсу
+ Удалять сокращенную сссылку
+ Получать все ссылки
+ Получать статистику по дням или месяцам

# Запросы и ответы

- Создание ссылки `POST /links`

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

- Изменение ссылки `POST /links/{link_id}`

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

- Переход по ссылке `GET /links/{hash}`

Если переданный хэш есть в базе данных, то произойдет redirect на оригинальный ресурс

- Удаление ссылки `GET /links/{link_id}`

Если link_id есть в базе данных, то ссылка будет удалена (выполняется soft delete)

- Получание всех ссылок авторизованного пользователя `GET /links?limit=5&offset=0`

Response example:
```json
{
    "count": 2,
    "links": [
        {
            "ID": 6,
            "CreatedAt": "2025-08-24T21:44:35.135527+03:00",
            "UpdatedAt": "2025-08-25T20:19:54.713286+03:00",
            "DeletedAt": null,
            "url": "https://testlink.com",
            "hash": "wegsA",
            "userId": 3,
            "Stats": null
        },
        {
            "ID": 7,
            "CreatedAt": "2025-08-26T12:24:10.71528+03:00",
            "UpdatedAt": "2025-08-26T12:24:10.71528+03:00",
            "DeletedAt": null,
            "url": "https://unknown.com",
            "hash": "RVNSb",
            "userId": 3,
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

- Получение статистики `GET /stats?from=date&to=date&by=day`

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
