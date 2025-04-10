package handlers

import (
	"log"
	"net/http"
	"strconv"
	"tuc2/db"
)

// LoginHandler — обрабатывает форму входа по никнейму
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ShowLoginForm(w, r)
		return
	}

	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	log.Println("Пытаемся войти: username =", username, "password =", password)

	user, err := db.GetUserByUsername(username)
	if err != nil {
		http.Error(w, "Пользователь не найден", http.StatusUnauthorized)
		log.Println("Пользователь не найден:", err)
		return
	}

	if user.Password != password {
		http.Error(w, "Неверный пароль", http.StatusUnauthorized)
		log.Println("Пароль не совпадает. Ожидали:", user.Password)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: strconv.Itoa(user.ID),
		Path:  "/",
		// HttpOnly: true, Secure: true // можно позже включить
	})

	log.Println("Вход успешный. Перенаправляем на /dashboard")

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
