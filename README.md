
Это микросервис регистрации и авторизации пользователей,
создан в целях учебного проекта.

Go + HTMX

tuc2 htmx

/tuc
├── handlers/             # HTTP-обработчики
│   ├── register.go       # Регистрация IndexHandler() RegisterHandler()
│   ├── login.go          # Авторизация LoginHandler()
│   ├── account.go        # Личный кабинет AccountHandler()
│   ├── htmx.go           # HTMX-фрагменты (формы) ShowLoginForm() ShowRegisterForm()
│
├── db/                  # Работа с БД
│   └── users.go         # Методы работы с пользователями AddUser() GetUserByEmail()
│
├── models/              # Структуры данных
│   └── user.go          # User и PageData
│
├── templates/           # HTML-шаблоны
│   ├── index.html       # Главная страница (форма входа/регистрации)
│   ├── login_form.html  # Форма входа (HTMX)
│   ├── register_form.html # Форма регистрации (HTMX)
│   └── dashboard.html   # Кабинет
│
├── main.go              # Точка входа, маршруты

Пакет handlers

register.go — всё, что связано с регистрацией

IndexHandler	Отображает основную страницу (index.html) с кнопками "Вход" / "Регистрация" и первым шаблоном формы
RegisterHandler	Обрабатывает форму регистрации: сохраняет пользователя в БД и перенаправляет в кабинет

login.go — авторизация

LoginHandler	Обрабатывает форму входа: проверяет пользователя в БД, сравнивает пароль, сохраняет сессию (в куки) и отправляет в кабинет

account.go — личный кабинет

AccountHandler	Проверяет наличие куки сессии, если всё ок — отображает личный кабинет (dashboard.html)

htmx.go — подгружаемые фрагменты для HTMX

ShowLoginForm	Отдаёт HTML-фрагмент формы входа (без всей страницы)
ShowRegisterForm	Отдаёт HTML-фрагмент формы регистрации

Пакет db

users.go — работа с пользователями в базе данных

AddUser	Добавляет нового пользователя в таблицу users
GetUserByEmail	Находит пользователя по email и возвращает структуру User

Пакет models
user.go — модели данных

User	Структура пользователя (email и хешированный пароль)
PageData	Данные для передачи в HTML-шаблоны (заголовок страницы, email и т.п.)

Шаблоны (templates/)

index.html	Главная страница с кнопками и контейнером под HTMX
login_form.html	Форма входа, загружается по HTMX
register_form.html	Форма регистрации, загружается по HTMX
dashboard.html	Личный кабинет пользователя

main.go — точка входа и маршруты

Настраивает базу данных и HTTP-маршруты
Обрабатывает: /, /login, /register, /account, /show-login, /show-register

Пользователь открывает / — видит кнопки "Вход / Регистрация"

HTMX загружает нужную форму — без перезагрузки

После отправки формы:

Регистрация → RegisterHandler → AddUser → куки → /account

Вход → LoginHandler → GetUserByEmail + проверка пароля → куки → /account

Кабинет загружается, если есть сессия