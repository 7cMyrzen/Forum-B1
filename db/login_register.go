package db

import (
	"database/sql"
	"forum/lib/func/image"
	"forum/lib/func/security"
	"forum/lib/func/time"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {
	return db
}

// Fonction pour ajouter un utilisateur à la base de données pour registerHandler
func AddUser(username, email, password string) (int, error) {
	GetDB()
	creationDate := time.GetCurrentTime()
	description := "Nouvel utilisateur"
	profilePicture, _ := image.ImageToBase64("static/images/userpfp.jpg")

	insert, err := db.Prepare("INSERT INTO users(username, email, password, description, profile_picture, created_at) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer insert.Close()

	_, err = insert.Exec(username, email, security.HashPassword(password), description, profilePicture, creationDate)
	if err != nil {
		return 0, err
	}

	id, _ := FindUser(username, password)

	return id, nil
}

// Fonction pour chercher un utilisateur et mot de passe dans la base de données pour loginHandler qui renvoie l'ID de l'utilisateur ou 0 si l'utilisateur n'existe pas
func FindUser(username, password string) (int, error) {
	GetDB()
	var id int
	err := db.QueryRow("SELECT id FROM users WHERE (username = ? OR email = ?) AND password = ?", username, username, security.HashPassword(password)).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Foction pour se connecter avec soit l'emal ou le nom d'utilisateur et le mot de passe qiu renvoie l'ID de l'utilisateur ou 0 si l'utilisateur n'existe pas
func AuthenticateUser(username, password string) (int, error) {
	GetDB()
	var id int
	err := db.QueryRow("SELECT id FROM users WHERE (username = ? OR email = ?) AND password = ?", username, username, security.HashPassword(password)).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
