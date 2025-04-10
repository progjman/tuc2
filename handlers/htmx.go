package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// ShowLoginForm — отдаёт HTML-фрагмент формы входа (для HTMX)
func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "login_form.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Ошибка загрузки формы входа", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Ошибка рендера формы входа", http.StatusInternalServerError)
	}
}

// ShowRegisterForm — отдаёт HTML-фрагмент формы регистрации (для HTMX)
func ShowRegisterForm(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "register_form.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Ошибка загрузки формы регистрации!", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Ошибка рендера формы регистрации", http.StatusInternalServerError)
	}
}

// LogoutHandler — обработчик для выхода пользователя
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Удаляем куки с сессионным токеном
	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  "",
		MaxAge: -1, // Устанавливаем MaxAge в отрицательное значение для удаления куки
		Path:   "/",
	})

	// Перенаправляем на главную страницу
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
