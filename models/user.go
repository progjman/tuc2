package models

// User представляет пользователя в системе
type User struct {
	ID       int    // ID пользователя в базе данных
	Username string // Имя пользователя
	Email    string // Почтовый адрес
	Password string // Хэшированный пароль
}

// PageData содержит данные, которые передаются в шаблон страницы
type PageData struct {
	Title string // Название страницы
	User  User   // Данные пользователя для отображения
}
