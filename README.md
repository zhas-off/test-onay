# test-onay

### Для локального запуска:
 
Клонировать сам проект

```bash
  git clone https://github.com/zhas-off/test-onay.git
```

Перейти в директорию проекта

```bash
  cd test-onay
```

Запустить приложение можно двумя способами

1 способ просто поднимает базу данных
```bash
  make run-db
  go run cmd/main.go
```

2 способ поднимает базу данных и само приложение
```bash
  make run
```

### Для отправки запросов прописаны

*GetUsers* для получения пользователей с маршрутом
```bash
http://localhost:8080/api/users
```

*PostUser* для создания пользователя с маршрутом
```bash
http://localhost:8080/api/user
// Пример отправки
{
    "fullName": "TestName",
    "age": 25,
    "address": "TestAdd"
}
```

*UpdateUser* для обновления данных пользователя с маршрутом
```bash
http://localhost:8080/api/user/:id
// Пример отправки
{
    "fullName": "TestNewName",
    "age": 20,
    "address": "TestNewAdd"
}
```

Также используется Middleware для проверки возраста в районе 18<age<100
