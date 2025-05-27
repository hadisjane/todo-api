# TodoApp

**TodoApp** — это простой API и консольное приложение для управления задачами (ToDo), написанное на чистом Go с использованием стандартной библиотеки [`net/http`](https://pkg.go.dev/net/http).

## Возможности

- Добавление задач
- Получение задачи по ID
- Удаление задачи
- Завершение задачи
- Получение списка всех задач
- REST API и консольный интерфейс

## Быстрый старт

### Запуск REST API

```sh
go run main.go
```

Сервер будет доступен по адресу: [http://localhost:8989](http://localhost:8989)

### Примеры REST-запросов

- Получить все задачи:
    ```
    GET /todos
    ```
- Добавить задачу:
    ```
    POST /todos
    Content-Type: application/json

    {
      "title": "Новая задача",
      "done": false
    }
    ```
- Получить задачу по ID:
    ```
    GET /todos/{id}
    ```
- Завершить задачу:
    ```
    PUT /todos/{id}
    ```
- Удалить задачу:
    ```
    DELETE /todos/{id}
    ```

### Пример структуры задачи

```json
{
  "id": 1,
  "title": "Пример задачи",
  "done": false,
  "created_at": "2024-06-01T12:00:00Z"
}
```

## Структура проекта

```
.
├── main.go
├── go.mod
├── controllers/
│   ├── controller.go
│   ├── console/
│   │   ├── routes.go
│   │   └── taskController.go
│   └── rest/
│       ├── errs.go
│       ├── handlers.go
│       └── routes.go
├── errs/
│   └── errs.go
├── models/
│   └── task.go
├── repositories/
│   ├── dummyData.go
│   └── taskRepo.go
├── services/
│   └── taskService.go
└── utils/
    └── utils.go
```

## Как работает

- Все данные хранятся в памяти (см. [`repositories/dummyData.go`](repositories/dummyData.go)).
- Логика работы с задачами реализована в [`services/taskService.go`](services/taskService.go).
- REST API реализован в [`controllers/rest`](controllers/rest/routes.go).
- Консольный интерфейс — в [`controllers/console`](controllers/console/routes.go).

## Требования

- Go 1.18+

## Авторы

- hadisjane

---

> Проект создан для обучения и демонстрации работы с чистым Go и стандартной библиотекой.