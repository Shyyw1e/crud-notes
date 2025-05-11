#crud-notes


2. GET /list — получить все заметки
Создайте новый запрос в Postman.

Выберите метод GET.

Введите URL без завершающего слэша:

--bash
http://localhost:8080/api/notes
Нажмите Send.
--
Если в памяти ещё нет заметок, вы увидите:

json
Копировать
Редактировать
[]
Если запрос возвращает 404 page not found, проверьте, что:

Вы именно на /api/notes, а не на /api/notes/

Сервер действительно слушает порт 8080

В Postman выбран именно HTTP (не HTTPS) и правильный хост

3. POST /api/notes — создать новую заметку
Создайте новый запрос.

Выберите метод POST.

URL такой же:

bash
Копировать
Редактировать
http://localhost:8080/api/notes
Перейдите на вкладку Headers и добавьте:

pgsql
Копировать
Редактировать
Content-Type: application/json
Перейдите на вкладку Body, выберите raw → JSON, и введите, например:

json
Копировать
Редактировать
{ "text": "Hello, Postman!" }
Нажмите Send.

Ожидаемый ответ:

http
Копировать
Редактировать
HTTP/1.1 201 Created
Content-Type: application/json

{ "id": 1, "text": "Hello, Postman!" }
4. GET /api/notes/{id} — получить одну заметку
Метод GET.

URL, например, чтобы получить заметку с id=1:

bash
Копировать
Редактировать
http://localhost:8080/api/notes/1
Send.

Ожидаемый ответ:

json
Копировать
Редактировать
{ "id": 1, "text": "Hello, Postman!" }
5. PUT /api/notes/{id} — обновить заметку
Метод PUT.

URL:

bash
Копировать
Редактировать
http://localhost:8080/api/notes/1
В Headers:

pgsql
Копировать
Редактировать
Content-Type: application/json
В Body (raw / JSON):

json
Копировать
Редактировать
{ "text": "Updated text" }
Send.

Ожидаемый ответ:

http
Копировать
Редактировать
HTTP/1.1 200 OK
Content-Type: application/json

{ "id": 1, "text": "Updated text" }
6. DELETE /api/notes/{id} — удалить заметку
Метод DELETE.

URL:

bash
Копировать
Редактировать
http://localhost:8080/api/notes/1
Send.

Ожидаемый ответ:

http
Копировать
Редактировать
HTTP/1.1 204 No Content
После этого повторный GET /api/notes/1 вернёт 404 Not Found.
