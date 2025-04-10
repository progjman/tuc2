package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"tuc2/db"
	"tuc2/models"

	"strconv"
)

// RegisterHandler — обрабатывает POST-запрос формы регистрации
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if username == "" || email == "" || password == "" {
		http.Error(w, "Все поля обязательны", http.StatusBadRequest)
		return
	}

	user := models.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	userID, err := db.AddUser(user)
	if err != nil {
		http.Error(w, "Ошибка при регистрации", http.StatusInternalServerError)
		log.Println("Ошибка добавления пользователя:", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: strconv.Itoa(userID),
		Path:  "/",
	})

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

}

// IndexHandler — отображает главную страницу с кнопками и контейнером для HTMX
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Данные для страницы
	data := struct {
		Title string
	}{
		Title: "Главная страница",
	}

	// Путь к шаблону
	tmplPath := filepath.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Ошибка загрузки главной страницы", http.StatusInternalServerError)
		log.Println("Ошибка парсинга index:", err)
		return
	}

	// Отображаем страницу
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Ошибка рендера страницы", http.StatusInternalServerError)
		log.Println("Ошибка рендера index:", err)
	}
}
