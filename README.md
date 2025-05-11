

````markdown
# Notes API

Простой REST-сервис для управления заметками в памяти.

## Запуск

```bash
go run main.go
````

Сервер будет доступен по адресу `http://localhost:8080`.

---

## Postman Collection

Вы можете импортировать следующие примеры запросов в Postman или вручную создать их по описанию.

### Переменные окружения

| Переменная    | Значение                   |
| ------------- | -------------------------- |
| `{{baseUrl}}` | `http://localhost:8080`    |
| `{{noteId}}`  | ID заметки (например, `1`) |

---

### 1. Получить все заметки

* **Метод**: `GET`
* **URL**: `{{baseUrl}}/api/notes`
* **Headers**:

  * `Accept: application/json`

#### Пример ответа (200 OK)

```json
[
  { "id": 1, "text": "First note" },
  { "id": 2, "text": "Another note" }
]
```

---

### 2. Создать новую заметку

* **Метод**: `POST`
* **URL**: `{{baseUrl}}/api/notes`
* **Headers**:

  * `Content-Type: application/json`
* **Body** (raw JSON):

```json
{
  "text": "My new note"
}
```

#### Пример ответа (201 Created)

```json
{
  "id": 3,
  "text": "My new note"
}
```

---

### 3. Получить заметку по ID

* **Метод**: `GET`
* **URL**: `{{baseUrl}}/api/notes/{{noteId}}`
* **Headers**:

  * `Accept: application/json`

#### Пример ответа (200 OK)

```json
{
  "id": 2,
  "text": "Another note"
}
```

Если заметка не найдена, возвращается `404 Not Found`.

---

### 4. Обновить текст заметки

* **Метод**: `PUT`
* **URL**: `{{baseUrl}}/api/notes/{{noteId}}`
* **Headers**:

  * `Content-Type: application/json`
* **Body** (raw JSON):

```json
{
  "text": "Updated note text"
}
```

#### Пример ответа (200 OK)

```json
{
  "id": 2,
  "text": "Updated note text"
}
```

---

### 5. Удалить заметку

* **Метод**: `DELETE`
* **URL**: `{{baseUrl}}/api/notes/{{noteId}}`

#### Пример ответа (204 No Content)

Нет тела ответа.

---

## Импорт коллекции в Postman

1. Откройте Postman.
2. Нажмите **Import** → **Raw Text**.
3. Вставьте URL или сохраните локально этот файл как `README.md`, затем при импорте выберите его.

Или создайте коллекцию вручную, добавив запросы, как описано выше.

> **Совет:** Заводите в окружении Postman переменные `baseUrl` и `noteId` для удобства тестирования.

---

Готово! Теперь вы можете быстро тестировать все эндпоинты вашего сервера через Postman.
