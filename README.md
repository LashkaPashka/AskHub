# AskHub

**AskHub** — это REST API сервис для управления вопросами и ответами.
Позволяет создавать, получать, удалять вопросы и ответы, а также работать с данными через удобные REST-эндпоинты.

---

## 📦 Версия

- API Version: 1.0
- Host: `localhost:8080`
- BasePath: `/`

---

## 🔧 Установка

1.  Клонируйте репозиторий:

    ```bash
    git clone https://github.com/LashkaPashka/AskHub.git
    cd AskHub
    ```

2.  Настройте переменные окружения для `docker-compose.yaml` в `.env`:

    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=root
    DB_NAME=askhub
    ```

3. Настройте конфиг приложения в `./AskHub/config/prod.yaml`:

    ```yaml
    env: "prod"
    storage_path: "host=postgres user=postgres password=root dbname=askhub port=5432 sslmode=disable"
    http_server:
      address: "0.0.0.0:8080"
      timeout: 4s
      idle_timeout: 30s
    ```

4. Настройте переменные окружения для миграций и запуска сервера `./AskHub/.env`:

    ```env
    migration_dir="./migrations"
    migration_dsn="host=postgres user=postgres password=root dbname=askhub port=5432 sslmode=disable"
    config_path="./config/prod.yaml"
    ```

5.  Соберите и запустите сервис:

    ```bash
    docker-compose up -d
    ```

API будет доступен по адресу: `http://localhost:8080`

## 🔄 Миграции
Проект содержит готовые миграции для создания таблиц questions и answers

## 🧪 Тестирование
Проект включает модульные тесты для:
- слоя Service
- слоя Storage

---

## 📚 Endpoints

### 1️⃣ Создать вопрос

`POST /questions`

Создаёт вопрос в базе данных.

**Пример запроса:**

```json
{
	"text": "What's java?"
} 
```

**Responses:**

-   `200`: `"Question was created successfully!"`
-   `400`: `Invalid request. Please check the submitted data.`
-   `500`: `Internal server error`

### 2️⃣ Получить вопрос

`GET /questions/{id}`

Получить конкретный вопрос.

**Пример ответа:**

```json
{
	"Questions": [
		{
			"ID": 1,
			"CreatedAt": "2025-11-21T19:17:30.346138Z",
			"UpdatedAt": "2025-11-21T19:17:30.346138Z",
			"DeletedAt": null,
			"text": "What's java?"
		}
	]
}
```

### 3️⃣ Получить все вопросы

`GET /questions`

Получает все вопросы.

**Пример ответа:**

```json
{
	"Questions": [
		{
			"ID": 1,
			"CreatedAt": "2025-11-21T19:17:30.346138Z",
			"UpdatedAt": "2025-11-21T19:17:30.346138Z",
			"DeletedAt": null,
			"text": "What's java?"
		},
    {
			"ID": 2,
			"CreatedAt": "2025-11-21T19:17:30.346138Z",
			"UpdatedAt": "2025-11-21T19:17:30.346138Z",
			"DeletedAt": null,
			"text": "What's python?"
		}
	]
}
```

### 4️⃣ Удалить вопрос
`DELETE /questions/{id}`

**Responses:**

-   `200`: `Question was deleted successfully!`
-   `400`: `Invalid request. Please check the submitted data.`
-   `500`: `Internal server error`

---

### 1️⃣ Создать ответ

`POST /questions/{id}/answers`

Создаёт ответ в базе данных, но перед созданием проверяет существует ли вопрос.

**Пример запроса:**

```json
{
	"user_id": "550e8400-e29b-41d4-a716-446655440000",
	"text": "java is programming language"
} 
```

**Responses:**

-   `200`: `Answer was successfully created!`
-   `400`: `Invalid request. Please check the submitted data.`
-   `500`: `Internal server error`

### 2️⃣ Получить ответ

`GET /answers/{id}`

Получить конкретный ответ.

**Пример ответа:**

```json
{
	"Answer": {
		"ID": 1,
		"CreatedAt": "2025-11-22T08:10:44.706157Z",
		"UpdatedAt": "2025-11-22T08:10:44.706157Z",
		"DeletedAt": null,
		"question_id": 1,
		"user_id": "550e8400-e29b-41d4-a716-446655440000",
		"text": "java is programming language"
	}
}
```

### 3️⃣ Удалить ответ

`DELETE /answers/{id}`

Удаляет ответ.

**Responses:**

-   `200`: `Question was deleted successfully!`
-   `400`: `Invalid request. Please check the submitted data.`
-   `500`: `Internal server error`

---

## 🔖 Модели

### Question

```
{
  ID        uint           `gorm:"primaryKey"`
  Text      string         `gorm:"type:text;not null"`
  Answers   []model.Answer `gorm:"constraint:OnDelete:CASCADE;"`
  CreatedAt time.Time
  UpdatedAt *time.Time
}
```

### Answer
```
{
  *gorm.Model
  QuestionID uint   `gorm:"not null" json:"question_id"`
  UserID     string `gorm:"type:uuid;not null" json:"user_id"`
  Text       string `gorm:"type:text;not null" json:"text"`
}
```

---

## ⚡ Технологии

-   Go (Golang)
-   Работа с БД через GORM.
-   PostgreSQL
-   Docker & Docker Compose
-   goose (миграции)
-   логгирование, тесты (slog/testing)
