package main

import (
	"log"
	"net/http"
	"tuc2/db"
	"tuc2/handlers"
)

func main() {
	// Инициализация базы данных
	db.InitDB()

	// Маршруты
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/dashboard", handlers.AccountHandler)
	http.HandleFunc("/htmx/register-form", handlers.ShowRegisterForm)
	http.HandleFunc("/htmx/login-form", handlers.ShowLoginForm)

	http.HandleFunc("/logout", handlers.LogoutHandler)

	// Запуск сервера
	log.Println("Сервер работает на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
