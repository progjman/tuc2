package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"tuc2/models"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // Подключаем драйвер PostgreSQL
)

// DB — глобальная переменная для работы с БД
var DB *sql.DB

// Инициализация базы данных
func InitDB() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Нет загрузки .env")
	}

	connStr := os.Getenv("DB_CONN") // Строка подключения к PostgreSQL

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	// Проверяем соединение с БД
	err = DB.Ping()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	// Создаём таблицу пользователей, если её нет
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatal("Ошибка создания таблицы пользователей:", err)
	}

}

// AddUser — добавление нового пользователя в БД
func AddUser(user models.User) (int, error) {
	var userID int
	err := DB.QueryRow(
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id",
		user.Username, user.Email, user.Password,
	).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

// GetUserByUsername — получение пользователя по никнейму
func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := DB.QueryRow("SELECT id, username, email, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("пользователь не найден")
		}
		return user, err
	}
	return user, nil
}

// GetUserByID — получение пользователя по ID
func GetUserByID(userID int) (models.User, error) {
	var user models.User
	err := DB.QueryRow("SELECT id, username, email, password FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("пользователь не найден")
		}
		return user, err
	}
	return user, nil
}
