# TodoApp API

**TodoApp** — это высокопроизводительное REST API для управления задачами, написанное на Go с использованием фреймворка Gin и PostgreSQL в качестве базы данных.

## 🚀 Возможности

- Создание задач с уникальными названиями
- Просмотр списка всех задач
- Просмотр задачи по ID
- Обновление статуса задачи (выполнено/не выполнено)
- Удаление задач
- Валидация входящих данных
- Обработка ошибок
- Подключение к PostgreSQL
- Автоматические миграции базы данных

## 🏗️ Архитектура

Проект следует чистой архитектуре с разделением на слои:

```
.
├── controller/      # Обработчики HTTP-запросов
├── service/         # Бизнес-логика
├── repository/      # Взаимодействие с базой данных
├── models/          # Модели данных
├── errs/            # Кастомные ошибки
├── db/              # Настройки базы данных и миграции
└── main.go          # Точка входа
```

### Требования

- Go 1.18+
- PostgreSQL

### Установка

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/hadisjane/todo-api.git
   cd todo-api
   ```

2. Настройте подключение к базе данных в файле `.env`:
   ```
   DB_DSN="host=localhost port=5432 user=postgres password=your_password dbname=your_db_name sslmode=disable"
   ```

3. Установите зависимости:
   ```bash
   go mod download
   ```

4. Запустите приложение:
   ```bash
   go run main.go
   ```

API будет доступно по адресу: [http://localhost:8989](http://localhost:8989)

## 📚 Документация API

### Получить все задачи
```
GET /todos
```

**Ответ:**
```json
[
  {
    "id": 1,
    "title": "Купить молоко",
    "done": false,
    "created_at": "2023-01-01T12:00:00Z"
  },
  ...
]
```

### Создать задачу
```
POST /todos
Content-Type: application/json

{
  "title": "Новая задача",
  "done": false
}
```

**Ответ:**
```json
{
  "id": 1,
  "title": "Новая задача",
  "done": false,
  "created_at": "2023-01-01T12:00:00Z"
}
```

### Получить задачу по ID
```
GET /todos/:id
```

**Ответ:**
```json
{
  "id": 1,
  "title": "Купить молоко",
  "done": false,
  "created_at": "2023-01-01T12:00:00Z"
}
```

### Обновить статус задачи
```
PUT /todos/:id
```

**Ответ:**
```json
{
  "id": 1,
  "title": "Купить молоко",
  "done": true,
  "created_at": "2023-01-01T12:00:00Z"
}
```

### Удалить задачу
```
DELETE /todos/:id
```

**Ответ:**
```
Status: 200 OK
```

## 🛠 Технологии

- [Gin](https://github.com/gin-gonic/gin) - Веб-фреймворк
- [sqlx](https://github.com/jmoiron/sqlx) - Расширение для работы с SQL
- [github.com/lib/pq](https://github.com/lib/pq) - Драйвер для подключения к PostgreSQL
- [godotenv](https://github.com/joho/godotenv) - Управление переменными окружения

> 🧠 Проект сделан для практики и изучения чистого Go, REST API и архитектурных подходов. А так же для понимания работы с базами данных и веб-серверами.