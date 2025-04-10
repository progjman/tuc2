package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"tuc2/db"
	"tuc2/models"
)

// AccountHandler — Личный кабинет
func AccountHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		log.Println("Нет куки session_token")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	userID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Println("Ошибка преобразования session_token:", cookie.Value)
		http.Error(w, "Неверный токен", http.StatusUnauthorized)
		return
	}

	user, err := db.GetUserByID(userID)
	if err != nil {
		log.Println("Ошибка получения пользователя по ID:", err)
		http.Error(w, "Ошибка получения пользователя", http.StatusInternalServerError)
		return
	}

	// Отображаем личный кабинет
	tmplPath := filepath.Join("templates", "dashboard.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Ошибка загрузки страницы личного кабинета", http.StatusInternalServerError)
		return
	}

	// Передаем данные пользователя в шаблон
	pageData := models.PageData{
		Title: "Личный кабинет",
		User:  user,
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		http.Error(w, "Ошибка рендера страницы личного кабинета", http.StatusInternalServerError)
	}
}
